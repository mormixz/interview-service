db = db.getSiblingDB('interview');
db.createCollection('interview');
db.interview.insertMany([
    {
        _id: new ObjectId("645f4882bcc2918b5e4605af"),
        description: "description interview 000",
        status: "Done",
        created_by: new ObjectId("645f2b47e0931117713e74ca"),
        created_at: new Date("January 1, 2023 01:00:00"),
        comments: []
    },
    {
        _id: new ObjectId("645f4882bcc2918b5e4605b0"),
        description: "description interview 001",
        status: "In Progress",
        created_by: new ObjectId("645f2b47e0931117713e74ca"),
        created_at: new Date("January 1, 2023 02:00:00"),
        comments: []
    },
    {
        _id: new ObjectId("645f4882bcc2918b5e4605b1"),
        description: "description interview 002",
        status: "To Do",
        created_by: new ObjectId("645f2b47e0931117713e74ca"),
        created_at: new Date("January 1, 2023 03:00:00"),
        comments: [
            {
                message: "comment 003",
                created_by: new ObjectId("645f2b47e0931117713e74cb"),
                created_at: new Date("January 2, 2023 11:00:00")
            },
            {
                message: "comment 002",
                created_by: new ObjectId("645f2b47e0931117713e74cc"),
                created_at: new Date("January 1, 2023 11:00:00")
            },
            {
                message: "comment 001",
                created_by: new ObjectId("645f2b47e0931117713e74cb"),
                created_at: new Date("January 1, 2023 08:00:00")
            }
        ]
    },
    {
        _id: new ObjectId("645f4882bcc2918b5e4605b2"),
        description: "description interview 003",
        status: "To Do",
        created_by: new ObjectId("645f2b47e0931117713e74ca"),
        created_at: new Date("January 1, 2023 07:00:00"),
        comments: [
            {
                message: "comment 001",
                created_by: new ObjectId("645f2b47e0931117713e74cb"),
                created_at: new Date("January 1, 2023 11:00:00")
            }
        ]
    },
    {
        _id: new ObjectId("645f4882bcc2918b5e4605b3"),
        description: "description interview 004",
        status: "To Do",
        created_by: new ObjectId("645f2b47e0931117713e74ca"),
        created_at: new Date("January 2, 2023 03:00:00"),
        comments: []
    },
    {
        _id: new ObjectId("645f4882bcc2918b5e4605b4"),
        description: "description interview 005",
        status: "To Do",
        created_by: new ObjectId("645f2b47e0931117713e74ca"),
        created_at: new Date("January 3, 2023 03:00:00"),
        comments: []
    },
    {
        _id: new ObjectId("645f4882bcc2918b5e4605b5"),
        description: "description interview 006",
        status: "To Do",
        created_by: new ObjectId("645f2b47e0931117713e74ca"),
        created_at: new Date("January 4, 2023 03:00:00"),
        comments: []
    }
]);

db.createCollection('users');
db.users.insertMany([
    {
        _id: new ObjectId("645f2b47e0931117713e74ca"),
        name: "โรบินฮู้ด",
        email: "user1@robinhood.co.th"
    },
    {
        _id: new ObjectId("645f2b47e0931117713e74cb"),
        name: "แบทแมน",
        email: "user2@robinhood.co.th"
    },
    {
        _id: new ObjectId("645f2b47e0931117713e74cc"),
        name: "แคทวูแมน",
        email: "user3@robinhood.co.th"
    }
]);
