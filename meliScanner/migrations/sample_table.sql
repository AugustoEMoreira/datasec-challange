CREATE TABLE `pii_table5` (
  `Name` int NOT NULL,
  `Passw` varchar(45) DEFAULT NULL,
  `Email` varchar(45) DEFAULT NULL,
  `first_name` varchar(45) DEFAULT NULL,
  `cc_number` varchar(45) DEFAULT NULL,
  `random_non_pii_field`varchar(45) DEFAULT NULL,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
