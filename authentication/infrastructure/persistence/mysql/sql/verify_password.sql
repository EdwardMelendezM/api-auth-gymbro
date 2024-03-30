SELECT COUNT(*) AS TOTAL
FROM core_users
WHERE username = ?
  AND password = ?;
