-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS answer (
   id            serial         not null,
   question_id   varchar(20)    not null,
   user_id       INT            not null,
   answer_text   varchar(20)    not null,
   likes         INT            default 0,
   created_at timestamp default current_timestamp,
   updated_at timestamp default current_timestamp,
   PRIMARY KEY(id)
   
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS answer;
-- +goose StatementEnd