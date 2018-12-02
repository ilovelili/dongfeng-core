-- define stored procedures
/* -------------------------------------------------------------------------------------------------------------------- */
DELIMITER $$

CREATE PROCEDURE spSelectNotifications
(
	IN UserID VARCHAR(12),
  IN AdminOnly INT(1)
)
BEGIN
    -- if admin only, we show all the categories
    IF ( AdminOnly = 1 ) THEN
      SELECT 
        n.*
      FROM
        `dongfeng_core`.`notifications` n      
      WHERE n.user_id in (UserID,  'AgentSmith')
      ORDER BY n.user_id ASC, n.created_at DESC ;

    ELSE

      SELECT 
        n.*
      FROM
        `dongfeng_core`.`notifications` n
      JOIN
        `dongfeng_core`.`categories` c
      ON n.category_id = c.id
      WHERE n.user_id in (UserID,  'AgentSmith') AND c.admin_only = 0
      ORDER BY n.user_id ASC, n.created_at DESC;

    END IF;
END$$

DELIMITER ;
/* -------------------------------------------------------------------------------------------------------------------- */
