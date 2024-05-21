CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	email VARCHAR(20) UNIQUE,
	password VARCHAR(255) NOT NULL,
	is_premium BOOLEAN DEFAULT FALSE,
	is_verified BOOLEAN DEFAULT FALSE,
	is_infinite_quota BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS members (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT,
	name VARCHAR(50) NOT NULL,
	gender VARCHAR(6),

	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS swipe_members (
	id BIGSERIAL PRIMARY KEY,
	member_id BIGINT NOT NULL,
	swiped_member_id BIGINT NOT NULL,
	is_liked BOOLEAN NOT NULL

	FOREIGN KEY (member_id) REFERENCES members(id),
	FOREIGN KEY (swiped_member_id) REFERENCES members(id)
);
