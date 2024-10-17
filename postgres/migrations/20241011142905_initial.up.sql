CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  email_verified BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  login_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS auth_providers (
  provider VARCHAR(255) NOT NULL PRIMARY KEY,
  display_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_identities (
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  auth_provider VARCHAR(255) NOT NULL REFERENCES auth_providers(provider) ON DELETE CASCADE,
  provider_id VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (user_id, auth_provider)
);

INSERT INTO auth_providers (provider, display_name) VALUES
  ('apple', 'Apple');

CREATE UNIQUE INDEX idx_user_identities_provider_id ON user_identities (provider_id);
