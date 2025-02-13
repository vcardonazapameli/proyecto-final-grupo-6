CREATE DATABASE IF NOT EXISTS meli_db;

USE meli_db;

CREATE TABLE countries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    country_name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE provinces (
    id INT AUTO_INCREMENT PRIMARY KEY,
    province_name VARCHAR(255) NOT NULL,
    id_country_fk INT,
    FOREIGN KEY (id_country_fk) REFERENCES countries(id)
);

CREATE TABLE localities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    locality_name VARCHAR(255) NOT NULL,
    province_id INT,
    FOREIGN KEY (province_id) REFERENCES provinces(id)
);

CREATE TABLE warehouses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    warehouse_code VARCHAR(255) NOT NULL UNIQUE,
    locality_id INT,
    FOREIGN KEY (locality_id) REFERENCES localities(id)
);

CREATE TABLE sellers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(255) UNIQUE NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    locality_id INT,
    FOREIGN KEY (locality_id) REFERENCES localities(id)
);

CREATE TABLE buyers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_card_number VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL
);

CREATE TABLE carriers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(255) NOT NULL UNIQUE,
    company_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    locality_id INT,
    FOREIGN KEY (locality_id) REFERENCES localities(id)
);

CREATE TABLE products_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE sections (
    id INT AUTO_INCREMENT PRIMARY KEY,
    section_number VARCHAR(255) NOT NULL UNIQUE,
    current_capacity INT NOT NULL,
    current_temperature DECIMAL(19,2) NOT NULL,
    maximum_capacity INT NOT NULL,
    minimum_capacity INT NOT NULL,
    minimum_temperature DECIMAL(19,2) NOT NULL,
    product_type_id INT,
    warehouse_id INT,
    FOREIGN KEY (product_type_id) REFERENCES products_types(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_code VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255) NOT NULL,
    expiration_rate DECIMAL(19,2) NOT NULL,
    recommended_freezing_temperature DECIMAL(19,2) NOT NULL,
    freezing_rate DECIMAL(19,2) NOT NULL,
    width DECIMAL(19,2) NOT NULL,
    height DECIMAL(19,2) NOT NULL,
    length DECIMAL(19,2) NOT NULL,
    net_weight DECIMAL(19,2) NOT NULL,
    product_type_id INT,
    seller_id INT,
    FOREIGN KEY (product_type_id) REFERENCES products_types(id),
    FOREIGN KEY (seller_id) REFERENCES sellers(id)
);

CREATE TABLE order_status (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE purchase_orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL UNIQUE,
    order_date DATETIME(6) NOT NULL,
    tracking_code VARCHAR(255) NOT NULL UNIQUE,
    buyer_id INT,
    carrier_id INT,
    order_status_id INT,
    wareHouse_id INT,
    FOREIGN KEY (buyer_id) REFERENCES buyers(id),
    FOREIGN KEY (carrier_id) REFERENCES carriers(id),
    FOREIGN KEY (order_status_id) REFERENCES order_status(id),
    FOREIGN KEY (wareHouse_id) REFERENCES warehouses(id)
);

CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_card_number VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    wareHouse_id INT,
    FOREIGN KEY (wareHouse_id) REFERENCES warehouses(id)
);

CREATE TABLE product_batches (
    id INT AUTO_INCREMENT PRIMARY KEY,
    batch_number VARCHAR(255) NOT NULL UNIQUE,
    current_quantity INT NOT NULL,
    current_temperature DECIMAL(19,2) NOT NULL,
    due_date DATETIME(6) NOT NULL,
    initial_quantity INT NOT NULL,
    manufacturing_date DATETIME(6) NOT NULL,
    manufacturing_hour DATETIME(6) NOT NULL,
    minimum_temperature DECIMAL(19,2) NOT NULL,
    product_id INT,
    section_id INT,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (section_id) REFERENCES sections(id)
);

CREATE TABLE product_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    last_update_date DATETIME(6) NOT NULL,
    purchase_price DECIMAL(19,2) NOT NULL,
    sale_price DECIMAL(19,2) NOT NULL,
    product_id INT,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE order_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    clean_liness_status VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    temperature DECIMAL(19,2) NOT NULL,
    product_record_id INT,
    purchase_order_id INT,
    FOREIGN KEY (product_record_id) REFERENCES product_records(id),
    FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id)
);

CREATE TABLE inbound_orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_date DATETIME(6) NOT NULL,
    order_number VARCHAR(255) NOT NULL UNIQUE,
    employe_id INT,
    product_batch_id INT,
    wareHouse_id INT,
    FOREIGN KEY (employe_id) REFERENCES employees(id),
    FOREIGN KEY (product_batch_id) REFERENCES product_batches(id),
    FOREIGN KEY (wareHouse_id) REFERENCES warehouses(id)
);