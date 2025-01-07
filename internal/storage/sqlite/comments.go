package sqlite

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Storage) CreateComment(ctx context.Context, ideaId, userId int64, text, creationDate string) (*emptypb.Empty, error){
	slog.Info("storage started CreateComment")

	stmt,err:= s.db.Prepare("INSERT INTO comments (idea_id,user_id,content,creation_date) VALUES(?,?,?,?)")
	if err!=nil{
		return nil, fmt.Errorf("storage CreateComment error: %v",err.Error())
	}
	_,err = stmt.ExecContext(ctx,ideaId,userId,text,creationDate)
	if err!=nil{
		return nil, fmt.Errorf("storage CreateComment error: %v",err.Error())
	}
	return nil,nil
}