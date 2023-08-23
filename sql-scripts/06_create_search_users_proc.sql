DROP PROCEDURE IF EXISTS `search_users`;

DELIMITER $$
CREATE PROCEDURE `search_users`(
  IN `p_name` VARCHAR(100),
  IN `p_job_title` VARCHAR(50),
  IN `p_department_name` VARCHAR(50),
  IN `p_phone` VARCHAR(50),
  IN `p_email` VARCHAR(50)
)
BEGIN
  declare name_val VARCHAR(110);
  declare job_title_val VARCHAR(55);
  declare department_name_val VARCHAR(55);
  declare phone_val VARCHAR(55);
  declare email_val VARCHAR(55);

  IF (p_name IS NULL) THEN
    SET name_val = '';
  ELSE
    SET name_val = p_name;
  END IF;

  SET name_val = CONCAT('%', name_val, '%');

  IF (p_job_title IS NULL) THEN
    SET job_title_val = '';
  ELSE
    SET job_title_val = p_job_title;
  END IF;

  SET job_title_val = CONCAT('%', job_title_val, '%');

  IF (p_department_name IS NULL) THEN
    SET department_name_val = '';
  ELSE
    SET department_name_val = p_department_name;
  END IF;

  SET department_name_val = CONCAT('%', department_name_val, '%');

  IF (p_phone IS NULL) THEN
    SET phone_val = '';
  ELSE
    SET phone_val = p_phone;
  END IF;

  SET phone_val = CONCAT('%', phone_val, '%');

  IF (p_email IS NULL) THEN
    SET email_val = '';
  ELSE
    SET email_val = p_email;
  END IF;

  SET email_val = CONCAT('%', email_val, '%');

  SELECT users.id, job_titles.job_title, first_name, last_name, gender, birth_date, departments.name AS department_name, badge_id, phone, email
  FROM users
  INNER JOIN job_titles ON job_titles.id = users.job_title_id
  INNER JOIN departments ON departments.id = users.department_id
  WHERE CONCAT(users.first_name, ' ', users.last_name) LIKE name_val AND job_title LIKE job_title_val AND departments.name LIKE department_name_val AND phone LIKE phone_val AND email LIKE email_val;
END$$
DELIMITER ;