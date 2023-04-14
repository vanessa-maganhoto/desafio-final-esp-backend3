CREATE DATABASE IF NOT EXISTS my_db;
USE my_db;

DROP TABLE IF EXISTS `dentists`;

CREATE TABLE `dentists` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(50) DEFAULT NULL,
  `surname` varchar(50) DEFAULT NULL,
  `registration` varchar(50) DEFAULT NULL
);

INSERT INTO dentists (name, surname, registration) VALUES
("Chris", "Martin", "1234A"),
("Jonny", "Buckland", "7654S"),
("Will", "Champion", "56434L"),
("Guy", "Berryman", "345678S"),
("Joao", "Borges Santos", "98017"),
("Fernanda", "Reis",  "99727"),
("Adriana", "Batista", "96336"),
("Luiz", "Freitas", "93280"),
("Paulo", "Mendes", "92144"),
("Ana", "Monteiro", "90050");

DROP TABLE IF EXISTS `patients`;

CREATE TABLE `patients` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(50) DEFAULT NULL,
  `surname` varchar(50) DEFAULT NULL,
  `documentNumber` varchar(50) DEFAULT NULL,
  `registrationDate` varchar(50) DEFAULT NULL
);

INSERT INTO patients (name, surname, documentNumber, registrationDate) VALUES
("Antonio", "Fernandes", "10003231011", "23/11/2018"),
("Pedro", "Soares", "81130075583", "23/01/2004"),
("Carlos", "Reis", "31011416034", "09/07/2012"),
("Antonia", "Andrade", "61301165565", "26/02/1999"),
("Maria", "Gomes", "21223330729", "10/10/2005"),
("Juliana", "Lopes", "31112353330", "30/04/2014"),
("Ana", "Ramos", "01215401205", "27/12/2007"),
("Marcos", "Monteiro", "62202328068", "02/08/2002"),
("Mariana", "Dias", "62220665569", "06/10/2022"),
("Aline", "Medeiros", "50103523251", "12/06/2015"),
("Lucas", "Martins", "31215321430", "07/01/2011"),
("Joana", "Felix", "41130074943", "07/07/2019"),
("Martina", "Souza", "32112531638", "28/02/2008"),
("Paula", "Benício", "11221215116", "16/10/2017"),
("Diego", "Montes", "81232648485", "26/01/2012"),
("Paula", "Matias", "42234055245", "10/09/2022"),
("Marcela", "Amorim", "02325070409", "11/09/2003"),
("Bruno", "Pereira", "72213404178", "06/06/2006"),
("Tiago", "Castro", "61120337763", "17/03/2019"),
("Jonas", "Rocha", "00032225105", "13/08/2014"),
("Luciana", "Castro", "82232273580", "30/11/2000"),
("Patricia", "Soares", "51111155755", "21/07/2020"),
("Ricardo", "Dias", "02211014208", "16/10/2021"),
("Antonieta", "Patrício", "01344322409", "27/05/2022"),
("Karina", "Santana", "31102325430", "06/10/2009");

DROP TABLE IF EXISTS `appointments`;

CREATE TABLE `appointments` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `idDentist` int NOT NULL,
  `idPatient` int NOT NULL,
  `apptDate` varchar(50) NOT NULL,
  `description` varchar(250) NOT NULL,
  FOREIGN KEY (idDentist) REFERENCES dentists (id),
  FOREIGN KEY (idPatient) REFERENCES patients (id)
);

INSERT INTO appointments (idDentist, idPatient, apptDate, description) VALUES
(2, 20, "30/05/2023 15:30", "lorem ipsum"),
(2, 21, "30/05/2023 15:30", "lorem ipsum"),
(6, 10, "30/05/2023 15:30", "lorem ipsum"),
(6, 16, "30/05/2023 15:30", "lorem ipsum"),
(7, 6, "30/05/2023 15:30", "lorem ipsum"),
(8, 11, "30/05/2023 15:30", "lorem ipsum"),
(9, 17, "30/05/2023 15:30", "lorem ipsum"),
(9, 18, "30/05/2023 15:30", "lorem ipsum");