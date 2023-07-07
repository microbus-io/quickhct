/*
Copyright 2023 Microbus LLC and various contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package quickhct

import (
	"fmt"
	"testing"
)

func Test_WithTone(t *testing.T) {
	for _, tone := range []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 95, 98, 100} {
		r, g, b := WithTone(0x71, 0x55, 0x73, tone)
		fmt.Printf("%3d %02x%02x%02x\n", tone, r, g, b)
	}
}
