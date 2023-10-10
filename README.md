# Golang - CI/CD Learn

### connect server

- sudo ssh -i "compute-service-key.pem" ec2-user@ec2-18-215-244-57.compute-1.amazonaws.com

### run container in server

- docker run -d -p 1234:1234 --name compute-service-app agungbhaskara/lms-belajar-docker:1.4
