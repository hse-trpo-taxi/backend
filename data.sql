-- Sample seed data for development. Safe to re-run: each INSERT
-- uses a "WHERE NOT EXISTS (SELECT 1 FROM <table>)" guard so it
-- won't duplicate rows when executed multiple times.

-- Clients
INSERT INTO clients (name, phone, email)
SELECT * FROM (VALUES
  ('Иван Иванов', '+70000000001', 'ivan@example.com'),
  ('Мария Петрова', '+70000000002', 'maria@example.com'),
  ('Алексей Смирнов', '+70000000003', 'alex@example.com')
) AS v(name, phone, email)
WHERE NOT EXISTS (SELECT 1 FROM clients);

-- Drivers
INSERT INTO drivers (name, phone, license_number, rating)
SELECT * FROM (VALUES
  ('Иван Иванов', '+79000000001', 'ABC12345', 4.7),
  ('Олег Кузнецов', '+79000000002', 'DEF67890', 4.2),
  ('Елена Крылова', '+79000000003', 'GHI23456', 4.9)
) AS v(name, phone, license_number, rating)
WHERE NOT EXISTS (SELECT 1 FROM drivers);

-- Cars: link to drivers by name (works for this seed dataset)
INSERT INTO cars (driver_id, brand, model, year, license_plate, color)
SELECT * FROM (VALUES
  ((SELECT id FROM drivers WHERE name='Иван Иванов' LIMIT 1), 'Toyota', 'Camry', 2018, 'A111AA', 'white'),
  ((SELECT id FROM drivers WHERE name='Олег Кузнецов' LIMIT 1), 'Lada', 'Vesta', 2017, 'B222BB', 'black'),
  ((SELECT id FROM drivers WHERE name='Елена Крылова' LIMIT 1), 'Kia', 'Rio', 2019, 'C333CC', 'silver')
) AS v(driver_id, brand, model, year, license_plate, color)
WHERE NOT EXISTS (SELECT 1 FROM cars);

-- Orders: some current and recent orders
INSERT INTO orders (driver_id, first_address, second_address, current, time, passengers, time_close, score, x, y)
SELECT * FROM (VALUES
  ((SELECT id FROM drivers WHERE name='Иван Иванов' LIMIT 1), 'ул. Ленина, 1', 'пр. Мира, 10', TRUE, 15, 2, 0, 5, 37.6173, 55.7558),
  ((SELECT id FROM drivers WHERE name='Олег Кузнецов' LIMIT 1), 'ул. Тверская, 5', 'ул. Садовая, 7', FALSE, 30, 1, 0, 4, 37.6156, 55.7520),
  ((SELECT id FROM drivers WHERE name='Елена Крылова' LIMIT 1), 'ул. Арбат, 12', 'ул. Никольская, 3', FALSE, 20, 3, 0, 5, 37.6090, 55.7526)
) AS v(driver_id, first_address, second_address, current, time, passengers, time_close, score, x, y)
WHERE NOT EXISTS (SELECT 1 FROM orders);

-- Driver schedules: next few days for the seeded drivers
INSERT INTO driver_schedules (driver_id, date)
SELECT * FROM (VALUES
  ((SELECT id FROM drivers WHERE name='Иван Иванов' LIMIT 1), CURRENT_DATE),
  ((SELECT id FROM drivers WHERE name='Иван Иванов' LIMIT 1), CURRENT_DATE + INTERVAL '1 day'),
  ((SELECT id FROM drivers WHERE name='Елена Крылова' LIMIT 1), CURRENT_DATE + INTERVAL '2 day')
) AS v(driver_id, date)
WHERE NOT EXISTS (SELECT 1 FROM driver_schedules);

-- Support requests seed
INSERT INTO support_requests (name, comment, time_reacting_score, quality_answer_score, date)
SELECT * FROM (VALUES
  ('Анастасия Смирнова', 'Проблема с оплатой, списали дважды', 4, 5, CURRENT_TIMESTAMP),
  ('Игорь Волков', 'Не успел водитель, опоздал на 20 минут', 3, 4, CURRENT_TIMESTAMP - INTERVAL '1 day')
) AS v(name, comment, time_reacting_score, quality_answer_score, date)
WHERE NOT EXISTS (SELECT 1 FROM support_requests);
