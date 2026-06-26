package fighters

// ProgressionTracker tracks a fighter's leveling progression
type ProgressionTracker struct {
	Fighter *Fighter
	XPMultiplier float64
	SkillTree map[string]SkillDefinition
}

// SkillDefinition defines when a skill unlocks
type SkillDefinition struct {
	Name          string
	UnlocksAtLevel int
	ManaCost      int
	PowerBonus    int
}

// DefaultSkillTree returns the standard skill progression
func DefaultSkillTree() map[string]SkillDefinition {
	return map[string]SkillDefinition{
		"power_strike": {
			Name:           "Power Strike",
			UnlocksAtLevel: 5,
			ManaCost:       10,
			PowerBonus:     8,
		},
		"defensive_stance": {
			Name:           "Defensive Stance",
			UnlocksAtLevel: 10,
			ManaCost:       5,
			PowerBonus:     0,
		},
		"dual_strike": {
			Name:           "Dual Strike",
			UnlocksAtLevel: 15,
			ManaCost:       15,
			PowerBonus:     12,
		},
		"heal": {
			Name:           "Heal",
			UnlocksAtLevel: 3,
			ManaCost:       20,
			PowerBonus:     0,
		},
	}
}

// NewProgressionTracker creates a new progression tracker
func NewProgressionTracker(fighter *Fighter) *ProgressionTracker {
	return &ProgressionTracker{
		Fighter:       fighter,
		XPMultiplier:  1.0,
		SkillTree:     DefaultSkillTree(),
	}
}

// RecalculateUnlockedSkills updates which skills should be unlocked at current level
func (pt *ProgressionTracker) RecalculateUnlockedSkills() {
	pt.Fighter.SkillsUnlocked = []string{"attack"}

	for skillName, skillDef := range pt.SkillTree {
		if pt.Fighter.Level >= skillDef.UnlocksAtLevel {
			if !pt.Fighter.HasSkill(skillName) {
				pt.Fighter.SkillsUnlocked = append(pt.Fighter.SkillsUnlocked, skillName)
			}
		}
	}
}

// GainXP applies experience gain with multipliers
func (pt *ProgressionTracker) GainXP(baseXP int) {
	xpGain := int(float64(baseXP) * pt.XPMultiplier)
	pt.Fighter.AddXP(xpGain)
	pt.RecalculateUnlockedSkills()
}

// SetXPMultiplier sets the multiplier for XP gains (for training mode, etc)
func (pt *ProgressionTracker) SetXPMultiplier(multiplier float64) {
	pt.XPMultiplier = multiplier
}

// GetXPToNextLevel returns how much XP is needed to level up
func (pt *ProgressionTracker) GetXPToNextLevel() int {
	xpPerLevel := 100
	return xpPerLevel - pt.Fighter.XP
}

// PrestigeReset resets the fighter to level 1 with bonus stats (advanced mechanic)
func (pt *ProgressionTracker) PrestigeReset() {
	oldLevel := pt.Fighter.Level

	// Add prestige bonus
	bonusPerLevel := oldLevel / 5
	pt.Fighter.BaseStats.Power += bonusPerLevel
	pt.Fighter.BaseStats.Defense += bonusPerLevel / 2
	pt.Fighter.BaseStats.Agility += bonusPerLevel / 2
	pt.Fighter.BaseStats.Spirit += bonusPerLevel

	// Reset level and XP
	pt.Fighter.Level = 1
	pt.Fighter.XP = 0

	// Recalculate unlocked skills
	pt.RecalculateUnlockedSkills()
}

// GetLevelProgress returns progress to next level as a percentage
func (pt *ProgressionTracker) GetLevelProgress() float64 {
	xpRequired := 100 + (pt.Fighter.Level * 50)
	return float64(pt.Fighter.XP) / float64(xpRequired) * 100
}
