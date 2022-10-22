insert into social.interests (interest, created_at) value ('Автомобили', now());
insert into social.interests (interest, created_at) value ('Куклы', now());
insert into social.interests (interest, created_at) value ('Сериалы', now());
insert into social.interests (interest, created_at) value ('Кино', now());
insert into social.interests (interest, created_at) value ('Музыка', now());
insert into social.interests (interest, created_at) value ('Программирование', now());

insert into social.user (public_id, pass_hash, email, first_name, last_name, middle_name, gender, town, created_at) values
        ('78758711-b581-46c7-8111-0abab5690d27', '', 'email1@email.ru', 'Вася', 'Васин', 'Романович', 2, 'Тула', now()),
        ('ab9e5ea9-e571-4ffa-ba5f-8530fc3916b6', '', 'email2@email.ru', 'Петя', 'Петин', 'Романович', 2, 'Москва', now()),
        ('53722bba-4030-453f-8578-dc1d3941069c', '', 'email3@email.ru', 'Маша', 'Маишна', 'Романовна', 3, 'Брянск', now()),
        ('7fa0c8e5-483f-44f4-b6ed-84409f0a559d', '', 'email4@email.ru', 'Ваня', 'Васин', 'Романович', 2, 'Сочи', now()),
        ('35533584-e171-46f3-a51b-3484b8bc921e', '', 'email5@email.ru', 'Саша', 'Сашина', 'Романовна', 3, 'Орёл', now()),
        ('b5b3ac59-557d-41c1-af66-6233355b981f', '', 'email6@email.ru', 'Коля', 'Колин', 'Романович', 2, 'Дубна', now()),
        ('fc35215c-d03a-44dc-a90f-b67f78279f49', '', 'email7@email.ru', 'Лена', 'Ленина', 'Романовна', 3, 'Городец', now()),
        ('67d55b4d-cfa8-40e6-8c56-72355136e241', '', 'email8@email.ru', 'Даша', 'Дашина', 'Романовна', 3, 'Ульяновск', now()),
        ('8edb153c-f7ea-4680-9c67-7db14b310faf', '', 'email9@email.ru', 'Катя', 'Катина', 'Романовна', 3, 'Владивосток', now()),
        ('f42da20f-016f-4d37-8058-db080a2a05cd', '', 'email10@email.ru', 'Гоша', 'Гошин', 'Романович', 2, 'Курган', now())
;

insert into social.user_interests_link (user_id, interest_id, created_at) values
        (1,2,now()),
        (1,3,now()),
        (1,5,now()),
        (2,1,now()),
        (2,5,now()),
        (3,2,now()),
        (3,5,now()),
        (3,6,now()),
        (4,6,now()),
        (4,3,now()),
        (5,4,now()),
        (6,2,now()),
        (6,1,now()),
        (6,3,now()),
        (6,6,now()),
        (7,1,now()),
        (7,5,now()),
        (8,6,now()),
        (9,4,now()),
        (9,3,now()),
        (10,2,now()),
        (10,6,now()),
        (10,3,now()),
        (10,4,now())
;
insert into social.user_friendship_link (user_id_a, user_id_b, comment, created_at) values
(1, 8, 'Хочу дружить', now()),
(1, 9, 'Давай дружить', now()),
(3, 5, 'Привет тебе', now()),
(3, 2, 'Привет!', now()),
(3, 8, 'Здорово!', now()),
(4, 5, 'Куку', now());
