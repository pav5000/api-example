package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pav5000/api-example/internal/handlers/createuser"
	"github.com/pav5000/api-example/internal/handlers/userbyid"
	"github.com/pav5000/api-example/internal/wrapper"
)

func main() {
	r := httprouter.New()

	userByIdHandler := userbyid.New()
	createUserHandler := createuser.New()

	// we pass a handle function here, pay the attention that we don't call it here like Handle()
	// wrapper allows us to convert a handle func with structures in and structures out to a httprouter handle func
	r.POST("/user_by_id", wrapper.Wrap(userByIdHandler.Handle))
	r.POST("/create_user", wrapper.Wrap(createUserHandler.Handle))

	log.Println("listening for requests...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
