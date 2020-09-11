package models

type AchievementEarned struct {
	CensusEventBase
	CharacterID   string `json:"character_id"`
	AchievementID string `json:"achievement_id"`
}
