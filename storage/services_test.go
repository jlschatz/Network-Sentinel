package storage

import "testing"

func TestStore_NewDatabaseConnection(t *testing.T) {

	s := NewStore()
	err := s.NewDatabaseConnection()
	if err != nil {
		t.Errorf("NewDatabaseConnection() error = %v", err)
	}
	s.CloseDatabaseConnection()
}

func TestStore_CloseDatabaseConnection(t *testing.T) {
	s := NewStore()
	s.NewDatabaseConnection()
	err := s.CloseDatabaseConnection()
	if err != nil {
		t.Errorf("CloseDatabaseConnection() error = %v", err)
	}
}

//func TestStore_SetTestEnvironmentVariables(t *testing.T) {
//	s := NewStore()
//	err := s.SetTestEnvironmentVariables()
//	if err != nil {
//		t.Errorf("SetTestEnvironmentVariables() error = %v", err)
//	}
//	s.CloseDatabaseConnection()
//}
