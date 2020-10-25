drop table if exists Endroit, quartier, categorie, Apero, Vote, Groupe_Utilisateur, Groupe,Utilisateur,Groupe,Authentification,Sessions cascade;

CREATE TABLE Categorie
(
    Id      SERIAL PRIMARY KEY,
    Libelle VARCHAR(50) NOT NULL
);


CREATE TABLE Quartier
(
    ID      SERIAL PRIMARY KEY,
    Libelle VARCHAR(50) NOT NULL
);



CREATE TABLE Utilisateur
(
    Nom             VARCHAR(50) NOT NULL,
    Id              SERIAL PRIMARY KEY,
    Mail            VARCHAR(50) DEFAULT '',
    NumTel          VARCHAR(15) NOT NULL UNIQUE,
    DateInscription DATE        NOT NULL
);

CREATE TABLE Authentification
(
    Id             SERIAL PRIMARY KEY,
    Pass           VARCHAR(500) NOT NULL,
    Utilisateur_Id int          NOT NULL,
    CONSTRAINT Utilisateur_Id_fkey FOREIGN KEY (Utilisateur_Id) REFERENCES Utilisateur (Id)
);

CREATE TABLE Sessions
(
    Id             SERIAL PRIMARY KEY,
    Utilisateur_Id INT NOT NULL,
    CONSTRAINT Utilisateur_Id_fkey FOREIGN KEY (Utilisateur_Id) REFERENCES Utilisateur (id)
);

CREATE TABLE Apero
(
    Id           SERIAL PRIMARY KEY,
    Nom          VARCHAR(50) NOT NULL,
    DateApero    DATE        NOT NULL,
    DateCreation DATE        NOT NULL,
    Createur_Id  int         NOT NULL,
    CONSTRAINT Createur_Id_fkey FOREIGN KEY (Createur_Id) REFERENCES Utilisateur (Id)
);


CREATE TABLE Endroit
(
    Id              SERIAL PRIMARY KEY,
    Appreciation    VARCHAR(200)  DEFAULT '',
    PrixMoyen       DECIMAL(5, 2) DEFAULT 0,
    Adresse         VARCHAR(50)   DEFAULT '',
    Ville           VARCHAR(50) NOT NULL,
    DateCreation    DATE        NOT NULL,
    Nom             VARCHAR(50) NOT NULL,
    noteCopiosite   INT           DEFAULT 1,
    noteDeliciosite INT           DEFAULT 1,
    noteCadre       INT           DEFAULT 1,
    noteAccueil     INT           DEFAULT 1,
    Quartier_Id     INT         NOT NULL,
    Categorie_Id    INT         NOT NULL,

    CONSTRAINT Id_Quartier_fkey FOREIGN KEY (Quartier_Id) REFERENCES Quartier (Id),
    CONSTRAINT Id_Categorie_fkey FOREIGN KEY (Categorie_Id) REFERENCES Categorie (Id)
);

CREATE TABLE Vote
(
    Id             SERIAL PRIMARY KEY,
    DateVote       Date NOT NULL,
    Endroit_Id     int  NOT NULL,
    Utilisateur_Id int  NOT NULL,
    Apero_Id       int  NOT NULL,
    CONSTRAINT Id_endroit_fkey FOREIGN KEY (Endroit_Id) REFERENCES Endroit (Id),
    CONSTRAINT Id_Utilisateur_fkey FOREIGN KEY (Utilisateur_Id) REFERENCES Utilisateur (Id),
    CONSTRAINT Id_Apero_fkey FOREIGN KEY (Apero_Id) REFERENCES Apero (Id)
);

CREATE TABLE Groupe
(
    Id             SERIAL PRIMARY KEY,
    Nom            VARCHAR(50) NOT NULL,
    Createur_Id int         NOT NULL,
    Apero_Id       int         NOT NULL,
    CONSTRAINT Id_Utilisateur_fkey FOREIGN KEY (Createur_Id) REFERENCES Utilisateur (Id)
);

CREATE TABLE Groupe_Utilisateur
(

    Groupe_Id      int NOT NULL,
    Utilisateur_Id int NOT NULL,
    CONSTRAINT Id_Utilisateur_fkey FOREIGN KEY (Utilisateur_Id) REFERENCES Utilisateur (Id),
    CONSTRAINT Id_Groupe_fkey FOREIGN KEY (Groupe_Id) REFERENCES Groupe (Id)

);




