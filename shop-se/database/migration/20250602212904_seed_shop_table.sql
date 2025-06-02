-- +goose Up
INSERT INTO  shops (
    id, name, owner_id, created_at, updated_at
) VALUES
      (
          'd4c3b2a1-f6e5-0987-dcba-9876543210a1',
          'Shop X',
          'a05b5ac4-53c4-11ee-8c99-0242ac120001',
          now(), now()
      ),
      (
          'd4c3b2a1-f6e5-0987-dcba-9876543210b2',
          'Shop Y',
          'a05b5ac4-53c4-11ee-8c99-0242ac120002',
          now(), now()
      );


-- +goose Down
DELETE
FROM shops
WHERE id IN ('d4c3b2a1-f6e5-0987-dcba-9876543210a1','d4c3b2a1-f6e5-0987-dcba-9876543210b2');
