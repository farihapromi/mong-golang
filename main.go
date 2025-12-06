package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("")
	r.POST("")
	r.DELETE("")
}
func getSession() *mgo.Session {

}
