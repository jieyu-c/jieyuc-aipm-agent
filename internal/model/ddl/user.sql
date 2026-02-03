-- 用户表 (users)
CREATE TABLE "aipm"."users" (
                                "id" bigserial NOT NULL,
                                "user_id" varchar(255) NOT NULL,
                                "phone" varchar(32) NOT NULL DEFAULT '',
                                "username" varchar(255) NOT NULL,
                                "password" varchar(255) NOT NULL,
                                "nickname" varchar(255) NOT NULL DEFAULT '',
                                "avatar" varchar(255) NOT NULL DEFAULT '',
                                "gender" smallint NOT NULL DEFAULT 0,
                                "status" smallint NOT NULL DEFAULT 0,
                                "created_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                "updated_at" timestamptz(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                "deleted_at" timestamptz(6),
                                PRIMARY KEY ("id"),
                                UNIQUE ("user_id"),
                                UNIQUE ("username"),
                                UNIQUE ("phone")
);

-- 添加字段注释
COMMENT ON COLUMN "aipm"."users"."id" IS '自增主键';
COMMENT ON COLUMN "aipm"."users"."user_id" IS '用户唯一ID';
COMMENT ON COLUMN "aipm"."users"."username" IS '用户名';
COMMENT ON COLUMN "aipm"."users"."password" IS '加密后的密码';
COMMENT ON COLUMN "aipm"."users"."nickname" IS '昵称';
COMMENT ON COLUMN "aipm"."users"."avatar" IS '头像URL';
COMMENT ON COLUMN "aipm"."users"."gender" IS '性别 0:未知 1:男 2:女';
COMMENT ON COLUMN "aipm"."users"."phone" IS '手机号码';
COMMENT ON COLUMN "aipm"."users"."status" IS '状态 0:正常 1:禁用';
COMMENT ON COLUMN "aipm"."users"."created_at" IS '创建时间';
COMMENT ON COLUMN "aipm"."users"."updated_at" IS '更新时间';
COMMENT ON COLUMN "aipm"."users"."deleted_at" IS '删除时间 (用于软删除)';

-- 添加表注释
COMMENT ON TABLE "aipm"."users" IS '用户表';

-- 创建一个触发器函数，用于在行更新时自动更新 updated_at 字段
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 将触发器绑定到 users 表
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON "aipm"."users"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();