CREATE TABLE `lockers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `locker_number` int UNIQUE NOT NULL,
  `location` varchar(255) NOT NULL,
  `status` ENUM ('used', 'available', 'malfunction') NOT NULL DEFAULT 'available',
  `nfc_sig` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `last_modified` datetime DEFAULT (now())
);

CREATE TABLE `locker_user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `locker_id` int NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `last_modified` datetime DEFAULT (now())
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password_hashed` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `role` ENUM ('admin', 'customer') NOT NULL DEFAULT 'customer',
  `created_at` datetime DEFAULT (now()),
  `last_modified` datetime DEFAULT (now())
);
