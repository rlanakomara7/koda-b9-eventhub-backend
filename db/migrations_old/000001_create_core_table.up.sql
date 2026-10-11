CREATE TYPE user_status AS ENUM (
    'active',
    'suspended'
);

CREATE TYPE community_status AS ENUM (
    'active',
    'non_active'
);

CREATE TYPE event_format AS ENUM (
    'in_person',
    'online'
);

CREATE TYPE event_status AS ENUM (
    'active',
    'full',
    'completed'
);


CREATE TABLE IF NOT EXISTS roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL UNIQUE
);


CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    role_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255),
    job VARCHAR(255),
    location VARCHAR(150),
    bio TEXT,
    status user_status DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
        REFERENCES roles(role_id)
);


CREATE TABLE IF NOT EXISTS communities (
    community_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    banner_url VARCHAR(255),
    status community_status DEFAULT 'active'
);


CREATE TABLE IF NOT EXISTS events (
    event_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    community_id INTEGER,
    title VARCHAR(150) NOT NULL,
    description TEXT,
    banner_url VARCHAR(255),
    event_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    format event_format NOT NULL,
    location VARCHAR(255),
    capacity INTEGER NOT NULL,
    status event_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_event_creator
        FOREIGN KEY (user_id)
        REFERENCES users(user_id),

    CONSTRAINT fk_event_community
        FOREIGN KEY (community_id)
        REFERENCES communities(community_id)
);


CREATE TABLE IF NOT EXISTS event_members (
    event_member_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    joined_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_event_member_event
        FOREIGN KEY (event_id)
        REFERENCES events(event_id),

    CONSTRAINT fk_event_member_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
);


CREATE TABLE IF NOT EXISTS password_resets (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT unique_reset_token
        UNIQUE (token),

    CONSTRAINT fk_user_reset
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
);


CREATE TABLE IF NOT EXISTS token_blacklist (
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);