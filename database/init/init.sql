-- Database and user are created by Docker environment variables
-- This script runs after database creation


CREATE TABLE author (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  -- email VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE blogs (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  author_id INTEGER,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY (author_id) REFERENCES author(id) ON DELETE SET NULL
);

CREATE TABLE file (
  id SERIAL PRIMARY KEY,
  blog_id INTEGER NOT NULL,
  data TEXT NOT NULL,
  FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE
);

CREATE TABLE tag (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE blog_tags (
  blog_id INTEGER NOT NULL,
  tag_id INTEGER NOT NULL,
  PRIMARY KEY (blog_id, tag_id),
  FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE
);

CREATE TABLE statistics (
  id SERIAL PRIMARY KEY,
  blog_id INTEGER NOT NULL,
  started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  ended_at TIMESTAMP,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE
)

-- Grant permissions to admin user
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin;