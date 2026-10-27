CREATE TYPE user_status AS ENUM (
    'active',
    'suspended'
);

CREATE TYPE community_status AS ENUM (
    'active',
    'non_active'
);

CREATE TYPE event_format AS ENUM (
    'in_person',
    'online'
);

CREATE TYPE event_status AS ENUM (
    'active',
    'full',
    'completed'
);