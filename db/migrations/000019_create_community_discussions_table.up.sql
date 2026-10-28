CREATE TABLE community_discussions (
    discussion_id SERIAL PRIMARY KEY,
    community_id INTEGER NOT NULL REFERENCES communities(community_id),
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT community_discussion_message_check
        CHECK (char_length(btrim(message)) BETWEEN 1 AND 1000)
);

CREATE INDEX idx_community_discussions_list
ON community_discussions (community_id, created_at DESC, discussion_id DESC);
