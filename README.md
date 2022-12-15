# ddd-to-do-list


Curl API

<!-- Get Activity -->
curl --location --request GET 'http://127.0.0.1:3030/activity-groups'
<!-- Get Activity By ID -->
curl --location --request GET 'http://127.0.0.1:3030/activity-groups/1'
<!-- Create Activity -->
curl --location --request POST 'http://localhost:3030/activity-groups' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Test 1",
    "email": "bagus@gmail.com"
}'
<!-- Update Activity -->
curl --location --request PUT 'http://127.0.0.1:3030/activity-groups/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "Bagus2@gmail.com",
    "title" : "Judul2"
}'
<!-- Delete Activity -->
curl --location --request DELETE 'http://127.0.0.1:3030/activity-groups/1'



<!-- Get Todo -->
curl --location --request GET 'http://127.0.0.1:3030/todo-items'
<!-- Get Todo By ID -->
curl --location --request GET 'http://127.0.0.1:3030/todo-items/1'
<!-- Create Todo -->
curl --location --request POST 'http://127.0.0.1:3030/todo-items' \
--header 'Content-Type: application/json' \
--data-raw '{
    "activity_group_id": 1,
    "title": "Todo 1",
    "priority" : "very_high"
}'
<!-- Update Todo -->
curl --location --request PUT 'http://127.0.0.1:3030/todo-items/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "activity_group_id" : 1,
    "title" : "Judul AKU",
    "is_active" : 1,
    "priority" : "low"
}'
<!-- Delete Todo -->
curl --location --request PUT 'http://127.0.0.1:3030/todo-items/1'
