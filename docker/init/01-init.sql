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
author_id varchar(255) NOT NULL,
project_name varchar(255) NOT NULL,
created_at TIMESTAMP NOT NULL,
last_modified TIMESTAMP NOT NULL,
tags text[] DEFAULT '{}' NOT NULL
);

INSERT INTO projects (project_id, author_id, project_name, created_at, last_modified, tags) VALUES (
'proj1_id', 'author1', 'proj1_name', NOW(), NOW(), ARRAY['web','go']);
INSERT INTO projects (project_id, author_id, project_name, created_at, last_modified, tags) VALUES (
'proj2_id', 'author2', 'proj2_name', NOW(), NOW(), ARRAY['db','sql']);
INSERT INTO projects (project_id, author_id, project_name, created_at, last_modified, tags) VALUES (
'proj3_id', 'author3', 'proj3_name', NOW(), NOW(), ARRAY['api','rest']);

CREATE TABLE IF NOT EXISTS posts (
  post_id varchar(255) PRIMARY KEY,
  author_id varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  last_modified TIMESTAMP NOT NULL,
  tags text[] DEFAULT '{}' NOT NULL
)

INSERT INTO projects (project_id, author_id, project_name, created_at, last_modified, tags) VALUES (
'prj:id1', 'author1', 'proj1_name', NOW(), NOW(), ARRAY['web','go']);
INSERT INTO projects (project_id, author_id, project_name, created_at, last_modified, tags) VALUES (
'proj:id2', 'author2', 'proj2_name', NOW(), NOW(), ARRAY['db','sql']);
INSERT INTO projects (project_id, author_id, project_name, created_at, last_modified, tags) VALUES (
'proj:id3', 'author3', 'proj3_name', NOW(), NOW(), ARRAY['api','rest']);
INSERT INTO blogs (blog_id, author_id, subject, created_at, last_modified, tags) VALUES (
'blg:id1', 'author1', 'blog1_name', NOW(), NOW(), ARRAY['intro','go']);
INSERT INTO blogs (blog_id, author_id, subject, created_at, last_modified, tags) VALUES (
'blg"id2', 'author2', 'blog2_name', NOW(), NOW(), ARRAY['tutorial','db']);
INSERT INTO blogs (blog_id, author_id, subject, created_at, last_modified, tags) VALUES (
'blg:id3', 'author3', 'blog3_name', NOW(), NOW(), ARRAY['advanced','web']);