-- Creation of products table
create table products(
	id serial primary key,
	name varchar,
	description varchar,
	price decimal,
	ammount integer
);

-- Populating product table
INSERT INTO products (name, description, price, ammount) VALUES
('Product 1', 'Description 1', 10.00, 10),
('Product 2', 'Description 2', 20.00, 20),
('Product 3', 'Description 3', 30.00, 30),
('Product 4', 'Description 4', 40.00, 40),
('Product 5', 'Description 5', 50.00, 50),
('Product 6', 'Description 6', 60.00, 60),
('Product 7', 'Description 7', 70.00, 70),
('Product 8', 'Description 8', 80.00, 80),
('Product 9', 'Description 9', 90.00, 90),
( 'Product 10', 'Description 10', 100.00, 100);