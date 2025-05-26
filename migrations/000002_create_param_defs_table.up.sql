CREATE TABLE IF NOT EXISTS steam_param_defs (
  key TEXT PRIMARY KEY,
  label TEXT NOT NULL,
  type TEXT NOT NULL,
  options JSONB DEFAULT '[]',
  default_value TEXT,
  help_text TEXT
);

-- Seed common params
INSERT INTO steam_param_defs(key, label, type, options, default_value, help_text) VALUES
  ('appid', 'App ID', 'number', '[]', '440', 'Steam application ID, e.g. 440 for TF2'),
  ('query', 'Search Term', 'string', '[]', '', 'Text to filter item names'),
  ('currency', 'Currency', 'enum', '[{"value":"1","label":"USD"},{"value":"3","label":"GBP"}]', '1', 'Currency code for price display');