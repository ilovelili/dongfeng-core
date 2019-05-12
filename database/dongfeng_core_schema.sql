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
  `email` varchar(125) NOT NULL COMMENT 'Email',
  `name` varchar(50) DEFAULT NULL COMMENT 'Full name',
  `avatar` varchar(255) DEFAULT NULL COMMENT 'Avatar',
  `settings` int(11) NOT NULL COMMENT 'Bitwise settings',
  `role` varchar(45) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT '1' COMMENT 'Enabled',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  KEY `idx_users_email` (`email`)
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
  `read` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `classes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(9) NOT NULL,
  `created_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `absences` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` varchar(10) NOT NULL,
  `date` varchar(10) NOT NULL,
  `class` varchar(10) NOT NULL,
  `name` varchar(10) NOT NULL,
  `created_by` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `date_class_name_UNIQUE` (`date`,`class`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `pupils` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` varchar(10),
  `class` varchar(10) NOT NULL,
  `name` varchar(10) NOT NULL,  
  `created_by` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `year_class_name_UNIQUE` (`year`,`class`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `teachers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` varchar(10),
  `class` varchar(100) NOT NULL,
  `name` varchar(10) NOT NULL,
  `email` varchar(100) NOT NULL,
  `role` varchar(10) NOT NULL,
  `created_by` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),  
  UNIQUE KEY `year_class_name_UNIQUE` (`year`,`class`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `holidays` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from` varchar(10) NOT NULL,
  `to` varchar(10) NOT NULL,
  `description` varchar(50) NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `from_to_UNIQUE` (`from`, `to`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `physique_height_to_weight_p_master` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `h_w` varchar(4) NOT NULL,
  `gender` varchar(4) NOT NULL,
  `age_min` double NOT NULL,
  `age_max` double NOT NULL,
  `p3` double NOT NULL,
  `p10` double NOT NULL,
  `p20` double NOT NULL,
  `p50` double NOT NULL,
  `p97` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `physique_age_height_weight_sd_master` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `h_w` varchar(4) NOT NULL,
  `gender` varchar(4) NOT NULL,
  `age` varchar(10) NOT NULL,
  `sdm2` double NOT NULL,
  `sdm1` double NOT NULL,
  `avg` double NOT NULL,
  `sd1` double NOT NULL,
  `sd2` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `physique_height_to_weight_p_master` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gender` varchar(4) NOT NULL,
  `height` double NOT NULL,
  `p3` double NOT NULL,
  `p10` double NOT NULL,
  `p20` double NOT NULL,
  `p50` double NOT NULL,
  `p80` double NOT NULL,
  `p97` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `physique_height_to_weight_sd_master` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gender` varchar(4) NOT NULL,
  `height` double NOT NULL,
  `sdm3` double NOT NULL,
  `sdm2` double NOT NULL,
  `sdm1` double NOT NULL,
  `sd0` double NOT NULL,
  `sd1` double NOT NULL,
  `sd2` double NOT NULL,
  `sd3` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=223 DEFAULT CHARSET=utf8;

CREATE TABLE `physique_bmi_master` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gender` varchar(4) NOT NULL,
  `age` varchar(10) NOT NULL,
  `avg` double NOT NULL,
  `1sd` double NOT NULL,
  `2sd` double NOT NULL,
  `3sd` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `physiques` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `year` varchar(10) NOT NULL,
  `class` varchar(10) NOT NULL,
  `name` varchar(10) NOT NULL,
  `gender` varchar(4) NOT NULL,
  `birth_date` varchar(10) NOT NULL,
  `exam_date` varchar(10) NOT NULL,
  `age` varchar(10) NOT NULL,
  `age_cmp` double NOT NULL,
  `height` double NOT NULL,
  `height_p` varchar(10) NOT NULL,
  `height_sd` varchar(10) NOT NULL,
  `weight` double NOT NULL,
  `weight_p` varchar(10) NOT NULL,
  `weight_sd` varchar(10) NOT NULL,
  `height_weight_p` varchar(10) NOT NULL,
  `height_weight_sd` varchar(10) NOT NULL,
  `bmi` double NOT NULL,
  `bmi_sd` varchar(10) NOT NULL,
  `fat_cofficient` double NOT NULL,
  `conclusion` varchar(45) DEFAULT NULL,
  `created_by` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `ingredients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `material` varchar(45) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `recipes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `ingredient_id` int(11) NOT NULL,
  `unit_amount` decimal(5,2) NOT NULL DEFAULT '0.00',
  `created_by` varchar(100) NOT NULL DEFAULT 'AgentSmith',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_ingredient_UNIQUE` (`name`,`ingredient_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `ingredient_nutritions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ingredient` varchar(50) NOT NULL,
  `protein_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `protein_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `fat_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `fat_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `carbohydrate_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `carbohydrate_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `heat_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `heat_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `calcium_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `calcium_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `iron_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `iron_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `zinc_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `zinc_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `va_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `va_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `vb1_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `vb1_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `vb2_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `vb2_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `vc_100g` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `vc_daily` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `ingredient_UNIQUE` (`ingredient`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `recipe_nutritions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `recipe` varchar(50) NOT NULL,
  `carbohydrate` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `dietaryfiber` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `protein` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `fat` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `heat` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `recipe_UNIQUE` (`recipe`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `menus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` varchar(10) NOT NULL,
  `recipe` varchar(45) NOT NULL,
  `breakfast_or_lunch` int(1) NOT NULL COMMENT '0: breakfast\n1: lunch\n2: snack',
  `junior_or_senior_class` int(1) NOT NULL COMMENT '0: junior\n1: senior',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `nutrition_allocations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `junior_or_senior` int(1) NOT NULL COMMENT '0: junior\n1: senior',
  `breakfast_or_lunch` int(1) NOT NULL COMMENT '0: breakfast\n1: lunch\n2: snack\n3: dinner',
  `allocation` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
