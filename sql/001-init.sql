BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_updated_at() RETURNS TRIGGER AS $update_updated_at$
    BEGIN
        NEW.updated_at = now();
        RETURN NEW;
    END;
$update_updated_at$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS users
(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    username text NOT NULL,
    email text,
    phone_number text,

    created_at timestamp with time zone NOT NULL default now(),
    signed_up_at timestamp with time zone,
    updated_at timestamp with time zone NOT NULL default now(),
    deleted_at timestamp with time zone
);

CREATE TRIGGER update_users BEFORE INSERT OR UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE UNIQUE INDEX IF NOT EXISTS idx_uniq_users_on_username ON users(username) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_uniq_users_on_email ON users(email) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_uniq_users_on_phone_number ON users(phone_number) WHERE deleted_at IS NULL;

-- Used to keep transaction of all shares of content back and forth.
-- Used when looking at share history between two users
CREATE TABLE IF NOT EXISTS shares
(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    content_id uuid NOT NULL,
    sender uuid NOT NULL,
    reciever uuid NOT NULL,
    is_read boolean NOT NULL DEFAULT false,

    sent_at timestamp with time zone NOT NULL DEFAULT now(),
    opened_at timestamp with time zone
);

-- Metadata about content being shared. Content gets passed along. The same article may be shared by different people
-- but passing along only works when person A sends article X to person B and B sends to C (linked by content.id)
CREATE TABLE IF NOT EXISTS content (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    description TEXT,
    image_url TEXT,
    content_type TEXT,
    url TEXT NOT NULL,
    domain TEXT NOT NULL,
    shares int NOT NULL DEFAULT 0 CHECK (shares >= 0),

    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now()
);

CREATE TRIGGER update_content BEFORE INSERT OR UPDATE ON content
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

-- Used to generate a user's feed. Comprised of references to shares.
-- Shares are never deleted (so you can see what you have sent other people)
-- User content can be deleted or archived
CREATE TABLE IF NOT EXISTS user_content (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    share_id uuid NOT NULL,
    content_id uuid NOT NULL,
    user_id uuid NOT NULL,
    sender_type text NOT NULL, -- [self, friend, us]

    created_at timestamp with time zone NOT NULL default now(),
    updated_at timestamp with time zone NOT NULL default now(),
    archived_at timestamp with time zone,
    deleted_at timestamp with time zone
);

CREATE TRIGGER update_user_content BEFORE INSERT OR UPDATE ON user_content
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

COMMIT;