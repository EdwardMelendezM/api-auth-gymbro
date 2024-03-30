SELECT enable as TOTAL
FROM core_users
WHERE username = ?
  AND password = ?;
