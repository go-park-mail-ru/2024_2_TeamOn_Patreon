### Start up
```bash
cd cmd/microservice
sudo docker-compose up
```

### Info
| **Services**           | **Info**                                                         | **Port** |
|------------------------|------------------------------------------------------------------|----------|
| **Auth**               | Сервер на http и на grpc принимает запросы на верификацию токена | 8081     |
| **Account**            |                                                                  | 8082     |
| **Author**             |                                                                  | 8083     |
| **Posts** (content)    |                                                                  | 8084     |
| **CustomSubscriptiom** |                                                                  | 8085     |
| **CSAT**               | Работает со своей БД                                             | 8086     |
| **Moderation**         |                                                                  | 8087     |

