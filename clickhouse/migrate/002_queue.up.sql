CREATE TABLE events_queue
(
    timestamp UInt64,
    item_id   String,
    action    String
) ENGINE = RabbitMQ SETTINGS
    rabbitmq_username = 'guest',
    rabbitmq_password = 'guest',
    rabbitmq_host_port = 'rabbitmq:5672',
    rabbitmq_exchange_name = 'events',
    rabbitmq_format = 'JSONEachRow',
    -- rabbitmq_num_consumers = 5,
    date_time_input_format = 'best_effort';