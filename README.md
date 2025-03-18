# bootstrap 💨

[![wakatime](https://wakatime.com/badge/user/965e81db-2a88-4564-b236-537c4a901130/project/5acc4f35-6871-45b9-8af6-1d5e246b808f.svg)](https://wakatime.com/badge/user/965e81db-2a88-4564-b236-537c4a901130/project/5acc4f35-6871-45b9-8af6-1d5e246b808f)
![Build Status](https://github.com/catalystgo/bootstrap/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/catalystgo/bootstrap)](https://goreportcard.com/report/github.com/catalystgo/bootstrap)

[![GitHub issues](https://img.shields.io/github/issues/catalystgo/bootstrap.svg)](https://github.com/catalystgo/bootstrap/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/catalystgo/bootstrap.svg)](https://github.com/catalystgo/bootstrap/pulls)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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
