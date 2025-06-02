-- +goose Up
INSERT INTO users (id, name, email, phone, password, role_id, created_at, updated_at, deleted_at)
VALUES ('a05b5ac4-53c4-11ee-8c99-0242ac120002', 'Roy Parsaoran', 'royparsaoran17@gmail.com', '+6281809134100', 'dec2c975f7d209f1fec604e71ab025ac561073e0', '84756252-dadd-46c4-836f-608a1e8d33ce', '2023-10-09 19:18:05.000000',
        '2023-10-09 19:18:05.000000', null); -- password : Edot123!

INSERT INTO users (id, name, email, phone, password, role_id, created_at, updated_at, deleted_at)
VALUES ('a05b5ac4-53c4-11ee-8c99-0242ac120001', 'Roy Parsaoran 2', 'royparsaoran18@gmail.com', '+6281809134101', 'dec2c975f7d209f1fec604e71ab025ac561073e0', '84756252-dadd-46c4-836f-608a1e8d33ce', '2023-10-09 19:18:05.000000',
        '2023-10-09 19:18:05.000000', null); -- password : Edot123!

-- +goose Down
DELETE
FROM users
WHERE id = 'a05b5ac4-53c4-11ee-8c99-0242ac120002';

