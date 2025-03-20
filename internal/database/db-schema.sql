-- This file initializes the database structure.
-- I don't like to use migration on small projects like this.

CREATE TABLE IF NOT EXISTS collections (
    id            TEXT PRIMARY KEY,   
    name          TEXT NOT NULL,      
    description   TEXT,               
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS folders (
    id            TEXT PRIMARY KEY,   
    collection_id TEXT NOT NULL,      
    name          TEXT NOT NULL,      
    position      INTEGER,            
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS requests (
    id            TEXT PRIMARY KEY,  
    collection_id TEXT,              
    folder_id     TEXT,              
    name          TEXT NOT NULL,      
    method        TEXT CHECK(method IN ('GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'OPTIONS', 'HEAD')),  
    url           TEXT NOT NULL,      
    body          TEXT,               
    notes         TEXT,               
    timeout       INTEGER DEFAULT 30000,  
    allow_redirects BOOLEAN DEFAULT TRUE, 
    ssl_verification BOOLEAN DEFAULT TRUE, 
    remove_referer_on_redirect BOOLEAN DEFAULT FALSE, 
    encode_url BOOLEAN DEFAULT TRUE, 
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
    FOREIGN KEY (folder_id) REFERENCES folders(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_headers (
    id          TEXT PRIMARY KEY,
    request_id  TEXT NOT NULL,
    key         TEXT NOT NULL,
    value       TEXT,
    FOREIGN KEY (request_id) REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_query_params (
    id          TEXT PRIMARY KEY,
    request_id  TEXT NOT NULL,
    key         TEXT NOT NULL,
    value       TEXT,
    FOREIGN KEY (request_id) REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_auth (
    id          TEXT PRIMARY KEY,
    request_id  TEXT NOT NULL,
    auth_type   TEXT CHECK(auth_type IN ('Bearer', 'Basic', 'API Key', 'OAuth2', 'JWT')),
    token       TEXT,   
    username    TEXT,   
    password    TEXT,   
    FOREIGN KEY (request_id) REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS environments (
    id          TEXT PRIMARY KEY,  
    name        TEXT NOT NULL,     
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS environment_variables (
    id            TEXT PRIMARY KEY,
    environment_id TEXT NOT NULL,
    key           TEXT NOT NULL,
    value         TEXT NOT NULL,
    FOREIGN KEY (environment_id) REFERENCES environments(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS collection_variables (
    id            TEXT PRIMARY KEY,
    collection_id TEXT NOT NULL,
    key           TEXT NOT NULL,
    value         TEXT NOT NULL,
    FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_history (
    id            TEXT PRIMARY KEY,
    request_id    TEXT NOT NULL,
    timestamp     DATETIME DEFAULT CURRENT_TIMESTAMP,
    history_type  TEXT CHECK(history_type IN ('auto', 'example')) DEFAULT 'auto',  
    method        TEXT NOT NULL,
    url           TEXT NOT NULL,
    headers       TEXT,  
    query_params  TEXT,  
    auth          TEXT,  
    body          TEXT,  
    status_code   INTEGER,  
    response      TEXT,  
    response_headers TEXT,  
    FOREIGN KEY (request_id) REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS request_scripts (
    id            TEXT PRIMARY KEY,
    request_id    TEXT NOT NULL,
    script_type   TEXT CHECK(script_type IN ('pre-request', 'test')), 
    script        TEXT NOT NULL, 
    FOREIGN KEY (request_id) REFERENCES requests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sync_metadata (
    id           TEXT PRIMARY KEY,
    entity_type  TEXT CHECK(entity_type IN ('requests', 'collections', 'variables', 'history')),
    entity_id    TEXT NOT NULL,
    last_updated DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_dirty     BOOLEAN DEFAULT FALSE 
);

CREATE TABLE IF NOT EXISTS keyboard_shortcuts (
    id          TEXT PRIMARY KEY,
    action      TEXT NOT NULL UNIQUE,  
    shortcut    TEXT NOT NULL          
);

CREATE TABLE IF NOT EXISTS user_settings (
    id            TEXT PRIMARY KEY DEFAULT 'singleton',  
    theme         TEXT CHECK(theme IN ('light', 'dark', 'system')) DEFAULT 'system',
    max_history   INTEGER DEFAULT 200,  
    language      TEXT DEFAULT 'en',
    font_family   TEXT DEFAULT 'default',
    font_size     INTEGER DEFAULT 14,
    working_dir   TEXT DEFAULT NULL  
);
