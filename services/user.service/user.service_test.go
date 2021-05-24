package user_service_test

import (
	"admin-server/models"
	userService "admin-server/services/user.service"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {

	t.Log("Creación de usuario.")

	user := models.User{
		Name:      "José Palma",
		Email:     "jcpalma@correo.com",
		Password:  "123456",
		Status:    "A",
		CreatedAt: time.Now().Unix(),
		Role:      "ADMIN_ROLE",
	}

	if err := userService.Create(user); err != nil {
		t.Error("La prueba de creación de usuario a fallado.")
		t.Error(err)
		t.Fail()
	} else {
		t.Log("Prueba de creación exitosa.")
	}

}
