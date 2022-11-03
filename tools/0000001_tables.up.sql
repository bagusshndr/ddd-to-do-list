CREATE TABLE IF NOT EXISTS `activities` (
  `id` int unsigned AUTO_INCREMENT,
  `email` varchar(128) NOT NULL,
  `title` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;