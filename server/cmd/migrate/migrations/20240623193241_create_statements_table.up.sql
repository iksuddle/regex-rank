create table if not exists statements (
    `id` int not null auto_increment,
    `problem_id` int not null,
    `match` char(1) not null,
    `literal` varchar(255) not null,

    primary key (`id`),
    foreign key (`problem_id`) references problems(`id`)
);
