BEGIN;

ALTER TABLE chats ADD COLUMN last_message_sent_at timestamp with time zone;

CREATE OR REPLACE FUNCTION update_chat_last_message() RETURNS TRIGGER AS $update_chat_last_message$
    BEGIN
        UPDATE chats c SET last_message_sent_at = now()
        WHERE c.id = NEW.chat_id;
        RETURN NEW;
    END
$update_chat_last_message$ LANGUAGE plpgsql;

CREATE TRIGGER update_chat_last_message
AFTER INSERT ON messages
FOR EACH ROW EXECUTE PROCEDURE update_chat_last_message();

DROP INDEX IF EXISTS idx_user_messages_on_user_id;

CREATE INDEX idx_unread_user_messages
ON user_messages(user_id, chat_id)
WHERE read_at IS NULL AND deleted_at IS NULL;

COMMIT;