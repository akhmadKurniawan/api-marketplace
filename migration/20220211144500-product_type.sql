
-- +migrate Up
CREATE TABLE `product_type` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(25),
  `description` text,
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `product_type` ADD FOREIGN KEY (`id`) REFERENCES `product` (`product_type`);

-- +migrate Down
