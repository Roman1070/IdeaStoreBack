package main

import (
	"fmt"
	"idea-store-auth/internal/config"
	"idea-store-auth/internal/middlewares"
	"log"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

const grpcHost = "localhost"
const clientAddr = "localhost:8000"

func main() {

	cfg := config.MustLoad()
	ideasClient, _ := NewIdeasClient(grpcIdeasAddress(cfg), cfg.Clients.Ideas.Timeout, cfg.Clients.Ideas.RetriesCount)
	authClient, _ := NewAuthClient(grpcAuthAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	boardsClient, _ := NewBoardsClient(grpcBoardsAddress(cfg), cfg.Clients.Boards.Timeout, cfg.Clients.Boards.RetriesCount)
	profilesClient, _ := NewProfilesClient(grpcProfilesAddress(cfg), cfg.Clients.Profiles.Timeout, cfg.Clients.Profiles.RetriesCount)

	router := mux.NewRouter()
	router.HandleFunc("/images/{name}", GetImages).Methods("GET", "OPTIONS")

	router.HandleFunc("/idea", ideasClient.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/idea", ideasClient.GetIdea).Methods("GET", "OPTIONS")
	router.HandleFunc("/ideas", ideasClient.GetAllIdeas).Methods("GET", "OPTIONS")

	router.HandleFunc("/board", boardsClient.CreateBoard).Methods("POST", "OPTIONS")
	router.HandleFunc("/board", boardsClient.GetBoard).Methods("GET", "OPTIONS")
	router.HandleFunc("/boards", boardsClient.GetAllBoards).Methods("GET", "OPTIONS")

	router.HandleFunc("/register", authClient.Regsiter).Methods("POST","OPTIONS")
	router.HandleFunc("/login", authClient.Login).Methods("POST","OPTIONS")
	
	router.HandleFunc("/profile", profilesClient.CreateProfile).Methods("POST","OPTIONS")
	router.HandleFunc("/profile", profilesClient.GetProfile).Methods("GET","OPTIONS")
	router.HandleFunc("/toggle-save-idea", profilesClient.ToggleSaveIdea).Methods("GET","OPTIONS")
	router.HandleFunc("/is-idea-saved", profilesClient.IsIdeaSaved).Methods("GET","OPTIONS")
	
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
func grpcProfilesAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.ProfilesMS.Port))
}
func GetUserIdByRequestWithCookie(r *http.Request) (int64,error){	
	for _,cookie := range r.Cookies(){
		fmt.Println(cookie.String())
	}
	tokenCookie,err:= r.Cookie("token")
	if err!=nil{
		slog.Error(err.Error())
		return -1, err
	}

	claims := jwt.MapClaims{}
	tokenStr :=tokenCookie.String()[6:]
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("yaro-gas"), nil
	})
	if err!=nil{
		slog.Error(err.Error())
		return -1, err
	}
	userId :=claims["uid"].(float64)
	userIdStr :=fmt.Sprint(userId)
	userIdInt, _ := strconv.ParseInt(userIdStr,10,64)
	return userIdInt,nil
}