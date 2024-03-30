SELECT status as TOTAL
FROM core_users
WHERE username = ?
  AND password = ?;
