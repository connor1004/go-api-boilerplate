DROP TABLE IF EXISTS `job_titles`;

CREATE TABLE `job_titles` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `job_title` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`job_title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;