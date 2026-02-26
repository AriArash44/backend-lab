UPDATE users
SET deleted_at = NOW(),
    updated_at = NOW()
WHERE username IN ('admin', 'testuser')
AND deleted_at IS NULL;