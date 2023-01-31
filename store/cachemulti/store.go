// Copyright (C) 2022, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package cachemulti

import (
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/lib/utils"
	sdkcachekv "github.com/cosmos/cosmos-sdk/store/cachekv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Compile-time check to ensure `Store` implements `storetypes.CacheMultiStore`.
var (
	_ storetypes.CacheMultiStore = (*Store)(nil)
	_ libtypes.Snapshottable     = (*Store)(nil)
)

// CACHEWRAPS
// `Store` is a wrapper around the Cosmos SDK `MultiStore` which supports snapshots and reverts.
// It stores revisions by cache-wrapping the cachekv stores on a call to `Snapshot`.
type Store struct {
	storetypes.MultiStore

	stores ds.Stack[map[storetypes.StoreKey]storetypes.CacheKVStore]
}

// `NewStoreFrom` creates and returns a new `Store` from a given MultiStore.
func NewStoreFrom(ms storetypes.MultiStore) *Store {
	stores := stack.New[map[storetypes.StoreKey]storetypes.CacheKVStore](32)
	stores.Push(make(map[storetypes.StoreKey]storetypes.CacheKVStore))
	return &Store{
		MultiStore: ms,
		stores:     stores,
	}
}

// `GetKVStore` shadows the SDK's `storetypes.MultiStore` function. Routes native module calls to
// read the dirty state during an eth tx. Any state that is modified by evm statedb, and using the
// context passed in to StateDB, will be routed to a tx-specific cache kv store.
func (s *Store) GetKVStore(key storetypes.StoreKey) storetypes.KVStore {
	// check if cache kv store already used
	curr := s.stores.Peek()
	if cacheKVStore, exists := curr[key]; exists {
		return cacheKVStore
	}

	// get kvstore from cachemultistore and set cachekv to memory
	curr[key] = sdkcachekv.NewStore(s.MultiStore.GetKVStore(key))
	return curr[key]
}

// `Snapshot` implements `libtypes.Snapshottable`.
func (s *Store) Snapshot() int {
	curr := s.stores.Peek()
	next := make(map[storetypes.StoreKey]storetypes.CacheKVStore)
	for key := range curr {
		next[key] = utils.MustGetAs[storetypes.CacheKVStore](curr[key].CacheWrap())
	}
	defer func() {
		s.stores.Push(next)
	}()

	return s.stores.Size()
}

// `Revert` implements `libtypes.Snapshottable`.
func (s *Store) RevertToSnapshot(id int) {
	s.stores.PopToSize(id)
	if id == 0 {
		s.stores.Push(make(map[storetypes.StoreKey]storetypes.CacheKVStore))
	}
}

// `Write` commits each of the individual cachekv stores to its corresponding parent kv stores.
//
// `Write` implements Cosmos SDK `storetypes.CacheMultiStore`.
func (s *Store) Write() {
	// to allow garbage collector to vibe
	for i := s.stores.Size() - 1; i >= 0; i-- {
		revision := s.stores.Pop()

		// Safe from non-determinism, since order in which
		// we write to the parent kv stores does not matter.
		//
		//#nosec:G705
		for key, store := range revision {
			store.Write()
			delete(revision, key)
		}
	}
}
