create table if not exists solutions (
    `id` int not null,
    `user_id` int not null,
    `problem_id` int not null,
    `literal` varchar(255) not null,

    primary key (`id`),
    foreign key (`user_id`) references users(`id`),
    foreign key (`problem_id`) references problems(`id`)
);
