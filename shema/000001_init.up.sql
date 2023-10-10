CREATE TABLE client_investments_info (
                                         id serial NOT NULL,
                                         bank_account varchar(30) default 0,
                                         investment_amount numeric default 0,
                                         investment_strategy varchar(50) default 'none',
                                         PRIMARY KEY (id)
);
CREATE TABLE manager_work_info (
                                   id serial NOT NULL,
                                   bank_account varchar(30) default '',
                                   capital_managment numeric default 0,
                                   profit_percent_day float4 default 0,
                                   PRIMARY KEY (id)
);

CREATE TABLE persons (
                         id serial NOT NULL UNIQUE,
                         Email varchar(50) default '',
                         Phone varchar(30) default '',
                         Address varchar(200) default '',
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

CREATE TABLE clients (
                         id serial NOT NULL UNIQUE,
                         manager_id int,
                         account_id int,
                         person_id int,
                         client_investments_info_id int,
                         PRIMARY KEY (id)
);
ALTER TABLE managers ADD CONSTRAINT managers_fk0 FOREIGN KEY (account_id) REFERENCES accounts(id);

ALTER TABLE managers ADD CONSTRAINT managers_fk1 FOREIGN KEY (person_id) REFERENCES persons(id);

ALTER TABLE managers ADD CONSTRAINT managers_fk2 FOREIGN KEY (work_info_id) REFERENCES manager_work_info(id);

ALTER TABLE clients ADD CONSTRAINT clients_fk0 FOREIGN KEY (account_id) REFERENCES accounts(id);

ALTER TABLE clients ADD CONSTRAINT clients_fk1 FOREIGN KEY (person_id) REFERENCES persons(id);

ALTER TABLE clients ADD CONSTRAINT clients_fk2 FOREIGN KEY (client_investments_info_id) REFERENCES client_investments_info(id);
ALTER TABLE clients ADD CONSTRAINT clients_fk3 FOREIGN KEY (manager_id) REFERENCES managers(id);