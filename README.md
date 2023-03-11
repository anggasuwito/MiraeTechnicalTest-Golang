# MiraeTechnicalTest-Golang

1. schema database berada di path migration
2. class dan objectnya berada pada package api/models , golang tidak murni mengadaptasi OOP
tetapi beberapa fungsi OOP bisa diterapkan digolang, contohnya seperti method function yang ada
di bagian paling bawah pada package api/models tersebut
3. jalankan command ```go run main.go```

```
EXAMPLE

REQUEST
localhost:8008/v1/order?page=1&limit=1

RESPONSE
{
    "data": [
        {
            "order_detail_id": 1,
            "quantity": 5,
            "unit_price": 1000,
            "discount": 0,
            "product": {
                "product_id": 1,
                "product_name": "Product A",
                "unit_price": 1000,
                "in_stock": "1"
            },
            "order": {
                "order_id": 1,
                "order_date": "2023-03-11",
                "purchase_order_number": "1",
                "ship_date": "2023-03-11",
                "freight_charge": 10,
                "taxes": 50,
                "payment_received": "1",
                "comment": "payed",
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
                }
            }
        }
    ],
    "err_response": null,
    "message": "success",
    "meta": {
        "current_page": 1,
        "last_page": 5,
        "total": 5,
        "per_page": 1
    },
    "status": true
}
```

untuk detail output lainnya ada di package screenshot_output