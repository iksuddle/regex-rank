create table if not exists users (
    `id` int not null,
    `username` varchar(40) not null,
    `avatar_url` varchar(255) not null,
    `created_at` datetime not null,

    primary key (`id`)
);
