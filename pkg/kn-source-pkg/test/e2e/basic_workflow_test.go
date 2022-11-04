// Copyright 2020 The Knative Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build e2e
// +build !eventing

package e2e

import (
	"os"
	"path/filepath"
	"testing"

	"gotest.tools/v3/assert"
	"knative.dev/client/lib/test"
	"knative.dev/client/pkg/util"
)

func TestBasicWorkflow(t *testing.T) {
	t.Parallel()

	currentDir, err := os.Getwd()
	assert.NilError(t, err)

	it, err := NewE2ETest("kn-source_pkg", filepath.Join(currentDir, "../.."), false)
	assert.NilError(t, err)
	defer func() {
		assert.NilError(t, it.KnTest().Teardown())
	}()

	r := test.NewKnRunResultCollector(t, it.KnTest())
	defer r.DumpIfFailed()

	err = it.KnPlugin().Install()
	assert.NilError(t, err)

	t.Log("kn-source_pkg list")
	it.knSourceList(t, r)

	t.Log("kn-source_pkg create 'source-name' with 'sink-name'")
	it.knSourceCreate(t, r, "source-name", "sink-name")

	t.Log("kn-source_pkg describe 'source-name'")
	it.knSourceDescribe(t, r, "source-name")

	t.Log("kn-source_pkg update 'source-name' with 'new-sink-name'")
	it.knSourceUpdate(t, r, "source-name", "new-sink-name")

	t.Log("kn-source_pkg delete 'source-name'")
	it.knSourceDelete(t, r, "source-name", "sink-name")

	err = it.KnPlugin().Uninstall()
	assert.NilError(t, err)
}

// Private
func (it *E2ETest) knSourceList(t *testing.T, r *test.KnRunResultCollector) {
	out := it.KnPlugin().Run("list")
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "list"))
}

func (it *E2ETest) knSourceCreate(t *testing.T, r *test.KnRunResultCollector, sourceName, sinkName string) {
	out := it.KnPlugin().Run("create", sourceName, "--sink", sinkName)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "create", sourceName, "namespace", it.KnTest().Namespace(), "sink", sinkName))
}

func (it *E2ETest) knSourceDescribe(t *testing.T, r *test.KnRunResultCollector, sourceName string) {
	out := it.KnPlugin().Run("describe", sourceName)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "describe", sourceName, "namespace", it.KnTest().Namespace()))
}

func (it *E2ETest) knSourceUpdate(t *testing.T, r *test.KnRunResultCollector, sourceName, sinkName string) {
	out := it.KnPlugin().Run("update", sourceName, "--sink", sinkName)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "update", sourceName, "namespace", it.KnTest().Namespace(), "sink", sinkName))
}

func (it *E2ETest) knSourceDelete(t *testing.T, r *test.KnRunResultCollector, sourceName, sinkName string) {
	out := it.KnPlugin().Run("delete", sourceName)
	r.AssertNoError(out)
	assert.Check(t, util.ContainsAllIgnoreCase(out.Stdout, "delete", sourceName, "namespace", it.KnTest().Namespace()))
}
