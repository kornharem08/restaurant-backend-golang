BEGIN;

DROP TABLE IF EXISTS schema_migrations; 

DROP TABLE IF EXISTS "users";

-- Rollback changes for the "food" table
DROP TABLE IF EXISTS food;

-- Rollback changes for the "food_type" table
DROP TABLE IF EXISTS food_type;

-- Rollback changes for the "ingredient" table
DROP TABLE IF EXISTS ingredient;

-- Rollback changes for the "order" table
DROP TABLE IF EXISTS "order";


COMMIT;