USE dongfeng_core;

INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (1, '幼儿成长档案', 0);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (2, '幼儿体格检查表', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (3, '仓库管理', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (4, '学校资产管理', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (5, '营养膳食', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (6, '出勤管理', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (7, '系统通知', 0);


INSERT INTO classes(`id`, `name`) VALUES ('01', '小一班');
INSERT INTO classes(`id`, `name`) VALUES ('02', '小二班');
INSERT INTO classes(`id`, `name`) VALUES ('03', '小三班');
INSERT INTO classes(`id`, `name`) VALUES ('04', '小四班');
INSERT INTO classes(`id`, `name`) VALUES ('11', '中一班');
INSERT INTO classes(`id`, `name`) VALUES ('12', '中二班');
INSERT INTO classes(`id`, `name`) VALUES ('13', '中三班');
INSERT INTO classes(`id`, `name`) VALUES ('21', '大一班');
INSERT INTO classes(`id`, `name`) VALUES ('22', '大二班');
INSERT INTO classes(`id`, `name`) VALUES ('23', '大三班');
