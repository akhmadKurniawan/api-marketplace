
-- +migrate Up

CREATE TABLE `shop` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(50),
  `description` text,
  `alamat` text,
  `logo` varchar(255),
  `saler_id` integer(11),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `shop` ADD FOREIGN KEY (`id`) REFERENCES `product` (`shop_id`);

ALTER TABLE `shop` ADD FOREIGN KEY (`saler_id`) REFERENCES `saler` (`id`);



-- +migrate Down
