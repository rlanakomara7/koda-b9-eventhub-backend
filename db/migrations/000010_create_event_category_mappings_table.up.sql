CREATE TABLE event_category_mappings (
    event_category_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,

    FOREIGN KEY (event_id)
        REFERENCES events(event_id),

    FOREIGN KEY (category_id)
        REFERENCES event_categories(category_id)
);