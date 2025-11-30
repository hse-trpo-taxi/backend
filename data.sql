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
-- Drivers (expanded schema: name, phone, license_number, rating, status, time_work, x, y, score_driver, passport, snils, inn)
INSERT INTO drivers (name, phone, license_number, rating, status, time_work, x, y, score_driver, passport, snils, inn)
SELECT * FROM (VALUES
  ('Иван Иванов', '+79000000001', 'ABC12345', 4.7, 'active', 8, 37.6173, 55.7558, 82.5, 123456, 11122333444, 7701234567),
  ('Олег Кузнецов', '+79000000002', 'DEF67890', 4.2, 'onOrder', 5, 37.6156, 55.7520, 78.3, 234567, 22233444555, 7702345678),
  ('Елена Крылова', '+79000000003', 'GHI23456', 4.9, 'active', 9, 37.6090, 55.7526, 90.1, 345678, 33344555666, 7703456789),
  ('Сергей Петров', '+79000000004', 'JKL34567', 3.9, 'offline', 0, 37.6200, 55.7570, 65.0, 456789, 44455666777, 7704567890),
  ('Анна Сидорова', '+79000000005', 'MNO45678', 4.5, 'active', 7, 37.6225, 55.7600, 85.2, 567890, 55566777888, 7705678901),
  ('Павел Иванов', '+79000000006', 'PQR56789', 4.0, 'atRequest', 3, 37.6050, 55.7550, 74.6, 678901, 66677888999, 7706789012),
  ('Марина Козлова', '+79000000007', 'STU67890', 4.8, 'active', 10, 37.6000, 55.7500, 88.8, 789012, 77788999000, 7707890123),
  ('Дмитрий Орлов', '+79000000008', 'VWX78901', 3.7, 'offline', 0, 37.5900, 55.7450, 69.4, 890123, 88899000111, 7708901234),
  ('Ольга Новикова', '+79000000009', 'YZA89012', 4.6, 'active', 6, 37.6005, 55.7475, 86.0, 901234, 99900111222, 7709012345),
  ('Игорь Лебедев', '+79000000010', 'BCD90123', 4.1, 'onOrder', 4, 37.6300, 55.7605, 77.0, 112345, 10111213141, 7710123456),
  ('Ксения Белова', '+79000000011', 'CDE01234', 4.4, 'active', 8, 37.6180, 55.7480, 80.5, 223456, 12131415161, 7721234567),
  ('Владимир Медведев', '+79000000012', 'DEF12345', 3.8, 'atRequest', 2, 37.6120, 55.7430, 70.2, 334567, 14151617181, 7732345678),
  ('Татьяна Федорова', '+79000000013', 'EFG23456', 4.3, 'active', 7, 37.6250, 55.7590, 79.9, 445678, 16171819202, 7743456789),
  ('Александр Соколов', '+79000000014', 'FGH34567', 4.0, 'active', 6, 37.6070, 55.7490, 75.1, 556789, 18192021222, 7754567890),
  ('Наталья Морозова', '+79000000015', 'GHI45678', 4.9, 'active', 12, 37.6190, 55.7555, 92.7, 667890, 20212223242, 7765678901),
  ('Роман Волков', '+79000000016', 'HIJ56789', 3.6, 'offline', 0, 37.6110, 55.7530, 66.8, 778901, 22232425262, 7776789012),
  ('Екатерина Сергеева', '+79000000017', 'IJK67890', 4.2, 'active', 5, 37.6140, 55.7510, 78.0, 889012, 24252627282, 7787890123),
  ('Максим Крылов', '+79000000018', 'JKL78901', 4.5, 'onOrder', 9, 37.6210, 55.7560, 84.4, 990123, 26272829292, 7798901234),
  ('Людмила Павлова', '+79000000019', 'KLM89012', 4.1, 'active', 6, 37.6165, 55.7540, 76.6, 101234, 28293031312, 7809012345),
  ('Григорий Захаров', '+79000000020', 'LMN90123', 3.9, 'atRequest', 3, 37.6185, 55.7535, 72.3, 121345, 30313233342, 7810123456)
) AS v(name, phone, license_number, rating, status, time_work, x, y, score_driver, passport, snils, inn)
WHERE NOT EXISTS (SELECT 1 FROM drivers);

-- Cars: link to drivers by name (works for this seed dataset)
INSERT INTO cars (driver_id, brand, model, year, license_plate, color, mileage, score_users, score_system)
SELECT * FROM (VALUES
  ((SELECT id FROM drivers WHERE name='Иван Иванов' LIMIT 1), 'Toyota', 'Camry', 2018, 'A111AA', 'white', 32000, 4.5, 4.2),
  ((SELECT id FROM drivers WHERE name='Олег Кузнецов' LIMIT 1), 'Lada', 'Vesta', 2017, 'B222BB', 'black', 85000, 4.0, 3.9),
  ((SELECT id FROM drivers WHERE name='Елена Крылова' LIMIT 1), 'Kia', 'Rio', 2019, 'C333CC', 'silver', 21000, 4.7, 4.6),
  ((SELECT id FROM drivers WHERE name='Сергей Петров' LIMIT 1), 'Renault', 'Logan', 2016, 'D444DD', 'gray', 98000, 3.8, 3.7),
  ((SELECT id FROM drivers WHERE name='Анна Сидорова' LIMIT 1), 'Hyundai', 'Solaris', 2018, 'E555EE', 'blue', 45000, 4.6, 4.5),
  ((SELECT id FROM drivers WHERE name='Павел Иванов' LIMIT 1), 'Ford', 'Focus', 2015, 'F666FF', 'red', 123000, 3.9, 3.8),
  ((SELECT id FROM drivers WHERE name='Марина Козлова' LIMIT 1), 'Volkswagen', 'Polo', 2020, 'G777GG', 'white', 15000, 4.8, 4.7),
  ((SELECT id FROM drivers WHERE name='Дмитрий Орлов' LIMIT 1), 'Skoda', 'Rapid', 2016, 'H888HH', 'black', 76000, 3.6, 3.5),
  ((SELECT id FROM drivers WHERE name='Ольга Новикова' LIMIT 1), 'Nissan', 'Almera', 2019, 'I999II', 'silver', 30000, 4.4, 4.3),
  ((SELECT id FROM drivers WHERE name='Игорь Лебедев' LIMIT 1), 'Lifan', 'Solano', 2014, 'J000JJ', 'green', 140000, 3.7, 3.6),
  ((SELECT id FROM drivers WHERE name='Ксения Белова' LIMIT 1), 'BMW', '3 Series', 2017, 'K121KK', 'black', 67000, 4.5, 4.2),
  ((SELECT id FROM drivers WHERE name='Владимир Медведев' LIMIT 1), 'Audi', 'A4', 2015, 'L131LL', 'gray', 110000, 3.6, 3.5),
  ((SELECT id FROM drivers WHERE name='Татьяна Федорова' LIMIT 1), 'Mitsubishi', 'Lancer', 2016, 'M141MM', 'white', 82000, 4.3, 4.0),
  ((SELECT id FROM drivers WHERE name='Александр Соколов' LIMIT 1), 'Opel', 'Astra', 2014, 'N151NN', 'silver', 130000, 3.9, 3.8),
  ((SELECT id FROM drivers WHERE name='Наталья Морозова' LIMIT 1), 'Kia', 'Ceed', 2020, 'O161OO', 'blue', 12000, 4.9, 4.8),
  ((SELECT id FROM drivers WHERE name='Роман Волков' LIMIT 1), 'Renault', 'Duster', 2018, 'P171PP', 'black', 90000, 3.5, 3.4),
  ((SELECT id FROM drivers WHERE name='Екатерина Сергеева' LIMIT 1), 'Hyundai', 'Solaris', 2019, 'Q181QQ', 'silver', 40000, 4.1, 4.0),
  ((SELECT id FROM drivers WHERE name='Максим Крылов' LIMIT 1), 'Toyota', 'Corolla', 2017, 'R191RR', 'white', 72000, 4.4, 4.1),
  ((SELECT id FROM drivers WHERE name='Людмила Павлова' LIMIT 1), 'Chevrolet', 'Niva', 2013, 'S202SS', 'brown', 150000, 3.8, 3.7),
  ((SELECT id FROM drivers WHERE name='Григорий Захаров' LIMIT 1), 'UAZ', 'Patriot', 2012, 'T212TT', 'green', 160000, 3.7, 3.6)
) AS v(driver_id, brand, model, year, license_plate, color, mileage, score_users, score_system)
WHERE NOT EXISTS (SELECT 1 FROM cars);

-- Orders: some current and recent orders
INSERT INTO orders (driver_id, first_address, second_address, current, time, passengers, time_close, score, x, y)
SELECT * FROM (VALUES
  ((SELECT id FROM drivers WHERE name='Иван Иванов' LIMIT 1), 'ул. Ленина, 1', 'пр. Мира, 10', TRUE, 15, 2, 0, 5, 37.6173, 55.7558),
  ((SELECT id FROM drivers WHERE name='Олег Кузнецов' LIMIT 1), 'ул. Тверская, 5', 'ул. Садовая, 7', FALSE, 30, 1, 0, 4, 37.6156, 55.7520),
  ((SELECT id FROM drivers WHERE name='Елена Крылова' LIMIT 1), 'ул. Арбат, 12', 'ул. Никольская, 3', FALSE, 20, 3, 0, 5, 37.6090, 55.7526)
) AS v(driver_id, first_address, second_address, current, time, passengers, time_close, score, x, y)
WHERE NOT EXISTS (SELECT 1 FROM orders);

-- Ensure each driver has between 5 and 15 orders. For drivers with fewer orders,
-- insert the missing number so total per driver is in [5,15].
-- This is safe to run multiple times: it only inserts when a driver has less than
-- the randomly chosen target number for that driver.
INSERT INTO orders (driver_id, first_address, second_address, current, time, passengers, time_close, score, x, y, created_at, updated_at)
SELECT d.id,
       'from: ' || d.name || ' #' || gs.i,
       'to: ' || d.name || ' #' || gs.i,
       FALSE,
       (10 + (random()*50))::int,
       (1 + floor(random()*4))::int,
       0,
       (1 + floor(random()*5))::int,
       d.x + (random() - 0.5) * 0.01,
       d.y + (random() - 0.5) * 0.01,
       CURRENT_TIMESTAMP - ((floor(random()*30))::int || ' days')::interval,
       CURRENT_TIMESTAMP - ((floor(random()*30))::int || ' days')::interval
FROM drivers d
LEFT JOIN LATERAL (SELECT count(*) AS cnt FROM orders o WHERE o.driver_id = d.id) existing ON TRUE
LEFT JOIN LATERAL (SELECT (floor(random()*11)+5)::int AS target) tgt ON TRUE
LEFT JOIN LATERAL (SELECT GREATEST(0, tgt.target - COALESCE(existing.cnt,0)) AS to_add) need ON TRUE
JOIN LATERAL generate_series(1, need.to_add) AS gs(i) ON need.to_add > 0;

-- For drivers that currently have status 'onOrder', mark exactly one of their orders as active.
-- We first clear current flags for those drivers, then set the most recent order to current = TRUE.
UPDATE orders o SET current = FALSE
FROM drivers d
WHERE o.driver_id = d.id AND d.status = 'onOrder';

WITH latest AS (
  SELECT o.id,
         o.driver_id,
         ROW_NUMBER() OVER (PARTITION BY o.driver_id ORDER BY o.created_at DESC, o.id DESC) rn
  FROM orders o
  JOIN drivers d ON o.driver_id = d.id
  WHERE d.status = 'onOrder'
)
UPDATE orders o SET current = TRUE
FROM latest l
WHERE o.id = l.id AND l.rn = 1;

-- Driver schedules: one entry per seeded driver (plain VALUES rows)
INSERT INTO driver_schedules (driver_id, date)
SELECT * FROM (VALUES
  ((SELECT id FROM drivers WHERE name='Иван Иванов' LIMIT 1), CURRENT_DATE + INTERVAL '0 day'),
  ((SELECT id FROM drivers WHERE name='Олег Кузнецов' LIMIT 1), CURRENT_DATE + INTERVAL '1 day'),
  ((SELECT id FROM drivers WHERE name='Елена Крылова' LIMIT 1), CURRENT_DATE + INTERVAL '2 day'),
  ((SELECT id FROM drivers WHERE name='Сергей Петров' LIMIT 1), CURRENT_DATE + INTERVAL '3 day'),
  ((SELECT id FROM drivers WHERE name='Анна Сидорова' LIMIT 1), CURRENT_DATE + INTERVAL '4 day'),
  ((SELECT id FROM drivers WHERE name='Павел Иванов' LIMIT 1), CURRENT_DATE + INTERVAL '5 day'),
  ((SELECT id FROM drivers WHERE name='Марина Козлова' LIMIT 1), CURRENT_DATE + INTERVAL '6 day'),
  ((SELECT id FROM drivers WHERE name='Дмитрий Орлов' LIMIT 1), CURRENT_DATE + INTERVAL '7 day'),
  ((SELECT id FROM drivers WHERE name='Ольга Новикова' LIMIT 1), CURRENT_DATE + INTERVAL '8 day'),
  ((SELECT id FROM drivers WHERE name='Игорь Лебедев' LIMIT 1), CURRENT_DATE + INTERVAL '9 day'),
  ((SELECT id FROM drivers WHERE name='Ксения Белова' LIMIT 1), CURRENT_DATE + INTERVAL '10 day'),
  ((SELECT id FROM drivers WHERE name='Владимир Медведев' LIMIT 1), CURRENT_DATE + INTERVAL '11 day'),
  ((SELECT id FROM drivers WHERE name='Татьяна Федорова' LIMIT 1), CURRENT_DATE + INTERVAL '12 day'),
  ((SELECT id FROM drivers WHERE name='Александр Соколов' LIMIT 1), CURRENT_DATE + INTERVAL '13 day'),
  ((SELECT id FROM drivers WHERE name='Наталья Морозова' LIMIT 1), CURRENT_DATE + INTERVAL '14 day'),
  ((SELECT id FROM drivers WHERE name='Роман Волков' LIMIT 1), CURRENT_DATE + INTERVAL '15 day'),
  ((SELECT id FROM drivers WHERE name='Екатерина Сергеева' LIMIT 1), CURRENT_DATE + INTERVAL '16 day'),
  ((SELECT id FROM drivers WHERE name='Максим Крылов' LIMIT 1), CURRENT_DATE + INTERVAL '17 day'),
  ((SELECT id FROM drivers WHERE name='Людмила Павлова' LIMIT 1), CURRENT_DATE + INTERVAL '18 day'),
  ((SELECT id FROM drivers WHERE name='Григорий Захаров' LIMIT 1), CURRENT_DATE + INTERVAL '19 day')
) AS v(driver_id, date)
WHERE NOT EXISTS (SELECT 1 FROM driver_schedules);

-- Support requests seed
INSERT INTO support_requests (name, comment, time_reacting_score, quality_answer_score, date)
SELECT * FROM (VALUES
  ('Анастасия Смирнова', 'Проблема с оплатой, списали дважды', 4, 5, CURRENT_TIMESTAMP),
  ('Игорь Волков', 'Не успел водитель, опоздал на 20 минут', 3, 4, CURRENT_TIMESTAMP - INTERVAL '1 day')
) AS v(name, comment, time_reacting_score, quality_answer_score, date)
WHERE NOT EXISTS (SELECT 1 FROM support_requests);
