CREATE EXTENSION IF NOT EXISTS pgcrypto;

--TABLE for credentials
CREATE TABLE credentials (
    id UUID  PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
--TABLE for users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    rocketpoints INT DEFAULT 0
);
--TABLE for daily steps
CREATE TABLE daily_steps (
    id UUID PRIMARY KEY,  -- Unique identifier for each entry
    user_id UUID NOT NULL,  -- The UUID of the user who recorded the steps
    steps_taken INT NOT NULL,  -- Number of steps taken on that day
    date DATE NOT NULL,  -- The date on which the steps were recorded
    UNIQUE (user_id, date)  -- Ensures one entry per user per day
);

CREATE TABLE image_store (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    image_name VARCHAR(255) NOT NULL,
    image_data BYTEA NOT NULL
);

CREATE TABLE settings (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    image_id UUID,
    step_goal INT DEFAULT 100000
);

CREATE TABLE challenges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description TEXT NOT NULL,
    points_reward INT NOT NULL
);

CREATE TABLE user_challenges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    challenge_id UUID NOT NULL,
    date DATE NOT NULL,
    is_completed BOOLEAN DEFAULT FALSE,
    UNIQUE (user_id, challenge_id, date)
);

CREATE TABLE friends (
    user_id UUID NOT NULL,
    friend_id UUID NOT NULL,
    PRIMARY KEY (user_id, friend_id),
    CHECK (user_id <> friend_id)  -- Prevent a user from adding themselves as a friend
);
