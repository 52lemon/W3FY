-- 用户表
drop table if exists `users`;
create table `users`(
    `id`          int(11)            not null auto_increment comment '主键',
    `username`    char(10)           not null comment '用户登录用户名',
    `nickname`    char(10)           not null default '' comment '用户名称',
    `password`    varchar(255)       not null comment '用户密码',
    `sex`         tinyint(1)         not null default 0 comment '性别,0-未知 1-男 2-女',
    `email`       char(100)          not null default '' comment 'email',
    `website`     varchar(255)       not null default '' comment '个人网站',
    `education`   enum("小学","初中","中专","高中","大专","本科","硕士","博士","其他") not null default '其他',
    `collage`     char(100)          not null default '' comment '毕业院校',
    `introduction` char(100)         not null default '' comment '个人简介',
    `github`      varchar(255)       not null default '' comment 'github地址',
    `avatar`      varchar(255)       not null default '' comment '用户头像',
    `coin`        int(11)            not null default 0 comment '平台币',
    `created_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    unique (`username`),
    unique (`email`)
)engine =innodb AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8 COMMENT ='用户表';