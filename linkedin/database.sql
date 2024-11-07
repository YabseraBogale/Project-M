create table UserEmail(
    FirstName varchar(30) not null,
    LastName varchar(30) not null,
    Country varchar(30) not null,
    Email varchar(30) not null,
    HQCompanyName varchar(30) not null
);

create table UserData(
	Email varchar(30) not null primary key,
    FirstName varchar(30) not null,
    LastName varchar(30) not null,
    Country varchar(30) not null,
    HQCompanyName varchar(30) not null,
	Sent varchar(30) not null 
);
