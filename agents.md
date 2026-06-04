# AGENTS.md - System Architecture Rules
# Role: System Architect Guidance by Director Goff

## Workspace Rules
- This is a monorepo setup using unified workspaces.
- Package boundaries must be respected: do not cross-import local packages without updating package.json configurations.
- Data Layer: SQLite is used as our dynamic information framework. Ensure schemas are tightly type-checked.

## Version Control Standards
- Every successful repair or modification must compile with zero linter or type-mismatches.
- Keep commits highly modular, descriptive, and localized to the modified package.
