# Lab 4: Code Review - Stun System Feature

## Branch Information

**Branch:** `feature/stun-system`  
**Base:** main (at Phase 1)  
**Feature:** Stun system expansion with effect interactions and resistance  
**Status:** Ready for review - all tests pass, but contains intentional issues for learning

## What Was Changed

### New Files
- `effects/stun_system.go` (271 lines) - Core stun system implementation
- `effects/stun_system_test.go` (~90 lines) - Test coverage (incomplete)

### What the Feature Does
- Manages stun effects on fighters
- Tracks stun duration and strength
- Calculates stun resistance based on fighter stats
- Allows checking if a fighter can act
- Provides UI information about active stuns

## Issues to Find During Review

### Issue 1: ARCHITECTURE - Tight Coupling ⭐ MAJOR
**Severity:** Medium  
**Location:** `effects/stun_system.go` lines 16-20, 32-38  
**Problem:**
```go
type StunEffect struct {
    FighterName string  // Reference to fighter, tight coupling
    ...
}
// Direct access in ApplyStun:
stun := &StunEffect{
    FighterName: fighter.Name,  // Using fighter object directly
    ...
}
```

**Why it's an issue:** 
- StunSystem directly accesses Fighter structs
- Can't easily swap fighter implementations
- Makes testing harder (need real fighters)
- Violates separation of concerns

**Suggested fix:** Use fighter ID/identifier, abstract interface between systems

---

### Issue 2: SECURITY - Missing Input Validation ⭐ MAJOR
**Severity:** High  
**Location:** `effects/stun_system.go` lines 32-38, 47-52  
**Problems:**
```go
func (ss *StunSystem) ApplyStun(fighter *fighters.Fighter, duration, strength int) {
    // No validation! What if duration is negative? What if strength is 1000?
    ...
}

func (ss *StunSystem) SetGlobalStunChance(chance int) {
    ss.GlobalStunChance = chance  // Could be -50 or 500!
}
```

**Why it's an issue:**
- No bounds checking on duration (could be negative, causing logic errors)
- No bounds checking on strength or chance (could be > 100%)
- Potential DOS if duration is extremely large
- Could cause undefined behavior

**Suggested fix:**
```go
func (ss *StunSystem) ApplyStun(fighter *fighters.Fighter, duration, strength int) error {
    if duration <= 0 || duration > 100 {
        return fmt.Errorf("invalid duration: %d", duration)
    }
    if strength < 0 || strength > 100 {
        return fmt.Errorf("invalid strength: %d", strength)
    }
    ...
}
```

---

### Issue 3: QA - Missing Edge Case Tests ⭐ MEDIUM
**Severity:** Medium  
**Location:** `effects/stun_system_test.go` lines ~95  
**Problem:** Test coverage is incomplete:
```go
// MISSING TESTS FOR:
// - What happens with negative duration?
// - What happens with extreme stat values?
// - What happens if we tick with nil fighters slice?
// - What happens if resistance calculation results in > 100%?
```

**Why it's an issue:**
- Edge cases aren't validated
- Could crash in production with unexpected input
- Tests pass but feature is fragile

**Suggested fix:** Add tests for:
- Negative duration/strength values (should error or be rejected)
- Nil fighter in TickStuns
- Edge case: resistance calculation overflow

---

### Issue 4: SRE - Missing Logging ⭐ MEDIUM
**Severity:** Low-Medium  
**Location:** `effects/stun_system.go` lines 32-38  
**Problem:**
```go
func (ss *StunSystem) ApplyStun(...) {
    // No logging! How do we debug in production?
    // No way to track when stuns are applied
    ss.ActiveStuns[fighter.Name] = stun
    // Comment says what should be here:
    // fmt.Printf("Stun applied to %s for %d turns\n", ...)
}
```

**Why it's an issue:**
- No visibility into system behavior
- Difficult to debug issues in production
- Hard to trace when stuns apply/expire
- Operations team can't monitor effectiveness

**Suggested fix:** Add logging (see comment in code for example)

---

### Issue 5: PERFORMANCE - O(n²) Algorithm ⭐ MEDIUM
**Severity:** Low (but important at scale)  
**Location:** `effects/stun_system.go` lines 56-67  
**Problem:**
```go
func (ss *StunSystem) TickStuns(fighters []*fighters.Fighter) {
    // O(n²) - for each fighter in list,
    // check map of active stuns
    // Better: pre-compute fighter → stun mapping
    for _, fighter := range fighters {  // O(n)
        if stun, exists := ss.ActiveStuns[fighter.Name] {  // O(1) lookup but
            // ... iterating through all fighters for all stuns is O(n)
        }
    }
}
```

**Why it's an issue:**
- Scales poorly with many stuns/fighters
- Could be O(1) with better data structure
- Not critical now, but will matter at scale

**Suggested fix:**
- Pre-index active stuns by fighter name
- Or maintain reverse mapping

---

## How to Review This Code

1. **Read the feature description** (above)
2. **Look for the 5 issues** listed above (they're marked with ⭐)
3. **Check the severity:**
   - MAJOR issues: Must fix before merging
   - MEDIUM issues: Should fix or document
4. **Verify:** Run tests and check they all pass
   ```bash
   git checkout feature/stun-system
   go test ./effects -v
   go run .  # Should work
   ```

## Expected Review Time

**20-30 minutes** to find and understand all issues:
- 5 min: Read the feature and new code
- 10-15 min: Identify the 5 main issues
- 5-10 min: Write review comments

## Success Criteria for Students

✅ Found at least 3 of the 5 issues  
✅ Understood why each is a problem  
✅ Proposed reasonable fixes  
✅ Ran tests to verify feature works  

## Related to Lab 1 Findings

Many of these issues relate to Lab 1:
- **Tight Coupling** = Issue #2 from Lab 1
- **String-Based Configuration** = Issue #3 from Lab 1 (FighterName as string)
- **Missing Validation** = Issue #4 from Lab 1

This reinforces lessons from the investigation phase.
