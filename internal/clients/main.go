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

const clientAddr = "176.114.67.252:8000"

func main() {

	cfg := config.MustLoad()
	ideasClient, _ := NewIdeasClient(common.GrpcIdeasAddress(cfg), cfg.Clients.Ideas.Timeout, cfg.Clients.Ideas.RetriesCount)
	authClient, _ := NewAuthClient(common.GrpcAuthAddress(cfg), cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	boardsClient, _ := NewBoardsClient(common.GrpcBoardsAddress(cfg), cfg.Clients.Boards.Timeout, cfg.Clients.Boards.RetriesCount)
	profilesClient, _ := NewProfilesClient(common.GrpcProfilesAddress(cfg), cfg.Clients.Profiles.Timeout, cfg.Clients.Profiles.RetriesCount)
	commentsClient, _ := NewCommentsClient(common.GrpcCommentsAddress(cfg), cfg.Clients.Comments.Timeout, cfg.Clients.Comments.RetriesCount)
	chatsClient, _ := NewChatsClient(common.GrpcChatsAddress(cfg), cfg.Clients.Chats.Timeout, cfg.Clients.Chats.RetriesCount)

	router := mux.NewRouter()
	router.HandleFunc("/app/files/{name}", GetImages).Methods(http.MethodGet)

	router.HandleFunc("/api/v1/idea", ideasClient.Create).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/idea", ideasClient.GetIdea).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/ideas", ideasClient.GetAllIdeas).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/ideas", ideasClient.GetIdeas).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/v1/board", boardsClient.CreateBoard).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/board", boardsClient.GetBoard).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/board", boardsClient.DeleteBoard).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/api/v1/my-boards", boardsClient.GetCurrentUsersBoards).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/boards", boardsClient.GetBoards).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/ideas-in-board", boardsClient.GetIdeasInBoard).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api/v1/register", authClient.Regsiter).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/login", authClient.Login).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/v1/profile", profilesClient.CreateProfile).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/profile", profilesClient.UpdateProfile).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/api/v1/profile", profilesClient.GetProfile).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/my-profile", profilesClient.GetCurrentProfile).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/toggle-save-idea", profilesClient.ToggleSaveIdea).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/is-idea-saved", profilesClient.IsIdeaSaved).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/get-saved-ideas", profilesClient.GetSavedIdeas).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/search-profiles", profilesClient.GetProfilesFromSearch).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/toggle-like-idea", profilesClient.ToggleLikeIdea).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/is-idea-liked", profilesClient.IsIdeaLiked).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api/v1/comments", commentsClient.GetComments).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/comment", commentsClient.CreateComment).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/v1/chats", chatsClient.GetChats).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/chat_ws", chatsClient.HandleChatWebSocket).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/messages", chatsClient.GetMessages).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/v1/message", chatsClient.SendMessage).Methods(http.MethodPost, http.MethodOptions)
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
		return []byte("yaro21u527"), nil
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
