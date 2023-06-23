package contracts

import "github.com/google/uuid"

type CreateOrderRequest struct {
	OrderType        string  `json:"order_type" validate:"required"`
	CurrentRank      string  `json:"current_rank" validate:"required"`
	CurrentLp        string  `json:"current_lp"`
	Server           string  `json:"server"`
	DesiredRank      string  `json:"desired_rank"`
	QType            string  `json:"q_type"`
	ChampionsRoles   bool    `json:"champions_roles"`
	PriorityBoost    bool    `json:"priority_boost"`
	StreamGames      bool    `json:"stream_games"`
	SoloOnly         bool    `json:"solo_only"`
	BonusWin         bool    `json:"bonus_win"`
	TotalPrice       float64 `json:"total_price"`
	WinsAmount       int64   `json:"wins_amount"`
	PlaysWithBooster bool    `json:"plays_with_booster"`
	PreviousRank     string  `json:"previous_rank"`
	GameMode         string  `json:"game_mode"`
	GamesAmount      int64   `json:"games_amount"`
}

type EditOrderRequest struct {
	ID               uuid.UUID `json:"id" validate:"required"`
	OrderType        string    `json:"order_type" validate:"required"`
	CurrentRank      string    `json:"current_rank" validate:"required"`
	CurrentLp        string    `json:"current_lp"`
	Server           string    `json:"server"`
	DesiredRank      string    `json:"desired_rank"`
	QType            string    `json:"q_type"`
	ChampionsRoles   bool      `json:"champions_roles"`
	PriorityBoost    bool      `json:"priority_boost"`
	StreamGames      bool      `json:"stream_games"`
	SoloOnly         bool      `json:"solo_only"`
	BonusWin         bool      `json:"bonus_win"`
	TotalPrice       float64   `json:"total_price"`
	WinsAmount       int64     `json:"wins_amount"`
	PlaysWithBooster bool      `json:"plays_with_booster"`
	PreviousRank     string    `json:"previous_rank"`
	GameMode         string    `json:"game_mode"`
	GamesAmount      int64     `json:"games_amount"`
}
