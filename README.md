# Gin MongoDB

### Database Environment Setup

You can follow [these](https://www.mongodb.com/basics/create-database) steps to create the database and user account
password.

1. Rename `.env.sample` to `.env`.
2. Replace below values with yours.
    - `<USER_NAME>`
    - `<PASSWORD>`
        - You have to use password from  *Security > Database Access*
          ![Image](https://i.ibb.co/YXDVXmc/Screen-Shot-2022-05-26-at-10-32-00-AM.png)
    - `<MONGO_CLUSTER>`
    - `<DB_NAME>`


### Run Project

Paste the below command to run the project

```
  go run main.go
```

#### API Endpoints

|Method|Endpoints|
|----|----|
|POST | /user | 
|GET | /user/:userId | 
|PUT | /user/:userId | 
|DELETE | /user/:userId |
|GET | /users |
