# prometheus
prometheus monitoring demo

#Docker to run the Prometheus server
docker run --name prometheus -d -p 9090:9090 -v prometheus:/etc/prometheus prom/prometheus:latest
