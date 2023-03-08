CREATE database db_commerce;
USE db_commerce;

CREATE TABLE customers (
    customer_id       int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    company_name      VARCHAR(50),
    first_name        VARCHAR(30),
    last_name         VARCHAR(50),
    billing_address   VARCHAR(255),
    city              VARCHAR(50),
    province          VARCHAR(20),
    zip_code          VARCHAR(20),
    email             VARCHAR(75),
    company_website   VARCHAR(200),
    phone_number      VARCHAR(30),
    fax_number        VARCHAR(30),
    ship_address      VARCHAR(255),
    ship_city         VARCHAR(50),
    ship_province     VARCHAR(50),
    ship_zip_code     VARCHAR(20),
    ship_phone_number VARCHAR(30)
);

CREATE TABLE employees (
    employee_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name  VARCHAR(50),
    last_name   VARCHAR(50),
    title       VARCHAR(50),
    work_phone  VARCHAR(30)
);

CREATE TABLE products (
    product_id   int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_name VARCHAR(50),
    unit_price   int,
    in_stock     CHAR(1)
);

CREATE TABLE shipping_methods (
    shipping_method_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    shipping_method    VARCHAR(20)
);

CREATE TABLE orders (
    order_id              int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    customer_id           int,
    employee_id           int,
    order_date            DATE,
    purchase_order_number VARCHAR(30),
    ship_date             DATE,
    shipping_method_id    int,
    freight_charge        int,
    taxes                 int,
    payment_received      CHAR(1),
    comment               VARCHAR(150),
    FOREIGN KEY (customer_id) REFERENCES customers (customer_id),
    FOREIGN KEY (employee_id) REFERENCES employees (employee_id),
    FOREIGN KEY (shipping_method_id) REFERENCES shipping_methods (shipping_method_id)
);

CREATE TABLE order_details (
    order_detail_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    order_id        int,
    product_id      int,
    quantity        int,
    unit_price      int,
    discount        int,
    FOREIGN KEY (order_id) REFERENCES orders (order_id),
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);