CREATE TABLE IF NOT EXISTS orders (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT REFERENCES users(id) NOT NULL ,
    total DECIMAL(10, 2) NOT NULL,
    status order_status NOT NULL DEFAULT 'pending',
    address TEXT  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
); 