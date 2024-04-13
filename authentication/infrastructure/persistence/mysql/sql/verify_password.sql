SELECT id as userId
FROM core_users
WHERE username = ?
  AND password = ?;
