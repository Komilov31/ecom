CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL NOT NULL PRIMARY KEY,
    order_id INT REFERENCES orders(id) NOT NULL,
    product_id INT NOT NULL REFERENCES products(id),
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);