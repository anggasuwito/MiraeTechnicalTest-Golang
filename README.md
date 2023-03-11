# MiraeTechnicalTest-Golang

1. schema database berada di path migration
2. class dan objectnya berada pada package api/models , golang tidak murni mengadaptasi OOP
tetapi beberapa fungsi OOP bisa diterapkan digolang, contohnya seperti method function yang ada
di bagian paling bawah pada package api/models tersebut
3. untuk running di local sesuaikan credentials yang ada pada file .env , kemudian jalankan command ```go run main.go```

```
EXAMPLE

REQUEST
localhost:8008/v1/order?page=1&limit=1

RESPONSE
{
    "data": [
        {
            "order_id": 1,
            "order_date": "2023-03-11",
            "purchase_order_number": "1",
            "ship_date": "2023-03-11",
            "freight_charge": 10,
            "taxes": 50,
            "payment_received": "1",
            "comment": "payed",
            "total_price": 2700,
            "final_total_price": 4060,
            "shipping_method": {
                "shipping_method_id": 1,
                "method": "JNE"
            },
            "employee": {
                "employee_id": 1,
                "first_name": "Employee A",
                "last_name": "Last Name",
                "title": "Manager",
                "work_phone": "0899999999"
            },
            "customer": {
                "customer_id": 1,
                "company_name": "company A",
                "first_name": "User",
                "last_name": "First",
                "billing_address": "Jakarta",
                "city": "Jakarta",
                "province": "DKI Jakarta",
                "zip_code": "12345",
                "email": "user@gmail.com",
                "company_website": "Company Web A",
                "phone_number": "08123123",
                "fax_number": "11111",
                "ship_address": "Bandung",
                "ship_city": "Bandung",
                "ship_province": "Bandung",
                "ship_zip_code": "22222",
                "ship_phone_number": "0800001111"
            },
            "details": [
                {
                    "order_detail_id": 1,
                    "quantity": 5,
                    "unit_price": 1000,
                    "discount": 10,
                    "price": 900,
                    "product": {
                        "product_id": 1,
                        "product_name": "Product A",
                        "unit_price": 1000,
                        "in_stock": "1"
                    }
                },
                {
                    "order_detail_id": 2,
                    "quantity": 6,
                    "unit_price": 2000,
                    "discount": 10,
                    "price": 1800,
                    "product": {
                        "product_id": 1,
                        "product_name": "Product A",
                        "unit_price": 1000,
                        "in_stock": "1"
                    }
                }
            ]
        }
    ],
    "err_response": null,
    "message": "success",
    "meta": {
        "current_page": 1,
        "last_page": 4,
        "total": 4,
        "per_page": 1
    },
    "status": true
}

REQUEST
localhost:8008/v1/order?page=2&limit=2

RESPONSE
{
    "data": [
        {
            "order_id": 3,
            "order_date": "2023-03-11",
            "purchase_order_number": "1",
            "ship_date": "2023-03-11",
            "freight_charge": 10,
            "taxes": 50,
            "payment_received": "1",
            "comment": "payed",
            "total_price": 9000,
            "final_total_price": 13510,
            "shipping_method": {
                "shipping_method_id": 1,
                "method": "JNE"
            },
            "employee": {
                "employee_id": 1,
                "first_name": "Employee A",
                "last_name": "Last Name",
                "title": "Manager",
                "work_phone": "0899999999"
            },
            "customer": {
                "customer_id": 1,
                "company_name": "company A",
                "first_name": "User",
                "last_name": "First",
                "billing_address": "Jakarta",
                "city": "Jakarta",
                "province": "DKI Jakarta",
                "zip_code": "12345",
                "email": "user@gmail.com",
                "company_website": "Company Web A",
                "phone_number": "08123123",
                "fax_number": "11111",
                "ship_address": "Bandung",
                "ship_city": "Bandung",
                "ship_province": "Bandung",
                "ship_zip_code": "22222",
                "ship_phone_number": "0800001111"
            },
            "details": [
                {
                    "order_detail_id": 5,
                    "quantity": 9,
                    "unit_price": 5000,
                    "discount": 10,
                    "price": 4500,
                    "product": {
                        "product_id": 2,
                        "product_name": "Product B",
                        "unit_price": 2000,
                        "in_stock": "1"
                    }
                },
                {
                    "order_detail_id": 6,
                    "quantity": 9,
                    "unit_price": 5000,
                    "discount": 10,
                    "price": 4500,
                    "product": {
                        "product_id": 2,
                        "product_name": "Product B",
                        "unit_price": 2000,
                        "in_stock": "1"
                    }
                }
            ]
        },
        {
            "order_id": 4,
            "order_date": "2023-03-11",
            "purchase_order_number": "1",
            "ship_date": "2023-03-11",
            "freight_charge": 10,
            "taxes": 50,
            "payment_received": "1",
            "comment": "payed",
            "total_price": 9000,
            "final_total_price": 13510,
            "shipping_method": {
                "shipping_method_id": 1,
                "method": "JNE"
            },
            "employee": {
                "employee_id": 1,
                "first_name": "Employee A",
                "last_name": "Last Name",
                "title": "Manager",
                "work_phone": "0899999999"
            },
            "customer": {
                "customer_id": 1,
                "company_name": "company A",
                "first_name": "User",
                "last_name": "First",
                "billing_address": "Jakarta",
                "city": "Jakarta",
                "province": "DKI Jakarta",
                "zip_code": "12345",
                "email": "user@gmail.com",
                "company_website": "Company Web A",
                "phone_number": "08123123",
                "fax_number": "11111",
                "ship_address": "Bandung",
                "ship_city": "Bandung",
                "ship_province": "Bandung",
                "ship_zip_code": "22222",
                "ship_phone_number": "0800001111"
            },
            "details": [
                {
                    "order_detail_id": 7,
                    "quantity": 9,
                    "unit_price": 5000,
                    "discount": 10,
                    "price": 4500,
                    "product": {
                        "product_id": 2,
                        "product_name": "Product B",
                        "unit_price": 2000,
                        "in_stock": "1"
                    }
                },
                {
                    "order_detail_id": 8,
                    "quantity": 9,
                    "unit_price": 5000,
                    "discount": 10,
                    "price": 4500,
                    "product": {
                        "product_id": 2,
                        "product_name": "Product B",
                        "unit_price": 2000,
                        "in_stock": "1"
                    }
                }
            ]
        }
    ],
    "err_response": null,
    "message": "success",
    "meta": {
        "current_page": 2,
        "last_page": 2,
        "total": 4,
        "per_page": 2
    },
    "status": true
}
```