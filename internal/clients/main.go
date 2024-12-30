package main

import (
	"fmt"
	"idea-store-auth/internal/config"
	"idea-store-auth/internal/middlewares"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const grpcHost = "localhost"
const clientAddr = "localhost:8000"

func main() {

	cfg := config.MustLoad()
	ideasClient, _ := NewIdeasClient(grpcIdeasAddress(cfg), cfg.Clients.Ideas.Timeout, cfg.Clients.Ideas.RetriesCount)
	authClient, _ := NewAuthClient(grpcAuthAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	boardsClient, _ := NewBoardsClient(grpcBoardsAddress(cfg), cfg.Clients.Boards.Timeout, cfg.Clients.Boards.RetriesCount)

	router := mux.NewRouter()
	router.HandleFunc("/idea", ideasClient.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/idea", ideasClient.GetIdea).Methods("GET", "OPTIONS")
	router.HandleFunc("/ideas", ideasClient.GetAllIdeas).Methods("GET", "OPTIONS")
	router.HandleFunc("/board", boardsClient.CreateBoard).Methods("POST", "OPTIONS")
	router.HandleFunc("/board", boardsClient.GetBoard).Methods("GET", "OPTIONS")
	router.HandleFunc("/boards", boardsClient.GetAllBoards).Methods("GET", "OPTIONS")
	router.HandleFunc("/images/{name}", GetImages).Methods("GET", "OPTIONS")
	router.HandleFunc("/register", authClient.Regsiter).Methods("POST","OPTIONS")
	router.HandleFunc("/login", authClient.Login).Methods("POST","OPTIONS")
	handler := middlewares.CorsMiddleware(router)
	fmt.Println("Server is listening...")

	log.Fatal(http.ListenAndServe(clientAddr, handler))
}

func grpcBoardsAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.BoardsMS.Port))
}
func grpcAuthAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.AuthMS.Port))
}
func grpcIdeasAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.IdeasMS.Port))
}