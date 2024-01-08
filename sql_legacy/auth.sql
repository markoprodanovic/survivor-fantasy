CREATE TABLE users (
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT,
    email TEXT NOT NULL,
    emailVerified INTEGER, -- Assuming timestamp in milliseconds
    image TEXT
);

CREATE TABLE accounts (
    userId TEXT NOT NULL,
    type TEXT NOT NULL,
    provider TEXT NOT NULL,
    providerAccountId TEXT NOT NULL,
    refresh_token TEXT,
    access_token TEXT,
    expires_at INTEGER,
    token_type TEXT,
    scope TEXT,
    id_token TEXT,
    session_state TEXT,
    PRIMARY KEY (provider, providerAccountId),
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE sessions (
    sessionToken TEXT NOT NULL PRIMARY KEY,
    userId TEXT NOT NULL,
    expires INTEGER NOT NULL, -- Assuming timestamp in milliseconds
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE verificationTokens (
    identifier TEXT NOT NULL,
    token TEXT NOT NULL,
    expires INTEGER NOT NULL, -- Assuming timestamp in milliseconds
    PRIMARY KEY (identifier, token)
);
