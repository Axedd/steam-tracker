-- Make sure column exists
ALTER TABLE steam_param_defs
ADD COLUMN IF NOT EXISTS appid INT;

-- Add the foreign key constraint
ALTER TABLE steam_param_defs
ADD CONSTRAINT fk_param_appid
FOREIGN KEY (appid)
REFERENCES appids(appid)
ON UPDATE CASCADE
ON DELETE CASCADE;