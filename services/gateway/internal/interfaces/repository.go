package interfaces

import (
	"context"
	"database/sql"

	"github.com/what-da-flac/wtf/openapi/models"
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
	InsertRole(ctx context.Context, role *models.Role) error
	UpdateRole(ctx context.Context, role *models.Role) error
	SelectRole(ctx context.Context, id string) (*models.Role, error)
	DeleteRole(ctx context.Context, id string) error
	ListRoles(ctx context.Context) ([]*models.Role, error)
	AddUser(ctx context.Context, role *models.Role, user *models.User) error
	RemoveUser(ctx context.Context, role *models.Role, user *models.User) error
	ListUsersInRole(ctx context.Context, roleId string) ([]*models.User, error)
	ListRolesForUser(ctx context.Context, user *models.User) ([]*models.Role, error)
}

type TorrentRepository interface {
	InsertTorrent(ctx context.Context, torrent *models.Torrent) error
	UpdateTorrent(ctx context.Context, torrent *models.Torrent) error
	SelectTorrent(ctx context.Context, id string) (*models.Torrent, error)
	ListTorrents(ctx context.Context, params models.GetV1TorrentsParams) ([]*models.Torrent, error)
	ListTorrentStatuses(ctx context.Context) []string
	InsertTorrentFile(ctx context.Context, file *models.TorrentFile) error
	DeleteTorrentFiles(ctx context.Context, torrentId string) error
	SelectTorrentFiles(ctx context.Context, id string) ([]*models.TorrentFile, error)
	DeleteTorrent(ctx context.Context, id string) error
}

type UserRepository interface {
	InsertUser(cxt context.Context, user *models.User) error
	UpdateUser(cxt context.Context, user *models.User) error
	SelectUser(cxt context.Context, id, email *string) (*models.User, error)
	ListUsers(cxt context.Context, request *models.UserListParams) ([]*models.User, error)
}
