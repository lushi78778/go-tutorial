# Dockerfile


# Build stage
FROM registry.jetbrains.team/p/writerside/builder/writerside-builder:232.10275 as build
WORKDIR /app
COPY . .
RUN /opt/builder/bin/idea.sh helpbuilderinspect -source-dir . -product Writerside/go-tutorial  -output-dir public/ || true

# Deploy stage
FROM ubuntu:20.04 as deploy
WORKDIR /app
RUN apt-get update -y && apt-get install -y unzip nginx
COPY --from=build /app/public /usr/share/nginx/html

# Start Nginx as the default command
CMD ["nginx", "-g", "daemon off;"]
