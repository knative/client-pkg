// Copyright © 2020 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package source

import (
	"testing"

	"gotest.tools/v3/assert"

	"knative.dev/client-pkg/pkg/kn-source-pkg/pkg/types"
)

func TestNewSourceCommand(t *testing.T) {
	sourceCmd := NewSourceCommand(&types.KnSourceParams{})
	assert.Assert(t, sourceCmd != nil)
	assert.Equal(t, sourceCmd.Use, "source")
	assert.Equal(t, sourceCmd.Short, "Knative eventing {{.Name}} source plugin")
	assert.Equal(t, sourceCmd.Long, "Manage your Knative {{.Name}} eventing sources")
}
