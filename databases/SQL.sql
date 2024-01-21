BEGIN;

SET TIME ZONE 'Asia/Bangkok';
-- ตารางอาหาร
CREATE TABLE food (
  id serial PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type INT NOT NULL,
  ingredients integer NOT NULL,
  price INT NOT NULL,
  size VARCHAR(255) NOT NULL,
  box_size VARCHAR(255) NOT NULL,
  image VARCHAR(255),
  description TEXT,
  discount INT,
  rating INT
);

-- ตารางประเภทอาหาร
CREATE TABLE food_type (
  id serial PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- ตารางวัตถุดิบ
CREATE TABLE ingredient (
  id serial PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  unit VARCHAR(255) NOT NULL,
  price INT NOT NULL
);

CREATE TABLE "order" (
  id serial PRIMARY KEY,
  order_number VARCHAR(255) NOT NULL,
  customer_name VARCHAR(255) NOT NULL,
  customer_phone VARCHAR(255) NOT NULL,
  customer_address VARCHAR(255) NOT NULL,
  order_date DATE NOT NULL,
  delivery_date DATE NOT NULL,
  order_total INT NOT NULL,
  discount INT NOT NULL,
  description TEXT,
  menu_list JSON NOT NULL
);

-- ความสัมพันธ์ระหว่างตาราง
ALTER TABLE food ADD FOREIGN KEY (type) REFERENCES food_type (id);

ALTER TABLE food ADD CONSTRAINT fk_ingredients FOREIGN KEY (ingredients) REFERENCES ingredient (id);

COMMIT;