# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- General folder structure.
- Database initialization mechanism on launch.
- New errors for collections repository layer.
- Database models based on the schema.
- Collection repository and dashboard service.

### Changed

- Database schema was revised and updated (seeding logic updated as well).
- Use a separate sqlite db for development mode instead of flushing & rewriting the main database.
- Sqlite db now runs in WAL mode for better concurrency.
- Use `sqlx` instead of the standard library `sql` for ease of use (struct scanning, etc.).
