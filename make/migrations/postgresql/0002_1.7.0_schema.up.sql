ALTER TABLE properties ALTER COLUMN v TYPE varchar(1024);
DELETE FROM properties where k='scan_all_policy';

create table job_log (
 log_id SERIAL NOT NULL,
 job_uuid varchar (64) NOT NULL,
 creation_time timestamp default CURRENT_TIMESTAMP,
 content text,
 primary key (log_id)
);

CREATE UNIQUE INDEX job_log_uuid ON job_log (job_uuid);

/*
Rename the dulicate names before adding "UNIQUE" constraint
*/
DO $$
BEGIN
    WHILE EXISTS (SELECT count(*) FROM replication_policy GROUP BY name HAVING count(*) > 1) LOOP
        UPDATE replication_policy AS r
        SET name = (
            /*
            truncate the name of it is too long after appending the sequence number
            */
            CASE WHEN (length(name)+length(v.seq::text)+1) > 256
            THEN
                substring(name from 1 for (255-length(v.seq::text))) || '_' || v.seq
            ELSE
                name || '_' || v.seq
            END
        )
        FROM (SELECT id, row_number() OVER)
    END LOOP
END $$