CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE CHECK (char_length(name) BETWEEN 3 AND 100),
    code TEXT NOT NULL UNIQUE CHECK (char_length(code) BETWEEN 3 AND 20),
    
    email TEXT NOT NULL CHECK (position('@' in email) > 1),
    phone TEXT CHECK (char_length(phone) BETWEEN 10 AND 15),
    
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);



CREATE TABLE tenant_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL UNIQUE REFERENCES tenants(id) ON DELETE CASCADE,

    legal_name TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL DEFAULT 'India',
    postal_code TEXT NOT NULL,

    gst_number TEXT UNIQUE,
    license_number TEXT UNIQUE,

    logo_url TEXT,
    website TEXT,

    emergency_contact TEXT,

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);



CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_set_updated_at_on_tenants
BEFORE UPDATE ON tenants
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trigger_set_updated_at_on_profiles
BEFORE UPDATE ON tenant_profiles
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

