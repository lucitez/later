CREATE TABLE IF NOT EXISTS domains (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    domain TEXT NOT NULL,
    content_type TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_domains BEFORE INSERT OR UPDATE ON domains
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

CREATE UNIQUE INDEX idx_uniq_domains_on_domain ON domains(domain)
WHERE deleted_at IS NULL;