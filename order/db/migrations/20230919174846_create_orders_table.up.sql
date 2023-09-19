CREATE TABLE `order` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `parent_order_id` int(11) DEFAULT NULL,
                         `user_id` int(11) NOT NULL,
                         `product_id` int(11) NOT NULL,
                         `quantity` int(11) NOT NULL,
                         `price` decimal(10,2) NOT NULL,
                         `created_at` datetime NOT NULL,
                         `updated_at` datetime NOT NULL,
                         `status` enum('created', 'pending','completed','cancelled', 'unknown') NOT NULL DEFAULT 'created',
                         PRIMARY KEY (`id`),
                         KEY `user_id` (`user_id`),
                         KEY `product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;