
////////////////////////////////////////////////////////////////
ติดตั้ง migrations
https://www.youtube.com/watch?v=G5xBzBKuBvM
https://github.com/golang-migrate/migrate/releases



1. export POSTGRESQL_URL='postgres://username:password@localhost:5432/example?sslmode=disable'
2. migrate create -ext sql -dir db/migrations -seq create_users_table //สร้างfile migrate ต้องเข้ามาใน folder backend ก่อน
3. migrate -database ${POSTGRESQL_URL} -path ./ up  เข้าไปอยู่ใน folder file migrate
4. migrate -database ${POSTGRESQL_URL} -path ./ down  เข้าไปอยู่ใน folder file migrate
--> ต้อง DROP TABLE schema_migrations

