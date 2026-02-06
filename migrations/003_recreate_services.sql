-- Older versions of the services table were created without a foreign key constraint
-- on the dog_id column resulting in orphaned services when a dog is deleted.

-- Backup the dog_services table
CREATE TABLE dog_services_backup AS SELECT * FROM dog_services;

-- Drop the dog_services table
DROP TABLE dog_services;

-- Delete any orphaned services from backup (i.e. services with no associated dog)
DELETE FROM dog_services_backup WHERE dog_id NOT IN (SELECT id FROM dogs);

-- Recreate the dog_services table
CREATE TABLE IF NOT EXISTS dog_services (
    id SERIAL PRIMARY KEY,
    dog_id INTEGER,
    service TEXT,
    quantity INTEGER,
    price REAL,
    FOREIGN KEY(dog_id) REFERENCES dogs(id) ON DELETE CASCADE
);

-- Restore the data from the backup (note: we don't restore id, let PostgreSQL auto-generate)
INSERT INTO dog_services (dog_id, service, quantity, price)
SELECT dog_id, service, quantity, price FROM dog_services_backup;

-- Drop the backup table
DROP TABLE dog_services_backup;
