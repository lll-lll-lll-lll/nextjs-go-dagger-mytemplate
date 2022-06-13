CREATE TABLE products 
(
    id SERIAL,
    name TEXT NOT NULL,
    price INT NOT NULL,
    CONSTRAINT product_pkey PRIMARY KEY (id)
)