-- MySQL dump 10.13  Distrib 9.0.1, for Win64 (x86_64)
--
-- Host: localhost    Database: toyota_chatbot
-- ------------------------------------------------------
-- Server version	9.0.1

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
-- Current Database: `toyota_chatbot`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `toyota_chatbot` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `toyota_chatbot`;

--
-- Table structure for table `responses`
--

DROP TABLE IF EXISTS `responses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `responses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `keywords` varchar(255) NOT NULL,
  `reply` text NOT NULL,
  `buttons` json DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `responses`
--

LOCK TABLES `responses` WRITE;
/*!40000 ALTER TABLE `responses` DISABLE KEYS */;
INSERT INTO `responses` VALUES (8,'dealership','You can find our dealerships in Colombo, Kandy, and Galle.','[{\"text\": \"Colombo\", \"value\": \"colombo\"}, {\"text\": \"Kandy\", \"value\": \"kandy\"}, {\"text\": \"Galle\", \"value\": \"galle\"}]','2024-10-21 08:18:03'),(9,'promotions','Toyota Lanka offers various promotions. Visit our promotions page for more details.',NULL,'2024-10-21 08:18:03'),(10,'hi','Hi! How can I assist you today?',NULL,'2024-10-22 04:03:44'),(11,'hello','Hello! How can I assist you today?',NULL,'2024-10-22 04:03:44'),(12,'hey','Hey! How can I assist you today?',NULL,'2024-10-22 04:03:44'),(13,'opening hours','Our showrooms are open from 8:30 AM to 5:30 PM, Monday to Saturday.',NULL,'2024-10-22 04:41:08'),(14,'contact, helpline, contact number, contact information','Toyota Lanka (Private) Limited, 337, Negombo Road, Wattala 11300, Sri Lanka. Tel: +9411 293 9000 - 6, Fax: +9411 293 9005, Email: info@toyota.lk.',NULL,'2024-10-22 04:41:08'),(15,'main office, main showroom','Toyota Lanka (Private) Limited, 337, Negombo Road, Wattala 11300, Sri Lanka. Tel: +9411 293 9000 - 6, Fax: +9411 293 9005, Email: info@toyota.lk.',NULL,'2024-10-22 04:41:08'),(16,'customer support, support, help','Our customer support team is available 24/7 to assist you. You can reach us at +9411 293 9000 or email support@toyota.lk.',NULL,'2024-10-25 13:53:57'),(17,'pricing, price, cost, vehicle cost','To get the latest pricing details for our vehicles, please visit our pricing page: <a href=\"https://www.toyota.lk/pricing\" target=\"_blank\">Pricing Toyota Lanka</a>.',NULL,'2024-10-25 13:53:57'),(18,'test drive, schedule test drive','You can schedule a test drive by visiting our nearest showroom or filling out the form on our website: <a href=\"https://www.toyota.lk/test-drive\" target=\"_blank\">Schedule Test Drive</a>.',NULL,'2024-10-25 13:53:57'),(19,'finance options, loans, leasing, finance','Toyota Lanka offers flexible finance options for all vehicle models. To learn more, visit: <a href=\"https://www.toyota.lk/finance-options\" target=\"_blank\">Finance Options</a>.',NULL,'2024-10-25 13:53:57'),(20,'insurance, vehicle insurance, insurance options','We provide comprehensive insurance options through our trusted partners. Check our insurance plans here: <a href=\"https://www.toyota.lk/insurance\" target=\"_blank\">Insurance Plans</a>.',NULL,'2024-10-25 13:53:57'),(21,'warranty, warranty information','Toyota Lanka offers comprehensive warranty packages for all vehicles. For details, visit: <a href=\"https://www.toyota.lk/warranty\" target=\"_blank\">Warranty Information</a>.',NULL,'2024-10-25 14:14:23'),(22,'finance, loan options, financing','We offer flexible financing options tailored to your needs. Explore our financing plans here: <a href=\"https://www.toyota.lk/finance\" target=\"_blank\">Finance Options</a>.',NULL,'2024-10-25 14:14:23'),(23,'trade-in, trade in, vehicle trade','Thinking of trading in your vehicle? Get an estimate on our website: <a href=\"https://www.toyota.lk/trade-in\" target=\"_blank\">Trade-in Program</a>.',NULL,'2024-10-25 14:14:23'),(24,'test drive, schedule test drive','Schedule a test drive for any model through our online booking system: <a href=\"https://www.toyota.lk/test-drive\" target=\"_blank\">Test Drive</a>.',NULL,'2024-10-25 14:14:23'),(25,'accessories, car accessories','Check out our range of Toyota accessories designed to enhance your driving experience: <a href=\"https://www.toyota.lk/accessories\" target=\"_blank\">Accessories</a>.',NULL,'2024-10-25 14:14:23');
/*!40000 ALTER TABLE `responses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_queries`
--

DROP TABLE IF EXISTS `user_queries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_queries` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_message` varchar(255) NOT NULL,
  `response_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `response_id` (`response_id`),
  CONSTRAINT `user_queries_ibfk_1` FOREIGN KEY (`response_id`) REFERENCES `responses` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_queries`
--

LOCK TABLES `user_queries` WRITE;
/*!40000 ALTER TABLE `user_queries` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_queries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vehicle_models`
--

DROP TABLE IF EXISTS `vehicle_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `vehicle_models` (
  `id` int NOT NULL AUTO_INCREMENT,
  `model_name` varchar(100) NOT NULL,
  `description` text NOT NULL,
  `link` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vehicle_models`
--

LOCK TABLES `vehicle_models` WRITE;
/*!40000 ALTER TABLE `vehicle_models` DISABLE KEYS */;
INSERT INTO `vehicle_models` VALUES (1,'Toyota Corolla','A compact sedan known for its reliability and efficiency.','https://www.toyota.com/corolla/','2024-10-17 08:19:08'),(2,'Toyota Camry','A midsize sedan with advanced technology and a comfortable ride.','https://www.toyota.com/camry/','2024-10-17 08:19:08'),(3,'Toyota RAV4','A compact SUV that offers versatility and space.','https://www.toyota.com/rav4/','2024-10-17 08:19:08'),(4,'Toyota Hilux','A tough and durable pickup truck suitable for any terrain.','https://www.toyota.com/hilux/','2024-10-17 08:19:08'),(5,'Toyota Supra','The Toyota Supra is a high-performance sports car with advanced technology and unmatched driving dynamics.','https://www.toyota.com/supra/','2024-10-25 13:54:03'),(6,'Toyota C-HR','The Toyota C-HR is a crossover SUV that combines style and agility, perfect for urban adventures.','https://www.toyota.com/c-hr/','2024-10-25 13:54:03'),(7,'Toyota Tacoma','The Toyota Tacoma is a midsize pickup truck built for durability and off-road capability.','https://www.toyota.com/tacoma/','2024-10-25 13:54:03'),(8,'Toyota Avalon','The Toyota Avalon is a premium sedan that offers a smooth ride and luxurious features.','https://www.toyota.com/avalon/','2024-10-25 13:54:03'),(9,'Toyota Sequoia','The Toyota Sequoia is a full-size SUV with spacious seating and advanced safety features, ideal for family trips.','https://www.toyota.com/sequoia/','2024-10-25 13:54:03'),(10,'Toyota 4Runner','The Toyota 4Runner is a robust SUV built for off-road adventures, offering ample space and reliability.','https://www.toyota.com/4runner/','2024-10-25 14:14:30'),(11,'Toyota C-HR','The Toyota C-HR combines bold styling with advanced features, making it a standout compact crossover.','https://www.toyota.com/c-hr/','2024-10-25 14:14:30'),(12,'Toyota Tundra','The Toyota Tundra is a powerful full-size pickup truck with towing capabilities and advanced technology.','https://www.toyota.com/tundra/','2024-10-25 14:14:30'),(13,'Toyota Venza','The Toyota Venza offers a blend of style, comfort, and hybrid efficiency in a mid-size crossover.','https://www.toyota.com/venza/','2024-10-25 14:14:30'),(14,'Toyota Supra','The Toyota Supra delivers a thrilling performance with its powerful engine and sleek sports car design.','https://www.toyota.com/supra/','2024-10-25 14:14:30');
/*!40000 ALTER TABLE `vehicle_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-25 20:42:39
