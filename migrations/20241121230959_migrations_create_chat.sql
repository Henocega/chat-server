-- +goose Up
-- +goose StatementBegin
create table "chat"
(
    id        serial primary key,
    usernames text[] not null
);

create table "message"
(
    id         serial primary key,
    "from"     text      not null,
    text       text      not null,
    created_at timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "chat";
drop table "message";
-- +goose StatementEnd
