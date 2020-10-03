drop table if exists Restaurant, quartier, categorie, note, sondage, votant, Sondage_restaurant cascade;

CREATE TABLE Categorie
(
    Id      SERIAL PRIMARY KEY,
    Libelle VARCHAR(50) not null

);


CREATE TABLE Quartier
(
    ID      SERIAL PRIMARY KEY,
    Libelle VARCHAR(50) not null

);


CREATE TABLE Apero
(
    Id           SERIAL PRIMARY KEY,
    Nom          VARCHAR(50) not null,
    DateApero    DATE,
    DateCreation DATE,
    Createur_Id int not null,
    CONSTRAINT Createur_Id_fkey FOREIGN KEY (Createur_Id) REFERENCES Utilisateur (Id)

);


CREATE TABLE Utilisateur
(
    Nom    VARCHAR(50) not null,
    Id     SERIAL PRIMARY KEY,
    Mail   VARCHAR(50),
    NumTel VARCHAR(15),
    Photo  BYTEA
);


CREATE TABLE Endroit
(
    Id              SERIAL PRIMARY KEY,
    Note            int           NULL,
    Appreciation    VARCHAR(200)  NULL,
    PrixMoyen       DECIMAL(5, 2) NULL,
    Adresse         VARCHAR(50)   NULL,
    Ville           VARCHAR(50),
    DateCreation    DATE          NULL,
    Nom             VARCHAR(50)   not NULL,
    noteCopiosite   int,
    noteDeliciosite int,
    noteCadre       int,
    noteAccueil     int,
    Quartier_Id     int           not null,
    Categorie_Id    int           not null,

    CONSTRAINT Id_Quartier_fkey FOREIGN KEY (Quartier_Id) REFERENCES Quartier (Id),
    CONSTRAINT Id_Categorie_fkey FOREIGN KEY (Categorie_Id) REFERENCES Categorie (Id)
);

CREATE TABLE Vote
(
    Id             SERIAL PRIMARY KEY,
    NbVotes        VARCHAR(50),
    DateVote Date,
    Endroit_Id     int,
    Utilisateur_Id int,
    Apero_Id       int,
    CONSTRAINT Id_endroit_fkey FOREIGN KEY (Endroit_Id) REFERENCES Endroit (Id),
    CONSTRAINT Id_Utilisateur_fkey FOREIGN KEY (Utilisateur_Id) REFERENCES Utilisateur (Id),
    CONSTRAINT Id_Apero_fkey FOREIGN KEY (Apero_Id) REFERENCES Apero (Id)
);

CREATE TABLE Groupe
(
    Id             SERIAL PRIMARY KEY,
    Nom            VARCHAR(50),
    Icone          bytea,
    Utilisateur_Id int,
    Apero_Id       int,
    CONSTRAINT Id_Utilisateur_fkey FOREIGN KEY (Utilisateur_Id) REFERENCES Utilisateur (Id)
);

CREATE TABLE Groupe_Utilisateur
(

    Groupe_Id      int,
    Utilisateur_Id int,
    CONSTRAINT Id_Utilisateur_fkey FOREIGN KEY (Utilisateur_Id) REFERENCES Utilisateur (Id),
    CONSTRAINT Id_Groupe_fkey FOREIGN KEY (Groupe_Id) REFERENCES Groupe (Id)

);




