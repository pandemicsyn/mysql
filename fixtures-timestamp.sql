-- Create the table
CREATE TABLE IF NOT EXISTS testable (
    created_at DATETIME
);

-- Insert a sample records - only works in strict mode off like via sql_mode=NO_ENGINE_SUBSTITUTION
INSERT INTO testable (created_at) VALUES ("2019-01-01T01:23:45.678Z");
