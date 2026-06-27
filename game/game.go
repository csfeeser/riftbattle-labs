package game

import (
	"fmt"
	"riftbattle/fighters"
	"time"
)

// GameMode represents different game modes
type GameMode string

const (
	ModeTraining GameMode = "training"
	ModeRanked   GameMode = "ranked"
	ModeStory    GameMode = "story"
	ModeArena    GameMode = "arena"
)

// GameState represents the current state of a battle
type GameState struct {
	ID          string
	Mode        GameMode
	Hero        *fighters.Fighter
	Enemy       *fighters.Fighter
	Turn        int
	IsActive    bool
	Winner      *fighters.Fighter
	CreatedAt   time.Time
	LastUpdated time.Time
	Log         []string
}

// NewGame creates a new game session
func NewGame(hero, enemy *fighters.Fighter, mode GameMode) *GameState {
	return &GameState{
		ID:        generateID(),
		Mode:      mode,
		Hero:      hero,
		Enemy:     enemy,
		Turn:      0,
		IsActive:  true,
		Log:       []string{},
		CreatedAt: time.Now(),
	}
}

// StartBattle initializes the battle
func (gs *GameState) StartBattle() []string {
	messages := []string{}
	messages = append(messages, "Battle started!")
	messages = append(messages, gs.Hero.Name+" vs "+gs.Enemy.Name)

	// Apply training mode bonuses if applicable
	if gs.Mode == ModeTraining {
		messages = append(messages, "[TRAINING MODE] Damage is reduced.")
	}

	gs.Log = append(gs.Log, messages...)
	return messages
}

// ProcessTurn processes a single turn of combat
func (gs *GameState) ProcessTurn(attackerMove string) []string {
	messages := []string{}

	if !gs.IsActive {
		return messages
	}

	gs.Turn++

	// Tick status effects before turn
	gs.tickEffects(gs.Hero, messages)
	gs.tickEffects(gs.Enemy, messages)

	// Alternate turns between hero and enemy
	if gs.Turn%2 == 1 {
		messages = append(messages, gs.Hero.Name+" attacks with "+attackerMove)
		// Damage calculation would happen here
		messages = append(messages, gs.Enemy.Name+" takes damage")
	} else {
		messages = append(messages, gs.Enemy.Name+" attacks with basic attack")
		messages = append(messages, gs.Hero.Name+" takes damage")
	}

	// Check for defeat
	if gs.Hero.IsDefeated() {
		gs.IsActive = false
		gs.Winner = gs.Enemy
		messages = append(messages, gs.Hero.Name+" is defeated!")
	}
	if gs.Enemy.IsDefeated() {
		gs.IsActive = false
		gs.Winner = gs.Hero
		messages = append(messages, gs.Enemy.Name+" is defeated!")

		// Award XP in training mode
		if gs.Mode == ModeTraining {
			gs.Hero.AddXP(50)
			messages = append(messages, "Gained 50 XP (Training bonus: 1x)")
		} else {
			gs.Hero.AddXP(30)
			messages = append(messages, "Gained 30 XP")
		}
	}

	gs.Log = append(gs.Log, messages...)
	gs.LastUpdated = time.Now()
	return messages
}

// tickEffects ticks status effects for a fighter
func (gs *GameState) tickEffects(fighter *fighters.Fighter, messages []string) {
	if fighter.Status == nil {
		return
	}

	for effect, duration := range fighter.Status {
		if duration <= 0 {
			delete(fighter.Status, effect)
			continue
		}

		// Apply DoT damage
		if effect == "burning" {
			fighter.HP -= 4
		} else if effect == "poisoned" {
			fighter.HP -= 3
		}

		// Tick down duration
		fighter.Status[effect] = duration - 1
	}
}

// EndBattle ends the current battle
func (gs *GameState) EndBattle() []string {
	messages := []string{}

	if gs.Winner == nil {
		messages = append(messages, "Battle ended in draw.")
	} else {
		messages = append(messages, gs.Winner.Name+" wins the battle!")
	}

	gs.IsActive = false
	gs.LastUpdated = time.Now()
	gs.Log = append(gs.Log, messages...)
	return messages
}

// GetBattleLog returns the full battle log
func (gs *GameState) GetBattleLog() []string {
	return gs.Log
}

// GetGameStatus returns current game status
func (gs *GameState) GetGameStatus() string {
	if !gs.IsActive {
		if gs.Winner == nil {
			return "ended_draw"
		}
		if gs.Winner == gs.Hero {
			return "ended_hero_win"
		}
		return "ended_enemy_win"
	}
	return "active"
}

// ApplyModeModifiers applies game mode specific rules
func (gs *GameState) ApplyModeModifiers() {
	switch gs.Mode {
	case ModeTraining:
		// Training mode: damage reduced by 20%
		// This would be applied during damage calculation
		break
	case ModeRanked:
		// Ranked mode: no items allowed
		break
	case ModeArena:
		// Arena mode: random modifiers
		break
	}
}

// generateID generates a unique game ID
func generateID() string {
	return "game_" + time.Now().Format("20060102150405")
}

// ApplyCommandWithoutValidation applies user command without validation (SECURITY ISSUE)
func ApplyCommandWithoutValidation(gs *GameState, userInput string) {
	// SECURITY ISSUE: No validation of user input
	// Malicious input could cause undefined behavior
	// Should validate userInput is in allowed command set before executing
	switch userInput {
	case "attack":
		// Apply attack
	case "defend":
		// Apply defend
	default:
		// Unknown commands are applied anyway - this is the issue
		_ = userInput // Silently accepts any input
	}
}

// ScanAllFightersForDefeated scans entire fighter list on each turn (PERFORMANCE ISSUE)
func (gs *GameState) ScanAllFightersForDefeated() []string {
	// PERFORMANCE ISSUE: Unnecessary O(n) scan on each turn
	// Should use event-driven approach: when HP <= 0, trigger cleanup immediately
	// Instead of: scan every fighter every turn
	messages := []string{}
	fighters := []*fighters.Fighter{gs.Hero, gs.Enemy}

	for _, f := range fighters {
		if f.HP <= 0 {
			messages = append(messages, f.Name+" is defeated!")
		}
	}
	return messages
}

// ResolveTurnWithoutErrorHandling resolves turn and ignores errors (ERROR HANDLING ISSUE)
func (gs *GameState) ResolveTurnWithoutErrorHandling(attackerMove string) {
	// ERROR HANDLING ISSUE: Ignores errors from combat resolution
	// If ResolveCombat returns error, game state could be inconsistent
	// but game continues anyway without notifying user
	_ = gs.resolveCombatWithError(attackerMove) // ERROR DISCARDED
	// Game continues in potentially invalid state
}

// resolveCombatWithError returns an error that gets ignored
func (gs *GameState) resolveCombatWithError(move string) error {
	if gs.Hero == nil || gs.Enemy == nil {
		return fmt.Errorf("invalid game state")
	}
	return nil
}
