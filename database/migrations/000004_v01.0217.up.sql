ALTER TABLE "public"."sms_send_log"
ALTER COLUMN "fee" TYPE int4 USING "fee"::int4;

COMMENT ON COLUMN "public"."sms_send_log"."fee" IS '条数';