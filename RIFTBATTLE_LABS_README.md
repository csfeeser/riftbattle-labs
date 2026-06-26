# RiftBattle: Complete Lab Preparation

This repository contains the RiftBattle codebase prepared for 5 comprehensive Claude Code labs teaching investigation, debugging, code review, and shipping workflows.

## Quick Status

✅ **Main branch:** 2,421 lines of code | 15+ tests | All passing  
✅ **Lab 1:** Code investigation ready  
✅ **Lab 2:** Bug injected and documented  
✅ **Lab 3:** Main branch ready to use as-is  
✅ **Lab 4:** Feature branch with code review issues  
✅ **Lab 5:** Feature branch with realistic commit history  

## Branches

### `main`
The production-ready expanded codebase.
- 8 packages: combat, effects, equipment, fighters, game, items, logging, modes, utils
- 2,421 lines of code (3,302 with tests)
- 15+ passing tests
- Intentional fragility for Lab 1 investigation

**Use for:** Labs 1 & 3

### `lab-2-with-bug`
The main branch with a deliberate bug injected: poisoned characters can heal (they shouldn't).

**Status:** 2 specific tests fail (TestHealingPotion_PoisonedCharacterCannotHeal, TestHealingPotion_MultiplePoisonStillBlocksHealing)  
**Use for:** Lab 2 (bug investigation and fixing)

### `feature/stun-system`
A new feature branch implementing an expanded stun system with 5 intentional code review issues embedded.

**Status:** All tests pass, but code has issues to discover  
**Issues to find:** Architecture coupling, missing validation, incomplete tests, no logging, inefficient algorithms  
**Use for:** Lab 4 (code review)

### `feature/equipment-system`
A complete equipment system feature with 5 realistic commits showing professional development workflow.

**Status:** All tests pass, production-ready  
**Commits:** Clear progression (data types → integration → UI → tests → bug fix)  
**Use for:** Lab 5 (shipping workflow and commit history review)

## Lab Structure

### Lab 1: Codebase Investigation
**Branch:** main  
**Task:** Identify 7 architectural weaknesses  
**Duration:** 40-60 minutes  
**Documentation:** `LAB_1_INVESTIGATION_NOTES.md`

Key findings to discover:
1. Circular dependencies
2. Tight coupling
3. String-based configuration  
4. Missing validation
5. Dead code
6. Naming inconsistencies
7. Inefficient algorithms (optional)

### Lab 2: Bug Hunt & Fixing
**Branch:** lab-2-with-bug  
**Task:** Discover and fix the poison-healing bug  
**Duration:** 30-40 minutes  
**Documentation:** `LAB_2_BUG_SPECIFICATION.md`

The bug: Poisoned characters can heal (they shouldn't)  
Fix: Add poison check to ApplyHealing() (3 lines)

### Lab 3: Code Assessment
**Branch:** main  
**Task:** Evaluate codebase quality using Lab 1 findings  
**Duration:** 20-30 minutes  
**Documentation:** Built into Lab 1 notes

No special preparation needed - main branch is ready.

### Lab 4: Code Review
**Branch:** feature/stun-system  
**Task:** Review a feature branch and identify 5 issues  
**Duration:** 45-60 minutes  
**Documentation:** `LAB_4_PREPARED_DIFF.md`

Issues by severity:
- MAJOR: Architecture coupling (Issue 1), Security validation (Issue 2)
- MEDIUM: Missing tests (Issue 3), No logging (Issue 4), Performance (Issue 5)

### Lab 5: Shipping Workflow
**Branch:** feature/equipment-system  
**Task:** Review commit history and evaluate production readiness  
**Duration:** 40-50 minutes  
**Documentation:** `LAB_5_BRANCH_STATE.md`

Content:
- 5 sequential commits showing feature development
- Clear commit messages
- Test additions at appropriate time
- Bug discovery and fix (Commit 5)

## Getting Started

### Prerequisites
- Go 1.22+
- Git
- Claude Code CLI

### Setup

```bash
# Clone the repository
git clone https://github.com/csfeeser/riftbattle.git
cd riftbattle

# Verify main branch works
git checkout main
go test ./...
go run .
```

### Running Each Lab

**Lab 1:**
```bash
git checkout main
# Follow LAB_1_INVESTIGATION_NOTES.md
```

**Lab 2:**
```bash
git checkout lab-2-with-bug
go test -run Poisoned ./items  # See failures
# Follow LAB_2_BUG_SPECIFICATION.md
```

**Lab 3:**
```bash
git checkout main
# Use findings from Lab 1
```

**Lab 4:**
```bash
git checkout feature/stun-system
# Review code using LAB_4_PREPARED_DIFF.md
go test ./...
```

**Lab 5:**
```bash
git checkout feature/equipment-system
git log --oneline | head -10
# Review commits using LAB_5_BRANCH_STATE.md
go test ./...
```

## File Structure

```
riftbattle/
├── main.go                        # Entry point
├── go.mod
├── LAB_1_INVESTIGATION_NOTES.md   # Lab 1 guide
├── LAB_2_BUG_SPECIFICATION.md     # Lab 2 guide
├── LAB_4_PREPARED_DIFF.md         # Lab 4 guide
├── LAB_5_BRANCH_STATE.md          # Lab 5 guide
├── RIFTBATTLE_LABS_README.md      # This file
│
├── combat/                        # Combat system
├── effects/                       # Status effects and stun system
├── equipment/                     # Equipment and rarity system
├── fighters/                      # Character progression and stats
├── game/                          # Game state management
├── items/                         # Consumables and inventory
├── logging/                       # Combat logging
├── modes/                         # Game modes
└── utils/                         # Utilities
```

## Codebase Metrics

| Metric | Value |
|--------|-------|
| Total Lines of Code | 2,421 |
| Total Lines (w/ tests) | 3,302 |
| Packages | 8 |
| Test Functions | 15+ |
| Test Pass Rate | 100% (main), Expected failures (lab-2) |
| Code Complexity | Moderate (intentional fragility) |
| Documentation | Comprehensive |

## Key Learning Outcomes

After all 5 labs, students will understand:

✅ How to investigate and assess codebases  
✅ How to debug and fix issues systematically  
✅ How to review code and identify design problems  
✅ How to evaluate shipping readiness  
✅ Professional development workflows with Git  
✅ Test-driven design and development  

## Instructor Notes

### Lab Timing
- Lab 1: 40-60 min (can vary based on depth)
- Lab 2: 30-40 min
- Lab 3: 20-30 min (shorter, builds on Lab 1)
- Lab 4: 45-60 min
- Lab 5: 40-50 min

**Total:** ~3-4 hours of content across 5 labs

### Customization
Each lab document includes:
- Difficulty adjustments
- Optional challenges
- Discussion questions
- Extension ideas

Feel free to adapt timing, scope, or focus areas for your audience.

### Success Criteria
See individual LAB_*_*.md files for specific success criteria for each lab.

## Troubleshooting

**Tests failing on main:** Run `go test ./...` - should see 0 failures  
**App doesn't run:** Verify Go 1.22+ and run `go mod tidy`  
**Lab 2 expectations:** lab-2-with-bug branch should have 2 failing tests (poison healing)  
**Lab 4 branch:** feature/stun-system should have all tests passing but contain code issues  
**Lab 5 branch:** feature/equipment-system should show 5 commits with `git log`

## Support

For detailed guidance on each lab, see:
- **Lab 1:** `LAB_1_INVESTIGATION_NOTES.md`
- **Lab 2:** `LAB_2_BUG_SPECIFICATION.md`
- **Lab 4:** `LAB_4_PREPARED_DIFF.md`
- **Lab 5:** `LAB_5_BRANCH_STATE.md`

Each document includes investigation strategies, expected findings, and discussion questions.

---

**Prepared:** June 26, 2026  
**Status:** Complete and ready for classroom use  
**Version:** 1.0
