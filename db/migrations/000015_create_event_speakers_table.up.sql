CREATE TABLE event_speakers (
    speaker_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    job VARCHAR(100),
    company VARCHAR(100),
    avatar_url VARCHAR(255),

    FOREIGN KEY (event_id)
        REFERENCES events(event_id)
);