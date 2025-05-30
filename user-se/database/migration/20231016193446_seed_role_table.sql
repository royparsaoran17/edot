-- +goose Up
INSERT INTO roles (id, name, created_at, updated_at, deleted_at)
VALUES ('84756252-dadd-46c4-836f-608a1e8d33ce', 'Super Admin', '2023-10-09 19:18:05.000000',
        '2023-10-09 19:18:05.000000', null);

INSERT INTO roles (id, name, created_at, updated_at, deleted_at)
VALUES ('10156252-dadd-46c4-836f-608a1e8d33ce', 'Customer', '2023-10-09 19:18:05.000000',
        '2023-10-09 19:18:05.000000', null);

-- +goose Down
DELETE
FROM roles
WHERE id = '84756252-dadd-46c4-836f-608a1e8d33ce';

DELETE
FROM roles
WHERE id = '10156252-dadd-46c4-836f-608a1e8d33ce';

