
-- +migrate Up
CREATE TABLE `transaction` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` integer(11),
  `product_id` varchar(50),
  `type` enum,
  `description` text,
  `amount` integer(25),
  `total_product` varchar(50),
  `created_at` timestamp,
  `updated_at` timestamp
);


ALTER TABLE `transaction` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

-- +migrate Down
