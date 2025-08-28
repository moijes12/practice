--- Find employee with second highest salary
--- Else return null
SELECT MAX(salary) AS secondhighestsalary
FROM Employee
WHERE salary < (SELECT MAX(salary) FROM Employee)
