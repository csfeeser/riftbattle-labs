package fighters

import "time"

type EquipSlot string

const (
	SlotMainHand EquipSlot = "main_hand"
	SlotOffHand  EquipSlot = "off_hand"
	SlotHead     EquipSlot = "head"
	SlotChest    EquipSlot = "chest"
	SlotLegs     EquipSlot = "legs"
	SlotFeet     EquipSlot = "feet"
	SlotAccessory EquipSlot = "accessory"
)

// Stats represents a fighter's core attributes
type Stats struct {
	Power     int
	Defense   int
	Agility   int
	Spirit    int
	CritChance int
}

// Fighter represents a combatant in the game
type Fighter struct {
	// Identity
	ID   string
	Name string
	Class string

	// Health & Resources
	HP    int
	MaxHP int
	MP    int
	MaxMP int

	// Base attributes
	Stats      Stats
	BaseStats  Stats

	// Equipment (NEW)
	Equipment map[EquipSlot]interface{} // Will hold equipment items

	// Combat state
	Status map[string]int

	// Progression (NEW)
	Level      int
	XP         int
	SkillsUnlocked []string

	// Inventory (NEW - reference to inventory package)
	InventoryID string

	// Game state
	TrainingMode bool
	LastActionTime time.Time

	// Legacy fields for compatibility
	WeaponType string
	Element    string
	ArmorType  string
}

// NewFighter creates a new fighter with the given parameters
func NewFighter(name, class string, hp int, weaponType, element, armorType string, stats Stats) *Fighter {
	maxMP := stats.Spirit * 2
	return &Fighter{
		Name:       name,
		Class:      class,
		HP:         hp,
		MaxHP:      hp,
		MP:         maxMP,
		MaxMP:      maxMP,
		Stats:      stats,
		BaseStats:  stats,
		Status:     map[string]int{},
		Level:      1,
		XP:         0,
		SkillsUnlocked: []string{"attack"}, // Everyone starts with basic attack
		Equipment: make(map[EquipSlot]interface{}),
		WeaponType: weaponType,
		Element:    element,
		ArmorType:  armorType,
	}
}

// AddXP adds experience points and checks for level up
func (f *Fighter) AddXP(amount int) bool {
	f.XP += amount
	xpPerLevel := 100
	if f.XP >= xpPerLevel {
		f.LevelUp()
		return true
	}
	return false
}

// LevelUp increases level and stat growth
func (f *Fighter) LevelUp() {
	f.Level++
	f.BaseStats.Power += 2
	f.BaseStats.Defense += 1
	f.BaseStats.Agility += 1
	f.BaseStats.Spirit += 2

	// Recalculate effective stats
	f.Stats = f.GetEffectiveStats()

	// Restore health on level up
	f.MaxHP += 10
	f.HP = f.MaxHP
	f.MaxMP = f.Stats.Spirit * 2
	f.MP = f.MaxMP

	// Unlock skills at certain levels
	if f.Level == 5 {
		f.SkillsUnlocked = append(f.SkillsUnlocked, "power_strike")
	}
	if f.Level == 10 {
		f.SkillsUnlocked = append(f.SkillsUnlocked, "defensive_stance")
	}

	f.XP = 0
}

// GetEffectiveStats calculates stats including equipment bonuses
func (f *Fighter) GetEffectiveStats() Stats {
	stats := f.BaseStats

	// Equipment bonuses would be applied here
	// For now, just return base stats
	// This will be expanded when equipment package is integrated

	return stats
}

// HasSkill checks if the fighter has unlocked a skill
func (f *Fighter) HasSkill(skill string) bool {
	for _, s := range f.SkillsUnlocked {
		if s == skill {
			return true
		}
	}
	return false
}

// CanPerformAction checks if fighter can perform an action
func (f *Fighter) CanPerformAction(actionName string) bool {
	// Can't act if stunned or defeated
	if status, ok := f.Status["stunned"]; ok && status > 0 {
		return false
	}
	if f.IsDefeated() {
		return false
	}

	// Check if skill is unlocked
	if actionName != "attack" && actionName != "defend" {
		return f.HasSkill(actionName)
	}

	return true
}

// IsDefeated checks if the fighter is defeated
func (f *Fighter) IsDefeated() bool {
	return f.HP <= 0
}

// CanAct checks if the fighter can act this turn
func (f *Fighter) CanAct() bool {
	return f.CanPerformAction("attack")
}
