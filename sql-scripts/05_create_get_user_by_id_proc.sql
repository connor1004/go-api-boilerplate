DROP PROCEDURE IF EXISTS `get_user_by_id`;

DELIMITER $$
CREATE PROCEDURE `get_user_by_id`(IN `p_id` BIGINT(21))
BEGIN
  SELECT users.id, job_titles.job_title, first_name, last_name, gender, birth_date, departments.name AS department_name, badge_id, phone, email
  FROM users
  INNER JOIN job_titles ON job_titles.id = users.job_title_id
  INNER JOIN departments ON departments.id = users.department_id
  WHERE users.id = p_id;
END$$
DELIMITER ;