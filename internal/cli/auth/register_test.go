// Copyright 2022 MongoDB Inc
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

//go:build unit
// +build unit

package auth

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		RegisterBuilder(),
		0,
		[]string{"gov", "noBrowser"},
	)
}

func Test_registerOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockAuthenticator(ctrl)
	mockConfig := mocks.NewMockLoginConfig(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)
	defer ctrl.Finish()
	buf := new(bytes.Buffer)

	loginOpts := &loginOpts{
		flow:       mockFlow,
		config:     mockConfig,
		NoBrowser:  true,
		SkipConfig: true,
	}

	opts := &RegisterOpts{
		*loginOpts,
	}

	opts.OutWriter = buf
	opts.Store = mockStore

	require.NoError(t, opts.Run())
	assert.Equal(t, `Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activating your account.
`, buf.String())
}