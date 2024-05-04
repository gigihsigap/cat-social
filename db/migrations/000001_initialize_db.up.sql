CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL
);

CREATE TYPE race_enum AS ENUM ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman');
CREATE TYPE sex_enum AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    race race_enum NOT NULL,
    sex sex_enum NOT NULL,
    age_in_months INT,
    description VARCHAR(200),
    image_urls TEXT[],
    user_id INT NOT NULL,
    has_matched BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    match_cat_id INT NOT NULL,
    user_cat_id INT NOT NULL,
    is_approved BOOLEAN DEFAULT NULL,
    message TEXT NOT NULL,
    issued_by INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT NULL,
    FOREIGN KEY (issued_by) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (match_cat_id) REFERENCES cats(id) ON DELETE CASCADE,
    FOREIGN KEY (user_cat_id) REFERENCES cats(id) ON DELETE CASCADE
);
