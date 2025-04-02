-- Create the superTokens user
CREATE USER "superTokens" WITH ENCRYPTED PASSWORD 'somePassword';

-- Create a database for superTokens user and assign ownership
CREATE DATABASE "superTokens_db" WITH OWNER "superTokens";

-- Create the app user
CREATE USER "app" WITH ENCRYPTED PASSWORD 'somePassword';

-- Create a database for app user and assign ownership
CREATE DATABASE "app_db" WITH OWNER "app";
