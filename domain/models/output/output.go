package output

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// FavoriteGamesRes : お気に入りゲームレスポンス
type FavoriteGamesRes struct {
	GameTitle string    `json:"game_title" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

// UserResBody : ユーザーレスポンス
type UserResBody struct {
	UserID        string             `json:"id" binding:"required"`
	UserName      string             `json:"user_name" binding:"required"`
	Icon          string             `json:"icon" binding:"required"`
	CreatedAt     time.Time          `json:"created_at" binding:"required"`
	FavoriteGames []FavoriteGamesRes `json:"favorite_games" binding:"required"`
}

// RoomResBody : ルームレスポンス
type RoomResBody struct {
	RoomID    string    `json:"room_id" binding:"required"`
	UserID    string    `json:"owner_id" binding:"required"`
	Icon      string    `json:"icon" binding:"required"`
	HardName  string    `json:"hard" binding:"required"`
	GameTitle string    `json:"title" binding:"required"`
	Capacity  int       `json:"capacity" binding:"required"`
	RoomText  string    `json:"text" binding:"required"`
	UserName  string    `json:"name" binding:"required"`
	Count     int       `json:"count" binding:"required"`
	IsLock    bool      `json:"is_lock" binding:"required"`
	CreatedAt time.Time `json:"created" binding:"required"`
	Start     null.Time `json:"start" binding:"required"`
}

// JoinUserRes : お気に入りゲームレスポンス
type JoinUserRes struct {
	UserID   string `json:"user_id" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Icon     string `json:"icon" binding:"required"`
}

// RoomDetailResBody : ルームレスポンス
type RoomDetailResBody struct {
	RoomID    string        `json:"room_id" binding:"required"`
	OwnerID   string        `json:"owner_id" binding:"required"`
	HardName  string        `json:"hard" binding:"required"`
	GameTitle string        `json:"title" binding:"required"`
	Capacity  int           `json:"capacity" binding:"required"`
	Count     int           `json:"count" binding:"required"`
	RoomText  string        `json:"text" binding:"required"`
	JoinUsers []JoinUserRes `json:"join_users" binding:"required"`
}
