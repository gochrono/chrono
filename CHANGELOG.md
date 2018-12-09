# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Added coloring to notes show command
- Updated help documentation for `start` and `notes` command
- Added `--at` flag to `stop` command to specify a time other than now for the new frame's end time.
- Added `projects` command which shows a unique list of all project names used
- Added `restart` command which starts a new frame, using the last frame's project & tags
- Added `--version` flag which works the same as the `version` command
- Added `delete` command which deletes a saved frame by either an index or UUID
- Added `cancel` command which stops tracking the current frame without saving it
- Added `frames` command which shows a list of frame ID's

### Changed
- Renamed `--start` flag to `--at` in `start` command

### Fixed
- Edge case where frames adjusted through the --round flag had negative times

## [1.0.0-beta] - 2018-10-11
### Added
- Core commands implemented
- Data storage format stabilized

[Unreleased]: https://github.com/gochrono/chrono/compare/v1.0.0-beta...HEAD

