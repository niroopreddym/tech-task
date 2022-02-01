GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;

-- lookup tables and the static data insertion part
CREATE TABLE Bank (
    bank_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    bank_uuid uuid,
    bank_name varchar(50),
    ifsc_code varchar(50),
    branch_name varchar(50)
);

CREATE TABLE Account (
    account_id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    bank_name varchar(50),
    branch_name varchar(50),
    account_holder_name varchar(50),
    identity_id uuid,
    first_name varchar(50),
    last_name varchar(50),
    acc_holder_addr varchar(200)
);