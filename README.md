# Golang + Redis

Run Go 
- go run . / go run main.go

Check go API
- curl localhost:8000/products
- curl localhost:8000/products/redis

Run Mariadb & Redis & Grafana & Influxdb
- docker-compose up mariadb redis grafana influxdb

Watch Grafana Dashboard
- Open Browser localhost:3000 (Grafana Dashboard)
- Click Dashboard
- Click k6 Load Testing Results

Run Load test Database
- docker-compose run --rm k6 run /script/test.js

Run Load test Redis
- docker-compose run --rm k6 run /script/test.js

