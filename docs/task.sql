/*
SELECT name, description, input_schema, output_schema, cron_frequent , enable
FROM job_task.task_type;

delete from job_task.task  where output is not null  

SELECT uuid, task_type_name, "input", "output", history, start_time, end_time, error, finish, created_time
FROM job_task.task;

SELECT uuid, task_type_name , input , output , start_time, end_time, error, finish, created_time,
         tt.description, tt.input_schema, tt.output_schema, tt.cron_frequent
         FROM job_task.task t 
         INNER JOIN job_task.task_type tt ON (t.task_type_name=tt.name)
         WHERE (finish is null OR finish ='N') AND (enable='S' OR enable IS NULL) AND task_type_name  ='HELLO '  $1      '
         
SELECT uuid, task_type_name, 
input , output, history, 
start_time, 
end_time, 
error, 
finish, 
created_time
FROM job_task.task WHERE ;



*/

-- CREATE SCHEMA job_task;

SET search_path TO job_task;

DROP TABLE TASK_TYPE CASCADE; 
DROP TABLE TASK CASCADE; 

 
CREATE TABLE TASK_TYPE ( 
  NAME CHAR(40) NOT NULL, 
  DESCRIPTION VARCHAR(220) NOT NULL, 
  INPUT_SCHEMA JSONB NULL, 
  OUTPUT_SCHEMA JSONB null,
  enable  CHAR(1) null , 
  CRON_FREQUENT VARCHAR(120)
); 
ALTER TABLE TASK_TYPE ADD CONSTRAINT TASK_TYPE_PK PRIMARY KEY (NAME); 
  
CREATE TABLE TASK ( 
  UUID uuid NOT NULL, 
  TASK_TYPE_NAME CHAR(40) NOT NULL, 
  INPUT JSONB NULL, 
  OUTPUT JSONB NULL, 
  HISTORY JSONB NULL, 
  START_TIME TIMESTAMP  NULL, 
  END_TIME TIMESTAMP NULL, 
  ERROR VARCHAR(220) NULL, 
  FINISH CHAR(1) NULL, 
  CREATED_TIME TIMESTAMP NULL 
); 
ALTER TABLE TASK ADD CONSTRAINT TASK_PK PRIMARY KEY (UUID); 
 
CREATE INDEX TASK_TASK_TYPE_NAME_FK ON TASK (TASK_TYPE_NAME); 

ALTER TABLE TASK ADD CONSTRAINT TASK_TYPE_REL FOREIGN KEY (TASK_TYPE_NAME) REFERENCES TASK_TYPE(NAME);
 














































 

-- End of generated script 











