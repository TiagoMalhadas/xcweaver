// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package reflection implements helpers for reflection code.
package reflection

import (
	"fmt"
	"reflect"
)

// Type returns the reflect.Type for T.
//
// This function is particularly useful when T is an interface
// and it is impossible to get a value with concrete type T.
func Type[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

// ComponentName returns the name of the component of type T.
// Note that T must be the interface type of the component, not its implementation type.
func ComponentName[T any]() string {
	t := Type[T]()
	return fmt.Sprintf("%s/%s", t.PkgPath(), t.Name())
}
