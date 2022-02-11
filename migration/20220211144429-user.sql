
-- +migrate Up
CREATE TABLE `user` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user` varchar(25),
  `role` integer(11),
  `last_login_at` timestamp,
  `created_at` timestamp,
  `updated_at` timestamp
);
-- +migrate Down
