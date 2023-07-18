// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package precompile

import (
	"errors"
	"fmt"
	"reflect"
)

// ValidateBasic checks if the precompile method argument and return types match
// the abi's argument return types. this is for overloaded Solidity functions and general
// validation.
func (m *Method) ValidateBasic() error {
	implMethod := m.execute
	abiMethod := m.abiMethod

	if len(m.abiMethod.Outputs) == 0 {
		// the Solidity compiler requires that precompiles must return at least one value.
		// see https://github.com/berachain/polaris/issues/491 for more information.

		//nolint:lll // error message.
		panic("The Solidity compiler requires all precompile functions to return at least one value. Consider returning a boolean.")
	}

	// first two args of Go precompile implementation are the receiver contract and the
	// context, so we skip those.
	if implMethod.Type.NumIn()-2 != len(abiMethod.Inputs) {
		return errors.New("number of arguments mismatch")
	}

	// last parameter of Go precompile implementation is an error (for reverts),
	// so we skip that.
	if implMethod.Type.NumOut()-1 != len(abiMethod.Outputs) {
		return errors.New("number of return types mismatch")
	}

	// if the function does not take any inputs, no need to check.
	// note again that for NumIn(), we check for 2 args, because the first two are the receiver and
	// ctx due to the nature of Go's `reflect` package.
	if implMethod.Type.NumIn() == 2 && len(abiMethod.Inputs) == 0 {
		return nil
	}

	// receiver is 0th param, context is 1st param, so skip those.
	for i := 2; i < implMethod.Type.NumIn(); i++ {
		implMethodParamType := implMethod.Type.In(i)
		abiMethodParamType := abiMethod.Inputs[i-2].Type.GetType()
		if err := validateArg(implMethodParamType, abiMethodParamType); err != nil {
			return fmt.Errorf("argument type mismatch: %v != %v", implMethodParamType, abiMethodParamType)
		}
	}

	for i := 0; i < implMethod.Type.NumOut()-1; i++ {
		implMethodReturnType := implMethod.Type.Out(i)
		abiMethodReturnType := abiMethod.Outputs[i].Type.GetType()
		if err := validateArg(implMethodReturnType, abiMethodReturnType); err != nil {
			return fmt.Errorf("return type mismatch: %v != %v", implMethodReturnType, abiMethodReturnType)
		}
	}

	return nil
}

// helper function for ValidateBasic. this function function uses reflection to see
// what types your implementation uses, and checks against the geth representation of the abi types.
func validateArg(implMethodVarType reflect.Type, abiMethodVarType reflect.Type) error {
	//nolint:exhaustive // nah, this is fine.
	switch implMethodVarType.Kind() {
	case abiMethodVarType.Kind(), reflect.Interface:
		// if the Go type matches the abi type, we're good.
		// if it's `any`, we leave it to the user to make sure that it is used/converted correctly.
		return nil
	case reflect.Struct:
		if err := validateStructFields(implMethodVarType, abiMethodVarType); err != nil {
			return err
		}
	case reflect.Slice, reflect.Array:
		for j := 0; j < abiMethodVarType.Len(); j++ {
			// if it is a slice/array of structs, then we need to check if the struct fields match
			if abiMethodVarType.Elem().Kind() == reflect.Struct {
				if err := validateStructFields(
					implMethodVarType.Elem(),
					abiMethodVarType.Elem(),
				); err != nil {
					return err
				}
				// any other case, we just check the elements.
			} else if implMethodVarType.In(j) != abiMethodVarType.In(j) {
				return fmt.Errorf("return type mismatch: %v != %v",
					implMethodVarType.Elem(),
					abiMethodVarType.Elem(),
				)
			}
		}
	default:
		return fmt.Errorf("return type mismatch: %v != %v", implMethodVarType, abiMethodVarType)
	}
	return nil
}

// this function checks to make sure that the struct fields match. if there is a nested struct, then
// we recurse until we reach the base case of a struct composing of only primitive types.
func validateStructFields(implMethodVarType reflect.Type,
	abiMethodVarType reflect.Type,
) error {
	if implMethodVarType == nil && abiMethodVarType == nil {
		return nil
	}
	if implMethodVarType.NumField() != abiMethodVarType.NumField() {
		return errors.New("number of return types mismatch")
	}
	for j := 0; j < implMethodVarType.NumField(); j++ {
		// if the field is a nested struct, then we recurse

		if implMethodVarType.Field(j).Type.Kind() == reflect.Struct &&
			abiMethodVarType.Field(j).Type.Kind() == reflect.Struct {
			if err := validateStructFields(
				implMethodVarType.Field(j).Type,
				abiMethodVarType.Field(j).Type,
			); err != nil {
				return err
			}
		} else if implMethodVarType.Field(j).Type != abiMethodVarType.Field(j).Type {
			return fmt.Errorf("return type mismatch: %v != %v",
				implMethodVarType.Field(j).Type,
				abiMethodVarType.Field(j).Type,
			)
		}
	}
	return nil
}
