PRAGMA foreign_keys = ON;

-- =========================
-- USERS
-- =========================
CREATE TABLE users (
    telegram_id TEXT PRIMARY KEY,
    lichess_username TEXT,
    is_whitelisted INTEGER DEFAULT 0,   -- 0 = false, 1 = true
    created_at TEXT,
    updated_at TEXT
);


-- =========================
-- ORGANIZER PROFILES
-- Only whitelisted users should be inserted here (enforced in app layer)
-- =========================
CREATE TABLE organizers (
    telegram_id TEXT PRIMARY KEY,
    name TEXT,
    logo_url TEXT,
    created_at TEXT,
    FOREIGN KEY (telegram_id) REFERENCES users(telegram_id)
);


CREATE TABLE presidents (
    telegram_id TEXT PRIMARY KEY,
    FOREIGN KEY (telegram_id) REFERENCES users(telegram_id)
);
-- =========================
-- TOURNAMENT DRAFTS
-- =========================
CREATE TABLE tournament_drafts (
    draft_uuid TEXT PRIMARY KEY,

    tournament_class TEXT,
    requested_by TEXT,

    name TEXT,
    format TEXT,
    time_control TEXT,
    start_datetime TEXT,
    end_datetime TEXT,
    timezone TEXT,
    capacity INTEGER,
    visibility TEXT,
    prizes TEXT,

    organizer_telegram_id TEXT,

    rationale TEXT,
    status TEXT,
    created_at TEXT,
    updated_at TEXT,

    FOREIGN KEY (requested_by) REFERENCES users(telegram_id),
    FOREIGN KEY (organizer_telegram_id) REFERENCES organizers(telegram_id)
);


-- =========================
-- AUDIT LOG
-- =========================
CREATE TABLE audit_log (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    entity_type TEXT,
    entity_id TEXT,
    action TEXT,
    performed_by TEXT,
    created_at TEXT,
    FOREIGN KEY (performed_by) REFERENCES users(telegram_id)
);
