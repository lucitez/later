CREATE TABLE user_sessions (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    expires_at timestamp with time zone NOT NULL,
    expired_at timestamp with time zone -- For manual expiry of tokens when user refreshes or logs out. Minimize hacker damage
);

CREATE INDEX idx_user_sessions_on_user_id_and_expired_at ON user_sessions(user_id, expired_at)