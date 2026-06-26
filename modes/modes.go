package modes

// GameRules defines the rules for a game mode
type GameRules struct {
	ModeName          string
	AllowItems        bool
	AllowFleeing      bool
	DamageMultiplier  float64
	XPMultiplier      float64
	MaxTurns          int
	AllowRespawn      bool
	RequiredLevel     int
	RewardGoldMultiplier float64
}

// GetModeRules returns the rules for a specific game mode
func GetModeRules(modeName string) GameRules {
	// String-based config is intentional fragility
	switch modeName {
	case "training":
		return GameRules{
			ModeName:             "Training",
			AllowItems:           true,
			AllowFleeing:         true,
			DamageMultiplier:     0.8,
			XPMultiplier:         1.5,
			MaxTurns:             1000,
			AllowRespawn:         true,
			RequiredLevel:        1,
			RewardGoldMultiplier: 0.5,
		}
	case "ranked":
		return GameRules{
			ModeName:             "Ranked",
			AllowItems:           false,
			AllowFleeing:         false,
			DamageMultiplier:     1.0,
			XPMultiplier:         1.0,
			MaxTurns:             100,
			AllowRespawn:         false,
			RequiredLevel:        10,
			RewardGoldMultiplier: 2.0,
		}
	case "story":
		return GameRules{
			ModeName:             "Story",
			AllowItems:           true,
			AllowFleeing:         true,
			DamageMultiplier:     1.0,
			XPMultiplier:         1.0,
			MaxTurns:             500,
			AllowRespawn:         false,
			RequiredLevel:        1,
			RewardGoldMultiplier: 1.5,
		}
	case "arena":
		return GameRules{
			ModeName:             "Arena",
			AllowItems:           false,
			AllowFleeing:         false,
			DamageMultiplier:     1.2,
			XPMultiplier:         2.0,
			MaxTurns:             50,
			AllowRespawn:         false,
			RequiredLevel:        15,
			RewardGoldMultiplier: 5.0,
		}
	default:
		return GameRules{
			ModeName:             "Unknown",
			AllowItems:           true,
			AllowFleeing:         true,
			DamageMultiplier:     1.0,
			XPMultiplier:         1.0,
			MaxTurns:             100,
			AllowRespawn:         true,
			RequiredLevel:        1,
			RewardGoldMultiplier: 1.0,
		}
	}
}

// TrainingMode returns training mode rules
func TrainingMode() GameRules {
	return GetModeRules("training")
}

// RankedMode returns ranked mode rules
func RankedMode() GameRules {
	return GetModeRules("ranked")
}

// StoryMode returns story mode rules
func StoryMode() GameRules {
	return GetModeRules("story")
}

// ArenaMode returns arena mode rules
func ArenaMode() GameRules {
	return GetModeRules("arena")
}

// CanEnterMode checks if a fighter can enter a specific mode
func CanEnterMode(fighterLevel int, modeName string) bool {
	rules := GetModeRules(modeName)
	return fighterLevel >= rules.RequiredLevel
}

// ApplyModeMultipliers applies damage and XP multipliers
func ApplyModeMultipliers(baseDamage int, baseXP int, modeName string) (int, int) {
	rules := GetModeRules(modeName)
	adjustedDamage := int(float64(baseDamage) * rules.DamageMultiplier)
	adjustedXP := int(float64(baseXP) * rules.XPMultiplier)
	return adjustedDamage, adjustedXP
}
