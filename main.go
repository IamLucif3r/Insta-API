package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/IamLucif3r/Insta-API/controllers"
)

var lock sync.Mutex

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	go safe_thread("Starting Application ...")
	time.Sleep(3 * time.Second)
	r.GET("/user/:id", uc.GetUser)
	r.GET("/posts/:id", uc.GetPost)
	r.POST("/user", uc.CreateUser)
	r.POST("/posts", uc.CreatePost)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	return s
}

// Safe-Thread Function
func safe_thread(name string) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Println(name)
	time.Sleep(1 * time.Second)
}
