CREATE TABLE events (
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