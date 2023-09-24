CREATE TABLE IF NOT EXISTS products (
           id SERIAL PRIMARY KEY,
           name VARCHAR(255) NOT NULL,
           created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS product_properties (
           id SERIAL PRIMARY KEY,
           product_id INT NOT NULL,
           name VARCHAR(255) NOT NULL,
           value VARCHAR(255),
           UNIQUE (product_id, name),
           FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

