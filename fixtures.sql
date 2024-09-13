-- Create the table
CREATE TABLE IF NOT EXISTS testable (
    created_at DATETIME
);

-- Insert a sample records
INSERT INTO testable (created_at) VALUES (NOW());
