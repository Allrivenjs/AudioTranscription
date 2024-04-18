package test

import (
	"AudioTranscription/serve/controllers"
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/models"
	"AudioTranscription/serve/repository"
	"github.com/gavv/httpexpect/v2"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

var conn db.Connection

// load after the test
func TestMain(m *testing.M) {
	// load the database
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panicln(err)
	}
	conn = db.NewConnection()
	models.AutoMigrate(conn)

	m.Run()
}

func Test_LoginTest(t *testing.T) {
	// register user into the database
	usersRepo := repository.NewUsersRepository(conn)
	user := &models.User{
		FirstName: "Test1",
		LastName:  "Test",
		Email:     "test2@gmail.com",
		Password:  "admin",
	}
	err := usersRepo.SaveOrUpdate(user)
	if err != nil {
		t.Errorf("Error saving user: %v", err)
	}
	e := httpexpect.WithConfig(httpexpect.Config{Reporter: httpexpect.NewAssertReporter(t)})
	up := controllers.SignIn{
		Email:    "test2@gmail.com",
		Password: "admin",
	}
	status := e.POST("/auth/signin", up).Expect().Status(200)
	t.Log(status)
}
