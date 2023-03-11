INSERT INTO customers (customer_id,company_name,first_name,last_name,billing_address,city,province,zip_code,email,company_website,phone_number,fax_number,ship_address,ship_city,ship_province,ship_zip_code,ship_phone_number) VALUES (1,'company A','User','First','Jakarta','Jakarta','DKI Jakarta','12345','user@gmail.com','Company Web A','08123123','11111','Bandung','Bandung','Bandung','22222','0800001111');

INSERT INTO employees (employee_id,first_name,last_name,title,work_phone) VALUES (1,'Employee A','Last Name','Manager','0899999999');

INSERT INTO shipping_methods (shipping_method_id,shipping_method) VALUES (1,'JNE');

INSERT INTO orders (order_id,customer_id,employee_id,order_date,purchase_order_number,ship_date,shipping_method_id,freight_charge,taxes,payment_received,comment) VALUES
    (1,1,1,now(),1,now(),1,10,50,'1','payed');

INSERT INTO products (product_id,product_name,unit_price,in_stock)
VALUES
(1,'Product A',1000,'1'),
(2,'Product B',2000,'1');


INSERT INTO order_details (order_detail_id,order_id,product_id,quantity,unit_price,discount)
VALUES
(1,1,1,5,1000,0),
(2,1,1,6,2000,0),
(3,1,2,7,3000,0),
(4,1,2,8,4000,0),
(5,1,2,9,5000,0);

