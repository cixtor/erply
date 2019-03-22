CREATE TABLE IF NOT EXISTS contacts (id INTEGER NOT NULL PRIMARY KEY, firstname TEXT, lastname TEXT, phone TEXT, address TEXT, email TEXT);

INSERT OR IGNORE INTO contacts VALUES (1, "John", "Smith", "6045551234", "350 W Georgia St, Vancouver, BC", "john@example.com");
INSERT OR IGNORE INTO contacts VALUES (2, "Alice", "Smith", "6045551234", "350 W Georgia St, Vancouver, BC", "alice@example.com");
