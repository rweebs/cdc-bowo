alter table transaction rename column text to message;
alter table transaction alter column message type varchar(255);
