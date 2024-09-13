-- +goose Up
-- CREATE TABLE IF NOT EXISTS employee (
--                           id UUID PRIMARY KEY,
--                           username VARCHAR(50) UNIQUE NOT NULL,
--                           first_name VARCHAR(50),
--                           last_name VARCHAR(50),
--                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
--
-- CREATE TYPE organization_type AS ENUM (
--     'IE',
--     'LLC',
--     'JSC'
--     );
--
-- CREATE TABLE IF NOT EXISTS organization (
--                               id UUID PRIMARY KEY,
--                               name VARCHAR(100) NOT NULL,
--                               description TEXT,
--                               type organization_type,
--                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
--
-- CREATE TABLE IF NOT EXISTS organization_responsible (
--                                           id SERIAL PRIMARY KEY,
--                                           organization_id uuid REFERENCES organization(id) ON DELETE CASCADE,
--                                           user_id uuid REFERENCES employee(id) ON DELETE CASCADE
-- );

CREATE TYPE service_type AS ENUM (
    'Construction',
    'Delivery',
    'Manufacture'
    );

CREATE TYPE tender_status AS ENUM (
    'Created',
    'Published',
    'Closed'
    );

CREATE TABLE IF NOT EXISTS tender (
                        id UUID PRIMARY KEY default 'uuid_generate_v4()',
                        name VARCHAR(100) NOT NULL,
                        description VARCHAR(500) NOT NULL,
                        service_type service_type,
                        status tender_status DEFAULT 'Created',
                        organization_id uuid REFERENCES organization(id) ON DELETE CASCADE,
                        user_id uuid  REFERENCES employee(id) ON DELETE CASCADE,
                        version INT DEFAULT 1,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE bild_status AS ENUM (
    'Created',
    'Published',
    'Canceled',
    'Approved',
    'Rejected'
    );

CREATE TYPE author_type AS ENUM (
    'Created',
    'Published',
    'Closed'
    );

CREATE TABLE IF NOT EXISTS bid (
                      id UUID PRIMARY KEY default 'uuid_generate_v4()',
                      name VARCHAR(100) NOT NULL,
                      description VARCHAR(500) NOT NULL,
                      status bild_status,
                      tender_id uuid REFERENCES tender(id) ON DELETE CASCADE,
                      author_type author_type,
                      author_id uuid REFERENCES employee(id) ON DELETE CASCADE,
                      version INT DEFAULT 1,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- Удаляем зависимые таблицы
DROP TABLE IF EXISTS organization_responsible;
DROP TABLE IF EXISTS bild;
DROP TABLE IF EXISTS tender;
DROP TABLE IF EXISTS employee;
DROP TABLE IF EXISTS organization;

-- Удаляем типы, когда они больше не используются
DROP TYPE IF EXISTS organization_type;
DROP TYPE IF EXISTS service_type;
DROP TYPE IF EXISTS tender_status;
DROP TYPE IF EXISTS bid_status;
DROP TYPE IF EXISTS author_type;


