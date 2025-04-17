--TABLE for users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL,
    rocketpoints INT DEFAULT 0,
    FOREIGN KEY (id) REFERENCES credentials(id)  -- Establishing the relationship to credentials table
);

CREATE TABLE daily_steps (
    id UUID PRIMARY KEY,  -- Unique identifier for each entry
    user_id UUID NOT NULL,  -- The UUID of the user who recorded the steps
    steps_taken INT NOT NULL,  -- Number of steps taken on that day
    date DATE NOT NULL,  -- The date on which the steps were recorded
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to users table
    UNIQUE (user_id, date)  -- Ensures one entry per user per day
);

CREATE TABLE settings (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    profile_image VARCHAR(225),
    step_goal INT DEFAULT 100000
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- CREATE TABLE runs (
-- username VARCHAR(255) NOT NULL REFERENCES users(username),
-- duration TIME NOT NULL,
-- distance real NOT NULL,
-- avg_speed real NOT NULL,
-- date DATE NOT NULL,
-- route GEOMETRY(LINESTRING, 4326) NOT NULL,
-- --LINESTRING is a type of geometry that represents a sequence of points in 2D space
-- --and 4326 is the SRID (Spatial Reference Identifier) for WGS 84, a standard for latitude and longitude coordinates
-- PRIMARY KEY (username, date)
-- --Maybe need to investigate if foreign key is needed here to link to users table
-- );
-- --TABLE for friends
-- --HOW TO HANDLE FRIENDS? AND REQUESTS?

-- --CREATE TABLE friendrequest
-- --CREATE TABLE friends
