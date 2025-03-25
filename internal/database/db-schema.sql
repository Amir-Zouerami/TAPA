CREATE TABLE IF NOT EXISTS collections (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    name                        TEXT NOT NULL,
    description                 TEXT,
    position                    INTEGER NOT NULL,
    created_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at                  DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_collections_position ON collections(position);

CREATE TABLE IF NOT EXISTS folders (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    collection_id               INTEGER NOT NULL,
    name                        TEXT NOT NULL,
    position                    INTEGER NOT NULL,
    created_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS collection_variables (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    collection_id               INTEGER NOT NULL,
    key                         TEXT NOT NULL,
    value                       TEXT NOT NULL,
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
    UNIQUE (collection_id, key)
);

CREATE TABLE IF NOT EXISTS requests (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    collection_id               INTEGER, -- Nullable to allow standalone requests
    folder_id                   INTEGER,
    position                    INTEGER NOT NULL,
    name                        TEXT NOT NULL,
    method                      TEXT CHECK(method IN ('GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS', 'HEAD')),
    url                         TEXT NOT NULL,
    body                        TEXT,
    body_format                 TEXT CHECK(body_format IN ('JSON', 'XML', 'form-data', 'raw')) DEFAULT 'JSON',
    notes                       TEXT,
    timeout                     INTEGER DEFAULT 30000,
    allow_redirects             BOOLEAN DEFAULT TRUE,
    ssl_verification            BOOLEAN DEFAULT TRUE,
    remove_referer_on_redirect  BOOLEAN DEFAULT FALSE,
    encode_url                  BOOLEAN DEFAULT TRUE,
    created_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
    FOREIGN KEY (folder_id)     REFERENCES folders(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_headers (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id                  INTEGER NOT NULL,
    key                         TEXT NOT NULL,
    value                       TEXT,
    FOREIGN KEY (request_id)    REFERENCES requests(id) ON DELETE CASCADE,
    UNIQUE (request_id, key)
);

CREATE TABLE IF NOT EXISTS request_query_params (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id                  INTEGER NOT NULL,
    key                         TEXT NOT NULL,
    value                       TEXT,
    FOREIGN KEY (request_id)    REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_cookies (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id                  INTEGER NOT NULL,
    key                         TEXT NOT NULL,
    value                       TEXT,
    FOREIGN KEY (request_id)    REFERENCES requests(id) ON DELETE CASCADE,
    UNIQUE (request_id, key)
);

CREATE TABLE IF NOT EXISTS environments (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    name                        TEXT NOT NULL,
    created_at                  DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS environment_variables (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    environment_id              INTEGER NOT NULL,
    key                         TEXT NOT NULL,
    value                       TEXT NOT NULL,
    FOREIGN KEY (environment_id) REFERENCES environments(id) ON DELETE CASCADE,
    UNIQUE (environment_id, key)
);

CREATE TABLE IF NOT EXISTS request_history (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id                  INTEGER NOT NULL,
    timestamp                   DATETIME DEFAULT CURRENT_TIMESTAMP,
    method                      TEXT NOT NULL,
    url                         TEXT NOT NULL,
    headers                     TEXT,
    query_params                TEXT,
    body                        TEXT,
    status_code                 INTEGER,
    response_time               INTEGER,
    data_volume                 INTEGER,
    FOREIGN KEY (request_id)    REFERENCES requests(id) ON DELETE CASCADE
);

-- Trigger to enforce a maximum of 300 history records
CREATE TRIGGER IF NOT EXISTS limit_request_history
AFTER INSERT ON request_history
BEGIN
    DELETE FROM request_history
    WHERE id IN (
        SELECT id FROM request_history
        ORDER BY timestamp ASC
        LIMIT (SELECT COUNT(*) - 300 FROM request_history)
    );
END;

CREATE TABLE IF NOT EXISTS request_examples (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id                  INTEGER NOT NULL,
    timestamp                   DATETIME DEFAULT CURRENT_TIMESTAMP,
    method                      TEXT NOT NULL,
    url                         TEXT NOT NULL,
    headers                     TEXT,
    query_params                TEXT,
    body                        TEXT,
    status_code                 INTEGER,
    response                    TEXT,
    response_headers            TEXT,
    response_cookies            TEXT,
    response_time               INTEGER,
    data_volume                 INTEGER,
    FOREIGN KEY (request_id)    REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_scripts (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id                  INTEGER NOT NULL,
    script                      TEXT NOT NULL,
    FOREIGN KEY (request_id)    REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS test_results (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    request_id                  INTEGER NOT NULL,
    test_name                   TEXT NOT NULL,
    result                      TEXT,
    created_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (request_id)    REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sync_metadata (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    entity_type                 TEXT CHECK(entity_type IN ('requests', 'collections', 'variables', 'history')),
    entity_id                   INTEGER NOT NULL,
    last_updated                DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_dirty                    BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS keyboard_shortcuts (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT,
    action                      TEXT NOT NULL UNIQUE,
    shortcut                    TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_settings (
    id                          INTEGER PRIMARY KEY AUTOINCREMENT DEFAULT 1,
    theme                       TEXT CHECK(theme IN ('light', 'dark', 'system')) DEFAULT 'dark',
    max_history                 INTEGER DEFAULT 200,
    language                    TEXT DEFAULT 'en',
    font_family                 TEXT DEFAULT 'default',
    font_size                   INTEGER DEFAULT 14
);

CREATE TABLE IF NOT EXISTS app_state (
    id                          INTEGER PRIMARY KEY DEFAULT 1,
    selected_environment        INTEGER,
    open_tabs                   TEXT,
    updated_at                  DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (selected_environment) REFERENCES environments(id) ON DELETE SET NULL
);