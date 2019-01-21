DROP TABLE IF EXISTS chatter_users;

CREATE TABLE chatter_users (
        id SERIAL NOT NULL UNIQUE,
        auto_user_id varchar(255) NOT NULL,
        message TEXT NOT NULL,
	time BIGINT NOT NULL
);
