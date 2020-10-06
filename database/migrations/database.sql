/*
*********************************************************************
Gechoplate
Go Echo Boilerplate for Rest APIs
*********************************************************************
Name: MySQL Sample Database
Version 1
*********************************************************************
*/

/*Schema for the table `user` */

CREATE TABLE `user` (
    `id`         int(11) NOT NULL,
    `last_name`  varchar(50) NULL,
    `first_name` varchar(50) NULL,
    `email`      varchar(50) NOT NULL,
    `password`   varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `user` */

insert  into `user`(`id`,`last_name`,`first_name`,`email`,`password`) values 
(1,'Doe','John','john.doe@gechoplate.com',
'$2y$12$f.u02sBB14r8R.mFUDu4k.HEf3gL/dJbdIuLnjHOuLgeipCEzrtry'); /* <-- This is 'password' in Bcrypt for User testing */