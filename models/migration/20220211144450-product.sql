
-- +migrate Up

CREATE TABLE `product` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `product_type` integer(21),
  `shop_id` integer(11),
  `name` varchar(50),
  `price` varchar(25),
  `description` text,
  `qty` varchar(25),
  `image` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);
-- +migrate Down
