CREATE TABLE IF NOT EXISTS "user" (
    id             SERIAL PRIMARY KEY,
    username       TEXT UNIQUE NOT NULL,
    email          TEXT UNIQUE NOT NULL,
    password       BYTEA NOT NULL,
    creation       DATE DEFAULT CURRENT_DATE NOT NULL,
    last_connection TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS "session" (
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER REFERENCES "user"(id) ON DELETE CASCADE NOT NULL,
    session_token UUID UNIQUE NOT NULL,
    device_token  TEXT NOT NULL, -- Removed UNIQUE (see notes below)
    created_at    TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS project (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    admin_id    INTEGER REFERENCES "user"(id) ON DELETE CASCADE NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS project_members (
    project_id  INTEGER REFERENCES project(id) ON DELETE CASCADE NOT NULL,
    member_id   INTEGER REFERENCES "user"(id) ON DELETE CASCADE NOT NULL,
    joined_at   TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    PRIMARY KEY (project_id, member_id)
);

CREATE TABLE IF NOT EXISTS "file" (
    id          SERIAL PRIMARY KEY,
    project_id  INTEGER REFERENCES project(id) ON DELETE CASCADE NOT NULL,
    file_path   TEXT NOT NULL, -- relative path from the server/assets/projectName dir
    created_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at  TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
