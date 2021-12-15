podman build -t zpeters/weather_exporter .
podman login docker.io
podman push zpeters/weather_exporter
