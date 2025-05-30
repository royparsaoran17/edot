-- +goose Up
INSERT INTO products (
    id, name, description, price, unit, sku, category, is_active, created_at, updated_at
) VALUES
      (
          'a1b2c3d4-e5f6-7890-abcd-1234567890aa',
          'Spicy Chips',
          'Crunchy spicy cassava chips with extra chili.',
          15000.00,
          'pack',
          'SPC-001',
          'snacks',
          true,
          now(), now()
      ),
      (
          'a1b2c3d4-e5f6-7890-abcd-1234567890bb',
          'Sweet Cookies',
          'Delicious handmade cookies with chocolate chips.',
          25000.00,
          'box',
          'SWC-002',
          'snacks',
          true,
          now(), now()
      ),
      (
          'a1b2c3d4-e5f6-7890-abcd-1234567890cc',
          'Mineral Water',
          'Fresh natural mineral water 600ml.',
          5000.00,
          'bottle',
          'MNW-003',
          'beverages',
          true,
          now(), now()
      );


-- +goose Down
DELETE
FROM products
WHERE id IN ('a1b2c3d4-e5f6-7890-abcd-1234567890aa','a1b2c3d4-e5f6-7890-abcd-1234567890bb','a1b2c3d4-e5f6-7890-abcd-1234567890cc');
