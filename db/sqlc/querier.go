// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAction(ctx context.Context, arg CreateActionParams) (Actions, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	CreateSubAction(ctx context.Context, arg CreateSubActionParams) (Subactions, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteAction(ctx context.Context, id int64) error
	DeleteAllSubAction(ctx context.Context, actionsID int64) error
	DeleteSubAction(ctx context.Context, id int64) error
	GetSession(ctx context.Context, id uuid.UUID) (Sessions, error)
	GetUserID(ctx context.Context, username string) (Users, error)
	ListAction(ctx context.Context, username string) ([]Actions, error)
	ListSubAction(ctx context.Context, actionsID int64) ([]Subactions, error)
	UpdateAction(ctx context.Context, arg UpdateActionParams) (Actions, error)
	UpdateSubAction(ctx context.Context, arg UpdateSubActionParams) (Subactions, error)
}

var _ Querier = (*Queries)(nil)
