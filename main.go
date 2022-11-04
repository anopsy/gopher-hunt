package main

import (
	"errors"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

/*
type Gopher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
*/
type User struct {
	ID       string `json:"id"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//Record   []Gopher `json:"record"`
	Points int `json:"points"`
}

/*
var gophers = []Gopher{
	{ID: "1", Name: "Gogo"},
	{ID: "2", Name: "Gofoo"},
	{ID: "3", Name: "Gobla"},
	{ID: "4", Name: "Gofee"},
	{ID: "5", Name: "Golee"},
	{ID: "6", Name: "Goblin"},
	{ID: "7", Name: "Goshee"},
	{ID: "8", Name: "Gocee"},
	{ID: "9", Name: "Gonee"},
	{ID: "10", Name: "Googoo"},
}
*/
var users = []User{
	{ID: "1", Nick: "GoMaster", Email: "gomaster@gmail.com", Password: "6969Olaboga", Points: 0},
	{ID: "2", Nick: "Swichu", Email: "swichu1111@gmail.com", Password: "ochuj420", Points: 0},
	{ID: "3", Nick: "Anopsy", Email: "anopsy28@gmail.com", Password: "2137jp2gmd", Points: 0},
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getRanking(c *gin.Context) {

	sort.Slice(users, func(i, j int) bool { return users[i].Points < users[i].Points })
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(id string) (*User, error) {
	for i, u := range users {
		if u.ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")

}

func addGopher(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}
	user, err := getUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
		return
	}
	user.Points++
	c.IndentedJSON(http.StatusOK, user)

}
func main() {
	router := gin.Default()
	router.POST("/users", createUser)
	router.GET("/users", getRanking)
	router.PUT("/user/:id", addGopher)

	router.Run("localhost:8000")
}
