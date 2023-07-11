# go-db-practice-gonews

### Контейнер для постгрес на основе dockerfile
```bash
    docker build -t my-postgres .
    docker run -d --name my-postgres-container -p 5432:5432 my-postgres
```
### Контейнер для монго
```bash
    docker run -d --name mongo -p 27017:52017 mongo:4.6.6
```
