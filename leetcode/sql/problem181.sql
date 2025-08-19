-- Find all employees who earn more than their managers
-- Runs on postgresql
select e.name as Employee
from Employee as e, Employee as m
where e.salary > m.salary and e.managerId = m.id
