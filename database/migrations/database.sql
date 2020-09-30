/*
*********************************************************************
API Go Boilerplate
*********************************************************************
Name: MySQL Sample Database
Version 1
*********************************************************************
*/

CREATE TABLE "user"
(
    "id"         SERIAL PRIMARY KEY,
    "last_name"  varchar(50) NULL,
    "first_name" varchar(50) NULL,
    "email"      varchar(50) NOT NULL,
    "password"   varchar(255) NOT NULL
);

