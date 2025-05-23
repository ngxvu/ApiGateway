CREATE TABLE logs (
 id SERIAL PRIMARY KEY,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 deleted_at TIMESTAMP,
 endpoint VARCHAR(255) NOT NULL,
 method VARCHAR(10) NOT NULL,
 params JSONB,
 results JSONB,
 ip_address VARCHAR(50) NOT NULL,
 status INT NOT NULL,
 elapsed DECIMAL(10, 2) NOT NULL
);