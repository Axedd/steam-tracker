
-- name: AddSteamParam :exec
INSERT INTO steam_param_defs (key, label, type, options, default_value, help_text, appid)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetGlobalParams :many
SELECT 
  key, label, type, options, default_value, help_text, appid
FROM steam_param_defs
WHERE appid IS NULL;

-- name: GetParamsByAppID :many
SELECT 
  key, label, type, options, default_value, help_text, appid
FROM steam_param_defs
WHERE appid = $1;

-- name: UpdateParamByKey :one
UPDATE steam_param_defs
SET key = $2,
    label = $3,
    options = $4,
    default_value = $5,
    help_text = $6,
    appid = $7
WHERE key = $1
RETURNING key, label, type, options, default_value, help_text, appid;

-- name: RemoveSteamParam :exec
DELETE FROM steam_param_defs
WHERE key = $1;