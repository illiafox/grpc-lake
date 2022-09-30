CREATE MATERIALIZED VIEW consumer TO events
AS
SELECT toDateTime(timestamp) AS timestamp, item_id, action
FROM events_queue;