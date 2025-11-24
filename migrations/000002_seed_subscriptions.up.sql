INSERT INTO subscriptions (
    id, service_name, price, user_id,
    start_date, end_date, created_at, updated_at
) VALUES
      ('7c4e6f8a-9b1c-4d2e-8f3a-1c5d7e9b2a4f', 'Netflix Premium', 1599, 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
       '2025-03-01 00:00:00'::timestamp, '2026-03-01 00:00:00'::timestamp, now(), now()),

      ('8d5f7a9b-1c2d-4e3f-9a0b-2d6e8f1a3c5b', 'Spotify Premium', 1099, 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
       '2025-01-01 00:00:00'::timestamp, '2025-12-01 00:00:00'::timestamp, now(), now()),

      ('9e6a8b1c-2d3e-4f5a-0b1c-3d7e9f2a4c6d', 'YouTube Premium', 1399, 'b2c3d4e5-f6a7-8901-bcde-f234567890ab',
       '2025-06-01 00:00:00'::timestamp, NULL, now(), now()),

      ('0f7b9c2d-3e4f-5a6b-1c2d-4e8f0a3c5d7e', 'Apple Music', 999, 'b2c3d4e5-f6a7-8901-bcde-f234567890ab',
       '2025-02-01 00:00:00'::timestamp, '2026-02-01 00:00:00'::timestamp, now(), now()),

      ('1a8c0d3e-4f5a-6b7c-2d3e-5f9a1c4d6e8f', 'Disney+', 899, 'c3d4e5f6-a7b8-9012-cdef-34567890abcd',
       '2025-09-01 00:00:00'::timestamp, NULL, now(), now()),

      ('2b9d1e4f-5a6b-7c8d-3e4f-6a0c2d5e7f9a', 'Amazon Prime', 1499, 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
       '2024-11-01 00:00:00'::timestamp, '2025-11-01 00:00:00'::timestamp, now(), now()),

      ('3c0e2f5a-6b7c-8d9e-4f5a-7b1d3e6f8a0c', 'HBO Max', 999, 'd4e5f6a7-b8c9-0123-def0-456789abcdef',
       '2025-07-01 00:00:00'::timestamp, '2025-12-01 00:00:00'::timestamp, now(), now()),

      ('4d1f3a6b-7c8d-9e0f-5a6b-8c2e4f7a9d1f', 'Adobe Creative Cloud', 5999, 'd4e5f6a7-b8c9-0123-def0-456789abcdef',
       '2025-04-01 00:00:00'::timestamp, '2026-04-01 00:00:00'::timestamp, now(), now()),

      ('5e2a4b7c-8d9e-0f1a-6b7c-9d3f5a8c0e2a', 'Notion Plus', 800, 'c3d4e5f6-a7b8-9012-cdef-34567890abcd',
       '2025-08-01 00:00:00'::timestamp, NULL, now(), now()),

      ('6f3b5c8d-9e0f-1a2b-7c8d-0e4a6c9e1f3b', 'ChatGPT Plus', 2000, 'b2c3d4e5-f6a7-8901-bcde-f234567890ab',
       '2025-10-01 00:00:00'::timestamp, NULL, now(), now());