DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS file_tag;

CREATE TABLE files(
    id INTEGER PRIMARY KEY,
    file_path TEXT UNIQUE NOT NULL,
    created datetime DEFAULT current_timestamp
);
CREATE TABLE tags (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE file_tag (
    id INTEGER PRIMARY KEY,
    file_id INTEGER,
    tag_id INTEGER,
    FOREIGN KEY(file_id) REFERENCES files(id), 
    FOREIGN KEY(tag_id) REFERENCES tag(id)
);