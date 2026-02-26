UPDATE users
SET 
    password_hash = crypt('Admin123!', gen_salt('bf')),
    deleted_at = NULL,
    updated_at = NOW()
WHERE username = 'admin';

INSERT INTO users (username, password_hash)
SELECT 'admin', crypt('Admin123!', gen_salt('bf'))
WHERE NOT EXISTS (
    SELECT 1 FROM users WHERE username = 'admin'
);


UPDATE users
SET 
    password_hash = crypt('Test123!', gen_salt('bf')),
    deleted_at = NULL,
    updated_at = NOW()
WHERE username = 'testuser';

INSERT INTO users (username, password_hash)
SELECT 'testuser', crypt('Test123!', gen_salt('bf'))
WHERE NOT EXISTS (
    SELECT 1 FROM users WHERE username = 'testuser'
);