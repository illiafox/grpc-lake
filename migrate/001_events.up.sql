CREATE TABLE IF NOT EXISTS events
(
    timestamp DATETIME,
    item_id   String,
    action    String
) ENGINE = MergeTree()
      PRIMARY KEY (timestamp, item_id)
      PARTITION BY toMonth(timestamp);

