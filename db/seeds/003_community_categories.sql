INSERT INTO community_categories (category_name)
VALUES
    ('Technology'),
    ('Programming'),
    ('AI'),
    ('Design'),
    ('Business'),
    ('Career')
ON CONFLICT (category_name) DO NOTHING;