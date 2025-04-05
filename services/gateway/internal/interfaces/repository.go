package interfaces

import (
	"context"
	"database/sql"

	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

//go:generate moq -out ../../mocks/repository.go -pkg mocks . Repository
type Repository interface {
	RoleRepository
	TorrentRepository
	UserRepository
}

type DB interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type RoleRepository interface {
	InsertRole(ctx context.Context, role *golang.Role) error
	UpdateRole(ctx context.Context, role *golang.Role) error
	SelectRole(ctx context.Context, id string) (*golang.Role, error)
	DeleteRole(ctx context.Context, id string) error
	ListRoles(ctx context.Context) ([]*golang.Role, error)
	AddUser(ctx context.Context, role *golang.Role, user *golang.User) error
	RemoveUser(ctx context.Context, role *golang.Role, user *golang.User) error
	ListUsersInRole(ctx context.Context, roleId string) ([]*golang.User, error)
	ListRolesForUser(ctx context.Context, user *golang.User) ([]*golang.Role, error)
}

type TorrentRepository interface {
	InsertTorrent(ctx context.Context, torrent *golang.Torrent) error
	UpdateTorrent(ctx context.Context, torrent *golang.Torrent) error
	SelectTorrent(ctx context.Context, id string) (*golang.Torrent, error)
	ListTorrents(ctx context.Context, params golang.GetV1TorrentsParams) ([]*golang.Torrent, error)
	ListTorrentStatuses(ctx context.Context) []string
	InsertTorrentFile(ctx context.Context, file *golang.TorrentFile) error
	DeleteTorrentFiles(ctx context.Context, torrentId string) error
	SelectTorrentFiles(ctx context.Context, id string) ([]*golang.TorrentFile, error)
	DeleteTorrent(ctx context.Context, id string) error
}

type UserRepository interface {
	InsertUser(cxt context.Context, user *golang.User) error
	UpdateUser(cxt context.Context, user *golang.User) error
	SelectUser(cxt context.Context, id, email *string) (*golang.User, error)
	ListUsers(cxt context.Context, request *golang.UserListParams) ([]*golang.User, error)
}
