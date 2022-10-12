-- Получаем все пользовательские интересы
select i.interest from social.user_interests_link uil
    join social.interests i on uil.interest_id = i.id
where uil.user_id = 2;

