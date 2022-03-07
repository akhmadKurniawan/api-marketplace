-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: app
-- ------------------------------------------------------
-- Server version	8.0.28-0ubuntu0.20.04.3

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
-- Table structure for table `costumers`
--

DROP TABLE IF EXISTS `costumers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `costumers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `name` varchar(25) DEFAULT NULL,
  `alamat` varchar(50) DEFAULT NULL,
  `no_hp` varchar(25) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `costumers`
--

LOCK TABLES `costumers` WRITE;
/*!40000 ALTER TABLE `costumers` DISABLE KEYS */;
INSERT INTO `costumers` VALUES (1,6,'test Cost','jl. Test','08212277000','2022-02-14 06:01:10','2022-02-14 06:01:10'),(6,1,'test test','jl. Test','08212277000','2022-02-14 13:55:49','2022-02-14 13:55:49'),(10,5,'test baru','jl. Test','08212277000','2022-02-14 15:38:19','2022-02-14 15:38:19'),(39,12,'test id double','jl. Test','08212277000','2022-02-15 08:02:23','2022-02-15 08:02:23'),(43,14,'test id double sip','jl. Test','08212277000','2022-02-15 14:12:49','2022-02-15 14:12:49'),(44,13,'test baru','jl. Test','08212277000','2022-02-16 06:47:43','2022-02-16 06:47:43'),(45,19,'cost baru','jl. Test','08212277000','2022-03-02 08:53:09','2022-03-02 08:53:09'),(68,41,'testing','Jl.Alamat','081322456','2022-03-04 03:13:54','2022-03-04 03:13:54');
/*!40000 ALTER TABLE `costumers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_types`
--

DROP TABLE IF EXISTS `product_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `image` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_types`
--

LOCK TABLES `product_types` WRITE;
/*!40000 ALTER TABLE `product_types` DISABLE KEYS */;
INSERT INTO `product_types` VALUES (1,'gadget','2022-02-14 07:59:36','2022-02-14 07:59:36',NULL),(2,'Fashion','2022-02-17 14:25:47','2022-02-17 14:25:47','http://localhost:5000/api/v1/images/9f64c6c1-f9b1-48ce-a5c6-8e480cf5ee27image.png'),(5,'Otomotifs','2022-02-18 06:41:36','2022-02-18 06:41:36','http://localhost:5000/api/v1/images/ebed8193-9b15-4ed5-b9e4-ef3028ed4aefimage.png'),(6,'Otomotifs','2022-02-18 06:41:37','2022-02-18 06:41:37','http://localhost:5000/api/v1/images/2222bf3f-d007-4ea3-a048-c639fd6e9bb9image.png'),(7,'Otomotifs','2022-02-18 06:41:38','2022-02-18 06:41:38','http://localhost:5000/api/v1/images/eef4fa42-0ba6-43d2-b857-78dbede5bc3fimage.png'),(8,'Otomotifs','2022-02-18 06:41:39','2022-02-18 06:41:39','http://localhost:5000/api/v1/images/ca871a0d-8e54-4448-b69d-17b0f4a5bae0image.png'),(9,'Otomotifsss','2022-02-18 06:44:18','2022-02-18 06:44:18','http://localhost:5000/api/v1/images/882522dd-216b-4a25-86ed-7b88f87f106cimage.png'),(10,'Otomotifssss','2022-02-18 06:44:26','2022-02-18 06:44:26','http://localhost:5000/api/v1/images/5747eecf-35ce-40d6-bb4c-3435396d31a2image.png');
/*!40000 ALTER TABLE `product_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_type` int DEFAULT NULL,
  `shop_id` int DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `price` int DEFAULT NULL,
  `description` text,
  `qty` int DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_shop_id` (`shop_id`)
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (29,1,3,'coklat',10000,'coklat desc',727,'http://localhost:5000/api/v1/images/a657e1af-2b84-4b28-afb2-6146f3778096image.png','2022-02-17 07:12:37','2022-03-03 09:21:07'),(54,1,12,'coklat',10000,'coklat desc',2,'http://localhost:5000/api/v1/images/99f21456-2c91-4869-955c-6ebf47079c65image.png','2022-02-17 10:15:46','2022-03-02 07:34:19'),(114,1,11,'coklat Keju',10000,'coklat desc',3,'http://localhost:5000/api/v1/images/7584b4c6-521f-4ad3-9bd3-ec51552e12aaimage.png','2022-02-18 07:51:14','2022-02-18 07:51:14'),(115,1,11,'Keju',10000,'coklat desc',3,'http://localhost:5000/api/v1/images/3b17ed80-32f7-4b9d-95af-8c5407514861image.png','2022-02-18 07:52:18','2022-02-18 07:52:18'),(116,1,11,'Keju coklat',10000,'coklat desc',3,'http://localhost:5000/api/v1/images/49957072-6522-4f35-be06-6ccda50654e8image.png','2022-02-18 08:43:37','2022-02-18 08:43:37'),(117,1,11,'Keju coklatt',10000,'coklat desc',3,'http://localhost:5000/api/v1/images/9592313a-8e1d-450f-9d7e-0822cd95cb8aimage.png','2022-02-18 08:44:46','2022-02-18 08:44:46'),(118,1,11,'coklat Keju baru',10000,'coklat desc baru',1,'http://localhost:5000/api/v1/images/7e251eeb-7388-499f-b0df-4ee08a93b6d3image.png','2022-03-02 07:26:09','2022-03-02 07:27:19');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sellers`
--

DROP TABLE IF EXISTS `sellers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sellers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `name` varchar(25) DEFAULT NULL,
  `alamat` varchar(50) DEFAULT NULL,
  `no_hp` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sellers`
--

LOCK TABLES `sellers` WRITE;
/*!40000 ALTER TABLE `sellers` DISABLE KEYS */;
INSERT INTO `sellers` VALUES (11,2,'test Cost','jl. Test','08212277000','2022-02-14 14:15:52','2022-02-14 14:15:52'),(26,1,'test shop 3','jl. Test','08212277000','2022-02-15 12:11:01','2022-02-15 12:11:01'),(33,16,'test','jl. Test','08212277000','2022-02-17 06:54:00','2022-02-17 06:54:00'),(34,19,'seller baru','jl. Test','08212277000','2022-03-02 07:21:41','2022-03-02 07:21:41'),(35,19,'seller baru lgi','jl. Test','08212277000','2022-03-02 07:23:43','2022-03-02 07:23:43'),(58,42,'testing','Jl.Alamat','081322456','2022-03-04 03:15:06','2022-03-04 03:15:06');
/*!40000 ALTER TABLE `sellers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shops`
--

DROP TABLE IF EXISTS `shops`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `shops` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `description` text NOT NULL,
  `alamat` text NOT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `seller_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shops`
--

LOCK TABLES `shops` WRITE;
/*!40000 ALTER TABLE `shops` DISABLE KEYS */;
INSERT INTO `shops` VALUES (2,'test baru 2','desc new','jl. Test','new logo',1,'2022-02-14 16:29:19','2022-02-14 16:29:19'),(3,'shop coba seller id','desc new','jl. Test','new logo',11,'2022-02-15 11:10:13','2022-02-15 11:10:13'),(4,'shop coba seller id 2','desc new','jl. Test','new logo',11,'2022-02-15 11:11:46','2022-02-15 11:11:46'),(6,'shop coba seller id 26','desc new','jl. Test','new logo',23,'2022-02-15 12:11:54','2022-02-15 12:11:54'),(11,'test','desc new','jl. Test','new logo',33,'2022-02-17 06:54:27','2022-02-17 06:54:27'),(12,'one','desc new','jl. Test','new logo',26,'2022-02-17 09:27:18','2022-02-17 09:27:18'),(13,'newShop','murah berkualitas','jl. gunung jati','http://localhost:5000/api/v1/images/095d35bd-0c83-42e5-ac80-6ad2cd7f33faimage.png',26,'2022-02-17 14:13:51','2022-02-17 14:13:51'),(14,'newShop','murah berkualitas','jl. gunung jati','http://localhost:5000/api/v1/images/d49c74f0-2f72-4e9c-a080-1f8155da34f8image.png',26,'2022-02-18 06:19:20','2022-02-18 06:19:20'),(15,'newShop','murah berkualitas','jl. gunung jati','http://localhost:5000/api/v1/images/7f8de4a3-9088-4056-b3ec-81092dd32214image.png',26,'2022-02-18 06:20:40','2022-02-18 06:20:40'),(16,'newShop','murah berkualitas','jl. gunung jati','http://localhost:5000/api/v1/images/99cd862d-f291-4583-b0e2-df31102bab62image.png',26,'2022-02-18 06:20:42','2022-02-18 06:20:42'),(17,'newShop baru','murah berkualitas','jl. gunung jati','http://localhost:5000/api/v1/images/89cbf5f9-fd3d-4579-a871-ea4c472c9752image.png',34,'2022-03-02 07:24:57','2022-03-02 07:24:57'),(18,'newShop baru','murah berkualitas','jl. gunung jati','http://localhost:5000/api/v1/images/bea8f75c-65ea-4e9c-91dd-f2538c7c26fbimage.png',26,'2022-03-04 03:36:58','2022-03-04 03:36:58');
/*!40000 ALTER TABLE `shops` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `product_id` varchar(50) DEFAULT NULL,
  `type` varchar(15) DEFAULT NULL,
  `description` text,
  `amount` varchar(50) DEFAULT NULL,
  `total_product` varchar(50) DEFAULT NULL,
  `Status` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `id_va` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=397 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (339,16,'29','Debit','buy food','30000','3','COMPLETED','2022-02-25 09:59:41','2022-02-25 10:03:15','VA_fixed-2022022539'),(340,16,'29','Kredit','buy food','30000','3','COMPLETED','2022-02-25 09:59:41','2022-02-25 10:03:15','VA_fixed-2022022539'),(341,16,'29','Debit','buy food','30000','3','COMPLETED','2022-03-01 06:35:32','2022-03-01 06:37:49','VA_fixed-2022030130'),(393,16,'29','Debit','buy food','-30000','-3','PENDING','2022-03-03 09:20:58','2022-03-03 09:20:58','VA_fixed-2022030357'),(394,2,'29','Kredit','buy food','30000','3','PENDING','2022-03-03 09:20:58','2022-03-03 09:20:58','VA_fixed-2022030357'),(395,16,'29','Debit','buy food','-30000','-3','PENDING','2022-03-03 09:21:07','2022-03-03 09:21:07','VA_fixed-2022030306'),(396,2,'29','Kredit','buy food','30000','3','PENDING','2022-03-03 09:21:07','2022-03-03 09:21:07','VA_fixed-2022030306');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_tokens`
--

DROP TABLE IF EXISTS `user_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_tokens` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_tokens`
--

LOCK TABLES `user_tokens` WRITE;
/*!40000 ALTER TABLE `user_tokens` DISABLE KEYS */;
INSERT INTO `user_tokens` VALUES (1,1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIk5hbWUiOiJ3YW4iLCJSb2xlIjozLCJleHAiOjE2NDY2MTU5MzcsImlhdCI6MTY0NjM4NTUzN30.uhnjN-XJwZrDeQhc0rP_j3KxXkDLQ1BHd0q8ef9KqsU','2022-02-14 06:00:36','2022-03-04 09:18:58'),(3,4,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjQsIk5hbWUiOiJ0ZXMgdXNlciIsIlJvbGUiOjEsImV4cCI6MTY0NTA3NzYyMCwiaWF0IjoxNjQ0ODQ3MjIwfQ.tHmzY1w4hMEFPM-BD3C-DrzwuxDKK60r36x3AVhABsA','2022-02-14 09:49:23','2022-02-14 14:00:21'),(4,5,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjUsIk5hbWUiOiJ0ZXN0dCIsIlJvbGUiOjEsImV4cCI6MTY0NTEyMzY5NywiaWF0IjoxNjQ0ODkzMjk3fQ.Ae7Sa0JuVf8Xtd3d_wrVjxG84OQapXQfY719wBS68ls','2022-02-14 14:16:58','2022-02-15 02:48:18'),(5,12,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEyLCJOYW1lIjoiY29zdHVtZXIiLCJSb2xlIjoxLCJleHAiOjE2NDU1Nzg5ODgsImlhdCI6MTY0NTM0ODU4OH0.iFT8PNB18NC0mCft2epOFtHpG3hNKj-JY6R7wxvptRE','2022-02-15 06:40:14','2022-02-20 09:16:28'),(7,14,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjE0LCJOYW1lIjoidGVzIGNvc3R1bW1lciIsIlJvbGUiOjEsImV4cCI6MTY0NTE2NDU1NywiaWF0IjoxNjQ0OTM0MTU3fQ.x8TcC47wmG4yFfZUY_wB7cc5Ox_h6NGxny6PAYI_3Ys','2022-02-15 14:09:17','2022-02-15 14:09:17'),(8,13,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEzLCJOYW1lIjoiY29zdHVtbWVyIHRlc3QiLCJSb2xlIjowLCJleHAiOjE2NDUzMzgyNjAsImlhdCI6MTY0NTEwNzg2MH0.lSzIrb9_UwuXIC9T1JkozRgU62NrgtmKNOFwQUWGveQ','2022-02-16 06:46:29','2022-02-17 14:24:20'),(9,16,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjE2LCJOYW1lIjoidGVzdCIsIlJvbGUiOjMsImV4cCI6MTY0NjU5MjExOCwiaWF0IjoxNjQ2MzYxNzE4fQ.IZG68mf5jwL3GpGgT1RYY5x-GixQBYf8N0fAqDC29Dk','2022-02-17 06:53:25','2022-03-04 02:41:58'),(10,17,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjE3LCJOYW1lIjoidGVzdHQiLCJSb2xlIjoxLCJleHAiOjE2NDU4MzkwNzUsImlhdCI6MTY0NTYwODY3NX0.B4k0j26ghHBvMxwc8VRPfuYABn1HwPzpjOCDbuwQS8k','2022-02-23 09:31:15','2022-02-23 09:31:15'),(11,19,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjE5LCJOYW1lIjoic2VsbGVyIGEiLCJSb2xlIjoyLCJleHAiOjE2NDY0MzYwNzIsImlhdCI6MTY0NjIwNTY3Mn0.qr7Pi3kWmKeOvmoXbKl7fbvpon0tVBoKSIj5aL_dl6I','2022-03-02 07:21:12','2022-03-02 07:21:12'),(12,24,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjI0LCJOYW1lIjoic2VsbGVyZCIsIlJvbGUiOjIsImV4cCI6MTY0NjQ0NDc3NCwiaWF0IjoxNjQ2MjE0Mzc0fQ.n_us7p9bTIrazB2U2n-lPXe3VVfXYOFhfs6PQVAFwgY','2022-03-02 09:46:14','2022-03-02 09:46:14'),(13,30,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjMwLCJOYW1lIjoiY29iYVVzZXIiLCJSb2xlIjoyLCJleHAiOjE2NDY1OTI2NTMsImlhdCI6MTY0NjM2MjI1M30.W0joWmFYot6n06RznbAuF0GeQ58WHvHycq3-xCm9r50','2022-03-04 02:42:34','2022-03-04 02:50:53'),(14,31,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjMxLCJOYW1lIjoiY29iYVVzZXIxIiwiUm9sZSI6MiwiZXhwIjoxNjQ2NTkzMDU1LCJpYXQiOjE2NDYzNjI2NTV9.SX5ez5O6a7AzO3JKH6B8sif73Ngn_dfPt_2Dqv4-ZXw','2022-03-04 02:57:36','2022-03-04 02:57:36'),(15,41,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjQxLCJOYW1lIjoiY29iYVVzZXIxMSIsIlJvbGUiOjEsImV4cCI6MTY0NjU5NTIwMywiaWF0IjoxNjQ2MzY0ODAzfQ.0E1aNdRhKCuS0uvX2NLnPPLtDQXJ0v9jSMnfVwFkMFc','2022-03-04 03:33:23','2022-03-04 03:33:23');
/*!40000 ALTER TABLE `user_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(100) DEFAULT NULL,
  `username` varchar(25) NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `role` int DEFAULT '1',
  `status` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `last_login_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,NULL,'wan','$2a$10$AOBdxX8JckONoP6qKQHyoeoi3i62gaPc8Wu9W.J2Mb3Oi7JihoyOK',3,NULL,'2022-03-04 09:18:58','2022-02-14 06:00:09','2022-03-04 09:18:58'),(4,NULL,'tes user','$2a$10$CW2dh/Al/XwXGfIFCM5EOOdFf.Tt85rQ9qRUAE2Ot/j.voTw3tXeq',1,NULL,'2022-02-14 14:00:21','2022-02-14 09:49:11','2022-02-14 14:00:21'),(12,NULL,'costumer','$2a$10$DWSlFI58GNSBpleluwBH5eRatpHBEouDJkDqzQOIMLLaft6bnywJ.',1,NULL,'2022-02-20 09:16:28','2022-02-15 06:33:01','2022-02-20 09:16:28'),(14,NULL,'tes costummer','$2a$10$NEBXn3X4xRaZR6pJOQcmDusJ/lBNXxERcdhs1U955WBXcq9OTbzsa',1,NULL,'2022-02-15 14:09:17','2022-02-15 14:08:56','2022-02-15 14:09:17'),(16,NULL,'test','$2a$10$LUe78t.FTbNu.xjINe2x3ea6s54/ihk0qgJ6MAUNY.VRU6ABMEcli',3,'','2022-03-04 02:41:58','2022-02-17 06:52:57','2022-03-04 02:41:58'),(30,'coba@gcoba.com','cobaUser','$2a$10$qs5iaQi72dJyNcTX5d4p7OideEpSPKs6349lwHIhBm9remdWMVOHq',2,'Inactivated','2022-03-04 02:50:53','2022-03-04 02:35:28','2022-03-04 02:50:53'),(31,'coba@gcoba.com','cobaUser1','$2a$10$06zOhztq4XfXZ6qgugoEeujFGj19TbDXn2XIJRFFW27YcGtsAEhr2',2,'A','2022-03-04 02:57:36','2022-03-04 02:57:01','2022-03-04 02:57:36'),(41,'coba@gcoba.com','cobaUser11','$2a$10$DaOV7uA.MGyCFIbleaja1.2J27zxtqB5KA1ddzsdzxHqs3ZTAqdvO',1,'A','2022-03-04 03:33:23','2022-03-04 03:13:54','2022-03-04 03:33:23'),(42,'coba@gcoba.com','cobaUser12','$2a$10$VL5bytENDFmhMHyS9BbTx.BOzZrv2h9Vw3tkqWavcinZuhPZ1IppO',2,'Inactivated','2022-03-04 03:15:06','2022-03-04 03:15:06','2022-03-04 03:15:06');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `walets`
--

DROP TABLE IF EXISTS `walets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `walets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `saldo` int DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `walets`
--

LOCK TABLES `walets` WRITE;
/*!40000 ALTER TABLE `walets` DISABLE KEYS */;
INSERT INTO `walets` VALUES (1,1,9990000,'2022-02-15 04:11:45','2022-03-03 03:57:19'),(2,16,98090000,'2022-02-18 09:20:58','2022-03-03 09:21:07'),(3,2,3680000,'2022-02-20 13:16:57','2022-03-03 09:21:07'),(4,12,998540000,'2022-02-20 09:16:57','2022-02-22 15:48:04');
/*!40000 ALTER TABLE `walets` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-07 11:22:22
