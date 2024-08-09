ALTER TABLE "public"."sms_template_config" RENAME COLUMN "template_no" TO "template_code";

COMMENT
ON COLUMN "public"."sms_template_config"."template_code" IS ''模版Code'';

ALTER TABLE "public"."sms_business_config" RENAME COLUMN "template_no" TO "template_code";

COMMENT
ON COLUMN "public"."sms_business_config"."template_code" IS ''模版Code'';

ALTER TABLE "public"."sms_template_config" RENAME COLUMN "third_party_template_no" TO "third_party_template_code";

COMMENT
ON COLUMN "public"."sms_template_config"."third_party_template_code" IS ''第三方模版Code'';


ALTER TABLE "public"."sms_app_config"
DROP
COLUMN "app_no",
  DROP
COLUMN "created_at",
  DROP
COLUMN "updated_at",
  DROP
COLUMN "deleted_at",
  ADD COLUMN "union_main_id" int8,
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "deleted_at" timestamp;

COMMENT
ON COLUMN "public"."sms_app_config"."union_main_id" IS ''所属主体id'';



ALTER TABLE "public"."sms_business_config"
DROP
COLUMN "created_at",
  DROP
COLUMN "updated_at",
  DROP
COLUMN "deleted_at",
  ADD COLUMN "union_main_id" int8,
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "deleted_at" timestamp;

COMMENT
ON COLUMN "public"."sms_business_config"."union_main_id" IS ''所属主体id'';



ALTER TABLE "public"."sms_send_log" RENAME COLUMN "app_no" TO "app_id";

ALTER TABLE "public"."sms_send_log"
DROP
COLUMN "status",
  DROP
COLUMN "created_at",
  DROP
COLUMN "updated_at",
  DROP
COLUMN "deleted_at",
  ADD COLUMN "union_main_id" int8,
  ADD COLUMN "form" int8,
  ADD COLUMN "type" int4 NOT NULL,
  ADD COLUMN "status" int4 NOT NULL,
  ADD COLUMN "meta_data" json,
  ADD COLUMN "sign_name" varchar(64) NOT NULL,
  ADD COLUMN "created_at" timestamp,
  ALTER
COLUMN "app_id" TYPE int8 USING "app_id"::int8,
  ALTER
COLUMN "app_id" SET NOT NULL,
  ALTER
COLUMN "business_no" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "fee" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "phone_number" TYPE varchar(16) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "phone_number" SET NOT NULL,
  ALTER
COLUMN "message" SET NOT NULL,
  ALTER
COLUMN "code" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "code" SET NOT NULL,
  ALTER
COLUMN "content" TYPE text COLLATE "pg_catalog"."default" USING "content"::text,
  ALTER
COLUMN "remark" TYPE text COLLATE "pg_catalog"."default" USING "remark"::text;

COMMENT
ON COLUMN "public"."sms_send_log"."union_main_id" IS ''所属主体id'';

COMMENT
ON COLUMN "public"."sms_send_log"."form" IS ''短信来源'';

COMMENT
ON COLUMN "public"."sms_send_log"."type" IS ''短信类型
：1验证
、2通知
、4业务
、8推广'';

COMMENT
ON COLUMN "public"."sms_send_log"."status" IS ''网关发送状态
：0失败
、1成功'';

COMMENT
ON COLUMN "public"."sms_send_log"."meta_data" IS ''网关返回元数据'';

COMMENT
ON COLUMN "public"."sms_send_log"."sign_name" IS ''签名名称'';



ALTER TABLE "public"."sms_app_config"
ALTER
COLUMN "app_name" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "app_name" SET NOT NULL,
  ALTER
COLUMN "available_number" SET NOT NULL,
  ALTER
COLUMN "current_limiting" SET NOT NULL,
  ALTER
COLUMN "remark" TYPE text COLLATE "pg_catalog"."default" USING "remark"::text;

COMMENT
ON COLUMN "public"."sms_app_config"."current_limiting" IS ''总条数'';


ALTER TABLE "public"."sms_business_config"
ALTER
COLUMN "app_no" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "app_no" SET NOT NULL,
  ALTER
COLUMN "business_name" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "business_name" SET NOT NULL,
  ALTER
COLUMN "business_no" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "business_no" SET NOT NULL,
  ALTER
COLUMN "template_code" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "business_desc" TYPE text COLLATE "pg_catalog"."default" USING "business_desc"::text,
  ALTER
COLUMN "remark" TYPE text COLLATE "pg_catalog"."default" USING "remark"::text;


ALTER TABLE "public"."sms_service_provider_config"
DROP
COLUMN "created_at",
  DROP
COLUMN "updated_at",
  DROP
COLUMN "deleted_at",
  ADD COLUMN "ext_json" json,
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "deleted_at" timestamp,
  ALTER
COLUMN "provider_no" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "provider_no" SET NOT NULL,
  ALTER
COLUMN "provider_name" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "provider_name" SET NOT NULL,
  ALTER
COLUMN "access_key_id" SET NOT NULL,
  ALTER
COLUMN "access_key_secret" SET NOT NULL,
  ALTER
COLUMN "sdk_app_id" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "region" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "remark" TYPE text COLLATE "pg_catalog"."default" USING "remark"::text;

COMMENT
ON COLUMN "public"."sms_service_provider_config"."ext_json" IS ''拓展字段'';


ALTER TABLE "public"."sms_sign_config"
DROP
COLUMN "created_at",
  DROP
COLUMN "updated_at",
  DROP
COLUMN "deleted_at",
  ADD COLUMN "ext_json" json,
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "deleted_at" timestamp,
  ALTER
COLUMN "sign_name" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "sign_name" SET NOT NULL,
  ALTER
COLUMN "provider_no" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "provider_no" SET NOT NULL,
  ALTER
COLUMN "provider_name" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "provider_name" SET NOT NULL,
  ALTER
COLUMN "remark" TYPE text COLLATE "pg_catalog"."default" USING "remark"::text,
  ALTER
COLUMN "status" SET NOT NULL,
  ALTER
COLUMN "audit_reply_msg" TYPE text COLLATE "pg_catalog"."default" USING "audit_reply_msg"::text;

COMMENT
ON COLUMN "public"."sms_sign_config"."ext_json" IS ''拓展字段'';


ALTER TABLE "public"."sms_sign_config"
DROP
COLUMN "created_at",
  DROP
COLUMN "updated_at",
  DROP
COLUMN "deleted_at",
  ADD COLUMN "union_main_id" int8,
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "deleted_at" timestamp;

COMMENT
ON COLUMN "public"."sms_sign_config"."union_main_id" IS ''关联主体ID'';



ALTER TABLE "public"."sms_template_config"
DROP
COLUMN "sign_name",
  DROP
COLUMN "created_at",
  DROP
COLUMN "updated_at",
  DROP
COLUMN "deleted_at",
  ADD COLUMN "ext_json" json,
  ADD COLUMN "union_main_id" int8,
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "deleted_at" timestamp,
  ALTER
COLUMN "template_code" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "template_code" SET NOT NULL,
  ALTER
COLUMN "template_name" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "template_name" SET NOT NULL,
  ALTER
COLUMN "template_content" TYPE text COLLATE "pg_catalog"."default" USING "template_content"::text,
  ALTER
COLUMN "template_content" SET NOT NULL,
  ALTER
COLUMN "third_party_template_code" TYPE varchar(64) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "provider_no" TYPE varchar(32) COLLATE "pg_catalog"."default",
  ALTER
COLUMN "provider_no" SET NOT NULL,
  ALTER
COLUMN "remark" TYPE text COLLATE "pg_catalog"."default" USING "remark"::text,
  ALTER
COLUMN "audit_reply_msg" TYPE text COLLATE "pg_catalog"."default" USING "audit_reply_msg"::text,
  ALTER
COLUMN "audit_at" TYPE timestamp USING "audit_at"::timestamp;

COMMENT
ON COLUMN "public"."sms_template_config"."ext_json" IS ''拓展字段'';

COMMENT
ON COLUMN "public"."sms_template_config"."union_main_id" IS ''关联主体ID'';