-- +goose Up
CREATE TABLE `object` (
    `id` VARBINARY(36) NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `content` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB, DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE `object`;
