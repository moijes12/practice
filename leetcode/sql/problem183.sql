-- Find Customers who have never ordered anything
SELECT name as Customers
FROM Customers
WHERE id not in (SELECT DISTINCT(customerId) FROM Orders)
