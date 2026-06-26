package fighters

// StatModifier represents a temporary stat boost
type StatModifier struct {
	StatName string
	Bonus    int
	Duration int
}

// ApplyModifier adds a temporary stat bonus
func (f *Fighter) ApplyModifier(modifier StatModifier) {
	switch modifier.StatName {
	case "power":
		f.Stats.Power += modifier.Bonus
	case "defense":
		f.Stats.Defense += modifier.Bonus
	case "agility":
		f.Stats.Agility += modifier.Bonus
	case "spirit":
		f.Stats.Spirit += modifier.Bonus
	case "crit":
		f.Stats.CritChance += modifier.Bonus
	}
}

// TickModifiers decrements stat modifier durations
func (f *Fighter) TickModifiers() {
	// This will be expanded when modifiers are stored
}

// CalculateStatBonus calculates the total stat bonus from equipment and effects
func (f *Fighter) CalculateStatBonus() Stats {
	bonus := Stats{}

	// Equipment bonuses would be calculated here
	// Tight coupling issue: game.go directly accesses Fighter.Stats
	// Should be abstracted away

	return bonus
}

// GetStatByName gets a stat value by name
func (f *Fighter) GetStatByName(statName string) int {
	switch statName {
	case "power":
		return f.Stats.Power
	case "defense":
		return f.Stats.Defense
	case "agility":
		return f.Stats.Agility
	case "spirit":
		return f.Stats.Spirit
	case "crit":
		return f.Stats.CritChance
	default:
		return 0
	}
}

// AdjustHealth adjusts the fighter's HP
func (f *Fighter) AdjustHealth(amount int) int {
	before := f.HP
	f.HP = clamp(f.HP+amount, 0, f.MaxHP)
	return f.HP - before
}

// RestoreMana restores MP
func (f *Fighter) RestoreMana(amount int) int {
	before := f.MP
	f.MP = clamp(f.MP+amount, 0, f.MaxMP)
	return f.MP - before
}

// clamp returns value bounded by min and max
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
