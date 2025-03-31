CREATE DATABASE userdata;
USE userdata;
--TABLE for users
CREATE TABLE users (
username VARCHAR(255) NOT NULL PRIMARY KEY,
email VARCHAR(255) NOT NULL,
firstname VARCHAR(255) NOT NULL,
lastname VARCHAR(255) NOT NULL,
rocketpoints INT DEFAULT 0,
);

CREATE TABLE runs (
username VARCHAR(255) NOT NULL REFERENCES users(username),
duration TIME NOT NULL,
distance DOUBLE NOT NULL,
avg_speed DOUBLE NOT NULL,
date DATE NOT NULL,
route GEOMETRY(LINESTRING, 4326) NOT NULL,
--LINESTRING is a type of geometry that represents a sequence of points in 2D space
--and 4326 is the SRID (Spatial Reference Identifier) for WGS 84, a standard for latitude and longitude coordinates
PRIMARY KEY (username, date)
--Maybe need to investigate if foreign key is needed here to link to users table
);
--TABLE for friends
--HOW TO HANDLE FRIENDS? AND REQUESTS?