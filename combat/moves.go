package combat

// Move represents an attack or ability
type Move struct {
	Name           string
	Power          int
	Element        string
	DamageType     string
	AppliesEffect  string
	EffectDuration int
	ManaCost       int
}

// MoveLibrary contains all defined moves
var MoveLibrary = map[string]Move{
	"slash": {
		Name:       "Slash",
		Power:      8,
		Element:    "none",
		DamageType: "physical",
	},
	"fireball": {
		Name:           "Fireball",
		Power:          10,
		Element:        "fire",
		DamageType:     "magic",
		AppliesEffect:  "burning",
		EffectDuration: 2,
		ManaCost:       15,
	},
	"ice_lance": {
		Name:           "Ice Lance",
		Power:          9,
		Element:        "ice",
		DamageType:     "magic",
		AppliesEffect:  "frozen",
		EffectDuration: 1,
		ManaCost:       12,
	},
	"shield_bash": {
		Name:           "Shield Bash",
		Power:          6,
		Element:        "none",
		DamageType:     "physical",
		AppliesEffect:  "stunned",
		EffectDuration: 1,
	},
	"power_strike": {
		Name:       "Power Strike",
		Power:      14,
		Element:    "none",
		DamageType: "physical",
		ManaCost:   10,
	},
	"heal": {
		Name:       "Heal",
		Power:      0,
		Element:    "none",
		DamageType: "heal",
		ManaCost:   20,
	},
	"poison_cloud": {
		Name:           "Poison Cloud",
		Power:          5,
		Element:        "poison",
		DamageType:     "magic",
		AppliesEffect:  "poisoned",
		EffectDuration: 3,
		ManaCost:       18,
	},
	"lightning_strike": {
		Name:           "Lightning Strike",
		Power:          12,
		Element:        "lightning",
		DamageType:     "magic",
		AppliesEffect:  "paralyzed",
		EffectDuration: 1,
		ManaCost:       16,
	},
}

// GetMove gets a move from the library
func GetMove(name string) *Move {
	if move, ok := MoveLibrary[name]; ok {
		return &move
	}
	return nil
}

// IsPhysicalMove checks if a move is physical damage
func IsPhysicalMove(move *Move) bool {
	return move.DamageType == "physical"
}

// IsMagicalMove checks if a move is magical damage
func IsMagicalMove(move *Move) bool {
	return move.DamageType == "magic"
}

// CanCast checks if fighter can cast this move
func CanCast(fighterMP int, move *Move) bool {
	return fighterMP >= move.ManaCost
}
