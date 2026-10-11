CREATE TABLE communities (
    community_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    banner_url VARCHAR(255),
    status community_status DEFAULT 'active'
);