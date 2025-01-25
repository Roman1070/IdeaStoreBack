package main

import (
	"fmt"
	common "idea-store-auth/cmd"
	"idea-store-auth/internal/config"
	"idea-store-auth/internal/middlewares"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

const clientAddr = "localhost:8000"

func main() {

	cfg := config.MustLoad()
	ideasClient, _ := NewIdeasClient(common.GrpcIdeasAddress(cfg), cfg.Clients.Ideas.Timeout, cfg.Clients.Ideas.RetriesCount)
	authClient, _ := NewAuthClient(common.GrpcAuthAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	boardsClient, _ := NewBoardsClient(common.GrpcBoardsAddress(cfg), cfg.Clients.Boards.Timeout, cfg.Clients.Boards.RetriesCount)
	profilesClient, _ := NewProfilesClient(common.GrpcProfilesAddress(cfg), cfg.Clients.Profiles.Timeout, cfg.Clients.Profiles.RetriesCount)
	commentsClient, _ := NewCommentsClient(common.GrpcCommentsAddress(cfg), cfg.Clients.Comments.Timeout, cfg.Clients.Comments.RetriesCount)
	chatsClient, _ := NewChatsClient(common.GrpcChatsAddress(cfg), cfg.Clients.Chats.Timeout, cfg.Clients.Chats.RetriesCount)

	router := mux.NewRouter()
	router.HandleFunc("/images/{name}", GetImages).Methods("GET", "OPTIONS")

	router.HandleFunc("/idea", ideasClient.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/idea", ideasClient.GetIdea).Methods("GET", "OPTIONS")
	router.HandleFunc("/ideas", ideasClient.GetAllIdeas).Methods("GET", "OPTIONS")
	router.HandleFunc("/ideas", ideasClient.GetIdeas).Methods("POST", "OPTIONS")

	router.HandleFunc("/board", boardsClient.CreateBoard).Methods("POST", "OPTIONS")
	router.HandleFunc("/board", boardsClient.GetBoard).Methods("GET", "OPTIONS")
	router.HandleFunc("/board", boardsClient.DeleteBoard).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/my-boards", boardsClient.GetCurrentUsersBoards).Methods("GET", "OPTIONS")
	router.HandleFunc("/boards", boardsClient.GetBoards).Methods("GET", "OPTIONS")
	router.HandleFunc("/ideas-in-board", boardsClient.GetIdeasInBoard).Methods("GET", "OPTIONS")

	router.HandleFunc("/register", authClient.Regsiter).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", authClient.Login).Methods("POST", "OPTIONS")

	router.HandleFunc("/profile", profilesClient.CreateProfile).Methods("POST", "OPTIONS")
	router.HandleFunc("/profile", profilesClient.UpdateProfile).Methods("PUT", "OPTIONS")
	router.HandleFunc("/profile", profilesClient.GetProfile).Methods("GET", "OPTIONS")
	router.HandleFunc("/my-profile", profilesClient.GetCurrentProfile).Methods("GET", "OPTIONS")
	router.HandleFunc("/toggle-save-idea", profilesClient.ToggleSaveIdea).Methods("GET", "OPTIONS")
	router.HandleFunc("/is-idea-saved", profilesClient.IsIdeaSaved).Methods("GET", "OPTIONS")
	router.HandleFunc("/get-saved-ideas", profilesClient.GetSavedIdeas).Methods("GET", "OPTIONS")

	router.HandleFunc("/comments", commentsClient.GetComments).Methods("GET", "OPTIONS")
	router.HandleFunc("/comment", commentsClient.CreateComment).Methods("POST", "OPTIONS")

	router.HandleFunc("/chats", chatsClient.GetChats).Methods("GET", "OPTIONS")
	handler := middlewares.CorsMiddleware(router)
	fmt.Println("Server is listening...")

	log.Fatal(http.ListenAndServe(clientAddr, handler))
}

func GetUserIdByRequestWithCookie(r *http.Request) (int64, error) {

	tokenCookie, err := r.Cookie("token")
	if err != nil {
		slog.Error(err.Error())
		return -1, err
	}

	claims := jwt.MapClaims{}
	tokenStr := tokenCookie.String()[6:]
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("yaro-gas"), nil
	})
	if err != nil {
		slog.Error(err.Error())
		return -1, err
	}
	userId := claims["uid"].(float64)
	userIdStr := fmt.Sprint(userId)
	userIdInt, _ := strconv.ParseInt(userIdStr, 10, 64)
	return userIdInt, nil
}
