package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	myAddr    string
	optUsers  []grpc.DialOption
	addrUsers string
)

func main() {
	log.SetOutput(os.Stdout)

	flag.StringVar(&myAddr, "addr", os.Getenv("GSERVER_ADDR"), "ip:port to bind")
	flag.Parse()

	if myAddr == "" {
		port := os.Getenv("SERVICES_DEFAULT_PORT")
		if port == "" {
			port = "8070"
		}
		myAddr = ":" + port
	}

	optUsers = append(optUsers, grpc.WithInsecure())
	addrUsers = os.Getenv("SERVICEUSERS_ADDR")
	if addrUsers == "" {
		addrs, err := net.LookupHost("serviceusers")
		log.Println(addrs)
		if err != nil {
			log.Println(err)
		}
		port := os.Getenv("SERVICES_DEFAULT_PORT")
		if port == "" {
			port = "8081"
		}
		if len(addrs) != 0 {
			addrUsers = addrs[0] + ":" + port
		} else {
			addrUsers = ":8080"
		}
	}
	log.Printf("addrUsers: %v\n", addrUsers)

	r := mux.NewRouter().StrictSlash(true)

	r.Handle("/users", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetUsers))).Methods(http.MethodGet)
	r.Handle("/users", handlers.LoggingHandler(os.Stdout, handlers.ContentTypeHandler(http.HandlerFunc(AddUser), "application/json"))).Methods(http.MethodPost)
	r.Handle("/users/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetUser))).Methods(http.MethodGet)
	r.Handle("/users/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(UpdateUser))).Methods(http.MethodPut)
	r.Handle("/users/{id:[0-9]+}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(RemoveUser))).Methods(http.MethodDelete)

	server := &http.Server{
		Addr:         myAddr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("I'm a new version!!!")
	log.Printf("Gserver started %v!!!", server.Addr)
	log.Fatal(server.ListenAndServe())
}
