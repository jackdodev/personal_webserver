CREATE TABLE IF NOT EXISTS blogs (                                                                  
  blog_id varchar(255) PRIMARY KEY,
  author_id varchar(255) NOT NULL,
  subject varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  last_modified TIMESTAMP NOT NULL,
  tags text[] DEFAULT '{}' NOT NULL
);

INSERT INTO blogs (blog_id, author_id, subject, created_at, last_modified, tags) VALUES (
'blog1_id', 'author1', 'blog1_name', NOW(), NOW(), ARRAY['intro','go']);
INSERT INTO blogs (blog_id, author_id, subject, created_at, last_modified, tags) VALUES (
'blog2_id', 'author2', 'blog2_name', NOW(), NOW(), ARRAY['tutorial','db']);
INSERT INTO blogs (blog_id, author_id, subject, created_at, last_modified, tags) VALUES (
'blog3_id', 'author3', 'blog3_name', NOW(), NOW(), ARRAY['advanced','web']);

CREATE TABLE IF NOT EXISTS projects (
project_id varchar(255) PRIMARY KEY,
name varchar(255) NOT NULL,
content_path varchar(255) NOT NULL,
created_at TIMESTAMP NOT NULL,
last_modified TIMESTAMP NOT NULL,
tags text[] DEFAULT '{}' NOT NULL
);

INSERT INTO projects (project_id, name, content_path, created_at, last_modified, tags) VALUES (
'proj1_id', 'proj1_name', 'proj1_content_path', NOW(), NOW(), ARRAY['web','go']);
INSERT INTO projects (project_id, name, content_path, created_at, last_modified, tags) VALUES (
'proj2_id', 'proj2_name', 'proj2_content_path', NOW(), NOW(), ARRAY['db','sql']);
INSERT INTO projects (project_id, name, content_path, created_at, last_modified, tags) VALUES (
'proj3_id', 'proj3_name', 'proj3_content_path', NOW(), NOW(), ARRAY['api','rest']);  