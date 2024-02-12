// Copyright 2024 The Wait4X Authors
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

package mysql

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"testing"
	"wait4x.dev/v2/checker"
)

// MySQLSuite is a test suite for MySQL checker
type MySQLSuite struct {
	suite.Suite
	container *mysql.MySQLContainer
}

// SetupSuite starts a MySQL container
func (s *MySQLSuite) SetupSuite() {
	var err error
	s.container, err = mysql.RunContainer(context.Background())
	s.Require().NoError(err)
}

// TearDownSuite stops the MySQL container
func (s *MySQLSuite) TearDownSuite() {
	err := s.container.Terminate(context.Background())
	s.Require().NoError(err)
}

// TestIdentity tests the identity of the MySQL checker
func (s *MySQLSuite) TestIdentity() {
	chk := New("user:password@tcp(localhost:3306)/dbname?tls=skip-verify")
	identity, err := chk.Identity()

	s.Require().NoError(err)
	s.Assert().Equal("localhost:3306", identity)
}

// TestInvalidIdentity tests the invalid identity of the MySQL checker
func (s *MySQLSuite) TestInvalidIdentity() {
	chk := New("xxx://127.0.0.1:3306")
	_, err := chk.Identity()

	s.Assert().ErrorContains(err, "default addr for network 'xxx:/' unknown")
}

// TestValidConnection tests the valid connection of the MySQL server
func (s *MySQLSuite) TestInvalidConnection() {
	var expectedError *checker.ExpectedError
	chk := New("user:password@tcp(localhost:8080)/dbname?tls=skip-verify")

	s.Assert().ErrorAs(chk.Check(context.Background()), &expectedError)
}

// TestValidAddress tests the valid address of the MySQL server
func (s *MySQLSuite) TestValidAddress() {
	ctx := context.Background()

	endpoint, err := s.container.ConnectionString(ctx)
	s.Require().NoError(err)

	chk := New(endpoint)
	s.Assert().Nil(chk.Check(ctx))
}

// TestMySQL runs the MySQL test suite
func TestMySQL(t *testing.T) {
	suite.Run(t, new(MySQLSuite))
}
