package database

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// Define a test suite
type DatabaseTestSuite struct {
	suite.Suite
	db *Database
}

// This will run once before all tests in the suite
func (s *DatabaseTestSuite) SetupSuite() {
	// Initialize resources that are shared by all tests
	InitializeTestEnvironment()
}

// This will run once after all tests in the suite
func (s *DatabaseTestSuite) TearDownSuite() {
	// Clean up resources used by the entire suite
	CleanupTestEnvironment()
}

// This will run before each test
func (s *DatabaseTestSuite) SetupTest() {
	// Create a fresh database instance for each test
	s.db = NewTestDatabase()
}

// This will run after each test
func (s *DatabaseTestSuite) TearDownTest() {
	// Close the database connection after each test
	s.db.Close()
}

// Test methods must begin with "Test"
func (s *DatabaseTestSuite) TestInsert() {
	err := s.db.Insert("key1", "value1")
	s.NoError(err, "insert should not fail")

	value, err := s.db.Get("key1")
	s.NoError(err, "get should not fail")
	s.Equal("value1", value, "retrieved value should match inserted value")
}

func (s *DatabaseTestSuite) TestDelete() {
	// Insert first
	err := s.db.Insert("key2", "value2")
	s.NoError(err, "insert should not fail")

	// Then delete
	err = s.db.Delete("key2")
	s.NoError(err, "delete should not fail")

	// Verify it's gone
	_, err = s.db.Get("key2")
	s.Error(err, "getting deleted key should fail")
}

// This function runs the suite
func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
