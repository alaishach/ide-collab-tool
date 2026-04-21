create table if not exists users (
  id serial primary key,
  username text unique not null,
  email text unique not null,
  password BYTEA not null,
  creation date default current_date not null,
  last_connection timestamptz default now() not null
);

create type session_source_type as ENUM (
  'Browser',
  'Mobile App',
  'Desktop App'
);

create table if not exists session_source (
  id serial primary key,
  source_type session_source_type not null,
  source_name text not null
);

create table if not exists sessions (
  id serial primary key,
  user_id integer references users (id) on delete cascade not null,
  session_token uuid unique not null
  -- session_source_id integer references session_source (id) on delete cascade not null
);

create table if not exists projects (
  id serial primary key,
  name text not null,
  admin_id integer references users (id) on delete cascade not null,
  created_at timestamptz default now() not null,
  description text
);

create table if not exists project_members (
  project_id integer references projects (id) on delete cascade not null,
  member_id integer references users (id) on delete cascade not null,
  joined_at timestamptz default now() not null,
  primary key (project_id, member_id)
);

create table if not exists files (
  id serial primary key,
  project_id integer references projects (id) on delete cascade not null,
  file_path text not null, -- relative path from the server/assets/projectName dir
  created_at timestamptz default now() not null,
  updated_at timestamptz default now() not null
);
