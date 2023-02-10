// Copyright (C) 2023, Berachain Foundation. All rights reserved.
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

package core

import (
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/params"
)

type StateFactory struct {
	config *params.ChainConfig

	host StargazerHostChain
}

func NewStateFactory(config *params.ChainConfig, host StargazerHostChain) *StateFactory {
	return &StateFactory{
		config: config,
		host:   host,
	}
}

func (sf *StateFactory) BuildStateProcessor(vmConfig vm.Config, commit bool) *StateProcessor {
	return NewStateProcessor(sf.host, state.NewStateDB(sf.host.GetStatePlugin()), vmConfig, commit)
}
