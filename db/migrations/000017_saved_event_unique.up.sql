ALTER TABLE saved_events
ADD CONSTRAINT saved_event_unique
UNIQUE (user_id, event_id);