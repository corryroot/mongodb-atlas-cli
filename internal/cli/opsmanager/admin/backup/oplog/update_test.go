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
// +build unit

package oplog

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOplogsUpdater(ctrl)
	defer ctrl.Finish()

	expected := &opsmngr.BackupStore{}

	opts := &UpdateOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().UpdateOplog(opts.ID, opts.NewBackupStore()).
		Return(expected, nil).
		Times(1)

	err := opts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateBuilder(t *testing.T) {
	cli.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.Output, flag.EncryptedCredentials, flag.MaxCapacityGB,
			flag.Assignment, flag.Label, flag.URI, flag.WriteConcern,
			flag.EncryptedCredentials, flag.SSL},
	)
}