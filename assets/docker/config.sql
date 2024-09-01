use store;
DROP TABLE IF EXISTS items;
CREATE TABLE items (
  id         INT AUTO_INCREMENT NOT NULL,
  name       VARCHAR(128) NOT NULL,
  qtd        INT NOT NULL,
  price      DECIMAL(20,2) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO items
  (name, qtd, price)
VALUES
  ('Banana', 5, 0.99),
  ('Apple', 6, 1.59),
  ('Orange', 2, 0.49),
  ('Grape', 7, 4.98);