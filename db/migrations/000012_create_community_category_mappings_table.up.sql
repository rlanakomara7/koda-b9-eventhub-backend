CREATE TABLE community_category_mappings (
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