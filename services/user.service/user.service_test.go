package user_service_test

import (
	"admin-server/models"
	userService "admin-server/services/user.service"
	"testing"
)

func TestCreate(t *testing.T) {

	t.Log("Creación de usuario.")
	user := models.NewUser()
	user.Name = "Test4"
	user.Email = "test4@correo.com"
	user.Password = []byte("123456")

	if err := userService.Create(user); err != nil {
		t.Error("La prueba de creación de usuario ha fallado.")
		t.Error(err)
		t.Fail()
	} else {
		t.Log("Prueba de creación exitosa.")
	}
}

func TestGetUserById(t *testing.T) {
	t.Log("Obtener un usuario por Id")

	_, err := userService.GetUserById("60adc562881169ecf1af90b8")
	if err != nil {
		t.Error("La prueba de buscar un usuario por id ha fallado.")
		t.Error(err)
		t.Fail()
	} else {
		t.Log("Prueba de obtener usuario exitosa.")
	}

}

func TestGetUsers(t *testing.T) {
	t.Log("Obtener todos los usuarios")

	_, err := userService.GetUsers()
	if err != nil {
		t.Error("La prueba de obtener todos los usuarios ha fallado.")
		t.Error(err)
		t.Fail()
	} else {
		t.Log("Prueba de obtener usuarios exitosa.")
	}
}

func BenchmarkCreate(t *testing.B) {
	user := models.NewUser()
	user.Name = "José Palma"
	user.Email = "jcpalma@correo.com"
	user.Password = []byte("123456")
	_ = userService.Create(user)
}
