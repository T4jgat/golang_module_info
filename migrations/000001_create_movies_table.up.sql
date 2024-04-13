CREATE TABLE module_info
(
    id              BIGSERIAL PRIMARY KEY,
    created_at      TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    module_name     VARCHAR(255)                NOT NULL,
    module_duration INTEGER                     NOT NULL,
    exam_type       VARCHAR(255)                NOT NULL,
    version         INTEGER                     NOT NULL DEFAULT 1
);

CREATE  FUNCTION update_updated_on_user_task()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_task_updated_on
    BEFORE UPDATE
    ON
        module_info
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_on_user_task();