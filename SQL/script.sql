drop  table if exists Restaurant, quartier, categorie, note, sondage, votant, Sondage_restaurant cascade;

CREATE TABLE Categorie(
   Id int  PRIMARY KEY,
   Libelle VARCHAR(50) not null

);


CREATE TABLE Quartier(
   ID int PRIMARY KEY,
   Libelle VARCHAR(50) not null

);


CREATE TABLE Note(
   Id int  PRIMARY KEY,
   NoteCopiosite INT,
   NoteDeliciosite INT,
   NoteCadre INT,
   NoteAccueil INT

);


CREATE TABLE Sondage(
   Id int PRIMARY KEY,
   Nom VARCHAR(50)

);


CREATE TABLE Votant(
   Nom VARCHAR(50),
   Id int PRIMARY KEY,
   Sondage_Id int not null,
CONSTRAINT Id_Sondage_fkey FOREIGN KEY (Sondage_Id)    REFERENCES sondage(Id)
);


CREATE TABLE Restaurant(
   Id int PRIMARY KEY,
   Note int NOT NULL,
   Appreciation VARCHAR(200),
   PrixMoyen DECIMAL(5,2),
   Adresse VARCHAR(50),
   Ville VARCHAR(50),
   DateCreation DATE,
   Nom VARCHAR(50),
   Note_Id int not null,
   Quartier_Id int not null,
   Categorie_Id  int not null,

CONSTRAINT Id_Note_fkey FOREIGN KEY (Note_Id) REFERENCES Note(Id),
CONSTRAINT Id_Quartier_fkey FOREIGN KEY (Quartier_Id) REFERENCES Quartier(Id),
CONSTRAINT Id_Categorie_fkey FOREIGN KEY (Categorie_Id) REFERENCES Categorie(Id)

);

CREATE TABLE Sondage_restaurant(

   NbVotes VARCHAR(50),
   NomSondage VARCHAR(50),
   DateApero Date,
   Restau_Id int not null,
   Sondage_Id int not null,
CONSTRAINT Id_Restaurants_fkey FOREIGN KEY (Restau_Id) REFERENCES Restaurant(Id),
CONSTRAINT Id_Sondage_fkey FOREIGN KEY (Sondage_Id) REFERENCES Sondage(Id),
PRIMARY KEY (Restau_id, Sondage_id)
);
