package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	"github.com/gorilla/mux"
	"github.com/mcarmonaa/lagrs/serviceusers/user"
)

// Users get the list of users from serviceusers
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// create request for rpc from the url's query string
	var numUsers, page int32 = 10, 1
	snum := r.FormValue("users_per_page")
	spage := r.FormValue("page")
	if snum != "" {
		n, err := strconv.Atoi(snum)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		numUsers = int32(n)
	}
	if spage != "" {
		p, err := strconv.Atoi(spage)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		page = int32(p)
	}
	req := &user.ListRequest{
		UsersPerPage: numUsers,
		Page:         page,
	}

	// get connection to serviceusers
	conn, err := grpc.Dial(addrUsers, optUsers...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%v : %v", addrUsers, err)
		return
	}
	defer conn.Close()

	// create a client against serviceusers
	cusers := user.NewUsersClient(conn)

	// rpc to get the users' list
	reply, err := cusers.List(context.Background(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if reply.GetStatus() != user.Status_OK {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// marshal the response to json content
	response, err := json.Marshal(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// send back a http response to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// AddUser create a new user calling serviceusers
func AddUser(w http.ResponseWriter, r *http.Request) {
	// read the body
	req := &user.AddRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get connection to serviceusers
	conn, err := grpc.Dial(addrUsers, optUsers...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer conn.Close()

	// create a client against serviceusers
	cusers := user.NewUsersClient(conn)

	// rpc for add a user
	reply, err := cusers.Add(context.Background(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if reply.GetStatus() != user.Status_OK {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// marshal the response to json content
	response, err := json.Marshal(reply.GetUser())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// send back a http response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetUser requests for a user to serviceusers
func GetUser(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["id"])
	// err never must be != nil
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := int32(num)

	// get connection to serviceusers
	conn, err := grpc.Dial(addrUsers, optUsers...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer conn.Close()

	// create a client against serviceusers
	cusers := user.NewUsersClient(conn)

	// create the request for the rpc
	req := &user.GetRequest{
		Id: id,
	}

	// rpc for add a user
	reply, err := cusers.Get(context.Background(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if reply.GetStatus() != user.Status_OK {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// marshal the response to json content
	response, err := json.Marshal(reply.GetUser())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// send back a http response to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateUser modifies the existing user's information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["id"])
	// err never must be != nil
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := int32(num)

	// read the body
	u := &user.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Poner el id de la url para que no se haga un update de otro user diferente
	u.Id = id
	req := &user.UpdateRequest{
		User: u,
	}

	// get connection to serviceusers
	conn, err := grpc.Dial(addrUsers, optUsers...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer conn.Close()

	// create a client against serviceusers
	cusers := user.NewUsersClient(conn)

	// rpc for update a user
	reply, err := cusers.Update(context.Background(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if reply.GetStatus() != user.Status_OK {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// marshal the response to json content
	response, err := json.Marshal(reply.GetUser())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// send back a http response to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// RemoveUser requests for remove a user to serviceusers
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	// get id
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["id"])
	// err never must be != nil
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := int32(num)

	// get connection to serviceusers
	conn, err := grpc.Dial(addrUsers, optUsers...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer conn.Close()

	// create a client against serviceusers
	cusers := user.NewUsersClient(conn)

	// create the request for the rpc
	req := &user.RemoveRequest{
		Id: id,
	}

	// rpc to remove a user
	reply, err := cusers.Remove(context.Background(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if reply.GetStatus() != user.Status_OK {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// marshal the response to json content
	response, err := json.Marshal(reply.GetUser())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// send back a http response to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
