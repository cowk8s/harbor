/* Fix retention_policy create_time, update_time for pgx driver */
ALTER TABLE retention_policy ALTER COLUMN create_time TYPE TIMESTAMP WITHOUT TIME ZONE USING (current_date + create_time);
ALTER TABLE retention_policy ALTER COLUMN update_time TYPE TIMESTAMP WITHOUT TIME ZONE USING (current_date + update_time);

/* create table of accessory */
CREATE TABLE IF NOT EXISTS artifact_accessory (
    id SERIAL PRIMARY KEY NOT NULL,
    /*
        the artifact id of the accessory itself.
    */
    artifact_id bigint,
    /*
    the subject artifact id of the accessory.
    */
    subject_artifact_id bigint,
    
)