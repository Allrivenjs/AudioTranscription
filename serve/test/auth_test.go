package test

import (
	"AudioTranscription/serve/controllers"
	"github.com/gavv/httpexpect/v2"
	"testing"
)

func Test_LoginTest(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{Reporter: httpexpect.NewAssertReporter(t)})
	up := controllers.SignIn{
		Email:    "test2@gmail.com",
		Password: "admin",
	}
	status := e.POST("/auth/signin", up).Expect().Status(200)
	t.Log(status)
}
