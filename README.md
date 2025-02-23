# Contact:
- **Mail**: *kiennv.contact@gmail.com*
- **Github**: *https://github.com/nguyenkien0703*

# Go lang project ecomerce:

## Pkg Documentation Usage:
- Handle http server: [gin](https://github.com/gin-gonic/gin)
- Handle manager configuration: [vipper](https://github.com/spf13/viper)
- Handle write logger:
    + [zap](https://github.com/uber-go/zap)
    + [lumberjack](https://github.com/natefinch/lumberjack) -  Manger file logging (size file, max age file, max backup file, ...)
- Handle database:
    + [gorm](https://github.com/go-gorm/gorm)
    + [sqlc](https://github.com/sqlc-dev/sqlc) - Fast generate code for SQL and handle ( gen model, exec query, ...)
    + [goose](https://github.com/pressly/goose) - Manage version of database - migration
    + [uuid google](https://github.com/google/uuid) - Google UUID for generate UUID
- Handle authentication:

- Handle middleware:

- Handle cache:
    + [redis](https://github.com/redis/go-redis)

- Handle message queue:
    + [kafka](https://github.com/segmentio/kafka-go)
        + [zookeeper](https://github.com/bitnami/containers)
        + [kafka-ui](https://github.com/provectus/kafka-ui)

- Handle dependencies injection wiht [wire](https://github.com/google/wire)

- Handle change data capture (CDC) with [debezium](https://debezium.io/)

- Handle send mail:
    + [sendgrid](https://sendgrid.com)
    + [smtp](https://pkg.go.dev/net/smtp)

- Handle monitoring:
    + [prometheus](https://prometheus.io/)
    + [grafana](https://grafana.com/)
        + [Go Processes](https://grafana.com/grafana/dashboards/6671-go-processes/)
        + [node_exporter](https://grafana.com/grafana/dashboards/1860-node-exporter-full/)
        + [oliver006/redis_exporter](https://grafana.com/grafana/dashboards/763)
        + [mysql_exporter](https://grafana.com/grafana/dashboards/14057-mysql/)
        + [kafka_exporter](https://grafana.com/grafana/dashboards/18276-kafka-dashboard/)
    + [node_exporter](https://www.google.com/search?q=node-exporter)