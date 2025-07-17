CREATE TABLE IF NOT EXISTS public.users (
    id            integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login         text NOT NULL UNIQUE,
    password_hash text NOT NULL,
    created_at    timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS public.ads (
    id          integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id     integer NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    title       text NOT NULL,
    description text NOT NULL,
    image_url   text,
    price       numeric(12,2) NOT NULL CHECK (price >= 0),
    created_at  timestamptz DEFAULT now() NOT NULL
);

CREATE INDEX IF NOT EXISTS ads_created_at_desc_idx ON public.ads (created_at DESC);
CREATE INDEX IF NOT EXISTS ads_price_idx           ON public.ads (price);