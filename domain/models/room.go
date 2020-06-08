package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"gopkg.in/guregu/null.v4"
)

// Room : roomテーブルモデル
type Room struct {
	RoomID      string
	UserID      string
	RoomText    string
	GameTitleID string
	GameHardID  string
	Capacity    int
	Start       null.Time
	IsLock      bool
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// NewRoom :
func NewRoom(uid, text, gtid, ghid string, cap int, s time.Time) (*Room, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	r := &Room{
		RoomID:      id.String(),
		UserID:      uid,
		RoomText:    text,
		GameTitleID: gtid,
		GameHardID:  ghid,
		Capacity:    cap,
		IsLock:      false,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	if s.IsZero() == true {
		r.Start = null.NewTime(s, false)
	} else {
		r.Start = null.NewTime(s, true)
	}
	return r, nil
}
