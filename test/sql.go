

CREATE TABLE `sys_project` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`user_id` bigint  NOT NULL COMMENT '用户ID',
`created_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`deleted_at` datetime  DEFAULT NULL,
`pro_uuid` varchar(191) COLLATE utf8mb4_general_ci NOT NULL COMMENT '项目UUID',
`pro_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '项目名称',
`app_key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'appKeyid',
`app_secret` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'AppSecret',
`settle_currency` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'settleCurrency',
`assemble_chain` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'assembleChain',
`assemble_address` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'assembleAddress',
`callback_domain` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'callbackDomain',
`callback_url` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'callbackUrl',
`status`  bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
PRIMARY KEY (`id`),
KEY `idx_sys_project_user_id` (`user_id`),
KEY `idx_sys_project_uuid` (`pro_uuid`),
KEY `idx_sys_project_assembleAddress` (`assemble_address`),
UNIQUE KEY unique_index_name (`user_id`, `assemble_chain`,`settle_currency`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




CREATE TABLE `user_wallet_address` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`merchantId`   bigint  NOT NULL COMMENT '用户ID',
`created_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`deleted_at` datetime DEFAULT NULL,
`name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Name',
`address` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Address',
`chain` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Chain',
`is_internal` bool COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'IsInternal',
PRIMARY KEY (`id`),
KEY `idx_user_id` (`merchantId`),
KEY `idx_address` (`address`),
KEY `idx_chain` (`chain`),
UNIQUE KEY unique_index_name (`merchantId`, `address`,`chain`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



CREATE TABLE `user_bank` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`merchantId` bigint    NOT NULL COMMENT '用户ID',
`created_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`deleted_at` datetime  DEFAULT NULL,
`region` varchar(191) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'region',
`remittanceType` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'remittanceType',
`enterpriseTitle` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'enterpriseTitle',
`bankTitle` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'bankTitle',
`bankCode` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'bankCode',
`fedWire` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'fedWire',
`receiverNumber` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'receiverNumber',
`receiverName` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'receiverName',
`receiverAddress` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'receiverAddress',
`referenceField` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'referenceField',
PRIMARY KEY (`id`),
KEY `idx_merchantId` (`merchantId`),
KEY `idx_senterpriseTitle` (`enterpriseTitle`),
KEY `idx_bankTitle` (`bankTitle`),
KEY `idx_bankCode` (`bankCode`),
UNIQUE KEY unique_index_name (`merchantId`,`enterpriseTitle`, `bankTitle`,`bankCode`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




CREATE TABLE `user_withdraw_order` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`merchantId` bigint  NOT NULL COMMENT '用户ID',
`merchantName` varchar(128)  COLLATE utf8mb4_general_ci  NOT NULL COMMENT 'merchantName',
`bankId` varchar(128)  COLLATE utf8mb4_general_ci  NOT NULL COMMENT 'bankId',
`bankTitle` varchar(256)  COLLATE utf8mb4_general_ci  NOT NULL COMMENT 'bankTitle',
`created_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`deleted_at` datetime  DEFAULT NULL,
`currency` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'currency',
`chain` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'chain',
`amount` DOUBLE(16, 5)  DEFAULT NULL COMMENT 'amount',
`fee`    DOUBLE(16, 5)  DEFAULT NULL COMMENT 'fee',
`remittanceFee` DOUBLE(16, 5)  DEFAULT NULL COMMENT 'remittanceFee',
`totalAmount` DOUBLE(16, 5)  DEFAULT NULL COMMENT 'totalAmount',
`status` int   DEFAULT 1  COMMENT 'status 1: 完成,2:待审核 3:审核完成 4: 进行中 5: 完成',
`statusName` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'statusName',
`txInfo` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'txInfo',
`txCertificationUrl` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'txCertificationUrl',
`txReference` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'txReference',
`inputNote`varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'inputNote',
`supplier` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'supplier',
`adminMemo` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'adminMemo',
`memo` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'memo',
`assemble_address` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'assembleAddress',
PRIMARY KEY (`id`),
KEY `idx_merchantId` (`merchantId`),
KEY `idx_bankId` (`bankId`),
KEY `idx_currency` (`currency`),
KEY `idx_chain` (`chain`),
UNIQUE KEY unique_index_name (`merchantId`, `bankId`,`currency`,`chain`,`created_at` )
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



