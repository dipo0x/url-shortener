# Go-fiber URL shortener

### Introduction

This is a RESTful API built using Golang, Fiber, Postgresql, RabbitMQ and Reflex. This API provides a simple implementation for scheduling jobs in Go.
</br>

### Setup

Clone the repository to your local machine.

```bash
git clone https://github.com/dipo0x/url-shortener
```

Ensure that you have Golang, PostgreSQL and RabbitMQ installed on your machine. Alternatively, you can use their cloud service.

Navigate to the root directory of the project in a terminal.

```bash
cd url-shortener
```

Run the following command to install the necessary dependencies

```bash
go install
```

After that, run this command to create a .env file with which you can get started.

```bash
bash build/scripts/setup.sh
```
After that, run this command to install migrate on your cmd so that you can create a migration for your models

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
Once the installation is successful, run this command to run the create table migration
```
make migrate-up
```

### Running Server

#### Locally

Run the following command to start the server and start up rabbitmq instance:

```bash
reflex -c .reflex
go run cmd/worker/main.go
```
For clarity sake, the first command spins up your development server while the second one spins up the RabbitMQ receiver. When you run the first command, you will see something similar to this:
<img width="425" height="196" alt="Screenshot 2025-10-08 at 12 39 14 AM" src="https://github.com/user-attachments/assets/e0958862-dbeb-486b-868b-17e1c668f2ed" />


When you run the second command, you will see something similar to this:
</br>

<img width="444" height="197" alt="Screenshot 2025-10-08 at 12 39 20 AM" src="https://github.com/user-attachments/assets/14dbb222-fb0f-425f-99ee-aaba55c8dd73" />
</br>
The server will run on http://localhost:8080 by default

</br>

### Test
To test the endpoints, run this command :
```bash
go test ./tests/
```
If all your tests are successful, you will see this: 
</br>
<img width="312" height="59" alt="Screenshot 2025-10-08 at 12 43 32 AM" src="https://github.com/user-attachments/assets/22e99388-db8d-452c-b59e-bc21d3792664" />

</br>
else, you will see something similar to this:
</br>
<img width="418" alt="Screenshot 2025-05-11 at 5 47 01 PM" src="https://github.com/user-attachments/assets/63ccc3cc-cf5b-46f5-a383-c02c1d8c9e24" />


## Available Endpoints

Base URL[dev]: 0.0.0.0:8080/\

When your server is running, call the base endpoint to ensure it is up, and you will receive a response like this:


<img width="926" alt="Screenshot 2025-05-10 at 11 20 43 PM" src="https://github.com/user-attachments/assets/e5ecd909-9dbe-48a0-956f-a851b85cb2e8" />


### Conclusion

You can find additional documentation for this API, including request and response signatures, by visiting https://documenter.getpostman.com/view/17975360/2sB2j999pk in your web browser.
