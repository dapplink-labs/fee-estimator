DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
        CREATE DOMAIN UINT256 AS NUMERIC
    CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
ELSE
    ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
    ALTER DOMAIN UINT256 ADD
        CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS gas_fee(
    guid            VARCHAR PRIMARY KEY,
    chain_name      VARCHAR UNIQUE NOT NULL,
    low             VARCHAR NOT NULL,
    normal          VARCHAR        NOT NULL,
    high            VARCHAR        NOT NULL,
    extra           VARCHAR        NOT NULL,
    timestamp       INTEGER        NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS gas_fee ON gas_fee (chain_name);


