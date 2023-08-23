DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` BIGINT(21) unsigned NOT NULL AUTO_INCREMENT,
  `job_title_id` INT(11) unsigned NOT NULL,
  `first_name` VARCHAR(50) NOT NULL,
  `last_name` VARCHAR(50) NOT NULL,
  `gender` ENUM('M', 'F') NOT NULL DEFAULT 'M',
  `birth_date` DATETIME NULL DEFAULT NULL,
  `department_id` INT(11) unsigned NOT NULL,
  `badge_id` BIGINT(21),
  `phone` VARCHAR(50) NOT NULL DEFAULT '',
  `email` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`badge_id`),
  FOREIGN KEY (`job_title_id`) REFERENCES `job_titles`(`id`),
  FOREIGN KEY (`department_id`) REFERENCES `departments`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;