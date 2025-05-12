package jobs

import (
	"encoding/json"

	"github.com/dipo0x/golang-url-shortener/internal/types"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const (
	TypeDeleteDocument = "mongo:delete_document"
)

func NewDeleteTask(documentID string) (*asynq.Task, error) {
	parsedID, _ := uuid.Parse(documentID)
	payload, err := json.Marshal(types.IDeletePayload{DocumentID: parsedID.String()})
	if err != nil {
		println(err)
		return nil, err
	}
	return asynq.NewTask(TypeDeleteDocument, payload), nil
}
