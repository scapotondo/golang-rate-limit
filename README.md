<p align="center">
    <h1 align="center">Rate-Limited Notification Service</h1>
    <p align="center"><a href="#features">Features</a> section describes in detail about service capabilities</p>
    </p>
<p align="center">

## Features

Send emails of different type with a rate limit associated per user and per type:

* Status type: not more than 2 per minute for each recipient
* News type: not more than 1 per day for each recipient
* Marketing type: not more than 1 per day for each recipient

## Local testing

### Run tests with coverage
There's a coverage.sh file that runs every test in the project located at `bin/coverage.sh` and needs permissions to run. It can be easily done executing the following command:

```
chmod +x coverage.sh
```

Now you can execute the file in the terminal

## Open API Specification

To update the Open API specification, follow these steps:
1. Write annotations in the code using the [Swag API operation](https://github.com/swaggo/swag#api-operation) format

2. install go swag locally. Follow swaggo's [getting started](https://github.com/swaggo/swag#getting-started)

3. Run the following command to generate/update the Open API specification:
    ```
    swag init
    ```

4. Run the golang-rate-limit service and navigate to the following URL to view the Open API Specification:
    ```
    http://localhost:8080/v1/docs/index.html
    ```

You will see the OAS with Swagger UI


## Creator

[Sebastian Capotondo](https://github.com/scapotondo) 
* Email: sebastian.capotondo@gmail.com