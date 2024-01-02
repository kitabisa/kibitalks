-- +migrate Up
CREATE TABLE IF NOT EXISTS `donations` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `amount` VARCHAR(100) NOT NULL,
  `payment_method_id` VARCHAR(45) NOT NULL,
  `campaign_id` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`id`)
);

-- +migrate Down
DROP table `donations`;