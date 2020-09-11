package models

type BattleRankUp struct {
	CensusEventBase
	CharacterID string `json:"character_id"`
	BattleRank  int    `json:"battle_rank,string"`
}
