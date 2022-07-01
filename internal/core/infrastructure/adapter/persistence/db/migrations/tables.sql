CREATE TABLE IF NOT EXISTS business (
    business_id     VARCHAR(36) PRIMARY KEY,
	name            VARCHAR(30) NOT NULL,
	description     VARCHAR(30) NOT NULL,
	public_code     VARCHAR(30),
    reg_date        DATE
);
