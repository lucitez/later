BEGIN;

CREATE TABLE IF NOT EXISTS hostnames (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    hostname TEXT NOT NULL,
    content_type TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_hostnames BEFORE UPDATE ON hostnames
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE UNIQUE INDEX idx_uniq_hostnames_on_hostname ON hostnames(hostname)
WHERE deleted_at IS NULL;

COMMIT;