package model

import (
	"github.com/KnutZuidema/golio/datadragon"
	"github.com/go-bongo/bongo"
)

type Champion struct {
	bongo.DocumentBase `bson:",inline"`
	Champ              datadragon.ChampionData
}

// type ChampionStatistics struct {
// 	Hp                   float64 `bson:"hp"`
// 	HpPerLevel           float64 `bson:"hpperlevel"`
// 	Mp                   float64 `bson:"mp"`
// 	MpPerLevel           float64 `bson:"mpperlevel"`
// 	MoveSpeed            float64 `bson:"movespeed"`
// 	Armor                float64 `bson:"armor"`
// 	ArmorPerLevel        float64 `bson:"armorperlevel"`
// 	SpellBlock           float64 `bson:"spellblock"`
// 	SpellBlockPerLevel   float64 `bson:"spellblockperlevel"`
// 	AttackRange          int     `bson:"attackrange"`
// 	HpRegen              float64 `bson:"hpregen"`
// 	HpRegenPerLevel      float64 `bson:"hpregenperlevel"`
// 	MpRegen              float64 `bson:"mpregen"`
// 	MpRegenPerLevel      float64 `bson:"mpregenperlevel"`
// 	Crit                 int     `bson:"crit"`
// 	CritPerLevel         int     `bson:"critperlevel"`
// 	AttackDamange        float64 `bson:"attackdamage"`
// 	AttackDamagePerLevel float64 `bson:"attackdamageperlevel"`
// 	AttackSpeedPerLevel  float64 `bson:"attackspeedperlevel"`
// 	AttackSpeed          float64 `bson:"attackspeed"`
// }

// func MakeInternalChampion(champion datadragon.ChampionData) Champion {
// 	inCh := Champion{
// 		Name:    champion.Name,
// 		Version: champion.Version,
// 		Key:     champion.Key,
// 		Title:   champion.Title,
// 		Stats: ChampionStatistics{
// 			Hp:         champion.Stats.HealthPoints,
// 			HpPerLevel: champion.Stats.HealthPointRegenerationPerLevel,
// 			Mp: champion.Stats.ManaPoints,
// 			MpPerLevel: champion.Stats.ManaPointsPerLevel,
// 			MoveSpeed: champion.Stats.MovementSpeed,
// 			Armor: champion.Stats.Armor,
// 			ArmorPerLevel: champion.Stats.ArmorPerLevel,
// 			SpellBlock: champion.Stats.SpellBlock,
// 			SpellBlockPerLevel: champion.,
// 		},
// 	}
// }
