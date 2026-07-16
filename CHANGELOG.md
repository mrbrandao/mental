# Changelog

All notable changes to mental are documented here.
Format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [0.1.2] - 2026-07-16

### Added
- Add -a/-p/-s save modes to mental mem save
- Add mental.toml config with mem/session/llm defaults
- Register built-ins and add --engine flag to mem
- Add skills/mental/ with agentskills.io SKILL.md


### Changed
- Rename mem engine package to memx

## [0.1.1] - 2026-07-16

### Added
- Rename module from ais to mental
- Add config package with XDG MENTAL_DIR resolution
- Move search under session command group
- Add mem, extensions, and status command stubs
- Add Extension interface and Manager registry
- Add mem extension config with YAML-driven layout
- Add mem types and layout path resolver
- Implement mental mem init
- Implement mental mem load
- Implement mental mem save
- Implement mental mem search
- Implement mental mem task add/done/list
- Add external extension runner and XDG discovery
- Add mental SKILL.md and memory section in AGENTS.md


### Documentation
- Add mental design specification
- Rename all ais references to mental
- Add developer guide for contributing and extensions

## [0.1.0] - 2026-07-13

### Added
- Add domain model types (Session, Query)
- Add Provider interface
- Add OpenCode SQLite provider with smart/fast/deep search
- Add output formatters (table/json/plain)
- Add cobra root command
- Add cobra search command and main entry point
- Add curl install script


### Documentation
- Add AGENTS.md with project rules and contribution guide
- Add README with badges and dev guide
- Add Apache 2.0 LICENSE
- Update coverage badge to 35.0%


### Fixed
- Install golangci-lint in CI via official action


