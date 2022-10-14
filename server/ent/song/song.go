// Code generated by ent, DO NOT EDIT.

package song

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the song type in the database.
	Label = "song"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldArtistID holds the string denoting the artist_id field in the database.
	FieldArtistID = "artist_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldPlays holds the string denoting the plays field in the database.
	FieldPlays = "plays"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeLikedBy holds the string denoting the liked_by edge name in mutations.
	EdgeLikedBy = "liked_by"
	// EdgePlaylists holds the string denoting the playlists edge name in mutations.
	EdgePlaylists = "playlists"
	// Table holds the table name of the song in the database.
	Table = "songs"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "songs"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "artist_id"
	// LikedByTable is the table that holds the liked_by relation/edge.
	LikedByTable = "likes"
	// LikedByInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikedByInverseTable = "likes"
	// LikedByColumn is the table column denoting the liked_by relation/edge.
	LikedByColumn = "song_id"
	// PlaylistsTable is the table that holds the playlists relation/edge.
	PlaylistsTable = "playlist_songs"
	// PlaylistsInverseTable is the table name for the PlaylistSong entity.
	// It exists in this package in order to avoid circular dependency with the "playlistsong" package.
	PlaylistsInverseTable = "playlist_songs"
	// PlaylistsColumn is the table column denoting the playlists relation/edge.
	PlaylistsColumn = "song_id"
)

// Columns holds all SQL columns for song fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldArtistID,
	FieldTitle,
	FieldURL,
	FieldPlays,
	FieldDuration,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultPlays holds the default value on creation for the "plays" field.
	DefaultPlays uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
