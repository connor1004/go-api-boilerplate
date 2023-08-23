DROP PROCEDURE IF EXISTS `add_user`;

DELIMITER $$
CREATE PROCEDURE `add_user`(
  IN `p_job_title` VARCHAR(50),
  IN `p_first_name` VARCHAR(50),
  IN `p_last_name` VARCHAR(50),
  IN `p_gender` ENUM('M', 'F'),
  IN `p_birth_date` VARCHAR(30),
  IN `p_department_name` VARCHAR(50),
  IN `p_badge_id` BIGINT(21),
  IN `p_phone` VARCHAR(50),
  IN `p_email` VARCHAR(50)
)
BEGIN  
  declare g_job_title_id INT(11);
  declare g_department_id INT(11);
  declare g_user_id BIGINT(21);    

  -- job_title check if a job_title with the given name already exists
  SELECT id INTO g_job_title_id FROM job_titles WHERE job_title = p_job_title LIMIT 1;

  IF (g_job_title_id IS NULL) THEN
    INSERT INTO job_titles(job_title) VALUES (p_job_title);
    set g_job_title_id = (select last_insert_id());
  END IF;

  -- department check if a department with the given name already exists
  SELECT id INTO g_department_id FROM departments WHERE name = p_department_name LIMIT 1;
  IF (g_department_id IS NULL) THEN
    INSERT INTO departments(name) VALUES (p_department_name);
    set g_department_id = (select last_insert_id());
  END IF;

  -- users check if a user with a given badge_id already exists
  SELECT id into g_user_id FROM users WHERE badge_id = p_badge_id LIMIT 1;
  IF (g_user_id IS NULL) THEN
    INSERT INTO users(job_title_id, first_name, last_name, gender, birth_date, department_id, badge_id, phone, email)
      VALUES (g_job_title_id, p_first_name, p_last_name, p_gender, p_birth_date, g_department_id, p_badge_id, p_phone, p_email);
    SELECT last_insert_id(), 1;
  ELSE
    SELECT 'badge_id already exists!', 0;
  END IF;
END$$
DELIMITER ;