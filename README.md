# go-register-mongodb

вы можете протестить данный код через postman, mongoDB кластер подключен в коде, так что можете не создавать свой кластер, а потестить через мой.
в коде идет проверка на то, есть ли пустые поля при регистрации, а также на повторение email (две одинаковые email сущности создать нельзя)

 регистрация: localhost:8080/register/ -Method Post
 
 как должен выглядеть JSON:
 {
 
    "nickName": "danik",
    "email": "danik@gmail.com",
    "password": "234"

}
 
 получение User по id: localhost:8080/users/id  - method Get 
 
 как должен выглядеть ответ:
 {
 
    "_id": "63c28009f33cae51c05d8152",   
    "nickName": "danua",
    "email": "sergei@gmail.com",
    "password": "2313"

}
