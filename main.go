package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:uuid", getUserByUUID)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
}

// user represents data about a record user.
type User struct {
	UUID     uuid.UUID `json:"UUID"`
	Username string    `json:"username"`
}

func init_users() []User {
	uuid1, _ := uuid.NewRandom()
	uuid2, _ := uuid.NewRandom()
	return []User{
		{UUID: uuid1, Username: "Alex"},
		{UUID: uuid2, Username: "George"},
	}
}

var users = init_users()

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser User

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new user to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// getUserByUUID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByUUID(c *gin.Context) {
	UUID, _ := uuid.Parse(c.Param("uuid"))
	// Loop over the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.UUID == UUID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
