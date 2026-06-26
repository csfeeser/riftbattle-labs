# Lab 2: Bug Hunt - Poison and Healing Interaction

## The Bug

**What's broken:** Poisoned characters can heal, ignoring their poison status.

**Expected behavior:** Poisoned characters should NOT be able to use healing potions.

**Current behavior:** Healing potions work normally even on poisoned characters.

## Investigation Guide

### Where to Look

The bug is in: **`items/healing.go`** → `ApplyHealing()` function

### Finding the Bug

1. **Run the failing test:**
   ```bash
   go test -run TestHealingPotion_PoisonedCharacterCannotHeal ./items -v
   ```
   
   Expected output: **FAIL** - HP should be 15 but was 27 after healing while poisoned

2. **Examine the test:**
   - File: `items/items_test.go` line ~42
   - The test applies poison, then tries to heal
   - Expects healing to be blocked (HP stays 15)
   - Actually gets healing (HP becomes 27)

3. **Find the root cause:**
   - Look at `ApplyHealing()` in `items/healing.go`
   - Missing: A check for the `"poisoned"` status
   - Compare with: `ApplyStatusEffect()` in `effects/status.go` to see how status checks work
   
4. **Minimal fix:**
   Add this check at the beginning of `ApplyHealing()`:
   ```go
   if _, poisoned := targetStatus["poisoned"]; poisoned {
       return 0  // Cannot heal while poisoned
   }
   ```

## Test Cases That Validate the Bug

| Test Name | Status | What It Tests |
|-----------|--------|---------------|
| TestHealingPotion_HealthyCharacterCanHeal | PASS | Healing works normally when not poisoned |
| TestHealingPotion_PoisonedCharacterCannotHeal | **FAIL** | Poison blocks healing ← THE BUG |
| TestHealingPotion_PoisonWearingOffAllowsHealing | PASS | Healing works after poison expires |
| TestHealingPotion_MultiplePoisonStillBlocksHealing | **FAIL** | Multiple poison stacks still block healing |

## Expected Investigation Time

**15-20 minutes** with guided constraints:
- ~2 minutes: Run tests and understand failure
- ~3-5 minutes: Find the test and understand what it expects
- ~5-8 minutes: Locate the bug in `ApplyHealing()`
- ~5 minutes: Implement and verify the fix

## Debugging Strategy

1. **Understand the symptom:**
   - Test expects HP=15 (no healing)
   - But gets HP=27 (12 HP healed)
   - This tells us healing happened when it shouldn't

2. **Narrow the location:**
   - `UseHealingPotion()` calls `ApplyHealing()`
   - The poison check should be in `ApplyHealing()`
   - Look for where poison status is checked elsewhere

3. **Make the fix:**
   - Add the poison check
   - Run the tests again - should all pass

4. **Verify the fix:**
   ```bash
   go test -run Poisoned ./items -v
   ```
   All poison-related tests should **PASS**

## Related Code

**How to check poison status elsewhere:**
- See `effects/status.go` → `ApplyStatusEffect()` for how status effects work
- See `ApplyStatusEffect()` → usage of `status` map with string keys

**How other restrictions work:**
- See `RestoreMana()` in same file for pattern of checking prerequisites

## Success Criteria

✅ All poison-related tests pass  
✅ Fix is minimal (3-5 lines)  
✅ No other tests break  
✅ Healing still works for non-poisoned characters  

## Discussion Questions

1. Why is the check at the beginning of the function important?
2. What if poison has multiple stacks? (Answer: Map keys are unique, so only one entry per status)
3. Should we also check other restrictive effects (e.g., silenced)? (Design decision)
4. Could this be caught by type checking instead of strings? (Yes - see Lab 1)

## Related Issues

This bug is an example of issues found in **Lab 1, Issue #3 (String-Based Configuration)** - Magic strings for status effects are error-prone.

The proper fix would be to use constants or enums instead of strings, which would catch this at compile time rather than requiring runtime tests.
