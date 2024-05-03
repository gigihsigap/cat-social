CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL
);

CREATE TYPE race_enum AS ENUM ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman');
CREATE TYPE sex_enum AS ENUM ('male', 'female');

CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    name VARCHAR(30) NOT NULL,
    race race_enum NOT NULL,
    sex sex_enum NOT NULL,
    age_in_month INT NOT NULL,
    image_urls TEXT[],
    description VARCHAR(200) NOT NULL,
    has_matched BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_cats_all_columns ON cats (name);

CREATE TYPE status_match_enum AS ENUM ('pending', 'approved', 'rejected');

CREATE TABLE cat_matches (
    id SERIAL PRIMARY KEY,
    issuer_id SERIAL NOT NULL,
    receiver_id SERIAL NOT NULL,
    match_cat_id SERIAL NOT NULL,
    user_cat_id SERIAL NOT NULL,
    message VARCHAR(120) NOT NULL,
    status status_match_enum NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL,
    FOREIGN KEY (issuer_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id) REFERENCES users(id)
);

CREATE INDEX idx_cat_matches_all_columns ON cat_matches (user_cat_id, match_cat_id,status);
