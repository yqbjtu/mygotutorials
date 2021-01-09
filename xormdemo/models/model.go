package models

import "time"

type PoUser struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Username  string    `json:"username" xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Fullname  string    `json:"fullname" xorm:"not null default '' comment('联系人') VARCHAR(50)"`
	Mobile    string    `json:"mobile" xorm:"not null default '' comment('手机号') VARCHAR(50)"`
	Address   string    `json:"address" xorm:"not null default '' comment('联系地址') VARCHAR(255)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

/*
DROP TABLE IF EXISTS `po_user`;
CREATE TABLE `po_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `fullname` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人',
  `mobile` varchar(50) NOT NULL DEFAULT '' COMMENT '手机号',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '联系地址',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
*/
