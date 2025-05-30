-- +goose Up
INSERT INTO  product_stocks (
    id, product_id, warehouse_id, quantity, created_at, updated_at
) VALUES
      (
          '11111111-aaaa-bbbb-cccc-000000000001',
          'a1b2c3d4-e5f6-7890-abcd-1234567890aa',
          'd4c3b2a1-f6e5-0987-dcba-9876543210aa',
          120,
          now(), now()
      ),
      (
          '11111111-aaaa-bbbb-cccc-000000000002',
          'a1b2c3d4-e5f6-7890-abcd-1234567890bb',
          'd4c3b2a1-f6e5-0987-dcba-9876543210aa',
          80,
          now(), now()
      ),
      (
          '11111111-aaaa-bbbb-cccc-000000000003',
          'a1b2c3d4-e5f6-7890-abcd-1234567890cc',
          'd4c3b2a1-f6e5-0987-dcba-9876543210bb',
          300,
          now(), now()
      );


-- +goose Down
DELETE
FROM warehouses
WHERE id IN ('11111111-aaaa-bbbb-cccc-000000000001','11111111-aaaa-bbbb-cccc-000000000002','11111111-aaaa-bbbb-cccc-000000000003');
