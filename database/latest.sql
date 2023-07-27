/*
 Navicat Premium Data Transfer

 Source Server         : 企迅云
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 127.0.0.1:5432
 Source Catalog        :
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 27/07/2023 10:28:25
*/


-- ----------------------------
-- Table structure for sms_app_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_app_config";
CREATE TABLE "public"."sms_app_config" (
  "id" int8 NOT NULL,
  "app_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "available_number" int4 NOT NULL,
  "total_number" int4 NOT NULL,
  "use_number" int4,
  "remark" text COLLATE "pg_catalog"."default",
  "status" int4,
  "union_main_id" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."sms_app_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."sms_app_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_app_config"."app_name" IS '应用名称';
COMMENT ON COLUMN "public"."sms_app_config"."available_number" IS '可用数量';
COMMENT ON COLUMN "public"."sms_app_config"."total_number" IS '总条数';
COMMENT ON COLUMN "public"."sms_app_config"."use_number" IS '已用数量';
COMMENT ON COLUMN "public"."sms_app_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_app_config"."status" IS '状态: 0禁用 1正常';
COMMENT ON COLUMN "public"."sms_app_config"."union_main_id" IS '所属主体id';

-- ----------------------------
-- Table structure for sms_business_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_business_config";
CREATE TABLE "public"."sms_business_config" (
  "id" int8 NOT NULL,
  "app_id" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "business_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "business_no" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "template_code" varchar(64) COLLATE "pg_catalog"."default",
  "business_desc" text COLLATE "pg_catalog"."default",
  "remark" text COLLATE "pg_catalog"."default",
  "status" int4,
  "union_main_id" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."sms_business_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."sms_business_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_business_config"."app_id" IS '应用ID';
COMMENT ON COLUMN "public"."sms_business_config"."business_name" IS '业务名称';
COMMENT ON COLUMN "public"."sms_business_config"."business_no" IS '业务编号';
COMMENT ON COLUMN "public"."sms_business_config"."template_code" IS '模版Code';
COMMENT ON COLUMN "public"."sms_business_config"."business_desc" IS '业务说明';
COMMENT ON COLUMN "public"."sms_business_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_business_config"."status" IS '状态: 0禁用 1正常';
COMMENT ON COLUMN "public"."sms_business_config"."union_main_id" IS '所属主体id';

-- ----------------------------
-- Table structure for sms_send_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_send_log";
CREATE TABLE "public"."sms_send_log" (
  "id" int8 NOT NULL,
  "app_id" int8,
  "business_no" varchar(32) COLLATE "pg_catalog"."default",
  "fee" int4,
  "phone_number" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "message" varchar(255) COLLATE "pg_catalog"."default",
  "code" varchar(64) COLLATE "pg_catalog"."default",
  "content" text COLLATE "pg_catalog"."default",
  "remark" text COLLATE "pg_catalog"."default",
  "union_main_id" int8,
  "form" int8,
  "type" int4 NOT NULL,
  "status" int4 NOT NULL,
  "meta_data" json,
  "sign_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6)
)
;
ALTER TABLE "public"."sms_send_log" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."sms_send_log"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_send_log"."app_id" IS '应用ID';
COMMENT ON COLUMN "public"."sms_send_log"."business_no" IS '业务编号';
COMMENT ON COLUMN "public"."sms_send_log"."fee" IS '条数';
COMMENT ON COLUMN "public"."sms_send_log"."phone_number" IS '发送手机号';
COMMENT ON COLUMN "public"."sms_send_log"."message" IS '接口响应消息';
COMMENT ON COLUMN "public"."sms_send_log"."code" IS '接口响应状态码';
COMMENT ON COLUMN "public"."sms_send_log"."content" IS '发送内容';
COMMENT ON COLUMN "public"."sms_send_log"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_send_log"."union_main_id" IS '所属主体id';
COMMENT ON COLUMN "public"."sms_send_log"."form" IS '短信来源';
COMMENT ON COLUMN "public"."sms_send_log"."type" IS '短信类型：1验证、2通知、4业务、8推广';
COMMENT ON COLUMN "public"."sms_send_log"."status" IS '网关发送状态：0失败、1成功';
COMMENT ON COLUMN "public"."sms_send_log"."meta_data" IS '网关返回元数据';
COMMENT ON COLUMN "public"."sms_send_log"."sign_name" IS '签名名称';

-- ----------------------------
-- Table structure for sms_service_provider_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_service_provider_config";
CREATE TABLE "public"."sms_service_provider_config" (
  "id" int8 NOT NULL,
  "provider_no" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "provider_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "access_key_id" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "access_key_secret" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "endpoint" varchar(255) COLLATE "pg_catalog"."default",
  "sdk_app_id" varchar(64) COLLATE "pg_catalog"."default",
  "region" varchar(32) COLLATE "pg_catalog"."default",
  "remark" text COLLATE "pg_catalog"."default",
  "status" int4,
  "ext_json" json,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "priority" int2,
  "is_default" bool DEFAULT false
)
;
ALTER TABLE "public"."sms_service_provider_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."sms_service_provider_config"."id" IS '渠道商id';
COMMENT ON COLUMN "public"."sms_service_provider_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."sms_service_provider_config"."provider_name" IS '渠道商名字';
COMMENT ON COLUMN "public"."sms_service_provider_config"."access_key_id" IS '身份标识';
COMMENT ON COLUMN "public"."sms_service_provider_config"."access_key_secret" IS '身份认证密钥';
COMMENT ON COLUMN "public"."sms_service_provider_config"."endpoint" IS '域名调用';
COMMENT ON COLUMN "public"."sms_service_provider_config"."sdk_app_id" IS '应用id';
COMMENT ON COLUMN "public"."sms_service_provider_config"."region" IS '地域';
COMMENT ON COLUMN "public"."sms_service_provider_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_service_provider_config"."status" IS '状态: 0禁用 1正常';
COMMENT ON COLUMN "public"."sms_service_provider_config"."ext_json" IS '拓展字段';
COMMENT ON COLUMN "public"."sms_service_provider_config"."priority" IS '优先级，使用默认选择优先级最高的';
COMMENT ON COLUMN "public"."sms_service_provider_config"."is_default" IS '是否默认：true是、false否 ，默认false';

-- ----------------------------
-- Table structure for sms_sign_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_sign_config";
CREATE TABLE "public"."sms_sign_config" (
  "id" int8 NOT NULL,
  "sign_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "provider_no" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "provider_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "remark" text COLLATE "pg_catalog"."default",
  "status" int4 NOT NULL,
  "audit_user_id" int8,
  "audit_reply_msg" text COLLATE "pg_catalog"."default",
  "audit_at" timestamptz(6),
  "ext_json" json,
  "union_main_id" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."sms_sign_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."sms_sign_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_sign_config"."sign_name" IS '短信签名名称';
COMMENT ON COLUMN "public"."sms_sign_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."sms_sign_config"."provider_name" IS '渠道商名字';
COMMENT ON COLUMN "public"."sms_sign_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_sign_config"."status" IS '状态: -1不通过 0待审核 1正常';
COMMENT ON COLUMN "public"."sms_sign_config"."audit_user_id" IS '审核者UserID';
COMMENT ON COLUMN "public"."sms_sign_config"."audit_reply_msg" IS '审核回复，仅审核不通过时才有值';
COMMENT ON COLUMN "public"."sms_sign_config"."audit_at" IS '审核时间';
COMMENT ON COLUMN "public"."sms_sign_config"."ext_json" IS '拓展字段';
COMMENT ON COLUMN "public"."sms_sign_config"."union_main_id" IS '关联主体ID';

-- ----------------------------
-- Table structure for sms_template_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_template_config";
CREATE TABLE "public"."sms_template_config" (
  "id" int8 NOT NULL,
  "template_code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "template_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "template_content" text COLLATE "pg_catalog"."default" NOT NULL,
  "third_party_template_code" varchar(64) COLLATE "pg_catalog"."default",
  "provider_no" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "remark" text COLLATE "pg_catalog"."default",
  "status" int4,
  "audit_user_id" int8,
  "audit_reply_msg" text COLLATE "pg_catalog"."default",
  "audit_at" timestamptz(6),
  "ext_json" json,
  "union_main_id" int8,
  "sign_name" varchar(64) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "type" int2
)
;
ALTER TABLE "public"."sms_template_config" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."sms_template_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_template_config"."template_code" IS '模版Code';
COMMENT ON COLUMN "public"."sms_template_config"."template_name" IS '模版名称';
COMMENT ON COLUMN "public"."sms_template_config"."template_content" IS '模版内容';
COMMENT ON COLUMN "public"."sms_template_config"."third_party_template_code" IS '第三方模版Code';
COMMENT ON COLUMN "public"."sms_template_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."sms_template_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_template_config"."status" IS '状态: 0禁用 1正常';
COMMENT ON COLUMN "public"."sms_template_config"."audit_user_id" IS '审核者UserID
审核者UserID
';
COMMENT ON COLUMN "public"."sms_template_config"."audit_reply_msg" IS '审核回复，仅审核不通过时才有值';
COMMENT ON COLUMN "public"."sms_template_config"."audit_at" IS '审核时间';
COMMENT ON COLUMN "public"."sms_template_config"."ext_json" IS '拓展字段';
COMMENT ON COLUMN "public"."sms_template_config"."union_main_id" IS '关联主体ID';
COMMENT ON COLUMN "public"."sms_template_config"."sign_name" IS '签名名称';
COMMENT ON COLUMN "public"."sms_template_config"."type" IS '业务场景类型：1注册，2登录，4找回用户名/修改用户名，8找回密码/重置密码，16设置手机号码';

-- ----------------------------
-- Primary Key structure for table sms_app_config
-- ----------------------------
ALTER TABLE "public"."sms_app_config" ADD CONSTRAINT "sms_app_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sms_business_config
-- ----------------------------
ALTER TABLE "public"."sms_business_config" ADD CONSTRAINT "sms_business_config_pkey1" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sms_send_log
-- ----------------------------
ALTER TABLE "public"."sms_send_log" ADD CONSTRAINT "sms_send_log_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sms_service_provider_config
-- ----------------------------
ALTER TABLE "public"."sms_service_provider_config" ADD CONSTRAINT "sms_service_provider_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sms_sign_config
-- ----------------------------
ALTER TABLE "public"."sms_sign_config" ADD CONSTRAINT "sms_business_config_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sms_template_config
-- ----------------------------
ALTER TABLE "public"."sms_template_config" ADD CONSTRAINT "sms_template_config_pkey" PRIMARY KEY ("id");
