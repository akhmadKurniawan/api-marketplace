
-- +migrate Up
CREATE TABLE `saler` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer(21),
  `name` varchar(25),
  `jenis_kelamin` enum,
  `alamat` varchar(50),
  `no_hp` varchar(50),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `saler` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

-- +migrate Down
