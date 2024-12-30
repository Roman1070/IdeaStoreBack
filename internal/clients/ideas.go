package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	ideasv1 "idea-store-auth/gen/go/idea"
	"idea-store-auth/internal/utils"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdeasClient struct {
	api ideasv1.IdeasClient
}
type getIdeaRequest struct{
	Id int64 `json:"id"`
}
func (c *IdeasClient) GetIdea(w http.ResponseWriter, r *http.Request){
	slog.Info("Client started to get idea")

	var req getIdeaRequest
	var err error
	req.Id, err = strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w,err.Error())
		return
	}

	resp,err := c.api.GetIdea(r.Context(),&ideasv1.GetRequest{IdeaId: req.Id})
	if err!=nil{
		slog.Error(err.Error())
		utils.WriteError(w,err.Error())
		return
	}
	m := protojson.MarshalOptions{EmitDefaultValues: true}
   
	result, err :=  m.Marshal(resp)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}
	slog.Info(string(result))
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func GetImages(w http.ResponseWriter, r *http.Request){
	
	file,err:= os.ReadFile("F:/Roman/WEB/IdeaStoreBack"+r.RequestURI)
	if err!=nil{
		slog.Error(err.Error())
		utils.WriteError(w,err.Error())
		return
	}
	
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/*")
	w.Write(file)
}

func (c *IdeasClient) Create(w http.ResponseWriter, r *http.Request) {
	slog.Info("Client started to create")
	tokenCookie,err:= r.Cookie("token")
	if err!=nil{
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}
	
	claims := jwt.MapClaims{}
	tokenStr :=tokenCookie.String()[6:]
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("yaro-gas"), nil
	})
	if err!=nil{
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}
	userId :=claims["uid"].(float64)
	userIdStr :=fmt.Sprint(userId)
	userIdInt, _ := strconv.ParseInt(userIdStr,10,64)
	
	r.ParseMultipartForm(12 << 20)
	defer r.Body.Close()
	file, h, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	ext:= filepath.Ext(h.Filename)
	hash :=md5.Sum([]byte(h.Filename))
	path:= "./Images/"+hex.EncodeToString(hash[:])+ext
	tmpfile, err := os.Create(path)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tmpfile.Close()
	
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	request := &ideasv1.CreateRequest{
		Image:       hex.EncodeToString(hash[:])+ext,
		Name:        r.Form.Get("name"),
		Description: r.Form.Get("description"),
		Link:        r.Form.Get("link"),
		Tags:        r.Form.Get("tags"),
		UserId: userIdInt,
	}
	
	createResponse, err := c.api.CreateIdea(r.Context(), request)
	if err != nil {
		slog.Error(err.Error())
		utils.WriteError(w, err.Error())
		return
	}

	result, err := json.Marshal(createResponse)
	if err != nil {
		utils.WriteError(w, err.Error())
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func (c *IdeasClient) GetAllIdeas(w http.ResponseWriter, r *http.Request){
	slog.Info("Client started to get ideas")
	resp,err:= c.api.GetAllIdeas(r.Context(),&emptypb.Empty{})
	if err!=nil{
		utils.WriteError(w,err.Error())
		return
	}
	result, err := json.Marshal(resp.Ideas)
	if err!=nil{
		utils.WriteError(w,err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func NewIdeasClient(addr string, timeout time.Duration, retriesCount int) (*IdeasClient, error) {
	const op = "client.ideas.New"

	retryOptions := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(
		grpcretry.UnaryClientInterceptor(retryOptions...),
	))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &IdeasClient{
		api: ideasv1.NewIdeasClient(cc),
	}, nil
}