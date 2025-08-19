-- Find all duplicate emails
SELECT DISTINCT(p.email) as Email
FROM Person as p, Person as q
WHERE p.id != q.id AND p.email = q.email
