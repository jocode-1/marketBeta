package queries

const GetUserByEmail = `
SELECT user_id, username, email, hashed_password, phone_number, user_address, profile_photo_url, 
       ip_address, is_verified, is_admin, is_vendor, role, status, updated_at, created_at 
FROM users 
WHERE email = $1
LIMIT 1;
`
