# FibonacciWS
Web Service for Calculating Fibonacci Numbers

1. This project provides a RESTful web service.

  a. It accepts a number, n, as input and returns the first n Fibonacci numbers, starting from 0. I.e. given n  = 5, it outputs the sequence [0, 1, 1, 2, 3].

  b. Given a negative number, it responds with an error.

2. It includes instructions to build and deploy the software. When run, it accepts requests and responding to them as appropriate.

3. It includes functional and unit tests.

4. It is written with the Go language.

5. Docker file is in the docker directory

6. To deploy with docker container on local machine:
   https://hub.docker.com/r/kkphk/fibonacci/

7. To deploy with docker container on AWS using Elastic Beanstalk:
   - AWS Elastic Beanstalk, select Create Web Server, Generic Docker environment and launch.
   - Deploy the aws-eb/fibonacci.zip file
   - Click on the URL.
   - Enter a forward slash and the Fibonacci number n at the end of the URL.
   - The Fibonacci numbers will be displayed on the web browser page.
