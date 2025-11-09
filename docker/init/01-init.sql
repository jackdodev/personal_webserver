CREATE TABLE IF NOT EXISTS blogs (                                                                  
  blog_id varchar(255) PRIMARY KEY,
  subject varchar(255) NOT NULL,
  content_path varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  last_modified TIMESTAMP NOT NULL
);

INSERT INTO blogs (blog_id, subject, content_path, created_at, last_modified) VALUES (
'blog1_id', 'blog1_name', 'blog1_content_path', NOW(), NOW());
INSERT INTO blogs (blog_id, subject, content_path, created_at, last_modified) VALUES (
'blog2_id', 'blog2_name', 'blog2_content_path', NOW(), NOW());
INSERT INTO blogs (blog_id, subject, content_path, created_at, last_modified) VALUES (
'blog3_id', 'blog3_name', 'blog3_content_path', NOW(), NOW());

CREATE TABLE IF NOT EXISTS projects (
project_id varchar(255) PRIMARY KEY,
name varchar(255) NOT NULL,
content_path varchar(255) NOT NULL,
created_at TIMESTAMP NOT NULL,
last_modified TIMESTAMP NOT NULL
);

INSERT INTO projects (project_id, name, content_path, created_at, last_modified) VALUES (
'proj1_id', 'proj1_name', 'proj1_content_path', NOW(), NOW());
INSERT INTO projects (project_id, name, content_path, created_at, last_modified) VALUES (
'proj2_id', 'proj2_name', 'proj2_content_path', NOW(), NOW());
INSERT INTO projects (project_id, name, content_path, created_at, last_modified) VALUES (
'proj3_id', 'proj3_name', 'proj3_content_path', NOW(), NOW());