CREATE TABLE IF NOT EXISTS Users(
    ID INT NOT NULL UNIQUE serial,
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Links(
    ID INT NOT NULL UNIQUE serial,
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID INT ,
    AddressId INT
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    FOREIGN KEY (AddressId) REFERENCES Address(ID) ,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Address(
    ID INT NOT NULL UNIQUE serial,
    Line1 VARCHAR (255) ,
    Line2 VARCHAR (255) ,
    City VARCHAR (255)  ,
    State VARCHAR (255),
    Zip   INT,
    PRIMARY KEY (ID)
);
