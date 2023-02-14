START TRANSACTION;

CREATE TABLE IF NOT EXISTS teachers(
  id INTEGER NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL UNIQUE,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS students (
  id INTEGER NOT NULL AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL UNIQUE,
  is_suspended boolean NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS classes (
  teacher_id INTEGER NOT NULL,
  student_id INTEGER NOT NULL,
  PRIMARY KEY (teacher_id, student_id),
  FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE,
  FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
);

COMMIT;