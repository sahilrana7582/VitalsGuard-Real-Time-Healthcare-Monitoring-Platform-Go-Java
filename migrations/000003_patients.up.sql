CREATE TABLE IF NOT EXISTS patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,

    full_name TEXT NOT NULL CHECK (char_length(full_name) <= 150),
    age       INTEGER NOT NULL,
    admission_reason TEXT NOT NULL,


    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);



CREATE TABLE IF NOT EXISTS patient_tables (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    patient_id UUID NOT NULL UNIQUE REFERENCES patients(id) ON DELETE CASCADE,

    gender TEXT NOT NULL CHECK (gender IN ('Male', 'Female', 'Other')),
    dob DATE NOT NULL,
    blood_group TEXT NOT NULL CHECK (
        blood_group IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-')
    ),

    contact_number TEXT NOT NULL CHECK (char_length(contact_number) <= 15),
    email TEXT UNIQUE CHECK (position('@' IN email) > 1),
    
    address TEXT,
    postal_code TEXT CHECK (char_length(postal_code) <= 20),

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);




CREATE TABLE IF NOT EXISTS vitals_info (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,

    heart_rate INTEGER NOT NULL CHECK (heart_rate BETWEEN 30 AND 220),
    blood_pressure_upper INTEGER CHECK (blood_pressure_upper BETWEEN 70 AND 250),
    blood_pressure_lower INTEGER CHECK (blood_pressure_lower BETWEEN 40 AND 150),
    respiratory_rate INTEGER CHECK (respiratory_rate BETWEEN 5 AND 60),
    oxygen_saturation INTEGER CHECK (oxygen_saturation BETWEEN 50 AND 100),
    temperature DECIMAL(4,1) CHECK (temperature BETWEEN 90.0 AND 110.0),
    glucose_level REAL,

    recorded_at TIMESTAMP NOT NULL DEFAULT now(), 

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);


