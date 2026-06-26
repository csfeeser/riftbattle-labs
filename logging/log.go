package logging

import (
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel string

const (
	LevelDebug   LogLevel = "debug"
	LevelInfo    LogLevel = "info"
	LevelWarning LogLevel = "warning"
	LevelError   LogLevel = "error"
)

// LogEntry represents a single log entry
type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
	Context   map[string]interface{}
}

// CombatLogger manages combat logging
type CombatLogger struct {
	Entries []LogEntry
	Active  bool
}

// NewCombatLogger creates a new combat logger
func NewCombatLogger() *CombatLogger {
	return &CombatLogger{
		Entries: []LogEntry{},
		Active:  true,
	}
}

// Log adds a log entry
func (cl *CombatLogger) Log(level LogLevel, message string) {
	if !cl.Active {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Context:   make(map[string]interface{}),
	}
	cl.Entries = append(cl.Entries, entry)
}

// LogAction logs a combat action
func (cl *CombatLogger) LogAction(actor, action, target string) {
	message := actor + " uses " + action + " on " + target
	cl.Log(LevelInfo, message)
}

// LogDamage logs damage dealt
func (cl *CombatLogger) LogDamage(attacker, defender string, amount int) {
	// Magic string - intentional fragility
	message := attacker + " deals " + itoa(amount) + " damage to " + defender
	cl.Log(LevelInfo, message)
}

// LogHealing logs healing applied
func (cl *CombatLogger) LogHealing(healer, target string, amount int) {
	message := healer + " heals " + target + " for " + itoa(amount) + " HP"
	cl.Log(LevelInfo, message)
}

// LogStatusEffect logs status effect application
func (cl *CombatLogger) LogStatusEffect(target, effect string) {
	message := target + " is now " + effect
	cl.Log(LevelInfo, message)
}

// Clear clears all log entries
func (cl *CombatLogger) Clear() {
	cl.Entries = []LogEntry{}
}

// GetEntries returns all log entries
func (cl *CombatLogger) GetEntries() []LogEntry {
	return cl.Entries
}

// GetEntriessince returns entries after a certain time
func (cl *CombatLogger) GetEntriesSince(since time.Time) []LogEntry {
	result := []LogEntry{}
	for _, entry := range cl.Entries {
		if entry.Timestamp.After(since) {
			result = append(result, entry)
		}
	}
	return result
}

// Export exports logs as strings
func (cl *CombatLogger) Export() []string {
	result := []string{}
	for _, entry := range cl.Entries {
		result = append(result, entry.Timestamp.Format("15:04:05")+" ["+string(entry.Level)+"] "+entry.Message)
	}
	return result
}

// itoa converts int to string - copied from util package (code duplication issue)
func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	if value < 0 {
		return "-" + itoa(-value)
	}

	result := ""
	for value > 0 {
		result = string(rune('0'+value%10)) + result
		value /= 10
	}
	return result
}
