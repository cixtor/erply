CREATE TABLE IF NOT EXISTS keys (id INTEGER NOT NULL PRIMARY KEY, username TEXT, password TEXT);

INSERT OR IGNORE INTO keys VALUES (1, "john@example.com", "85EC496B-7EC4-4478-B27B-94B381B4030F");
INSERT OR IGNORE INTO keys VALUES (2, "alice@example.com", "7CDBA0B1-8ACF-4E7F-9998-51038DF231E6");

CREATE TABLE IF NOT EXISTS contacts (id INTEGER NOT NULL PRIMARY KEY, firstname TEXT, lastname TEXT, phone TEXT, address TEXT, email TEXT);

INSERT OR IGNORE INTO contacts VALUES (1, "John", "Smith", "6045551234", "350 W Georgia St, Vancouver, BC", "john@example.com");
INSERT OR IGNORE INTO contacts VALUES (2, "Alice", "Smith", "6045551234", "350 W Georgia St, Vancouver, BC", "alice@example.com");
