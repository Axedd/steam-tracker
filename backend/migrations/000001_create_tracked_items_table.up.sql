--  migrate.exe -path migrations -database $Env:DATABASE_URL up
CREATE TABLE IF NOT EXISTS tracked_items (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  query TEXT NOT NULL,
  filters JSONB NOT NULL,
  sent_ids TEXT[] NOT NULL DEFAULT '{}',
  active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);