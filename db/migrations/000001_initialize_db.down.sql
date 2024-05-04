-- Drop indexes
DROP INDEX IF EXISTS idx_cat_matches_all_columns;
DROP INDEX IF EXISTS idx_cats_all_columns;

-- Drop foreign key constraints
ALTER TABLE matches DROP CONSTRAINT IF EXISTS matches_issuer_id_fkey;
ALTER TABLE matches DROP CONSTRAINT IF EXISTS matches_receiver_id_fkey;
ALTER TABLE cats DROP CONSTRAINT IF EXISTS cats_user_id_fkey;

-- Drop tables
DROP TABLE IF EXISTS matches;
DROP TABLE IF EXISTS cats;
DROP TABLE IF EXISTS users;

-- Drop types
DROP TYPE IF EXISTS status_match_enum;
DROP TYPE IF EXISTS sex_enum;
DROP TYPE IF EXISTS race_enum;