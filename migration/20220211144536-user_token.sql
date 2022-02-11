
-- +migrate Up
CREATE TABLE `user_token` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer(11),
  `token` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `user_token` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

-- +migrate Down
