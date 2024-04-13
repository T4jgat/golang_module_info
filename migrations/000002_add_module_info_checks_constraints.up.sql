ALTER TABLE module_info
    ADD CONSTRAINT module_name_not_empty CHECK (LENGTH(module_name) > 0),
    ADD CONSTRAINT module_duration_positive CHECK (module_duration > 0),
    ADD CONSTRAINT version_positive CHECK (version > 0);