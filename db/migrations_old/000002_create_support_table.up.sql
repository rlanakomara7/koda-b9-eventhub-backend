CREATE TABLE IF NOT EXISTS event_categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS event_category_mappings (
    event_category_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,

    FOREIGN KEY (event_id)
        REFERENCES events(event_id),

    FOREIGN KEY (category_id)
        REFERENCES event_categories(category_id)
);

CREATE TABLE IF NOT EXISTS community_categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS community_category_mappings (
    community_category_id SERIAL PRIMARY KEY,
    community_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,

    CONSTRAINT fk_community_category_community
        FOREIGN KEY (community_id)
        REFERENCES communities(community_id),

    CONSTRAINT fk_community_category_category
        FOREIGN KEY (category_id)
        REFERENCES community_categories(category_id)
);

CREATE TABLE IF NOT EXISTS community_members (
    member_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    community_id INTEGER NOT NULL,
    joined_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_community_member_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id),

    CONSTRAINT fk_community_member_community
        FOREIGN KEY (community_id)
        REFERENCES communities(community_id)
);

CREATE TABLE IF NOT EXISTS saved_events (
    saved_event_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,

    CONSTRAINT fk_saved_event_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id),

    CONSTRAINT fk_saved_event_event
        FOREIGN KEY (event_id)
        REFERENCES events(event_id)
);

CREATE TABLE IF NOT EXISTS event_speakers (
    speaker_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    job VARCHAR(100),
    company VARCHAR(100),
    avatar_url VARCHAR(255),

    FOREIGN KEY (event_id)
        REFERENCES events(event_id)
);

CREATE TABLE IF NOT EXISTS event_discussions (
    discussion_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (event_id)
        REFERENCES events(event_id),

    FOREIGN KEY (user_id)
        REFERENCES users(user_id)
);