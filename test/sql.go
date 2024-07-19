

CREATE TABLE `sys_project` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`user_id` bigint  NOT NULL COMMENT '用户ID',
`created_at` datetime(3) DEFAULT NULL,
`updated_at` datetime(3) DEFAULT NULL,
`deleted_at` datetime(3) DEFAULT NULL,
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
UNIQUE KEY unique_index_name (`user_id`, `pro_uuid`,`assemble_address`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



CREATE TABLE `sys_project` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`user_id` bigint  NOT NULL COMMENT '用户ID',
`created_at` datetime(3) DEFAULT NULL,
`updated_at` datetime(3) DEFAULT NULL,
`deleted_at` datetime(3) DEFAULT NULL,
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
UNIQUE KEY unique_index_name (`user_id`, `pro_uuid`,`assemble_address`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

