CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name VARCHAR(255),
    price BIGINT NOT NULL DEFAULT 0,
    user_id UUID NOT NULL DEFAULT gen_random_uuid(),
    start_date TIMESTAMP NOT NULL DEFAULT NOW(),
    end_date TIMESTAMP DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_subscriptions_id on subscriptions(id);