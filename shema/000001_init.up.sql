CREATE TABLE manager_work_info (
                                   id serial NOT NULL,
                                   bank_account varchar(30) UNIQUE,
                                   capital_managment numeric,
                                   profit_percent_day float4,
                                   PRIMARY KEY (id)
);

CREATE TABLE persons (
                         id serial NOT NULL UNIQUE,
                         Email varchar(50) UNIQUE,
                         Phone varchar(30) UNIQUE,
                         Address varchar(200) UNIQUE,
                         PRIMARY KEY (id)
);

CREATE TABLE accounts (
                          id serial NOT NULL UNIQUE,
                          login varchar(30) NOT NULL UNIQUE,
                          password_hash varchar(100) NOT NULL,
                          PRIMARY KEY (id)
);

CREATE TABLE managers (
                          id serial NOT NULL UNIQUE,
                          account_id int,
                          person_id int,
                          work_info_id int,
                          PRIMARY KEY (id)
);

ALTER TABLE managers ADD CONSTRAINT managers_fk0 FOREIGN KEY (account_id) REFERENCES accounts(id);

ALTER TABLE managers ADD CONSTRAINT managers_fk1 FOREIGN KEY (person_id) REFERENCES persons(id);

ALTER TABLE managers ADD CONSTRAINT managers_fk2 FOREIGN KEY (work_info_id) REFERENCES manager_work_info(id);