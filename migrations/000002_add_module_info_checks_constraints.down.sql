-- This is a migration file to drop constraints for the module_info table

-- Drop constraints
ALTER TABLE module_info
    DROP CONSTRAINT IF EXISTS module_name_not_empty,
    DROP CONSTRAINT IF EXISTS module_duration_positive,
    DROP CONSTRAINT IF EXISTS version_positive;
