package effects

import (
	"riftbattle/fighters"
	"testing"
)

func TestNewStunSystem(t *testing.T) {
	ss := NewStunSystem()

	if ss == nil {
		t.Fatal("expected StunSystem to be created")
	}
	if len(ss.ActiveStuns) != 0 {
		t.Error("expected no active stuns on creation")
	}
}

func TestApplyStun(t *testing.T) {
	ss := NewStunSystem()
	stats := fighters.Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := fighters.NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	ss.ApplyStun(fighter, 2, 5)

	if ss.ActiveStuns["Hero"].Duration != 2 {
		// ISSUE 2 (Security): This test passes even with invalid input (no validation)
		t.Error("expected stun duration 2")
	}
}

func TestCanAct(t *testing.T) {
	ss := NewStunSystem()
	stats := fighters.Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := fighters.NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	if !ss.CanAct("Hero") {
		t.Error("expected hero to be able to act initially")
	}

	ss.ApplyStun(fighter, 1, 5)
	if ss.CanAct("Hero") {
		t.Error("expected hero to not be able to act while stunned")
	}
}

func TestGetStunDuration(t *testing.T) {
	ss := NewStunSystem()
	stats := fighters.Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := fighters.NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	duration := ss.GetStunDuration("Hero")
	if duration != 0 {
		t.Errorf("expected duration 0 for non-stunned fighter, got %d", duration)
	}

	ss.ApplyStun(fighter, 3, 5)
	duration = ss.GetStunDuration("Hero")
	if duration != 3 {
		t.Errorf("expected duration 3, got %d", duration)
	}
}

func TestTickStuns(t *testing.T) {
	ss := NewStunSystem()
	stats := fighters.Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := fighters.NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	ss.ApplyStun(fighter, 2, 5)

	// First tick: duration 2 -> 1
	ss.TickStuns([]*fighters.Fighter{fighter})

	if ss.GetStunDuration("Hero") != 1 {
		t.Error("expected stun to tick down to 1")
	}

	// Second tick: duration 1 -> 0 (stun is deleted when duration <= 0)
	ss.TickStuns([]*fighters.Fighter{fighter})

	if !ss.CanAct("Hero") {
		t.Error("expected hero to be able to act after stun expires")
	}
}

// ISSUE: Missing test for edge case - what happens if we tick with nil fighters slice?
// ISSUE: Missing test for resistance with extremely high/low stats
// ISSUE: Missing test for GlobalStunChance with invalid values

func TestSetGlobalStunChance(t *testing.T) {
	ss := NewStunSystem()

	// No validation - this is valid per the code but doesn't make sense
	ss.SetGlobalStunChance(-100)
	ss.SetGlobalStunChance(500)

	// Both pass without error
	if ss.GlobalStunChance != 500 {
		t.Error("expected global stun chance to be set")
	}
}

func TestResistStun(t *testing.T) {
	ss := NewStunSystem()
	stats := fighters.Stats{Power: 5, Defense: 4, Agility: 3, Spirit: 6, CritChance: 10}
	fighter := fighters.NewFighter("Hero", "Warrior", 50, "sword", "fire", "heavy", stats)

	// High spirit should give good stun resistance
	result := ss.ResistStun(fighter, 1)

	// ISSUE: This test only checks that function doesn't crash
	// Doesn't actually validate the resistance calculation
	if result != (result == result) {
		t.Error("unexpected result")
	}
}
