USE dongfeng_core;

INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (1, '幼儿成长档案', 0);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (2, '幼儿体格检查表', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (3, '仓库管理', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (4, '学校资产管理', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (5, '营养膳食', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (6, '出勤管理', 1);
INSERT INTO categories(`id`, `description`, `admin_only`) VALUES (7, '系统通知', 0);

insert into holidays (`from`, `to`, `description`) values ('2018-01-01', '2018-01-01', '元旦');
insert into holidays (`from`, `to`, `description`) values ('2018-01-26', '2018-02-22', '寒假');
insert into holidays (`from`, `to`, `description`) values ('2018-02-15', '2018-02-21', '春节');
insert into holidays (`from`, `to`, `description`) values ('2018-04-05', '2018-04-07', '清明节');
insert into holidays (`from`, `to`, `description`) values ('2018-04-29', '2018-05-01', '劳动节');
insert into holidays (`from`, `to`, `description`) values ('2018-06-16', '2018-06-18', '端午节');
insert into holidays (`from`, `to`, `description`) values ('2018-07-01', '2018-08-31', '暑假');
insert into holidays (`from`, `to`, `description`) values ('2018-09-22', '2018-09-24', '中秋节');
insert into holidays (`from`, `to`, `description`) values ('2018-10-01', '2018-10-07', '国庆节');
insert into holidays (`from`, `to`, `description`) values ('2018-12-30', '2018-12-31', '元旦');

insert into holidays (`from`, `to`, `description`) values ('2019-01-01', '2019-01-01', '元旦');
insert into holidays (`from`, `to`, `description`) values ('2019-01-24', '2019-02-19', '寒假');
insert into holidays (`from`, `to`, `description`) values ('2019-02-04', '2019-02-10', '春节');
insert into holidays (`from`, `to`, `description`) values ('2019-04-05', '2019-04-07', '清明节');
insert into holidays (`from`, `to`, `description`) values ('2019-05-01', '2019-05-04', '劳动节');
insert into holidays (`from`, `to`, `description`) values ('2019-06-07', '2019-06-09', '端午节');
insert into holidays (`from`, `to`, `description`) values ('2019-07-01', '2019-08-31', '暑假');
insert into holidays (`from`, `to`, `description`) values ('2019-09-13', '2019-09-15', '中秋节');
insert into holidays (`from`, `to`, `description`) values ('2019-10-01', '2019-10-07', '国庆节');
insert into holidays (`from`, `to`, `description`) values ('2019-12-30', '2019-12-31', '元旦');
