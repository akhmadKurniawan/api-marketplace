
-- +migrate Up
CREATE TABLE `user` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `username` varchar(25),
  `password` varchar(255),
  `role` integer(11),
  `last_login_at` timestamp,
  `created_at` timestamp,
  `updated_at` timestamp
);
-- +migrate Down
