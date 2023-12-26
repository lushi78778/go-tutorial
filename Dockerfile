# Dockerfile


# Build stage
FROM registry.jetbrains.team/p/writerside/builder/writerside-builder:232.10275 as build
WORKDIR /app
ENV DISPLAY=:99
RUN Xvfb :99 &
COPY . .
RUN /opt/builder/bin/idea.sh helpbuilderinspect -source-dir . -product Writerside/go-tutorial --runner gitlab -output-dir public/ || true
RUN test -e public/webHelpGO-TUTORIAL2-all.zip

# Test stage
FROM openjdk:18-jdk-alpine as test
WORKDIR /app
RUN apk add --no-cache curl
COPY --from=build /app/public /app/public
RUN curl -o wrs-checker.jar -L https://packages.jetbrains.team/maven/p/writerside/maven/com/jetbrains/writerside/writerside-ci-checker/1.0/writerside-ci-checker-1.0.jar && java -jar wrs-checker.jar report.json Writerside/go-tutorial

# Deploy stage
FROM ubuntu:20.04 as deploy
WORKDIR /app
RUN apt-get update -y && apt-get install -y unzip nginx
COPY --from=test /app/public /usr/share/nginx/html
# You may need to configure Nginx based on your specific requirements, such as adding a custom configuration file or setting up SSL

# Start Nginx as the default command
CMD ["nginx", "-g", "daemon off;"]
