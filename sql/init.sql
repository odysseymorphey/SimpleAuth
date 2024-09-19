CREATE TABLE IF NOT EXISTS refresh_tokens (
    id bigserrial PRIMARY KEY,
    GUID TEXT NOT NULL,
    user_ip varchar(20) NOT NULL,
    hashed_token TEXT NOT NULL,
    pair_id TEXT NOT NULL,
);