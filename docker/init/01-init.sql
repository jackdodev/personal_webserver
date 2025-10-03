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