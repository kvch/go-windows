// Copyright 2018 Elasticsearch Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build windows

package windows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTickCount64(t *testing.T) {
	millis, err := GetTickCount64()
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, millis > 0, "millis (%d) must be greater than 0", millis)
}
