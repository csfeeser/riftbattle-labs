package equipment

// Rarity represents equipment rarity levels
type Rarity string

const (
	RarityCommon    Rarity = "common"
	RarityUncommon  Rarity = "uncommon"
	RarityRare      Rarity = "rare"
	RarityEpic      Rarity = "epic"
	RarityLegendary Rarity = "legendary"
)

// EquipmentType represents the type of equipment
type EquipmentType string

const (
	TypeWeapon    EquipmentType = "weapon"
	TypeArmor     EquipmentType = "armor"
	TypeAccessory EquipmentType = "accessory"
)

// EquipSlot represents where equipment can be equipped
type EquipSlot string

const (
	SlotMainHand  EquipSlot = "main_hand"
	SlotOffHand   EquipSlot = "off_hand"
	SlotHead      EquipSlot = "head"
	SlotChest     EquipSlot = "chest"
	SlotLegs      EquipSlot = "legs"
	SlotFeet      EquipSlot = "feet"
	SlotAccessory EquipSlot = "accessory"
)

// Equipment represents an equippable item
type Equipment struct {
	ID              string
	Name            string
	EquipmentType   EquipmentType
	Slot            EquipSlot
	Rarity          Rarity
	RequiredLevel   int
	PowerBonus      int
	DefenseBonus    int
	AgilityBonus    int
	SpiritBonus     int
	CritBonus       int
	EffectResist    map[string]int // Effect resistances (e.g., "poison": 25)
	WeaponClass     string         // For weapons: "sword", "staff", "dagger"
	DamageType      string         // For weapons: "physical", "magic", "special"
	ArmorClass      string         // For armor: "light", "medium", "heavy"
	SpecialAbility  string         // Rare+ equipment might have special abilities
}

// NewWeapon creates a new weapon
func NewWeapon(name, weaponClass, damageType string, power int, level int, rarity Rarity) *Equipment {
	return &Equipment{
		Name:           name,
		EquipmentType:  TypeWeapon,
		Slot:           SlotMainHand,
		Rarity:         rarity,
		RequiredLevel:  level,
		PowerBonus:     power,
		WeaponClass:    weaponClass,
		DamageType:     damageType,
		EffectResist:   make(map[string]int),
	}
}

// NewArmor creates a new armor piece
func NewArmor(name, slot, armorClass string, defense, spirit int, level int, rarity Rarity) *Equipment {
	return &Equipment{
		Name:           name,
		EquipmentType:  TypeArmor,
		Slot:           EquipSlot(slot),
		Rarity:         rarity,
		RequiredLevel:  level,
		DefenseBonus:   defense,
		SpiritBonus:    spirit,
		ArmorClass:     armorClass,
		EffectResist:   make(map[string]int),
	}
}

// NewAccessory creates a new accessory
func NewAccessory(name string, power, defense, agility, spirit int, level int, rarity Rarity) *Equipment {
	return &Equipment{
		Name:           name,
		EquipmentType:  TypeAccessory,
		Slot:           SlotAccessory,
		Rarity:         rarity,
		RequiredLevel:  level,
		PowerBonus:     power,
		DefenseBonus:   defense,
		AgilityBonus:   agility,
		SpiritBonus:    spirit,
		EffectResist:   make(map[string]int),
	}
}

// GetRarityBonus returns the stat multiplier for this equipment's rarity
func (e *Equipment) GetRarityBonus() float64 {
	switch e.Rarity {
	case RarityCommon:
		return 1.0
	case RarityUncommon:
		return 1.15
	case RarityRare:
		return 1.35
	case RarityEpic:
		return 1.6
	case RarityLegendary:
		return 2.0
	default:
		return 1.0
	}
}

// GetTotalPowerBonus calculates power bonus with rarity multiplier
func (e *Equipment) GetTotalPowerBonus() int {
	return int(float64(e.PowerBonus) * e.GetRarityBonus())
}

// CanEquipAt checks if equipment can be equipped at a given level
func (e *Equipment) CanEquipAt(level int) bool {
	return level >= e.RequiredLevel
}

// AddEffectResistance adds resistance to an effect
func (e *Equipment) AddEffectResistance(effect string, resistance int) {
	e.EffectResist[effect] = resistance
}

// GetEffectResistance returns the resistance to an effect
func (e *Equipment) GetEffectResistance(effect string) int {
	if val, ok := e.EffectResist[effect]; ok {
		return val
	}
	return 0
}
