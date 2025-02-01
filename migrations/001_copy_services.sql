-- Insert into services and capture the mapping
INSERT INTO dog_services (dog_id, service, quantity, price)
SELECT id, service, quantity, price 
FROM dogs;
