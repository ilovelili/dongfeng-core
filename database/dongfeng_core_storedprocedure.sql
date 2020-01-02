-- define stored procedures
/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$

CREATE PROCEDURE spSelectNotifications
(
	IN UserID VARCHAR(12),
  IN AdminOnly INT(1),
  IN ExcludeRead INT(1)
)
BEGIN
    -- if admin only, we show all the categories
    IF ( AdminOnly = 1 ) THEN

      IF ( ExcludeRead = 1 ) THEN
        SELECT 
          n.*
        FROM
          `dongfeng_zhonglou`.`notifications` n      
        WHERE n.user_id in (UserID,  'AgentSmith')
        AND n.read = 0
        ORDER BY n.user_id ASC, n.created_at DESC ;
      ELSE
        SELECT 
          n.*
        FROM
          `dongfeng_zhonglou`.`notifications` n      
        WHERE n.user_id in (UserID,  'AgentSmith')        
        ORDER BY n.user_id ASC, n.created_at DESC ;
      END IF;

    ELSE

      IF ( ExcludeRead = 1 ) THEN
        SELECT 
          n.*
        FROM
          `dongfeng_zhonglou`.`notifications` n
        JOIN
          `dongfeng_zhonglou`.`categories` c
        ON n.category_id = c.id
        WHERE n.user_id in (UserID,  'AgentSmith') 
        AND c.admin_only = 0
        AND n.read = 0
        ORDER BY n.user_id ASC, n.created_at DESC;
      ELSE
        SELECT 
          n.*
        FROM
          `dongfeng_zhonglou`.`notifications` n
        JOIN
          `dongfeng_zhonglou`.`categories` c
        ON n.category_id = c.id
        WHERE n.user_id in (UserID,  'AgentSmith') 
        AND c.admin_only = 0        
        ORDER BY n.user_id ASC, n.created_at DESC;
      END IF;

    END IF;
END$$

DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeleteRecipes
(
    IN RecipeName VARCHAR(50)
)
BEGIN
    DELETE FROM `dongfeng_zhonglou`.`recipes` WHERE name = RecipeName;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeletePupils
(
    IN Year VARCHAR(10),
    IN Cls VARCHAR(10)
)
BEGIN
    DELETE FROM `dongfeng_zhonglou`.`pupils` WHERE year = Year AND class = Cls;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeleteTeachers
(
    IN Year VARCHAR(10)
)
BEGIN
    DELETE FROM `dongfeng_zhonglou`.`teachers` WHERE year = Year;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeleteClasses()
BEGIN
    TRUNCATE TABLE `dongfeng_zhonglou`.`classes`;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeleteAbsence(
    IN Year VARCHAR(10),
    IN Cls VARCHAR(10),
    IN AbsenceDate VARCHAR(10),
    IN Name VARCHAR(10)
)
BEGIN
    DELETE FROM `dongfeng_zhonglou`.`absences` WHERE year = Year AND class = Cls AND date = AbsenceDate AND name = Name;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeleteAbsences(
    IN Year VARCHAR(10),
    IN Cls VARCHAR(10),
    IN AbsenceDate VARCHAR(10)
)
BEGIN
    DELETE FROM `dongfeng_zhonglou`.`absences` WHERE year = Year AND class = Cls AND date = AbsenceDate;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeletePhysiques
(
    IN Year VARCHAR(10),
    IN Cls VARCHAR(10)
)
BEGIN
    DELETE FROM `dongfeng_zhonglou`.`physiques` WHERE year = Year AND class = Cls;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

-- show
SHOW PROCEDURE STATUS WHERE db = 'dongfeng_zhonglou';