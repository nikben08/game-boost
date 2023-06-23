package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	UserId          string    `gorm:"user_id"`
	OrderType       string    `gorm:"order_type"`
	CurrentRank     string    `gorm:"current_rank"`
	CurrentLp       string    `gorm:"current_lp"`
	Server          string    `gorm:"server"`
	DesiredRank     string    `gorm:"desired_rank"`
	QType           string    `gorm:"q_type"`
	ChampionsRoles  bool      `gorm:"champions_roles"`
	PriorityBoost   bool      `gorm:"priority_boost"`
	StreamGames     bool      `gorm:"stream_games"`
	SoloOnly        bool      `gorm:"solo_only"`
	BonusWin        bool      `gorm:"bonus_win"`
	TotalPrice      float64   `gorm:"total_price"`
	WinsAmount      int64     `gorm:"wins_amount"`
	PlayWithBooster bool      `gorm:"play_with_booster"`
	PreviousRank    string    `gorm:"previous_rank"`
	GameMode        string    `gorm:"game_mode"`
	GamesAmount     int64     `gorm:"games_amount"`
}
