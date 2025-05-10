# Go-fiber URL shortener

### Introduction

This is a RESTful API built using Golang, Fiber, MongoDB, Redis, Async and Reflex. This API provides a simple implementation to scheduling jobs in golang.
</br>

### Setup

Clone the repository to your local machine.

```bash
git clone https://github.com/dipo0x/golang-url-shortener
```

Ensure that you have Golang, Redis and MongoDB installed on your machine. Alternatively, you can use MongoDB Cloud Atlas and Redis cloud service

Navigate to the root directory of the project in a terminal.

```bash
cd golang-url-shortener
```

Run the following command to install the necessary dependencies

```bash
go install
```

After that, run this command to create a .env file with which youcan get started with.

```bash
bash setup.sh
```

</br>

### Running Server

#### Locally

Run the following command to start the server:

```bash
reflex -c .reflex
```

Run the following command to view your redis jobs details on a dashboard:

```bash
asynqmon --redis-addr=localhost:6379
```
<img width="612" alt="Screenshot 2025-05-10 at 11 17 00 PM" src="https://github.com/user-attachments/assets/56d6a976-d511-459a-abc4-919605b575db" />

and when you run your asynqmon start command, you will see this:
<img width="470" alt="Screenshot 2025-05-10 at 11 19 39 PM" src="https://github.com/user-attachments/assets/b430595e-83a0-4d78-a541-f9c3aa7de20f" />



The server will run on http://localhost:8080 by default

</br>

## Available Endpoints

Base URL[dev]: 0.0.0.0:8080/\

When your server is running, call the base endpoint to ensure it is up, and you will receive a response like this:


<img width="926" alt="Screenshot 2025-05-10 at 11 20 43 PM" src="https://github.com/user-attachments/assets/e5ecd909-9dbe-48a0-956f-a851b85cb2e8" />


### Conclusion

You can find additional documentation for this API, including request and response signatures, by visiting https://documenter.getpostman.com/view/17975360/2sB2j999pk in your web browser.
