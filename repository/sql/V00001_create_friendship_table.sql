drop table if exists social.user_friendship_link;

create table if not exists social.user_friendship_link
(
    id bigint primary key auto_increment comment 'внутренний сквозной идентификатор',
    user_id_a bigint not null comment 'идентификатор пользователя запросившего дружбу',
    user_id_b bigint not null comment 'идентификатор пользователя получающего приглашение',

    comment varchar(1024) comment 'комментарий к запросу на дружбу',
    created_at timestamp not null comment 'время создания',
    approved_at timestamp comment 'время подтверждения'
);

create unique index ufl_ua_ub on social.user_friendship_link (user_id_a, user_id_b);
create unique index ufl_ub_ua on social.user_friendship_link (user_id_b, user_id_a);