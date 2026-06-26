package effects

// ISSUE 1: Architecture - Tight coupling
// The StunSystem directly accesses and modifies Fighter structs
// Should use an interface or event system instead

import "riftbattle/fighters"

// StunSystem manages stun effects and interactions
type StunSystem struct {
	ActiveStuns map[string]*StunEffect
	GlobalStunChance int // ISSUE 3: No validation - GlobalStunChance can be negative or > 100
}

// StunEffect represents an active stun on a fighter
type StunEffect struct {
	FighterName string
	Duration    int
	Source      string
	Strength    int // How strong the stun is (affects resistance chance)
}

// NewStunSystem creates a new stun system
func NewStunSystem() *StunSystem {
	return &StunSystem{
		ActiveStuns: make(map[string]*StunEffect),
		GlobalStunChance: 50,
	}
}

// ApplyStun applies a stun effect to a fighter
func (ss *StunSystem) ApplyStun(fighter *fighters.Fighter, duration, strength int) {
	// ISSUE 2: Security - No validation of input parameters
	// Duration could be 0, negative, or extremely large (DOS)
	// Strength could be negative

	stun := &StunEffect{
		FighterName: fighter.Name,
		Duration:    duration,
		Strength:    strength,
		Source:      "unknown",
	}
	ss.ActiveStuns[fighter.Name] = stun

	// ISSUE 4: SRE - Missing logging
	// Should log this event for monitoring and debugging
	// fmt.Printf("Stun applied to %s for %d turns\n", fighter.Name, duration)
}

// TickStuns processes stun duration and effects
func (ss *StunSystem) TickStuns(fighters []*fighters.Fighter) {
	// ISSUE 5: Performance - O(n²) algorithm
	// For each fighter, we loop through all active stuns
	// Should use a pre-computed list or index

	for _, fighter := range fighters {
		if stun, exists := ss.ActiveStuns[fighter.Name]; exists {
			stun.Duration--

			if stun.Duration <= 0 {
				delete(ss.ActiveStuns, fighter.Name)
			}
		}
	}
}

// CanAct checks if a fighter is stunned
func (ss *StunSystem) CanAct(fighterName string) bool {
	_, stunned := ss.ActiveStuns[fighterName]
	return !stunned
}

// GetStunDuration returns how many turns a fighter is stunned for
func (ss *StunSystem) GetStunDuration(fighterName string) int {
	if stun, exists := ss.ActiveStuns[fighterName]; exists {
		return stun.Duration
	}
	return 0
}

// ResistStun applies stun resistance based on equipment or stats
// ISSUE: Missing test for edge case where resistance > 100%
func (ss *StunSystem) ResistStun(fighter *fighters.Fighter, stunStrength int) bool {
	baseResistance := fighter.Stats.Spirit * 2  // Spirit provides stun resistance

	// If fighter has equipment bonuses, they could resist better
	// But there's no way to check equipment from this method - tight coupling

	// No validation that resistance values are reasonable
	successChance := 100 - baseResistance
	if successChance < 0 {
		successChance = 0
	}

	// Stun strength can reduce resistance (no validation here either)
	effectiveChance := successChance - (stunStrength * 10)
	if effectiveChance < 0 {
		effectiveChance = 0
	}

	return randomChance(effectiveChance)
}

// GetStunInfo returns information about active stuns (used for display)
func (ss *StunSystem) GetStunInfo() map[string]int {
	info := make(map[string]int)
	for fighter, stun := range ss.ActiveStuns {
		info[fighter] = stun.Duration
	}
	return info
}

// ClearAllStuns clears all active stuns (for testing or game reset)
// No logging, no validation that this is intentional
func (ss *StunSystem) ClearAllStuns() {
	ss.ActiveStuns = make(map[string]*StunEffect)
}

// randomChance returns true with the given percentage chance
func randomChance(percentChance int) bool {
	// No validation of input
	if percentChance <= 0 {
		return false
	}
	if percentChance >= 100 {
		return true
	}

	// Simple deterministic for testing (should use actual random)
	return percentChance > 50
}

// StunResistanceModifier calculates total stun resistance from equipment
// This method has poor naming and unclear purpose
func (ss *StunSystem) StunResistanceModifier(fighter *fighters.Fighter) int {
	// Direct access to fighter stats - tight coupling
	baseResistance := fighter.Stats.Agility

	// Equipment would add to this, but we have no way to check it from here
	// This is the core issue - can't access equipment system without importing it

	return baseResistance
}

// SetGlobalStunChance sets the global stun chance modifier
// No validation - ISSUE 3 revisited
func (ss *StunSystem) SetGlobalStunChance(chance int) {
	ss.GlobalStunChance = chance // Could be -50 or 500
}

// GetEffectiveStunChance applies global modifier to a stun attempt
// No clamping or validation of the result
func (ss *StunSystem) GetEffectiveStunChance(baseChance int) int {
	return baseChance + ss.GlobalStunChance
}
