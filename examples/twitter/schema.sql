DROP TABLE IF EXISTS public.twitter_users;

CREATE TABLE public.twitter_users (
    id bigint NOT NULL,
    screen_name character varying(24) NOT NULL,
    name character varying(24) NOT NULL,
    description character varying,
    protected bool NOT NULL,
    verified bool NOT NULL,

    twitter_data jsonb NOT NULL,
    twitter_created_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

DROP TABLE IF EXISTS public.twitter_tweets;

CREATE TABLE public.twitter_tweets (
    id bigint NOT NULL,
    user_id bigint NOT NULL,

    text character varying NOT NULL,
    source character varying(24) NOT NULL,
    in_reply_to_status_id bigint,
    in_reply_to_user_id bigint,
    in_reply_to_screen_name character varying(24) NOT NULL,

    quoted_status_id bigint,
    is_quote_status bool NOT NULL,

    twitter_data jsonb NOT NULL,
    twitter_created_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);
