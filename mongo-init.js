db.createUser(
    {
        user: "root",
        pwd: "root",
        roles: [
            {
                role: "readWrite",
                db: "fakered"
            }
        ]
    }
);

db.createCollection("users");
db.createCollection("persons");