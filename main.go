package main

import (
	"admin-server/database"
	"admin-server/models"
	userService "admin-server/services/user.service"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Backend Server 'AdminPro'")

	user := models.User{
		Name:      "José Palma",
		Email:     "jcpalma@correo.com",
		Password:  "123456",
		Status:    "A",
		CreatedAt: time.Now().Unix(),
		Role:      "ADMIN_ROLE",
	}

	if err := userService.Create(user); err != nil {
		fmt.Printf("No se pudo crear el usario:\n%s\n", err)
	}

	defer database.Disconnect()
}
