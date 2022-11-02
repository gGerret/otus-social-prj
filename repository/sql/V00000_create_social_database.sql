-- drop schema if exists social;

create schema if not exists social;

create table if not exists social.user
(
    id bigint primary key auto_increment comment 'внутренний сквозной идентификатор',
    public_id varchar(64) unique not null comment 'публичный уникальный идентификатор',
    pass_hash varchar(255) comment 'хеш пароля пользователя',
    email varchar(255) unique not null comment 'email польщователя',
    first_name varchar(32) not null comment 'имя',
    last_name varchar(32) comment 'фамилия',
    middle_name varchar(32) comment 'отчество',
    gender int not null comment 'id пола из таблицы social.gender',
    town varchar(50) comment 'город',
    created_at timestamp not null comment 'время создания записи',
    updated_at timestamp comment 'время последнего обновления записи',
    deleted_at timestamp comment 'время удаления записи, если установлено, считаем, что пользователь удалён'

) comment 'Таблица с информацией о пользователях';

create table if not exists social.gender
(
    id int primary key auto_increment comment 'внутренний сквозной идентификатор',
    code varchar(10) unique not null comment 'внутренний код пола',
    short_desc varchar(10) not null comment 'краткое обозначения пола',
    full_desc varchar(10) not null comment 'полное обозначения пола'
) comment 'Таблица - справочник полов';

insert ignore into social.gender values (1, 'NOT', 'Не скажу', 'Не скажу');
insert ignore into social.gender values (2, 'MALE', 'муж.','мужской');
insert ignore into social.gender values (3, 'FEMALE', 'жен.', 'женский');


create table if not exists social.interests
(
    id bigint primary key auto_increment comment 'внутренний сквозной идентификатор',
    interest varchar(255) not null unique comment 'Интерес - строка с текстом, уникальная в рамках системы',
    created_at timestamp not null comment 'время создания'
) comment 'Таблица уникальных интересов';

create table if not exists social.user_interests_link
(
    user_id bigint not null comment 'Ссылка на social.user.id',
    interest_id bigint not null comment 'Ссылка на social.interests.id',
    primary key (user_id, interest_id),

    created_at timestamp not null comment 'время создания'
) comment 'Таблица связей пользователей и их интересов';