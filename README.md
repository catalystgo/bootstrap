# bootstrap 💨

💨 A bootstrap library to initialize external service's clients without boilerplate 💨

Supported clients:

- Kafka
- Postgres
- Redis

---

### Kafka 🚀

```bash
go get github.com/catalystgo/bootstrap/kafka
```

example: **[code](./kafka/example)**

---

### Postgres 🚀

```bash
go get github.com/catalystgo/bootstrap/postgres
```

example: **[code](./postgres/example)**

---

### Redis 🚀

```bash
go get github.com/catalystgo/bootstrap/redis
```

example: **[code](./redis/example)**

---

### Cache 🚀

```bash
go get github.com/catalystgo/bootstrap/cache
```

example: **[code](./cache/example)**

---

### UI 😎

To see the UI, run the following command:

```bash
docker compose up
```

- [Kafka](http://localhost:8084)
- [RabbitMQ](http://localhost:15672)
- [MinIO](http://localhost:9001)
- [CockroachDB](http://localhost:8080)

## Milestones 🚀

- [ ] add memcached client support
- [ ] add pgbouncer client support
