CREATE ROLE connected_roots_role WITH LOGIN SUPERUSER PASSWORD 'Password1';
CREATE DATABASE connected_roots;
GRANT ALL PRIVILEGES ON DATABASE connected_roots TO connected_roots_role;

\connect connected_roots connected_roots_role;
CREATE SCHEMA IF NOT EXISTS sc_connected_roots;
