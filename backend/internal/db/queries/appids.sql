-- name: ListAppIDs :many
SELECT
    appid,
    name,
    logo_url
FROM appids
ORDER BY name;

-- name: GetAppIDByID :one
SELECT
  appid,
  name,
  logo_url
FROM appids
WHERE appid = $1;

-- name: CreateAppID :one
INSERT INTO appids (appid, name, logo_url)
VALUES ($1, $2, $3)
RETURNING
  appid,
  name,
  logo_url;

-- name: DeleteAppID :exec
DELETE FROM appids
WHERE appid = $1;