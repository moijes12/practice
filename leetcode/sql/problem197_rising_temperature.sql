SELECT w1.id
FROM Weather As w1
WHERE w1.temperature > (
    SELECT w2.temperature
    FROM Weather as w2
    WHERE w2.recordDate = w1.recordDate - INTERVAL '1 day'
)
