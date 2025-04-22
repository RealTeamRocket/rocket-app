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
