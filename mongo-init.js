db.createUser(
    {
        user: "root",
        pwd: "root",
        roles: [
            {
                role: "readWrite",
                db: "fakerAPI"
            }
        ]
    }
);

db.createCollection("users");
db.createCollection("persons");