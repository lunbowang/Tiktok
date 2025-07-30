set names utf8mb4;
set foreign_key_checks =0;

-- 1. 用户表（无外键依赖，基础表）
CREATE TABLE `user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
                        `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
                        `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
                        `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录创建时间',
                        `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
                        PRIMARY KEY (`id`) USING BTREE,
                        UNIQUE INDEX `name` (`name`) USING BTREE COMMENT '用户名保证唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


-- 6. 关系表（依赖 user 表的 id 作为外键）
CREATE TABLE `relation` (
                            `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'follow关系id',
                            `user_id` int(11) NOT NULL COMMENT '用户id',
                            `following_id` int(11) NOT NULL COMMENT 'user_id关注的用户id',
                            `followed` int(11) NOT NULL COMMENT '默认0表示未关注，1表示已关注',
                            `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录创建时间',
                            `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
                            PRIMARY KEY (`id`) USING BTREE,
                            INDEX `fk_relation_user_1` (`user_id`) USING BTREE,
                            INDEX `fk_relation_user_2` (`following_id`) USING BTREE,
                            CONSTRAINT `fk_relation_user_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                            CONSTRAINT `fk_relation_user_2` FOREIGN KEY (`following_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;