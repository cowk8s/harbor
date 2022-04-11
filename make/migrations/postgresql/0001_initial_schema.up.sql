create table access (
 access_id SERIAL PRIMARY KEY NOT NULL,
 access_code char(1),
 comment varchar (30)
);

insert into access (access_code, comment) values 
('M', 'Management access for project'),
('R', 'Read access for project'),
('W', 'Write access for project'),
('D', 'Delete access for project'),
('S', 'Search access for project');

create table role (
 role_id SERIAL PRIMARY KEY NOT NULL,
 role_mask int DEFAULT 0 NOT NULL,
 role_code varchar(20),
 name varchar (20)
);

/*
role mask is used for future enhancement when a project member can have multi-roles
currently set to 0
*/

insert into role (role_code, name) values 
('MDRWS', 'projectAdmin'),
('RWS', 'developer'),
('RS', 'guest');

create table harbor_user (
 user_id SERIAL PRIMARY KEY NOT NULL,
 username varchar(255),
 email varchar(255),
 password varchar(40) NOT NULL,
 realname varchar (255) NOT NULL,
 comment varchar (30),
 deleted boolean DEFAULT false NOT NULL,
 reset_uuid varchar(40) DEFAULT NULL,
 salt varchar(40) DEFAULT NULL,
 sysadmin_flag boolean DEFAULT false NOT NULL,
 creation_time timestamp default CURRENT_TIMESTAMP,
 update_time timestamp default CURRENT_TIMESTAMP,
 UNIQUE (username),
 UNIQUE (email)
);

CREATE FUNCTION update_update_time_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
  BEGIN
    NEW.update_time = NOW();
    RETURN NEW;
  END;
$$;

CREATE TRIGGER harbor_user_update_time_at_modtime BEFORE UPDATE ON harbor_user FOR EACH ROW EXECUTE PROCEDURE update_update_time_at_column();

insert into harbor_user (username, password, realname, comment, deleted, sysadmin_flag, creation_time, update_time) values
('admin', '', 'system admin', 'admin user',false, true, NOW(), NOW()),
('anonymous', '', 'anonymous user', 'anonymous user', true, false, NOW(), NOW());

create table project (
 project_id SERIAL PRIMARY KEY NOT NULL,
 owner_id int NOT NULL,
 /*
 The max length of name controlled by API is 30, 
 and 11 is reserved for marking the deleted project.
 */
 name varchar (255) NOT NULL,
 creation_time timestamp default CURRENT_TIMESTAMP,
 update_time timestamp  default CURRENT_TIMESTAMP,
 deleted boolean DEFAULT false NOT NULL,
 FOREIGN KEY (owner_id) REFERENCES harbor_user(user_id),
 UNIQUE (name)
);

CREATE TRIGGER project_update_time_at_modtime BEFORE UPDATE ON project FOR EACH ROW EXECUTE PROCEDURE update_update_time_at_column();

insert into project (owner_id, name, creation_time, update_time) values 
(1, 'library', NOW(), NOW());

create table project_member (
  
)

insert into role (role_code, name) values 
('MDRWS', 'projectAdmin'),
('RWS', 'developer'),
('RS', 'guest');

create table admin_job (
 id SERIAL NOT NULL,
 job_name varchar(64) NOT NULL,
 job_kind varchar(64) NOT NULL,
 cron_str varchar(256),
 status varchar(64) NOT NULL,
 job_uuid varchar(64),
 creation_time timestamp default CURRENT_TIMESTAMP,
 update_time timestamp default CURRENT_TIMESTAMP,
 deleted boolean DEFAULT false NOT NULL,
 PRIMARY KEY(id)
);

CREATE TRIGGER admin_job_update_time_at_modtime BEFORE UPDATE ON admin_job FOR EACH ROW EXECUTE PROCEDURE update_update_time_at_column();

CREATE INDEX admin_job_status ON admin_job (status);
CREATE INDEX admin_job_uuid ON admin_job (job_uuid);

CREATE TABLE IF NOT EXISTS alembic_version (
    version_num varchar(32) NOT NULL
);

insert into alembic_version values ('1.6.0');