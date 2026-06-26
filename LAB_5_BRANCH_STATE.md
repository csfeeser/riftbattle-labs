# Lab 5: Shipping Workflow - Equipment System Feature

## Branch Information

**Branch:** `feature/equipment-system`  
**Base:** main (at Phase 1)  
**Feature:** Complete equipment system with rarity bonuses and fighter integration  
**Status:** Production-ready - all tests pass, realistic commit history

## Commit History

This branch demonstrates a realistic feature development workflow with 5 commits:

### Commit 1: `feat(equipment): Add equipment rarity bonus system`
**What:** Foundation - data types and rarity-based bonuses  
**Files:** `equipment/rarity_bonuses.go`  
**Why:** Establishes the core concept that rarity provides stat multipliers

```bash
git log --oneline | head -1
# de59ea9 feat(equipment): Add equipment rarity bonus system
```

---

### Commit 2: `feat(equipment): Integrate equipment system with fighter stats`
**What:** Integration - equipment bonuses apply to fighters  
**Files:** `equipment/integration.go`  
**Why:** Connects equipment to fighters so bonuses actually matter

Key functions:
- `EquipItem()` - puts equipment on fighters
- `CalculateFighterStatBonus()` - calculates total bonuses from all equipment
- `IsValidSlotForType()` - ensures items go in valid slots

---

### Commit 3: `feat(equipment): Add equipment display and comparison functions`
**What:** UX - players can see and compare equipment  
**Files:** `equipment/display.go`  
**Why:** Makes equipment system user-friendly

Key functions:
- `FormatEquipmentName()` - displays items with rarity prefix
- `DescribeEquipment()` - detailed item info
- `CompareEquipment()` - helps players choose better gear
- `ListEquippedItems()` - inventory view

---

### Commit 4: `test(equipment): Add rarity bonus system tests`
**What:** Quality - comprehensive test coverage  
**Files:** `equipment/rarity_bonuses_test.go`  
**Why:** Ensures bonuses are calculated correctly

Test coverage:
- All rarity levels (common → legendary)
- Bonus calculations are accurate
- Edge case: unknown rarity returns common

---

### Commit 5: `fix(equipment): Improve EquipItem to properly handle slot conflicts`
**What:** Bug fix - prevents accidental item loss  
**Files:** `equipment/integration.go` (modified)  
**Why:** Real-world issue discovered during development

**Before:** Equipping new item just overwrites old item (it's lost)  
**After:** Returns the unequipped item so nothing is lost

---

## Expected Reviewer Experience

### What Reviewers Should See

1. **Logical progression:** Each commit builds on previous ones
2. **Small, focused commits:** Each commit does one thing well
3. **Testing at appropriate points:** Tests come after core features
4. **Bug fixes included:** Real-world issues discovered and fixed
5. **Clear commit messages:** Anyone can understand the evolution

### Questions Reviewers Might Ask

✅ "Can we approve this for production?"  
Evaluate based on:
- Are all tests passing?
- Is the commit history clean?
- Does the feature work end-to-end?
- Are there obvious bugs or design issues?

### Expected Review Time

**20-30 minutes** to review all commits:
- 5 min: Read commit messages to understand intent
- 10 min: Review code for correctness and style
- 5 min: Verify tests pass
- 5 min: Check for production readiness

## Testing the Feature

```bash
# Switch to the feature branch
git checkout feature/equipment-system

# Run all tests
go test ./...

# Verify app runs
go run .

# View the commit history
git log --oneline | head -10
```

## The Intentional Imperfection

**What:** Commit 5 shows a real bug that was discovered  
**Where:** Equipment swapping could lose items if not done carefully  
**Lesson:** Even good features have edge cases; testing and iteration matter  
**Fix:** Return the unequipped item instead of losing it  

This demonstrates:
- Bugs can happen in well-written code
- Code review and testing catch them
- Fixes should be minimal and focused
- Commit messages explain the "why"

## Success Criteria for Students

✅ Feature is production-ready  
✅ All tests pass  
✅ Commit history is clean and understandable  
✅ Code changes are focused and logical  
✅ The "imperfection" (Commit 5) shows good problem-solving  

## Shipping Checklist (For Instructors)

Before marking Lab 5 complete:
- [ ] All 5 commits are present on the branch
- [ ] `go test ./equipment` passes
- [ ] `go test ./...` passes
- [ ] `go run .` runs without error
- [ ] Commit messages follow conventional format
- [ ] Code is readable and maintainable
- [ ] Tests cover the main functionality
- [ ] Student understands the equipment system

## Related Concepts

This workflow demonstrates:
- **Atomic commits:** Each commit is complete and testable
- **Incremental development:** Build features in small steps
- **TDD principles:** Tests validate functionality (Commit 4)
- **Real-world bug fixing:** Find and fix bugs discovered during dev (Commit 5)
- **Code review preparation:** Well-structured changes are easy to review
