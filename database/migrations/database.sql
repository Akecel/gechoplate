/*
*********************************************************************
Gechoplate
Go Echo Boilerplate for Rest APIs
*********************************************************************
Name: MySQL Sample Database
Version 1
*********************************************************************
*/

/*Schema for the table `users` */

CREATE TABLE `users` (
    `last_name`  varchar(50),
    `first_name` varchar(50),
    `email`      varchar(50) NOT NULL,
    `password`   varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `users` */

insert  into `users`(`id`,`last_name`,`first_name`,`email`,`password`) values 
(1,'Doe','John','john.doe@gechoplate.com',
'$2y$12$f.u02sBB14r8R.mFUDu4k.HEf3gL/dJbdIuLnjHOuLgeipCEzrtry'); /* <-- This is 'password' in Bcrypt for User testing */