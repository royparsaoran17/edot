-- +goose Up
INSERT INTO  warehouses (
    id, name, is_active, created_at, updated_at
) VALUES
      (
          'd4c3b2a1-f6e5-0987-dcba-9876543210aa',
          'Central Warehouse',
          true,
          now(), now()
      ),
      (
          'd4c3b2a1-f6e5-0987-dcba-9876543210bb',
          'East Warehouse',
          true,
          now(), now()
      );


-- +goose Down
DELETE
FROM warehouses
WHERE id IN ('d4c3b2a1-f6e5-0987-dcba-9876543210aa','d4c3b2a1-f6e5-0987-dcba-9876543210bb');
