CREATE DATABASE link_shortener;

CREATE TABLE IF NOT EXISTS shorteners (
    uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    original_url varchar(256) NOT NULL,
    shorten_url varchar(256) NOT NULL,
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS shortener_logs (
    shortener_id uuid PRIMARY KEY,
    new_url varchar(256) NOT NULL,
    created_at timestamptz DEFAULT NOW()
);