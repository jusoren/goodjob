package goodjob_test

import (
	"os"
	"testing"

	"github.com/jusoren/goodjob"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var client *goodjob.Client

func TestMain(m *testing.M) {
	var err error

	// Establish a connection to the database
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&goodjob.Job{}, &goodjob.Execution{})

	// Create a driver
	driver := goodjob.NewDriverGorm(db)

	// Create a client
	client = goodjob.NewClient(driver)

	code := m.Run()

	// Close the connection to the database
	os.Remove("test.db")

	os.Exit(code)
}
