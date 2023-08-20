CREATE TABLE `datastorage` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `host` varchar(256) NOT NULL,
  `port` int NOT NULL,
  `user` varchar(256) NOT NULL,
  `pass` varchar(256) NOT NULL,
  `classification` json DEFAULT NULL,
  `db` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `rules` (
  `idrules` int NOT NULL AUTO_INCREMENT,
  `regex` varchar(512) DEFAULT NULL,
  `annotation` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`idrules`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
