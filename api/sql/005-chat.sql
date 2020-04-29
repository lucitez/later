CREATE TABLE chats (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    group_id uuid,
    user1_id uuid,
    user2_id uuid,

    CHECK (user1_id != user2_id),

    created_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE OR REPLACE FUNCTION user_chat_uniqueness_check() RETURNS TRIGGER AS $user_chat_uniqueness_check$
    BEGIN
        IF EXISTS (
            SELECT * FROM chats
            WHERE user2_id = NEW.user1_id
            AND user1_id = NEW.user2_id
            AND deleted_at IS NULL
        )
        THEN RAISE EXCEPTION 'existing chat for this user pair';
        END IF;
        RETURN NEW;
    END
$user_chat_uniqueness_check$ LANGUAGE plpgsql;

CREATE TRIGGER check_user_chat_uniqueness
BEFORE INSERT ON chats
FOR EACH ROW EXECUTE PROCEDURE user_chat_uniqueness_check();

CREATE UNIQUE INDEX idx_uniq_chats_on_group_id ON chats(group_id)
WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX idx_uniq_chats_on_user_ids ON chats(user1_id, user2_id)
WHERE deleted_at IS NULL;

CREATE TABLE messages (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    chat_id uuid NOT NULL,
    sent_by uuid NOT NULL,
    message text,
    content_id uuid,

    created_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE INDEX idx_messages_on_chat_id
ON messages(chat_id)
WHERE deleted_at IS NULL;

-- Used to populate a user's convnersation feed. Select distinct chat_id from
-- user_messages where user_id=$1 order by created_at [where read_at is null];
CREATE TABLE user_messages (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    message_id uuid NOT NULL,
    chat_id uuid NOT NULL,

    read_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE INDEX idx_user_messages_on_user_id
ON user_messages(user_id)
WHERE deleted_at IS NULL;