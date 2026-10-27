CREATE TABLE event_discussions (
    discussion_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (event_id)
        REFERENCES events(event_id),

    FOREIGN KEY (user_id)
        REFERENCES users(user_id)
);find db/migrations -type f -size 0 -print