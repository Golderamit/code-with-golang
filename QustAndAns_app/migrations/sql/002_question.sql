-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS question
(
    id         serial             not null,
    user_id INT,
    QuestionTitle  varchar(20)        not null,
    created_at timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp,

    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS question;
-- +goose StatementEnd
