
-- +migrate Up
CREATE TABLE `costumer` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer(21),
  `name` varchar(25),
  `jenis_kelamin` enum,
  `alamat` varchar(50),
  `no_hp` varchar(25),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `costumer` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

-- +migrate Down
