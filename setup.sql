-- user
drop table if exists user;
create table user (
    `id` int unsigned not null auto_increment comment '主鍵',
    `created_at` datetime default null comment '創建時間',
    `updated_at` datetime default null comment '更新時間',
    `username` nvarchar(191) not null unique comment '使用者名稱',
    `password` nvarchar(191) not null comment '密碼',
    primary key (`id`)
);
insert into user (id, created_at, updated_at, username, password)
values (1, now(), now(), 'admin', '$2a$10$3zL.F5U8QrWPy6y4Z.WIouEwTZwBtyTA9la5EFzrEaUMimi5ngQc2');

-- role
drop table if exists role;
create table role (
    `id` int unsigned not null auto_increment comment '主鍵',
    `created_at` datetime default null comment '創建時間',
    `updated_at` datetime default null comment '更新時間',
    `code` nvarchar(30) not null unique comment '角色代碼',
    `name` nvarchar(30) not null comment '角色名稱',
    `description` nvarchar(191) comment '說明',
    primary key (`id`)
);
insert into role (id, created_at, updated_at, code, name, description)
values (1, now(), now(), 'root', 'root', '請勿刪除');

-- user_role
drop table if exists user_role;
create table user_role (
    user_id int not null comment 'user id',
    role_id int not null comment 'role id',
    primary key (user_id, role_id)
);
insert into user_role (user_id, role_id)
values (1, 1);
