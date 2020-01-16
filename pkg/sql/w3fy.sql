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

-- 帖子表
drop table if exists `topic`;
create table `topic`(
    `id`       int(11)             not null auto_increment comment '主键',
    `tag`      char(10)            not null default '其他' comment '帖子标签',
    `uid`      int(11)             not null comment '发帖人',
    `title`    char(120)           not null comment '帖子标题',
    `content`  text           not null  comment '帖子内容',
    `is_deleted` tinyint(11)   not null default 0 comment '逻辑删除帖子,0-否,1-是',
    `created_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    primary key (`id`),
    constraint i2u_id_fk foreign key (`uid`) references `users`(`id`)
)engine = innodb auto_increment=1 default charset =utf8 comment ='帖子表';

-- 评论表
drop table if exists `comment`;
create table `comment`(
    `id`         int(11)           not null auto_increment comment '主键',
    `top_id`     int(11)           not null comment '帖子id',
    `father_id`  int(11)           not null default 0 comment '层主id,默认为0给题主评论',
    `from_id`    int(11)           not null comment '评论人',
    `to_id`      int(11)           not null comment '被评论人,默认为给题主评论',
    `comments`   text              not null comment '评论内容',
    `is_deleted` int(11)          not null comment '逻辑删除帖子,0-否,1-是',
    `created_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    primary key (`id`),
    constraint  c2u_id_fk foreign key (`top_id`) references `topic`(`id`)
)engine = innodb auto_increment=1 default charset =utf8 comment ='评论表';

-- 粉丝/关注表
drop table if exists `relation`;
create table `relation`(
    `id`         int(11)          not null auto_increment comment '主键',
    `from_id`    int(11)          not null comment '关注人',
    `to_id`      int(11)          not null comment '被关注人',
    `created_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    primary key (`id`),
    key (`from_id`),
    key (`to_id`)
)engine = innodb auto_increment=1 default charset =utf8 comment '粉丝/关注表';

-- 节点表
drop table if exists `tags`;
create table `tags`(
    `id`               int(11)          not null auto_increment comment '主键',
    `name`             char(20)         not null comment '标签名字',
    `created_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    primary key (`id`),
    unique (`name`)
)engine = innodb auto_increment=1 default charset =utf8 comment '节点表';

-- 节点收藏表
drop table if exists `taglikes`;
create table `taglikes`(
    `id`         int(11)           not null auto_increment comment '主键',
    `uid`        int(11)           not null comment '用户id',
    `tname`      char(20)          not null comment '节点名',
    `created_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    primary key (`id`),
    constraint t2t_id_fk foreign key (`tname`) references `tags`(`name`)
)engine = innodb auto_increment=1 default charset =utf8 comment '节点收藏表';

-- 帖子收藏表
drop table if exists `topiclikes`;
create table `topiclikes`(
    `id`                int(11)         not null auto_increment comment '主键',
    `uid`               int(11)         not null comment '用户id',
    `tid`               int(11)         not null comment '帖子id',
    `updated_at`   timestamp      NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    primary key (`id`),
    key (`id`),
    constraint l2t_id_fk foreign key (`tid`) references `topic`(`id`)
)engine = innodb auto_increment=1 default charset = utf8 comment '帖子收藏表';

-- 管理员表
drop table if exists `admins`;
create table `admins`(
    `id`              int(11)           not null auto_increment comment '主键',
    `nickname`       char(80)     NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar`         varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `username`       char(18)     NOT NULL DEFAULT '' COMMENT 'Username',
    `password`       varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
    `created_at` timestamp    NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '写入时间',
    `updated_at` timestamp    NOT NULL DEFAULT '0000-00-00 00:00:00'
        COMMENT '更新时间',
    `deleted_at` timestamp    NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
)engine = innodb auto_increment=1 default charset =utf8 comment '管理员表';

-- 插入管理员账户
insert into admins (`nickname`,`avatar`,`username`,`password`)values ("alien","https://cdn.v2ex.com/avatar/6097/2a28/459612_large.png?m=1576816395","alien","root");

