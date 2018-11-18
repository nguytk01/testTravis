create extension "uuid-ossp";
drop table if exists Task;
drop table if exists Account;
create table Account (
  email varchar(30) Unique Primary Key,
  timeZone integer,
  encryptedPasswordHash varchar(30),
  firstName varchar(100),
  lastName varchar(100),
  middleName varchar(100)
);

create table Task(
  taskId UUID Unique Primary key,
  taskName varchar(20),
  timeSpent real,
	taskDate time,
  email varchar(30) REFERENCES Account(email)
);