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
          `dongfeng_core`.`notifications` n      
        WHERE n.user_id in (UserID,  'AgentSmith')
        AND n.read = 0
        ORDER BY n.user_id ASC, n.created_at DESC ;
      ELSE
        SELECT 
          n.*
        FROM
          `dongfeng_core`.`notifications` n      
        WHERE n.user_id in (UserID,  'AgentSmith')        
        ORDER BY n.user_id ASC, n.created_at DESC ;
      END IF;

    ELSE

      IF ( ExcludeRead = 1 ) THEN
        SELECT 
          n.*
        FROM
          `dongfeng_core`.`notifications` n
        JOIN
          `dongfeng_core`.`categories` c
        ON n.category_id = c.id
        WHERE n.user_id in (UserID,  'AgentSmith') 
        AND c.admin_only = 0
        AND n.read = 0
        ORDER BY n.user_id ASC, n.created_at DESC;
      ELSE
        SELECT 
          n.*
        FROM
          `dongfeng_core`.`notifications` n
        JOIN
          `dongfeng_core`.`categories` c
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
    DELETE FROM `dongfeng_core`.`recipes` WHERE name = RecipeName;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeletePupils
(
    IN Year VARCHAR(10),
    IN Class VARCHAR(10)
)
BEGIN
    DELETE FROM `dongfeng_core`.`pupils` WHERE year = Year AND class = Class;
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
    DELETE FROM `dongfeng_core`.`teachers` WHERE year = Year;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeleteClasses()
BEGIN
    TRUNCATE TABLE `dongfeng_core`.`classes`;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$
CREATE PROCEDURE spDeleteAbsences(
    IN Year VARCHAR(10),
    IN Class VARCHAR(10),
    IN Date VARCHAR(10)
)
BEGIN
    DELETE FROM `dongfeng_core`.`absences` WHERE year = Year AND class = Class AND date = Date;
END$$
DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */

-- show
SHOW PROCEDURE STATUS WHERE db = 'dongfeng_core';