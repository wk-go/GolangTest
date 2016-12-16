CREATE TABLE "post"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "title" TEXT(50) NOT NULL,
    "content" TEXT(65535) NOT NULL,
    "decimal_test" decimal(10,2) Default 0.00 NOT NULL,
    "user_id" INTEGER Default 0 NOT NULL,
    "Tags" INTEGER Default 0 NOT NULL
)



CREATE TABLE "profile"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "age" INTEGER Default 0 NOT NULL,
    "user" INTEGER Default 0 NOT NULL
)



CREATE TABLE "tags"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "name" TEXT(50) NOT NULL
)



CREATE TABLE "user"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "name" TEXT(50) NOT NULL,
    "profile_id" INTEGER Default 0 NOT NULL
)



