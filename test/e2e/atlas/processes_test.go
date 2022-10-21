// Copyright 2020 MongoDB Inc
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
//go:build e2e || (atlas && processes)

package atlas_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
	exec "golang.org/x/sys/execabs"
)

func TestProcesses(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProjectAndCluster("processes")

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	var processes []*mongodbatlas.Process

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			processesEntity,
			"list",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if err := json.Unmarshal(resp, &processes); assert.NoError(t, err) {
			require.NotEmpty(t, processes)
		}
	})

	t.Run("describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			processesEntity,
			"describe",
			processes[0].ID,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var process *mongodbatlas.Process
		if err := json.Unmarshal(resp, &process); assert.NoError(t, err) {
			assert.Equal(t, process.ID, processes[0].ID)
		}
	})
}