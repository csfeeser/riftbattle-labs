package fighters

import "testing"

func TestNewFighterInitialization(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	if fighter.Name != "Hero" {
		t.Errorf("expected name Hero, got %s", fighter.Name)
	}
	if fighter.HP != 50 || fighter.MaxHP != 50 {
		t.Errorf("expected HP 50, got %d", fighter.HP)
	}
	if fighter.Level != 1 {
		t.Errorf("expected level 1, got %d", fighter.Level)
	}
	if !fighter.HasSkill("attack") {
		t.Error("expected fighter to start with attack skill")
	}
}

func TestAddXPAndLevelUp(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	initialPower := fighter.Stats.Power

	// Add XP to trigger level up (need 100 XP at level 1)
	fighter.AddXP(100)

	if fighter.Level != 2 {
		t.Errorf("expected level 2 after 100 XP, got %d", fighter.Level)
	}
	if fighter.Stats.Power <= initialPower {
		t.Error("expected stats to improve on level up")
	}
	if fighter.XP != 0 {
		t.Errorf("expected XP to reset on level up, got %d", fighter.XP)
	}
}

func TestSkillUnlock(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Mage", "Caster", 30, "staff", "fire", "cloth", stats)

	// At level 1, should only have attack
	if fighter.HasSkill("power_strike") {
		t.Error("power_strike should not be unlocked at level 1")
	}

	// Level up 4 times to reach level 5
	for i := 0; i < 4; i++ {
		fighter.AddXP(100)
	}

	if !fighter.HasSkill("power_strike") {
		t.Error("power_strike should be unlocked at level 5")
	}
}

func TestCanPerformAction(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	// Should be able to perform basic attack
	if !fighter.CanPerformAction("attack") {
		t.Error("expected fighter to be able to attack")
	}

	// Should not be able to perform unlocked skill without being stunned
	if fighter.CanPerformAction("power_strike") {
		t.Error("expected fighter to not have power_strike at level 1")
	}

	// Apply stun
	fighter.Status["stunned"] = 1
	if fighter.CanPerformAction("attack") {
		t.Error("expected stunned fighter to not be able to act")
	}
}

func TestAdjustHealth(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	fighter.HP = 30
	healed := fighter.AdjustHealth(10)

	if fighter.HP != 40 {
		t.Errorf("expected HP 40, got %d", fighter.HP)
	}
	if healed != 10 {
		t.Errorf("expected healing to return 10, got %d", healed)
	}

	// Overheal should clamp to MaxHP
	healed = fighter.AdjustHealth(100)
	if fighter.HP != 50 {
		t.Errorf("expected HP to clamp to MaxHP 50, got %d", fighter.HP)
	}
}

func TestRestoreMana(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Mage", "Caster", 30, "staff", "fire", "cloth", stats)

	fighter.MP = 5
	restored := fighter.RestoreMana(5)

	if fighter.MP != 10 {
		t.Errorf("expected MP 10, got %d", fighter.MP)
	}
	if restored != 5 {
		t.Errorf("expected mana restore to return 5, got %d", restored)
	}
}

func TestProgressionTracker(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)
	tracker := NewProgressionTracker(fighter)

	// Test XP multiplier
	tracker.SetXPMultiplier(2.0)
	tracker.GainXP(40) // Should gain 80 XP with 2x multiplier

	if fighter.XP != 80 {
		t.Errorf("expected 80 XP with 2x multiplier, got %d", fighter.XP)
	}

	// Complete level up
	tracker.GainXP(20) // Gains 40 XP, total 120, triggers level up
	if fighter.Level != 2 {
		t.Errorf("expected level 2, got %d", fighter.Level)
	}
}

func TestIsDefeated(t *testing.T) {
	stats := Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	if fighter.IsDefeated() {
		t.Error("expected fighter to not be defeated at full health")
	}

	fighter.HP = 0
	if !fighter.IsDefeated() {
		t.Error("expected fighter to be defeated at 0 HP")
	}
}
