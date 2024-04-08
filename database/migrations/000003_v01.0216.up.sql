ALTER TABLE "public"."sms_send_log"
    ALTER COLUMN "message" DROP NOT NULL,
ALTER COLUMN "code" DROP NOT NULL;


ALTER TABLE "public"."sms_send_log"
    ALTER COLUMN "app_id" DROP NOT NULL;

ALTER TABLE "public"."sms_template_config"
DROP COLUMN "created_at",
  DROP COLUMN "updated_at",
  DROP COLUMN "deleted_at",
  ADD COLUMN "sign_name" varchar(64),
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "deleted_at" timestamp;

COMMENT ON COLUMN "public"."sms_template_config"."sign_name" IS '签名名称';