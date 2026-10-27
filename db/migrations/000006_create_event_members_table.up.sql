CREATE TABLE event_members (
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