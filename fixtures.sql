-- Create the table
CREATE TABLE IF NOT EXISTS testable (
    created_at DATETIME
);

-- Insert a sample records
--INSERT INTO testable (created_at) VALUES ("2019-01-01T01:23:45.678Z");
INSERT INTO testable (created_at) VALUES (NOW());