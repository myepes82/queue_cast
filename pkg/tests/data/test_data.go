package test_data

const (
	ComplexJsonObject = `
{
    "user": {
        "id": 123456789,
        "username": "exampleUser",
        "email": "exampleUser@example.com",
        "profile": {
            "first_name": "Example",
            "last_name": "User",
            "age": 30,
            "gender": "other",
            "address": {
                "street": "123 Example St",
                "city": "Exampleville",
                "state": "EX",
                "zip_code": "12345",
                "country": "Exampleland"
            },
            "phone_numbers": [
                {"type": "home", "number": "123-456-7890"},
                {"type": "work", "number": "098-765-4321"}
            ]
        },
        "preferences": {
            "language": "en",
            "timezone": "America/Example",
            "notifications": {
                "email": true,
                "sms": false,
                "push": true
            }
        },
        "activity_log": [
            {
                "timestamp": "2023-07-01T12:00:00Z",
                "action": "login",
                "ip_address": "192.0.2.1",
                "device": "desktop"
            },
            {
                "timestamp": "2023-07-02T08:30:00Z",
                "action": "logout",
                "ip_address": "192.0.2.1",
                "device": "desktop"
            },
            {
                "timestamp": "2023-07-02T10:00:00Z",
                "action": "purchase",
                "details": {
                    "item_id": "987654321",
                    "item_name": "Example Product",
                    "quantity": 2,
                    "price_per_unit": 49.99,
                    "total_price": 99.98
                }
            }
        ]
    },
    "friends": [
        {
            "id": 234567890,
            "username": "friend1",
            "profile": {
                "first_name": "Friend",
                "last_name": "One",
                "age": 25,
                "gender": "female",
                "address": {
                    "street": "456 Friend St",
                    "city": "Friendville",
                    "state": "FR",
                    "zip_code": "67890",
                    "country": "Friendland"
                }
            }
        },
        {
            "id": 345678901,
            "username": "friend2",
            "profile": {
                "first_name": "Friend",
                "last_name": "Two",
                "age": 28,
                "gender": "male",
                "address": {
                    "street": "789 Friend Ave",
                    "city": "Friendtown",
                    "state": "FT",
                    "zip_code": "11223",
                    "country": "Friendland"
                }
            }
        }
    ],
    "messages": [
        {
            "from": "friend1",
            "to": "exampleUser",
            "timestamp": "2023-07-02T12:00:00Z",
            "content": "Hey! How are you?"
        },
        {
            "from": "exampleUser",
            "to": "friend1",
            "timestamp": "2023-07-02T12:05:00Z",
            "content": "I'm good, thanks! How about you?"
        },
        {
            "from": "friend2",
            "to": "exampleUser",
            "timestamp": "2023-07-02T14:00:00Z",
            "content": "Let's meet up for lunch."
        },
        {
            "from": "exampleUser",
            "to": "friend2",
            "timestamp": "2023-07-02T14:10:00Z",
            "content": "Sure, where do you want to go?"
        }
    ]
}
`
)
