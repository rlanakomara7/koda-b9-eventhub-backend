CREATE TABLE IF NOT EXISTS event_discussions (
    discussion_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events(event_id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT check_event_discussion_message
    CHECK (CHAR_LENGTH(TRIM(message)) BETWEEN 1 AND 1000)
);

CREATE INDEX IF NOT EXISTS idx_event_discussions_event_id
ON event_discussions(event_id, created_at);
