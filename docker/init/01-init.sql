CREATE TABLE IF NOT EXISTS posts (
  post_id varchar(255) PRIMARY KEY,
  type varchar(255) NOT NULL,
  author_id varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  last_modified TIMESTAMP NOT NULL,
  tags text[] DEFAULT '{}' NOT NULL
)

INSERT INTO posts (post_id, type, author_id, title, created_at, last_modified, tags) VALUES (
'prj:id1', 'project', 'author1', 'proj1_name', NOW(), NOW(), ARRAY['web','go']);
INSERT INTO posts (post_id, type, author_id, title, created_at, last_modified, tags) VALUES (
'proj:id2','project', 'author2', 'proj2_name', NOW(), NOW(), ARRAY['db','sql']);
INSERT INTO posts (post_id, type, author_id, title, created_at, last_modified, tags) VALUES (
'proj:id3', 'project', 'author3', 'proj3_name', NOW(), NOW(), ARRAY['api','rest']);
INSERT INTO posts (post_id, type, author_id, title, created_at, last_modified, tags) VALUES (
'blg:id1', 'blog', 'author1', 'blog1_name', NOW(), NOW(), ARRAY['intro','go']);
INSERT INTO posts (post_id, type, author_id, title, created_at, last_modified, tags) VALUES (
'blg"id2', 'blog', 'author2', 'blog2_name', NOW(), NOW(), ARRAY['tutorial','db']);
INSERT INTO posts (post_id, type, author_id, title, created_at, last_modified, tags) VALUES (
'blg:id3', 'blog', 'author3', 'blog3_name', NOW(), NOW(), ARRAY['advanced','web']);