# RiftBattle: A Go Combat Game Engine

RiftBattle is a turn-based combat game engine written in Go. This codebase is designed as a teaching resource for professional software development workflows—investigation, debugging, code review, and shipping practices.

## What Is RiftBattle?

A combat game where characters engage in turn-based battles. Features include:
- **Character Progression:** Leveling system with experience points, stat growth, and skill unlocks
- **Combat System:** Turn-based damage calculation with elemental effects and critical strikes
- **Status Effects:** Poison, burn, stun, and other effects that persist across turns
- **Equipment System:** Weapons and armor with rarity-based stat bonuses
- **Inventory Management:** Item collection, consumption, and effect management
- **Game Modes:** Training, ranked, story, and arena modes with different rules
- **Combat Logging:** Full action history and battle statistics

## Codebase Overview

**Metrics:**
- 2,421 lines of code (3,302 with tests)
- 8 packages
- 15+ tests
- Pure Go (no external dependencies for core logic)

**Package Structure:**

```
riftbattle/
├── fighters/       # Character data, progression, stats, leveling
├── equipment/      # Equipment items, rarity, bonuses, slot management
├── items/          # Consumables (potions, antidotes), inventory
├── effects/        # Status effects (poison, burn, stun), damage-over-time
├── combat/         # Damage calculation, moves, turn resolution
├── game/           # Game state, battle management, mode rules
├── logging/        # Combat action logging and export
├── modes/          # Game mode definitions and multipliers
├── utils/          # Math utilities, validation helpers
└── main.go         # Entry point
```

## Branches

Each branch represents a specific state of the codebase for teaching purposes.

### `main` — Production-Ready Codebase
**Use for:** Labs 1 & 3 (investigation and assessment)

Complete, working implementation with intentional architectural fragility:
- All tests pass
- App runs without errors
- Contains discoverable design issues for investigation:
  - Tight coupling between packages
  - Missing validation in critical functions
  - Dead code and unused utilities
  - String-based configuration (should be typed)
  - Circular dependencies in some areas
  - Naming inconsistencies

### `lab-2-with-bug` — Bug Injection for Lab 2
**Use for:** Lab 2 (bug hunt and diagnosis)

Same as main, but with a single deliberate bug injected:
- **Bug:** Poisoned characters can heal (they shouldn't)
- **Location:** `items/healing.go` - poison check is removed
- **Test Impact:** 2 tests fail:
  - `TestHealingPotion_PoisonedCharacterCannotHeal`
  - `TestHealingPotion_MultiplePoisonStillBlocksHealing`
- **Discovery difficulty:** Medium (~15-20 minutes)
- **Fix:** Add 3 lines of poison-checking logic back

### `feature/stun-system` — Code Review Target for Lab 4
**Use for:** Lab 4 (code review from multiple perspectives)

Expanded stun system with 5 intentional code quality issues:
- All tests pass ✅
- Feature is functional but has design problems
- Contains embedded issues for review:
  1. **Architecture:** Tight coupling between stun system and fighter state
  2. **Security:** Missing validation on stun duration (could be infinite)
  3. **QA:** Incomplete test coverage for edge cases
  4. **SRE:** No logging when stuns are applied or tick down
  5. **Performance:** O(n²) algorithm for stun resolution
- Students review from 5 different perspectives

### `feature/equipment-system` — Shipping Workflow for Lab 5
**Use for:** Lab 5 (shipping readiness and commit history)

Complete equipment system with 5 sequential commits:
- All tests pass
- Production-ready implementation
- Clean commit history showing realistic development:
  1. Equipment data types and constructors
  2. Fighter integration (equipment bonuses)
  3. Display and comparison functions
  4. Comprehensive test suite
  5. Bug fix (prevent equipment loss on unequip)
- Students evaluate commit quality and shipping readiness

## Getting Started

### Prerequisites
- Go 1.22 or later
- Git

### Setup
```bash
# Clone the repository
git clone https://github.com/csfeeser/riftbattle-labs.git
cd riftbattle-labs

# Verify the main branch works
git checkout main
go run .
go test ./...
```

### Checking Out Different Branches
```bash
# For Lab 1 & 3 (investigation and assessment)
git checkout main

# For Lab 2 (bug hunt)
git checkout lab-2-with-bug

# For Lab 4 (code review)
git checkout feature/stun-system

# For Lab 5 (shipping workflow)
git checkout feature/equipment-system
```

## Lab Documentation

Each lab has detailed guidance in the repo root:
- **LAB_1_INVESTIGATION_NOTES.md** — What to discover on main branch
- **LAB_2_BUG_SPECIFICATION.md** — Bug details and investigation guide
- **LAB_4_PREPARED_DIFF.md** — Code review issues and how to find them
- **LAB_5_BRANCH_STATE.md** — Shipping workflow evaluation criteria
- **RIFTBATTLE_LABS_README.md** — Complete overview and setup

## Running Tests

```bash
# All tests
go test ./...

# Specific package
go test ./fighters

# Specific test
go test -run TestAddXP ./fighters

# Lab 2 specific (see expected failures)
go test -run Poisoned ./items
```

## Code Quality

**Main branch:**
- 100% test pass rate
- Intentional architectural issues for discovery
- No external dependencies
- Realistic codebase complexity

**Feature branches:**
- All tests pass despite embedded issues
- Issues are subtle and require careful review
- Realistic commit history (Lab 5)

## Architecture Notes

**Intentional Fragility** (for learning):
- The main branch deliberately couples packages to teach modularity
- Some functions are defined but never used (dead code)
- String-based effect management instead of typed constants
- Missing validation on critical functions
- Naming inconsistencies between modules

This fragility is **intentional** and designed to teach investigation skills. Production code should not follow these patterns.
