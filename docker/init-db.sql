-- Create the superTokens user
CREATE USER "superTokens" WITH ENCRYPTED PASSWORD 'somePassword';

-- Create a database for superTokens user and assign ownership
CREATE DATABASE "superTokens_db" WITH OWNER "superTokens";

-- Create the api user
CREATE USER "api" WITH ENCRYPTED PASSWORD 'somePassword';

-- Create a database for api user and assign ownership
CREATE DATABASE "api_db" WITH OWNER "api";
