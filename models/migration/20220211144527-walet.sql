
-- +migrate Up

CREATE TABLE `walet` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer(11),
  `saldo` integer(20),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `walet` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);


-- +migrate Down
