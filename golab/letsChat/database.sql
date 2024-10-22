create table addisCheretaUser(
	UID int not null  primary key,
	FirstName varchar(30) not null,
	MiddleName varchar(30) not null,
	Photo varchar(30) not null,
	LastName varchar(30) not null,
	Email Text not null,
	Phonenumber varchar(30) not null,
	UserLocation varchar(30) not null,
	Password Text not null
);

create table Verfication(
	UID int references addisCheretaUser,
	code  int not null  primary key,
	dateStored date not null ,
	status varchar(8) not null 
);

Create table Item(
	IID int not null primary key,
	Name text not null,
	Description text not null,
	status varchar(15) not null,
	Photo varchar(30) not null,
	Category varchar(30) not null,
	startingPrice float not null
);

create table Seller(
	UID int REFERENCES addisCheretaUser,
	IID int REFERENCES Item,
	Description Text,
	Rating int not null
);

create Table Buyer(
	UID int references addisCheretaUser,
	IID int references Item
);

create table Auction(
	AID int not null primary key,
	IID int REFERENCES Item,
	Description Text not null,
	startTime date not null,
	EndTime date not null,
	Status varchar(15) not null,
	ReservePrice float not null
);

create Table FinalAuctionMessage(
	UID int REFERENCES addisCheretaUser,
	startTime date not null,
	messageTalke Text not null,
	Photo varchar(30) not null
);
