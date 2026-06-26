# Lab 1: Codebase Investigation Guide

## Objective
Investigate the riftbattle codebase and identify architectural weaknesses, design flaws, and areas for improvement.

## Expected Findings

Students should discover the following fragilities and design issues:

### 1. **Circular Dependencies** (Architecture Issue)
**Location:** `combat/damage.go` (previously imported effects package)  
**What to find:** Tight coupling between damage and effects systems  
**Why it's fragile:** Makes it difficult to test combat in isolation; changes in one system risk breaking the other  
**Expected discovery time:** 10-15 minutes (look at imports and function signatures)

### 2. **Tight Coupling** (Architecture Issue)
**Location:** `game/game.go`  
**What to find:** Game directly mutates Fighter structs (HP, Status)  
**Why it's fragile:** Can't easily swap fighter implementations or add new fighter types  
**Expected discovery time:** 15-20 minutes (examine how Game modifies Fighter fields)

### 3. **String-Based Configuration** (Code Quality Issue)
**Location:** Multiple files (`effects/status.go`, `combat/damage.go`, `modes/modes.go`)  
**What to find:** Magic strings everywhere for effect names, damage types, equipment rarities  
Examples:
- Status effects: `"burning"`, `"poisoned"`, `"frozen"` (effects/status.go line ~65)
- Damage types: `"physical"`, `"magic"`, `"special"` (combat/damage.go line ~48)
- Rarities: `"common"`, `"epic"`, `"legendary"` (equipment/equipment.go line ~7)

**Why it's fragile:** Typos are easy to make and hard to catch; refactoring is error-prone  
**Expected discovery time:** 10-15 minutes (scan for duplicate string literals)

### 4. **Missing Validation** (Security/QA Issue)
**Location:** `equipment/stats.go`, `items/inventory.go`, `equipment/equipment.go`  
**What to find:**
- Equipment can be equipped to invalid slots without validation
- No bounds checking on inventory size
- No validation that equipment level requirements are respected

**Why it's fragile:** Players could theoretically equip two main-hand weapons; inventory could overflow  
**Expected discovery time:** 15-20 minutes (examine Equip, AddItem functions)

### 5. **Dead Code** (Maintenance Issue)
**Location:** `utils/validation.go`  
**What to find:**
- `IsValidStatus()` function is defined but never called
- `CalculateArmorReduction()` function defined but never used
- `MaxInventorySlots` constant defined but never used

**Why it's fragile:** Dead code obscures the real intent; increases codebase complexity unnecessarily  
**Expected discovery time:** 5-10 minutes (use IDE "find usages" feature)

### 6. **Naming Inconsistencies** (Code Quality Issue)
**Location:** Various files  
**What to find:**
- Some files use "Effect" and some use "Status" for the same concept
- Some functions use "armor" and some use "defence"
- Equipment uses "EquipSlot" enum but also takes string slot parameters

**Why it's fragile:** Confusing for new team members; prone to bugs in attribute names  
**Expected discovery time:** 10-15 minutes (compare function signatures across packages)

### 7. **Inefficient Algorithms** (Performance Issue - Optional)
**Location:** `effects/stun_system.go`, `equipment/stats.go`  
**What to find:**
- O(n²) loop in StunSystem.TickStuns (loops through fighters for each stun)
- CalculateTotalBonuses iterates through all equipment slots every time

**Why it's fragile:** Performance degrades with more equipment/fighters  
**Expected discovery time:** 20-25 minutes (analyze nested loops)

## Investigation Strategy

1. **Start broad:** Run `go test ./...` and examine test names for hints about functionality
2. **Follow imports:** Trace imports to find tight coupling
3. **Search for patterns:** Look for repeated code (copy-paste), duplicate strings
4. **Read function signatures:** Look for parameters that indicate tight coupling
5. **Examine error handling:** Notice where validation is missing
6. **Check dead code:** Use IDE "find usages" to identify unused functions

## Success Criteria

Students should identify **at least 4 of the 7 major issues** listed above.

## Related Labs

- **Lab 2** tests poison-healing interaction directly (related to #3 and #4)
- **Lab 4** code review will focus on similar issues to #1, #2, #4
- **Lab 5** shipping process will encourage clean code practices

## Discussion Questions

1. How would you decouple the game and fighter systems?
2. What's the best way to eliminate magic strings?
3. How should equipment validation work?
4. What refactoring would eliminate the circular dependencies?
5. Should dead code be removed immediately or documented?
