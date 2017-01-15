package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mcarmonaa/lagrs/serviceusers/user"
)

var (
	//mu     sync.Mutex
	//lastID int32 = 1

	db *gorm.DB
)

type service struct {
	users []*user.User
	list  sync.Mutex
}

func (s *service) List(ctx context.Context, r *user.ListRequest) (*user.ListReply, error) {
	log.Printf("RPC List: %v", r)

	var users []user.User
	pusers := make([]*user.User, 0)
	var count *int32
	db.Table("users").Count(&count)
	//db.Find(&users).Limit(r.GetUsersPerPage()).Offset(*count / r.GetUsersPerPage())
	db.Order("id asc").Find(&users)
	for i := range users {
		pusers = append(pusers, &users[i])
	}
	reply := &user.ListReply{
		Status:         user.Status_OK,
		UsersRequested: r.GetUsersPerPage(),
		PageRequested:  r.GetPage(),
		UsersReplied:   strconv.FormatInt(int64(len(pusers)), 10),
		PageReplied:    "1",
		Pages:          strconv.FormatInt(int64(*count/r.GetUsersPerPage()), 10),
		Users:          pusers,
	}
	return reply, nil
}

func (s *service) Add(ctx context.Context, r *user.AddRequest) (*user.AddReply, error) {
	log.Printf("RPC Add: %v", r)

	u := &user.User{
		Name: r.Name,
		Age:  r.Age,
		Mail: r.Mail,
		Role: r.Role,
	}
	db.Create(u)
	reply := &user.AddReply{
		Status: user.Status_OK,
		User:   u,
	}
	return reply, nil
}

func (s *service) Get(ctx context.Context, r *user.GetRequest) (*user.GetReply, error) {
	log.Printf("RPC Get: %v", r)

	u := &user.User{}
	db.First(u, r.GetId())
	reply := &user.GetReply{}
	if u.GetId() == r.GetId() {
		reply.Status = user.Status_OK
		reply.User = u
	} else {
		reply.Status = user.Status_ERR
	}
	return reply, nil
}

func (s *service) Update(ctx context.Context, r *user.UpdateRequest) (*user.UpdateReply, error) {
	log.Printf("RPC Update: %v", r)

	u := &user.User{}
	db.Find(u, r.GetUser().GetId())
	reply := &user.UpdateReply{}
	if u.GetId() == 0 {
		reply.Status = user.Status_ERR
		return reply, nil
	}
	db.Save(r.GetUser())
	reply.Status = user.Status_OK
	u.Reset()
	db.Find(u, r.GetUser().GetId())
	reply.User = u
	return reply, nil
}

func (s *service) Remove(ctx context.Context, r *user.RemoveRequest) (*user.RemoveReply, error) {
	log.Printf("RPC Remove: %v", r)

	uDel := &user.User{}
	db.Find(uDel, r.GetId())
	reply := &user.RemoveReply{}
	if uDel.GetId() == 0 {
		reply.Status = user.Status_ERR
		return reply, nil
	}
	db.Delete(uDel)
	u := &user.User{}
	db.Find(u, r.GetId())
	if u.GetId() != 0 {
		reply.Status = user.Status_ERR
		return reply, nil
	}
	reply.Status = user.Status_OK
	reply.User = uDel
	return reply, nil
}

func main() {
	log.SetOutput(os.Stdout)

	var err error
	// export DB_URL="postgres://manuel:pass@localhost:5432/manuel?sslmode=disable"
	// dbengine://user:pass@host:port/dbname?sslmode=disable ---> https://godoc.org/github.com/lib/pq
	log.Println(os.Getenv("DB_URL"))
	log.Println(os.Getenv("TEST"))
	db, err = gorm.Open("postgres", os.Getenv("DB_URL"))
	//db, err = gorm.Open("postgres", "user=manuel dbname=manuel password=pass sslmode=disable ")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&user.User{})

	service := &service{users: make([]*user.User, 0)}
	port := os.Getenv("SERVICES_DEFAULT_PORT")
	if port == "" {
		port = "8081"
	}
	port = ":" + port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUsersServer(grpcServer, service)
	log.Printf("ServiceUsers started %v!!!", lis.Addr())
	log.Fatal(grpcServer.Serve(lis))
}
