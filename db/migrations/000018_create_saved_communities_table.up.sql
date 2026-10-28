CREATE TABLE saved_communities (
    saved_community_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    community_id INTEGER NOT NULL REFERENCES communities(community_id),
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT saved_community_unique UNIQUE (user_id, community_id)
);
