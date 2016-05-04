// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Author: Radu Berinde (radu@cockroachlabs.com)

package sql

import (
	"testing"

	"github.com/cockroachdb/cockroach/client"
	"github.com/cockroachdb/cockroach/server/testingshim"
	"github.com/cockroachdb/cockroach/util/leaktest"
)

// Temporary proof-of-concept test that uses the testingshim to set up a test
// server from the sql package.
func TestPOC(t *testing.T) {
	defer leaktest.AfterTest(t)()

	s := testingshim.NewTestServer()
	if err := s.Start(); err != nil {
		t.Fatal(err)
	}
	defer s.Stop()

	kvClient := s.ClientDB().(*client.DB)
	err := kvClient.Put("testkey", "testval")
	if err != nil {
		t.Fatal(err)
	}
	kv, err := kvClient.Get("testkey")
	if err != nil {
		t.Fatal(err)
	}
	if kv.PrettyValue() != `"testval"` {
		t.Errorf(`Invalid Get result: %s, expected "testval"`, kv.PrettyValue())
	}
}
