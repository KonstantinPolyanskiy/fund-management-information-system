CREATE TABLE managers (
                          Id SERIAL PRIMARY KEY NOT NULL,
                          Name VARCHAR(50) NOT NULL,
                          Surname VARCHAR(50) NOT NULL,
                          Address VARCHAR(100),
                          Email VARCHAR(100),
                          Phone VARCHAR(20) NOT NULL UNIQUE,
                          Login VARCHAR(50) NOT NULL,
                          PasswordHash VARCHAR(100) NOT NULL
);
CREATE TABLE clients (
                         Id SERIAL PRIMARY KEY NOT NULL,
                         Name VARCHAR(50) NOT NULL,
                         Surname VARCHAR(50) NOT NULL,
                         Address VARCHAR(100) NOT NULL,
                         Phone VARCHAR(20) NOT NULL,
                         Email VARCHAR(100) NOT NULL,
                         Login VARCHAR(50) NOT NULL,
                         PasswordHash VARCHAR(100) NOT NULL,
                         ManagerId INT NOT NULL,
                         FOREIGN KEY (ManagerId) REFERENCES managers(Id)
);