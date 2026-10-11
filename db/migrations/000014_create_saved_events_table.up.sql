CREATE TABLE saved_events (
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