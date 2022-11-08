drop table if exists social.user_page;

create table if not exists social.user_page (
    id bigint primary key auto_increment comment 'внутренний сквозной идентификатор',
    public_id varchar(64) unique not null comment 'публичный уникальный идентификатор',
    user_id bigint not null comment 'Ссылка на social.user.id',
    image_link text comment 'Ссылка на картинку пользователя',
    page_text text comment 'Текст, написанный пользователем о себе',
    created_at timestamp not null comment 'время создания записи',
    updated_at timestamp comment 'время последнего обновления записи',
    deleted_at timestamp comment 'время удаления записи, если установлено, считаем, что страница удалена'
) comment 'Таблица для персональных страниц пользователей';