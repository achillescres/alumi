Ну че
golang, ent/postgresql
docker run -e ... ваши данные postgres
Написать данные в .env

docker build --tag alu . -f Dockerfile
docker run --name alu -d --net=host alu 
PS host если нужно на локалке к бд подлкючаться
