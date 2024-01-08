CREATE TABLE episodes (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    episode_number INTEGER NOT NULL,
    episode_date DATE NOT NULL
);


CREATE TABLE episode_points (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    episode_id INTEGER NOT NULL,
    cast_id INTEGER NOT NULL,
    points INTEGER NOT NULL,
    FOREIGN KEY (episode_id) REFERENCES episodes(id),
    FOREIGN KEY (cast_id) REFERENCES players(id)
);

CREATE TABLE players (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INTEGER NOT NULL,
    tribe_id INTEGER,
    eliminated BOOLEAN DEFAULT false NOT NULL,
    FOREIGN KEY (tribe_id) REFERENCES tribes(id)
);

CREATE TABLE tribes (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    colour TEXT NOT NULL
);

CREATE TABLE user_picks (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT,
    is_admin BOOLEAN DEFAULT false
);