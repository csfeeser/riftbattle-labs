package combat

// ActionType represents the type of action taken
type ActionType string

const (
	ActionAttack  ActionType = "attack"
	ActionDefend  ActionType = "defend"
	ActionCast    ActionType = "cast"
	ActionItem    ActionType = "item"
	ActionFlee    ActionType = "flee"
	ActionStance  ActionType = "stance"
)

// Action represents an action a fighter takes in combat
type Action struct {
	ActionType   ActionType
	SourceName   string
	TargetName   string
	MoveName     string
	ItemID       string
	Parameters   map[string]interface{}
	Timestamp    int64
}

// NewAttackAction creates an attack action
func NewAttackAction(sourceName, targetName, moveName string) *Action {
	return &Action{
		ActionType: ActionAttack,
		SourceName: sourceName,
		TargetName: targetName,
		MoveName:   moveName,
		Parameters: make(map[string]interface{}),
	}
}

// NewDefendAction creates a defend action
func NewDefendAction(sourceName string) *Action {
	return &Action{
		ActionType: ActionDefend,
		SourceName: sourceName,
		Parameters: make(map[string]interface{}),
	}
}

// NewCastAction creates a spell casting action
func NewCastAction(sourceName, targetName, spellName string) *Action {
	return &Action{
		ActionType: ActionCast,
		SourceName: sourceName,
		TargetName: targetName,
		MoveName:   spellName,
		Parameters: make(map[string]interface{}),
	}
}

// NewItemAction creates an item use action
func NewItemAction(sourceName, targetName, itemID string) *Action {
	return &Action{
		ActionType: ActionItem,
		SourceName: sourceName,
		TargetName: targetName,
		ItemID:     itemID,
		Parameters: make(map[string]interface{}),
	}
}

// NewFleeAction creates a flee action
func NewFleeAction(sourceName string) *Action {
	return &Action{
		ActionType: ActionFlee,
		SourceName: sourceName,
		Parameters: make(map[string]interface{}),
	}
}

// IsControlAction checks if action is control-type (stun, paralyze)
func IsControlAction(action *Action) bool {
	move := GetMove(action.MoveName)
	if move == nil {
		return false
	}
	return move.AppliesEffect == "stunned" || move.AppliesEffect == "paralyzed"
}

// ValidateAction checks if an action is valid
func ValidateAction(action *Action) bool {
	// Missing validation - intentional fragility
	// Should validate:
	// 1. Source can act (not stunned, not dead)
	// 2. Move exists
	// 3. MP is sufficient
	// 4. Target exists
	// 5. Target is valid (not already dead, etc)

	return true
}
