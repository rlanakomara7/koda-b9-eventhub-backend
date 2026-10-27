
-- DUMMY COMMUNITIES

INSERT INTO communities (
    name,
    description,
    banner_url,
    status
)
VALUES
(
    'Bandung Go Community',
    'Komunitas developer Go untuk belajar backend, microservices, concurrency, dan software engineering.',
    'https://images.unsplash.com/photo-1517245386807-bb43f82c33c4',
    'active'
),
(
    'AI Indonesia Community',
    'Komunitas teknologi yang membahas Artificial Intelligence, Machine Learning, dan perkembangan teknologi AI.',
    'https://images.unsplash.com/photo-1677442136019-21780ecad995',
    'active'
),
(
    'UI UX Designer Hub',
    'Komunitas untuk UI UX Designer yang ingin berbagi pengalaman, portfolio, design system, dan product design.',
    'https://images.unsplash.com/photo-1558655146-d09347e92766',
    'active'
),
(
    'Digital Business Network',
    'Komunitas entrepreneur dan profesional untuk berdiskusi mengenai bisnis digital, startup, dan strategi pemasaran.',
    'https://images.unsplash.com/photo-1556761175-b413da4baf72',
    'active'
),
(
    'Career Development Indonesia',
    'Komunitas pengembangan karier untuk mahasiswa, fresh graduate, dan profesional.',
    'https://images.unsplash.com/photo-1521737711867-e3b97375f902',
    'active'
)
ON CONFLICT DO NOTHING;

-- COMMUNITY CATEGORY MAPPING

INSERT INTO community_category_mappings (
    community_id,
    category_id
)
SELECT
    c.community_id,
    cc.category_id
FROM communities c, community_categories cc
WHERE c.name = 'Bandung Go Community'
AND cc.category_name = 'Programming'
AND NOT EXISTS (
    SELECT 1
    FROM community_category_mappings m
    WHERE m.community_id = c.community_id
    AND m.category_id = cc.category_id
);


INSERT INTO community_category_mappings (
    community_id,
    category_id
)
SELECT
    c.community_id,
    cc.category_id
FROM communities c, community_categories cc
WHERE c.name = 'Bandung Go Community'
AND cc.category_name = 'Technology'
AND NOT EXISTS (
    SELECT 1
    FROM community_category_mappings m
    WHERE m.community_id = c.community_id
    AND m.category_id = cc.category_id
);


INSERT INTO community_category_mappings (
    community_id,
    category_id
)
SELECT
    c.community_id,
    cc.category_id
FROM communities c, community_categories cc
WHERE c.name = 'AI Indonesia Community'
AND cc.category_name = 'AI'
AND NOT EXISTS (
    SELECT 1
    FROM community_category_mappings m
    WHERE m.community_id = c.community_id
    AND m.category_id = cc.category_id
);


INSERT INTO community_category_mappings (
    community_id,
    category_id
)
SELECT
    c.community_id,
    cc.category_id
FROM communities c, community_categories cc
WHERE c.name = 'UI UX Designer Hub'
AND cc.category_name = 'Design'
AND NOT EXISTS (
    SELECT 1
    FROM community_category_mappings m
    WHERE m.community_id = c.community_id
    AND m.category_id = cc.category_id
);


INSERT INTO community_category_mappings (
    community_id,
    category_id
)
SELECT
    c.community_id,
    cc.category_id
FROM communities c, community_categories cc
WHERE c.name = 'Digital Business Network'
AND cc.category_name = 'Business'
AND NOT EXISTS (
    SELECT 1
    FROM community_category_mappings m
    WHERE m.community_id = c.community_id
    AND m.category_id = cc.category_id
);


INSERT INTO community_category_mappings (
    community_id,
    category_id
)
SELECT
    c.community_id,
    cc.category_id
FROM communities c, community_categories cc
WHERE c.name = 'Career Development Indonesia'
AND cc.category_name = 'Career'
AND NOT EXISTS (
    SELECT 1
    FROM community_category_mappings m
    WHERE m.community_id = c.community_id
    AND m.category_id = cc.category_id
);


-- =========================================
-- DUMMY EVENTS
-- organizer menggunakan user pertama
-- =========================================

INSERT INTO events (
    user_id,
    community_id,
    title,
    description,
    banner_url,
    event_date,
    start_time,
    end_time,
    format,
    location,
    capacity,
    status
)
VALUES

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'Bandung Go Community' LIMIT 1),
    'Go Concurrency Workshop',
    'Workshop mendalam mengenai goroutine, channel, worker pool, context, dan concurrency pattern menggunakan Go.',
    'https://images.unsplash.com/photo-1515879218367-8466d910aaa4',
    '2026-10-15',
    '09:00',
    '16:00',
    'in_person',
    'Bandung Creative Hub',
    100,
    'active'
),

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'Bandung Go Community' LIMIT 1),
    'Building REST API with Go',
    'Belajar membangun REST API menggunakan Go, Gin, PostgreSQL, JWT Authentication, dan clean architecture.',
    'https://images.unsplash.com/photo-1461749280684-dccba630e2f6',
    '2026-10-20',
    '13:00',
    '17:00',
    'online',
    'Zoom Meeting',
    150,
    'active'
),

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'AI Indonesia Community' LIMIT 1),
    'Generative AI for Beginners',
    'Pengenalan Generative AI, Large Language Model, prompt engineering, dan implementasinya di dunia kerja.',
    'https://images.unsplash.com/photo-1620712943543-bcc4688e7485',
    '2026-10-25',
    '09:00',
    '12:00',
    'online',
    'Google Meet',
    200,
    'active'
),

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'AI Indonesia Community' LIMIT 1),
    'Machine Learning Bootcamp',
    'Bootcamp dasar Machine Learning meliputi data preparation, supervised learning, evaluation, dan deployment.',
    'https://images.unsplash.com/photo-1555949963-aa79dcee981c',
    '2026-11-01',
    '08:00',
    '16:00',
    'in_person',
    'Jakarta Convention Center',
    120,
    'active'
),

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'UI UX Designer Hub' LIMIT 1),
    'UI UX Portfolio Review',
    'Sesi review portfolio UI UX bersama praktisi product design dan diskusi cara membangun portfolio profesional.',
    'https://images.unsplash.com/photo-1561070791-2526d30994b5',
    '2026-11-05',
    '13:00',
    '16:00',
    'online',
    'Zoom Meeting',
    80,
    'active'
),

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'Digital Business Network' LIMIT 1),
    'Digital Business Strategy 2026',
    'Seminar mengenai strategi bisnis digital, customer acquisition, branding, dan pertumbuhan startup.',
    'https://images.unsplash.com/photo-1552664730-d307ca884978',
    '2026-11-10',
    '09:00',
    '15:00',
    'in_person',
    'Jakarta Business Center',
    150,
    'active'
),

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'Career Development Indonesia' LIMIT 1),
    'Career Preparation for Fresh Graduate',
    'Workshop persiapan karier meliputi CV, LinkedIn, interview, personal branding, dan strategi mencari pekerjaan.',
    'https://images.unsplash.com/photo-1551836022-d5d88e9218df',
    '2026-11-15',
    '10:00',
    '14:00',
    'online',
    'Google Meet',
    250,
    'active'
),

(
    (SELECT user_id FROM users ORDER BY user_id LIMIT 1),
    (SELECT community_id FROM communities WHERE name = 'Career Development Indonesia' LIMIT 1),
    'Tech Career Talk',
    'Diskusi bersama profesional teknologi mengenai Backend Engineer, Frontend Engineer, Data Scientist, dan DevOps.',
    'https://images.unsplash.com/photo-1516321318423-f06f85e504b3',
    '2026-11-20',
    '13:00',
    '17:00',
    'in_person',
    'Bandung Digital Valley',
    100,
    'active'
);



-- EVENT CATEGORY MAPPING

INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Go Concurrency Workshop'
AND ec.category_name = 'Programming'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Go Concurrency Workshop'
AND ec.category_name = 'Technology'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Building REST API with Go'
AND ec.category_name = 'Programming'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Generative AI for Beginners'
AND ec.category_name = 'AI'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Machine Learning Bootcamp'
AND ec.category_name = 'AI'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'UI UX Portfolio Review'
AND ec.category_name = 'Design'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Digital Business Strategy 2026'
AND ec.category_name = 'Business'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Career Preparation for Fresh Graduate'
AND ec.category_name = 'Career'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);


INSERT INTO event_category_mappings (
    event_id,
    category_id
)
SELECT
    e.event_id,
    ec.category_id
FROM events e, event_categories ec
WHERE e.title = 'Tech Career Talk'
AND ec.category_name = 'Career'
AND NOT EXISTS (
    SELECT 1
    FROM event_category_mappings m
    WHERE m.event_id = e.event_id
    AND m.category_id = ec.category_id
);