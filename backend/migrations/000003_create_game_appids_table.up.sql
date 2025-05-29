-- 000003_create_appids_table.up.sql

CREATE TABLE IF NOT EXISTS appids (
  appid     INT      PRIMARY KEY,
  name      TEXT     NOT NULL UNIQUE,
  logo_url  TEXT     NOT NULL
);

-- Seed common apps
INSERT INTO appids(appid, name, logo_url) VALUES
  (440,  'Team Fortress 2',    'https://cdn.fastly.steamstatic.com/steamcommunity/public/images/apps/440/e3f595a92552da3d664ad00277fad2107345f743.jpg'),
  (730,  'Counter-Strike 2',   'https://cdn.fastly.steamstatic.com/steamcommunity/public/images/apps/730/8dbc71957312bbd3baea65848b545be9eae2a355.jpg'),
  (252490, 'Rust',             'https://cdn.fastly.steamstatic.com/steamcommunity/public/images/apps/252490/820be4782639f9c4b64fa3ca7e6c26a95ae4fd1c.jpg')
ON CONFLICT (appid) DO NOTHING;