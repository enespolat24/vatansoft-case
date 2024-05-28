package database

import (
	"fmt"
	"os"
	"testing"
	"vatansoft-case/internal/model"

	"github.com/ory/dockertest/v3" // test container
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		fmt.Printf("Could not connect to Docker: %s", err)
		os.Exit(1)
	}

	resource, err := pool.Run("mysql", "latest", []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_DATABASE=testdb"})
	if err != nil {
		fmt.Printf("Could not start resource: %s", err)
		os.Exit(1)
	}

	defer func() {
		if err := pool.Purge(resource); err != nil {
			fmt.Printf("Could not purge resource: %s", err)
		}
	}()

	code := m.Run()
	os.Exit(code)
}

func TestNewDbConnection(t *testing.T) {
	conn := NewDbConnection()
	if conn == nil {
		t.Errorf("DbConnection should not be nil")
	}
}

func TestAutoMigrate(t *testing.T) {
	conn := NewDbConnection()

	err := conn.AutoMigrate()
	if err != nil {
		t.Errorf("AutoMigrate returned an error: %v", err)
	}

	var count int64
	result := conn.GetDbInstance().Model(&model.User{}).Count(&count)
	if result.Error != nil {
		t.Errorf("Count operation returned an error: %v", result.Error)
	}
	if count != 0 {
		t.Errorf("Initial user count should be 0, got: %d", count)
	}
}
