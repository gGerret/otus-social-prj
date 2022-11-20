# otus-social-prj
Социальная сеть, разработанная на курсе Отус Highload Architect

## Запуск
* Колнировать репозиторий
  ```
  git clone https://github.com/gGerret/otus-social-prj && cd ./otus-social-prj && go mod tidy
  cp ./bin/config.debug.json ./config.json
  ```
* В файле `./config.json` изменить параметры подключения к БД
  ```
  "db": {
    "username": "social_svc",
    "password": "social_sql_passw0rd",
    "database": "social", //не менять
    "hostname": "localhost",
    "port": 3306,
    "net": "tcp",
    "ssl_mode": false
  }
  ```
* В БД MySQL создать базу данных (схему) с именем `social`. 
* Выполнить в созданной БД скрипт `./sql/V00000_create_social_database.sql` для создания структуры БД  
* Выполнить скрипт `./sql/test/test_fill_db.sql` для наполнения таблиц тестовыми данными 
* Запустить сервис
  ```
  go run main.go
  ```

## Тестирование
[Коллекция Postman](https://github.com/gGerret/otus-social-prj/blob/main/test/api/Gerret's%20Social%20Project%20(Otus%20HLA).postman_collection.json) для тестирования API 