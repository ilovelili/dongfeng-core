-- create database schema
CREATE DATABASE IF NOT EXISTS dongfeng_core;
USE dongfeng_core;

CREATE TABLE `settings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL COMMENT 'Setting name',
  `description` varchar(100) DEFAULT NULL COMMENT 'Description',
  `value` int(11) NOT NULL COMMENT 'Bitwise value',
  `enabled` tinyint(4) NOT NULL DEFAULT '1' COMMENT 'Enabled or not',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `value_UNIQUE` (`value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
  `id` varchar(12) NOT NULL COMMENT 'User ID',
  `email` varchar(255) NOT NULL COMMENT 'Email',
  `name` varchar(100) DEFAULT NULL COMMENT 'Full name',  
  `avatar` varchar(255) DEFAULT NULL COMMENT 'Avatar',
  `settings` int(11) NOT NULL COMMENT 'Bitwise settings', -- The default setting need to be discussed
  `enabled` tinyint(4) NOT NULL DEFAULT '1' COMMENT 'Enabled',  
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  KEY `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_friends` (
  `user_id` varchar(12) NOT NULL COMMENT 'User id',
  `friend_id` varchar(12) NOT NULL COMMENT 'Friend id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Created time',
  PRIMARY KEY (`user_id`,`friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `description` varchar(255) NOT NULL COMMENT 'Description',
  `admin_only` tinyint(4) NOT NULL COMMENT 'Is the category visible to admin only',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `notifications` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(12) NOT NULL,
  `custom_code` varchar(6) NOT NULL,
  `category_id` int(11) NOT NULL,
  `details` text NOT NULL,
  `link` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `classes` (
  `id` varchar(5) NOT NULL,
  `name` varchar(9) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `class_code_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `attendances` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` varchar(10) NOT NULL,
  `class_id` varchar(10) NOT NULL,
  `name` varchar(10) NOT NULL,  
  `created_by` varchar(45) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `date_class_name_UNIQUE` (`date`,`class_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `namelists` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` varchar(10) NOT NULL,
  `class_id` varchar(10) NOT NULL,
  `name` varchar(10) NOT NULL,  
  `created_by` varchar(45) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `year_class_name_UNIQUE` (`year`,`class_id`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `recipes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `ingredient_id` int(11) NOT NULL,
  `created_by` varchar(45) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_ingredient_UNIQUE` (`name`,`ingredient_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
