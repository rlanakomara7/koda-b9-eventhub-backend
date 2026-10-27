CREATE TABLE community_members (
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