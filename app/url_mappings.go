package app

import (
	"github.com/Vineeth-BekalWork/bookstores_users-api/controllers/ping"
	"github.com/Vineeth-BekalWork/bookstores_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
