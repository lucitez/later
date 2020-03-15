BEGIN;

CREATE TABLE friends (
    id              uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id         uuid NOT NULL,
    friend_user_id  uuid NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX idx_uniq_friends_on_user_id_and_friend_id
ON friends (user_id, friend_user_id)
WHERE deleted_at IS NULL;

CREATE TRIGGER update_friends
BEFORE INSERT OR UPDATE ON friends
FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE TABLE friend_requests (
    id uuid                 PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    sent_by_user_id         uuid NOT NULL,
    recipient_user_id       uuid NOT NULL,

    created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    accepted_at TIMESTAMP WITH TIME ZONE,
    declined_at TIMESTAMP WITH TIME ZONE,
    deleted_at  TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX idx_uniq_friend_requests_on_sent_by_and_recipient
ON friend_requests(sent_by_user_id, recipient_user_id)
WHERE declined_at IS NULL
AND deleted_at IS NULL;

CREATE TRIGGER update_friend_requests
BEFORE INSERT OR UPDATE ON friend_requests
FOR EACH ROW EXECUTE FUNCTION update_updated_at();

COMMIT;