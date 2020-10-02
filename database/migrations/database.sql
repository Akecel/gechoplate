/*
*********************************************************************
Gechoplate
Go Echo Boilerplate for Rest APIs
*********************************************************************
Name: MySQL Sample Database
Version 1
*********************************************************************
*/

CREATE TABLE "user"
(
    "id"        int(11) NOT NULL,
    "last_name"  varchar(50) NULL,
    "first_name" varchar(50) NULL,
    "email"      varchar(50) NOT NULL,
    "password"   varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

