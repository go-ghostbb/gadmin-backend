create table user (
    `id` int unsigned not null auto_increment comment '主鍵',
    `created_at` datetime default null comment '創建時間',
    `updated_at` datetime default null comment '更新時間',
    `username` nvarchar(191) not null unique comment '使用者名稱',
    `password` nvarchar(191) not null comment '密碼',
    primary key (`id`)
)