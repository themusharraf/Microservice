db.createUser(
        {
            user: "root",
            pwd: "aSLJqN22pBluITGH",
            roles: [
                {
                    role: "readWrite",
                    db: "testdb"
                }
            ]
        }
);