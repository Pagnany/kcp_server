-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: kcp
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `kcp`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `kcp` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `kcp`;

--
-- Table structure for table `anwesenheit`
--

DROP TABLE IF EXISTS `anwesenheit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `anwesenheit` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_mitglied` int DEFAULT NULL,
  `id_veranstaltung` int DEFAULT NULL,
  `anwesend` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `anwesenheit`
--

LOCK TABLES `anwesenheit` WRITE;
/*!40000 ALTER TABLE `anwesenheit` DISABLE KEYS */;
INSERT INTO `anwesenheit` VALUES (1,1,1,1),(2,2,1,0);
/*!40000 ALTER TABLE `anwesenheit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mitglieder`
--

DROP TABLE IF EXISTS `mitglieder`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mitglieder` (
  `id` int NOT NULL AUTO_INCREMENT,
  `vname` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mitglieder`
--

LOCK TABLES `mitglieder` WRITE;
/*!40000 ALTER TABLE `mitglieder` DISABLE KEYS */;
INSERT INTO `mitglieder` VALUES (1,'Eric','Schmale','Nani'),(2,'Hendrik','Poggemann','Hühnerbauer'),(3,'Simon','Heskamp','Hessi'),(4,'Erik','Pröhl','Kurt'),(5,'Nils','Berger','Bils'),(6,'Julian','Roß','Rossi'),(7,'Nicolai','Altemeyer','Bauer'),(8,'Jona','Niemeyer','Blond'),(9,'Jan','Poggemann','Dungi'),(10,'Leon','Twenning','Twente'),(11,'Pascal','Elling','Palle-Qualle'),(12,'Moritz','Niehaus','Neverhome'),(13,'Florian','Mülder','Juri'),(14,'Niklas','Sunderdiek','Niki'),(15,'Philipp','Budde','Flippi');
/*!40000 ALTER TABLE `mitglieder` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `strafen`
--

DROP TABLE IF EXISTS `strafen`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `strafen` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_strafe_typ` int DEFAULT NULL,
  `id_mitglied` int DEFAULT NULL,
  `preis` float DEFAULT NULL,
  `datum` date DEFAULT NULL,
  `anzahl` float DEFAULT NULL,
  `id_veranstaltung` int DEFAULT NULL,
  `bezeich` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `strafen`
--

LOCK TABLES `strafen` WRITE;
/*!40000 ALTER TABLE `strafen` DISABLE KEYS */;
INSERT INTO `strafen` VALUES (20,1,1,0,NULL,4,1,''),(21,5,1,0,NULL,2,1,''),(22,13,1,0,NULL,4,1,''),(25,1,12,0,NULL,2,1,''),(26,13,12,0,NULL,4,1,''),(27,2,12,0,NULL,1,1,''),(29,1,11,0,NULL,5,1,''),(30,5,11,0,NULL,1,1,''),(31,13,11,0,NULL,3,1,''),(32,1,8,0,NULL,8,1,''),(33,5,8,0,NULL,2,1,''),(34,13,8,0,NULL,4,1,''),(35,13,14,0,NULL,4,1,''),(36,1,14,0,NULL,4,1,''),(37,5,14,0,NULL,1,1,''),(38,1,5,0,NULL,6,1,''),(39,5,5,0,NULL,5,1,''),(40,13,5,0,NULL,3,1,''),(41,2,5,0,NULL,2,1,''),(47,12,5,0,NULL,1,1,''),(53,1,2,0,NULL,7,1,''),(54,5,2,0,NULL,2,1,''),(55,13,2,0,NULL,4,1,''),(56,2,2,0,NULL,3,1,''),(57,1,13,0,NULL,6,1,''),(58,5,13,0,NULL,1,1,''),(59,13,13,0,NULL,4,1,''),(60,1,4,0,NULL,2,1,''),(61,5,4,0,NULL,1,1,''),(62,13,4,0,NULL,4,1,''),(63,2,4,0,NULL,1,1,''),(64,1,15,0,NULL,10,1,''),(65,5,15,0,NULL,3,1,''),(66,13,15,0,NULL,2,1,''),(67,2,15,0,NULL,1,1,''),(69,0,12,9,NULL,1,1,''),(70,0,11,24.3,NULL,1,1,''),(71,0,8,9.9,NULL,1,1,''),(72,0,14,8.1,NULL,1,1,''),(73,0,5,16.8,NULL,1,1,''),(74,0,2,15.9,NULL,1,1,''),(75,0,1,9.9,NULL,1,1,''),(76,0,13,14.4,NULL,1,1,''),(77,0,4,15.6,NULL,1,1,'');
/*!40000 ALTER TABLE `strafen` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `strafen_typ`
--

DROP TABLE IF EXISTS `strafen_typ`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `strafen_typ` (
  `id` int NOT NULL AUTO_INCREMENT,
  `bezeichnung` varchar(255) DEFAULT NULL,
  `preis` float DEFAULT NULL,
  `aktiv` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `strafen_typ`
--

LOCK TABLES `strafen_typ` WRITE;
/*!40000 ALTER TABLE `strafen_typ` DISABLE KEYS */;
INSERT INTO `strafen_typ` VALUES (1,'Pumpe',0.3,1),(2,'Klingel',1.5,1),(4,'Kugel fallen lassen',5,1),(5,'Pumpe Technisch',0.6,1),(6,'Kugel bringen lassen (Bahn)',2,1),(7,'Kugel bringen lassen (Abwesend)',5,1),(8,'Kugel falsch bringen',2,1),(9,'Kugel erlaufen',10,1),(10,'Kugel versuchen zu erlaufen',2.5,1),(11,'Läufer behindern',10,1),(12,'Lustwurf',1,1),(13,'Neunen',1,1),(14,'Auf andere Bahn werfen',10,1),(15,'Objekt hinter Klingel treffen',15,1),(16,'Kranz',20,1);
/*!40000 ALTER TABLE `strafen_typ` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `veranstaltungen`
--

DROP TABLE IF EXISTS `veranstaltungen`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `veranstaltungen` (
  `id` int NOT NULL AUTO_INCREMENT,
  `bezeichnung` varchar(255) DEFAULT NULL,
  `datum` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `veranstaltungen`
--

LOCK TABLES `veranstaltungen` WRITE;
/*!40000 ALTER TABLE `veranstaltungen` DISABLE KEYS */;
INSERT INTO `veranstaltungen` VALUES (1,'Kegeln','2024-04-12'),(2,'Kegeln','2024-03-15');
/*!40000 ALTER TABLE `veranstaltungen` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-05-03 22:26:18
