CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,

    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    name TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    UNIQUE (tenant_id, email)
);


CREATE TRIGGER trigger_set_updated_at_on_profiles
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();