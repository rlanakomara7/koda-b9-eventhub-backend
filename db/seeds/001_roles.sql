INSERT INTO roles (role_name)
VALUES
    ('attendee'),
    ('organizer'),
    ('admin')
ON CONFLICT (role_name) DO NOTHING;