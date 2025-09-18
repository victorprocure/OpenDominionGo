--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5 (Debian 17.5-1.pgdg120+1)
-- Dumped by pg_dump version 17.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: achievements; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.achievements (
    id integer NOT NULL,
    name character varying(191),
    description character varying(191),
    icon character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: achievements_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.achievements_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: achievements_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.achievements_id_seq OWNED BY public.achievements.id;


--
-- Name: bounties; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.bounties (
    id integer NOT NULL,
    round_id integer NOT NULL,
    source_realm_id integer NOT NULL,
    source_dominion_id integer NOT NULL,
    target_dominion_id integer NOT NULL,
    collected_by_dominion_id integer,
    type character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    reward boolean DEFAULT false NOT NULL
);


--
-- Name: bounties_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.bounties_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: bounties_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.bounties_id_seq OWNED BY public.bounties.id;


--
-- Name: council_posts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.council_posts (
    id integer NOT NULL,
    council_thread_id integer NOT NULL,
    dominion_id integer NOT NULL,
    body text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


--
-- Name: council_posts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.council_posts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: council_posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.council_posts_id_seq OWNED BY public.council_posts.id;


--
-- Name: council_threads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.council_threads (
    id integer NOT NULL,
    realm_id integer NOT NULL,
    dominion_id integer NOT NULL,
    title character varying(191) NOT NULL,
    body text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    last_activity timestamp(0) without time zone
);


--
-- Name: council_threads_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.council_threads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: council_threads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.council_threads_id_seq OWNED BY public.council_threads.id;


--
-- Name: daily_rankings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.daily_rankings (
    id integer NOT NULL,
    round_id integer NOT NULL,
    dominion_id integer NOT NULL,
    dominion_name character varying(191) NOT NULL,
    race_name character varying(191) NOT NULL,
    realm_number integer NOT NULL,
    realm_name character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    key character varying(191) NOT NULL,
    value integer DEFAULT 0 NOT NULL,
    rank integer,
    previous_rank integer
);


--
-- Name: daily_rankings_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.daily_rankings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: daily_rankings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.daily_rankings_id_seq OWNED BY public.daily_rankings.id;


--
-- Name: dominion_history; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dominion_history (
    id integer NOT NULL,
    dominion_id integer NOT NULL,
    event character varying(191) NOT NULL,
    delta text NOT NULL,
    created_at timestamp(3) without time zone,
    ip character varying(191) NOT NULL,
    device character varying(191) NOT NULL
);


--
-- Name: dominion_history_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dominion_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dominion_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dominion_history_id_seq OWNED BY public.dominion_history.id;


--
-- Name: dominion_journals; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dominion_journals (
    id integer NOT NULL,
    dominion_id integer NOT NULL,
    content text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: dominion_journals_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dominion_journals_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dominion_journals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dominion_journals_id_seq OWNED BY public.dominion_journals.id;


--
-- Name: dominion_queue; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dominion_queue (
    dominion_id integer NOT NULL,
    source character varying(191) NOT NULL,
    resource character varying(191) NOT NULL,
    hours integer NOT NULL,
    amount integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: dominion_spells; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dominion_spells (
    dominion_id integer NOT NULL,
    duration integer NOT NULL,
    cast_by_dominion_id integer,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    spell_id integer NOT NULL
);


--
-- Name: dominion_techs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dominion_techs (
    id integer NOT NULL,
    dominion_id integer NOT NULL,
    tech_id integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: dominion_techs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dominion_techs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dominion_techs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dominion_techs_id_seq OWNED BY public.dominion_techs.id;


--
-- Name: dominion_tick; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dominion_tick (
    id integer NOT NULL,
    dominion_id integer NOT NULL,
    prestige integer DEFAULT 0 NOT NULL,
    peasants integer DEFAULT 0 NOT NULL,
    morale integer DEFAULT 0 NOT NULL,
    spy_strength double precision DEFAULT '0'::double precision NOT NULL,
    wizard_strength double precision DEFAULT '0'::double precision NOT NULL,
    resource_platinum integer DEFAULT 0 NOT NULL,
    resource_food integer DEFAULT 0 NOT NULL,
    resource_food_production integer DEFAULT 0 NOT NULL,
    resource_lumber integer DEFAULT 0 NOT NULL,
    resource_lumber_production integer DEFAULT 0 NOT NULL,
    resource_mana integer DEFAULT 0 NOT NULL,
    resource_mana_production integer DEFAULT 0 NOT NULL,
    resource_ore integer DEFAULT 0 NOT NULL,
    resource_gems integer DEFAULT 0 NOT NULL,
    resource_tech integer DEFAULT 0 NOT NULL,
    resource_boats double precision DEFAULT '0'::double precision NOT NULL,
    military_draftees integer DEFAULT 0 NOT NULL,
    military_unit1 integer DEFAULT 0 NOT NULL,
    military_unit2 integer DEFAULT 0 NOT NULL,
    military_unit3 integer DEFAULT 0 NOT NULL,
    military_unit4 integer DEFAULT 0 NOT NULL,
    military_spies integer DEFAULT 0 NOT NULL,
    military_wizards integer DEFAULT 0 NOT NULL,
    military_archmages integer DEFAULT 0 NOT NULL,
    land_plain integer DEFAULT 0 NOT NULL,
    land_mountain integer DEFAULT 0 NOT NULL,
    land_swamp integer DEFAULT 0 NOT NULL,
    land_cavern integer DEFAULT 0 NOT NULL,
    land_forest integer DEFAULT 0 NOT NULL,
    land_hill integer DEFAULT 0 NOT NULL,
    land_water integer DEFAULT 0 NOT NULL,
    discounted_land integer DEFAULT 0 NOT NULL,
    building_home integer DEFAULT 0 NOT NULL,
    building_alchemy integer DEFAULT 0 NOT NULL,
    building_farm integer DEFAULT 0 NOT NULL,
    building_smithy integer DEFAULT 0 NOT NULL,
    building_masonry integer DEFAULT 0 NOT NULL,
    building_ore_mine integer DEFAULT 0 NOT NULL,
    building_gryphon_nest integer DEFAULT 0 NOT NULL,
    building_tower integer DEFAULT 0 NOT NULL,
    building_wizard_guild integer DEFAULT 0 NOT NULL,
    building_temple integer DEFAULT 0 NOT NULL,
    building_diamond_mine integer DEFAULT 0 NOT NULL,
    building_school integer DEFAULT 0 NOT NULL,
    building_lumberyard integer DEFAULT 0 NOT NULL,
    building_forest_haven integer DEFAULT 0 NOT NULL,
    building_factory integer DEFAULT 0 NOT NULL,
    building_guard_tower integer DEFAULT 0 NOT NULL,
    building_shrine integer DEFAULT 0 NOT NULL,
    building_barracks integer DEFAULT 0 NOT NULL,
    building_dock integer DEFAULT 0 NOT NULL,
    starvation_casualties text,
    updated_at timestamp(0) without time zone,
    highest_land_achieved integer DEFAULT 0 NOT NULL,
    calculated_networth integer DEFAULT 0 NOT NULL,
    resource_food_decay integer DEFAULT 0 NOT NULL,
    resource_lumber_decay integer DEFAULT 0 NOT NULL,
    resource_mana_decay integer DEFAULT 0 NOT NULL,
    resource_boat_production double precision DEFAULT '0'::double precision NOT NULL,
    expiring_spells text,
    military_assassins integer DEFAULT 0 NOT NULL,
    improvement_science integer DEFAULT 0 NOT NULL,
    improvement_keep integer DEFAULT 0 NOT NULL,
    improvement_forges integer DEFAULT 0 NOT NULL,
    improvement_walls integer DEFAULT 0 NOT NULL,
    resilience integer DEFAULT 0 NOT NULL,
    fireball_meter integer DEFAULT 0 NOT NULL,
    lightning_bolt_meter integer DEFAULT 0 NOT NULL
);


--
-- Name: dominion_tick_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dominion_tick_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dominion_tick_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dominion_tick_id_seq OWNED BY public.dominion_tick.id;


--
-- Name: dominions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.dominions (
    id integer NOT NULL,
    user_id integer,
    round_id integer NOT NULL,
    realm_id integer NOT NULL,
    race_id integer NOT NULL,
    name character varying(191) NOT NULL,
    prestige integer NOT NULL,
    peasants integer NOT NULL,
    peasants_last_hour integer DEFAULT 0 NOT NULL,
    draft_rate integer NOT NULL,
    morale integer NOT NULL,
    spy_strength numeric(6,3) NOT NULL,
    wizard_strength numeric(6,3) NOT NULL,
    resource_platinum integer NOT NULL,
    resource_food integer NOT NULL,
    resource_lumber integer NOT NULL,
    resource_mana integer NOT NULL,
    resource_ore integer NOT NULL,
    resource_gems integer NOT NULL,
    resource_tech integer NOT NULL,
    resource_boats numeric(10,4) NOT NULL,
    improvement_science integer NOT NULL,
    improvement_keep integer NOT NULL,
    improvement_spires integer NOT NULL,
    improvement_forges integer NOT NULL,
    improvement_walls integer NOT NULL,
    improvement_harbor integer NOT NULL,
    military_draftees integer NOT NULL,
    military_unit1 integer NOT NULL,
    military_unit2 integer NOT NULL,
    military_unit3 integer NOT NULL,
    military_unit4 integer NOT NULL,
    military_spies integer NOT NULL,
    military_wizards integer NOT NULL,
    military_archmages integer NOT NULL,
    land_plain integer NOT NULL,
    land_mountain integer NOT NULL,
    land_swamp integer NOT NULL,
    land_cavern integer NOT NULL,
    land_forest integer NOT NULL,
    land_hill integer NOT NULL,
    land_water integer NOT NULL,
    building_home integer NOT NULL,
    building_alchemy integer NOT NULL,
    building_farm integer NOT NULL,
    building_smithy integer NOT NULL,
    building_masonry integer NOT NULL,
    building_ore_mine integer NOT NULL,
    building_gryphon_nest integer NOT NULL,
    building_tower integer NOT NULL,
    building_wizard_guild integer NOT NULL,
    building_temple integer NOT NULL,
    building_diamond_mine integer NOT NULL,
    building_school integer NOT NULL,
    building_lumberyard integer NOT NULL,
    building_forest_haven integer NOT NULL,
    building_factory integer NOT NULL,
    building_guard_tower integer NOT NULL,
    building_shrine integer NOT NULL,
    building_barracks integer NOT NULL,
    building_dock integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    daily_platinum boolean DEFAULT false NOT NULL,
    daily_land boolean DEFAULT false NOT NULL,
    ruler_name character varying(191),
    pack_id integer,
    discounted_land integer DEFAULT 0 NOT NULL,
    council_last_read timestamp(0) without time zone,
    royal_guard_active_at timestamp(0) without time zone,
    elite_guard_active_at timestamp(0) without time zone,
    monarchy_vote_for_dominion_id integer,
    stat_attacking_success integer DEFAULT 0 NOT NULL,
    stat_defending_success integer DEFAULT 0 NOT NULL,
    stat_espionage_success integer DEFAULT 0 NOT NULL,
    stat_spell_success integer DEFAULT 0 NOT NULL,
    stat_total_platinum_production integer DEFAULT 0 NOT NULL,
    stat_total_food_production integer DEFAULT 0 NOT NULL,
    stat_total_lumber_production integer DEFAULT 0 NOT NULL,
    stat_total_mana_production integer DEFAULT 0 NOT NULL,
    stat_total_ore_production integer DEFAULT 0 NOT NULL,
    stat_total_gem_production integer DEFAULT 0 NOT NULL,
    stat_total_tech_production integer DEFAULT 0 NOT NULL,
    stat_total_boat_production double precision DEFAULT '0'::double precision NOT NULL,
    stat_total_land_explored integer DEFAULT 0 NOT NULL,
    stat_total_land_conquered integer DEFAULT 0 NOT NULL,
    last_tick_at timestamp(0) without time zone,
    stat_total_platinum_stolen integer DEFAULT 0 NOT NULL,
    stat_total_food_stolen integer DEFAULT 0 NOT NULL,
    stat_total_lumber_stolen integer DEFAULT 0 NOT NULL,
    stat_total_mana_stolen integer DEFAULT 0 NOT NULL,
    stat_total_ore_stolen integer DEFAULT 0 NOT NULL,
    stat_total_gems_stolen integer DEFAULT 0 NOT NULL,
    spy_mastery integer DEFAULT 0 NOT NULL,
    wizard_mastery integer DEFAULT 0 NOT NULL,
    stat_assassinate_draftees_damage integer DEFAULT 0 NOT NULL,
    stat_assassinate_wizards_damage integer DEFAULT 0 NOT NULL,
    stat_magic_snare_damage integer DEFAULT 0 NOT NULL,
    stat_sabotage_boats_damage integer DEFAULT 0 NOT NULL,
    stat_disband_spies_damage integer DEFAULT 0 NOT NULL,
    stat_fireball_damage integer DEFAULT 0 NOT NULL,
    stat_lightning_bolt_damage integer DEFAULT 0 NOT NULL,
    stat_earthquake_hours integer DEFAULT 0 NOT NULL,
    stat_great_flood_hours integer DEFAULT 0 NOT NULL,
    stat_insect_swarm_hours integer DEFAULT 0 NOT NULL,
    stat_plague_hours integer DEFAULT 0 NOT NULL,
    highest_land_achieved integer DEFAULT 250 NOT NULL,
    stat_attacking_failure integer DEFAULT 0 NOT NULL,
    stat_defending_failure integer DEFAULT 0 NOT NULL,
    stat_espionage_failure integer DEFAULT 0 NOT NULL,
    stat_spell_failure integer DEFAULT 0 NOT NULL,
    stat_spies_executed integer DEFAULT 0 NOT NULL,
    stat_wizards_executed integer DEFAULT 0 NOT NULL,
    stat_spells_reflected integer DEFAULT 0 NOT NULL,
    forum_last_read timestamp(0) without time zone,
    protection_ticks_remaining integer DEFAULT 0 NOT NULL,
    locked_at timestamp(0) without time zone,
    calculated_networth integer DEFAULT 0 NOT NULL,
    settings text,
    stat_cyclone_damage integer DEFAULT 0 NOT NULL,
    stat_wonder_damage integer DEFAULT 0 NOT NULL,
    stat_wonders_destroyed integer DEFAULT 0 NOT NULL,
    stat_total_land_lost integer DEFAULT 0 NOT NULL,
    stat_spies_lost integer DEFAULT 0 NOT NULL,
    stat_wizards_lost integer DEFAULT 0 NOT NULL,
    town_crier_last_seen timestamp(0) without time zone,
    wonders_last_seen timestamp(0) without time zone,
    stat_total_platinum_spent_construction integer DEFAULT 0 NOT NULL,
    stat_total_lumber_spent_construction integer DEFAULT 0 NOT NULL,
    stat_total_platinum_spent_exploration integer DEFAULT 0 NOT NULL,
    stat_total_platinum_spent_investment integer DEFAULT 0 NOT NULL,
    stat_total_lumber_spent_investment integer DEFAULT 0 NOT NULL,
    stat_total_mana_spent_investment integer DEFAULT 0 NOT NULL,
    stat_total_ore_spent_investment integer DEFAULT 0 NOT NULL,
    stat_total_gems_spent_investment integer DEFAULT 0 NOT NULL,
    stat_total_platinum_spent_rezoning integer DEFAULT 0 NOT NULL,
    stat_total_platinum_spent_training integer DEFAULT 0 NOT NULL,
    stat_total_lumber_spent_training integer DEFAULT 0 NOT NULL,
    stat_total_mana_spent_training integer DEFAULT 0 NOT NULL,
    stat_total_ore_spent_training integer DEFAULT 0 NOT NULL,
    stat_total_gems_spent_training integer DEFAULT 0 NOT NULL,
    stat_total_food_decay integer DEFAULT 0 NOT NULL,
    stat_total_lumber_decay integer DEFAULT 0 NOT NULL,
    stat_total_mana_decay integer DEFAULT 0 NOT NULL,
    ai_enabled boolean DEFAULT false NOT NULL,
    ai_config text,
    abandoned_at timestamp(0) without time zone,
    black_guard_active_at timestamp(0) without time zone,
    black_guard_inactive_at timestamp(0) without time zone,
    stat_assassinate_draftees_damage_received integer DEFAULT 0 NOT NULL,
    stat_assassinate_wizards_damage_received integer DEFAULT 0 NOT NULL,
    stat_magic_snare_damage_received integer DEFAULT 0 NOT NULL,
    stat_sabotage_boats_damage_received integer DEFAULT 0 NOT NULL,
    stat_disband_spies_damage_received integer DEFAULT 0 NOT NULL,
    stat_fireball_damage_received integer DEFAULT 0 NOT NULL,
    stat_lightning_bolt_damage_received integer DEFAULT 0 NOT NULL,
    stat_earthquake_hours_received integer DEFAULT 0 NOT NULL,
    stat_great_flood_hours_received integer DEFAULT 0 NOT NULL,
    stat_insect_swarm_hours_received integer DEFAULT 0 NOT NULL,
    stat_plague_hours_received integer DEFAULT 0 NOT NULL,
    stat_spells_deflected integer DEFAULT 0 NOT NULL,
    military_assassins integer DEFAULT 0 NOT NULL,
    hourly_activity bytea,
    daily_actions integer DEFAULT 0 NOT NULL,
    stat_total_investment integer DEFAULT 0 NOT NULL,
    stat_bounties_collected integer DEFAULT 0 NOT NULL,
    resilience integer DEFAULT 0 NOT NULL,
    fireball_meter integer DEFAULT 0 NOT NULL,
    lightning_bolt_meter integer DEFAULT 0 NOT NULL,
    valor integer DEFAULT 0 NOT NULL,
    stat_military_unit1_lost integer DEFAULT 0 NOT NULL,
    stat_military_unit2_lost integer DEFAULT 0 NOT NULL,
    stat_military_unit3_lost integer DEFAULT 0 NOT NULL,
    stat_military_unit4_lost integer DEFAULT 0 NOT NULL,
    racial_value double precision DEFAULT '0'::double precision NOT NULL,
    chaos integer DEFAULT 0 NOT NULL,
    stat_incite_chaos_damage integer DEFAULT 0 NOT NULL,
    stat_incite_chaos_damage_received integer DEFAULT 0 NOT NULL,
    stat_spies_charmed integer DEFAULT 0 NOT NULL
);


--
-- Name: dominions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.dominions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: dominions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.dominions_id_seq OWNED BY public.dominions.id;


--
-- Name: failed_jobs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.failed_jobs (
    id bigint NOT NULL,
    connection text NOT NULL,
    queue text NOT NULL,
    payload text NOT NULL,
    exception text NOT NULL,
    failed_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: failed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.failed_jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: failed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.failed_jobs_id_seq OWNED BY public.failed_jobs.id;


--
-- Name: forum_posts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.forum_posts (
    id integer NOT NULL,
    forum_thread_id integer NOT NULL,
    dominion_id integer NOT NULL,
    body text NOT NULL,
    flagged_for_removal boolean DEFAULT false NOT NULL,
    flagged_by text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


--
-- Name: forum_posts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.forum_posts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: forum_posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.forum_posts_id_seq OWNED BY public.forum_posts.id;


--
-- Name: forum_threads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.forum_threads (
    id integer NOT NULL,
    round_id integer NOT NULL,
    dominion_id integer NOT NULL,
    title character varying(191) NOT NULL,
    body text NOT NULL,
    flagged_for_removal boolean DEFAULT false NOT NULL,
    flagged_by text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    last_activity timestamp(0) without time zone
);


--
-- Name: forum_threads_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.forum_threads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: forum_threads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.forum_threads_id_seq OWNED BY public.forum_threads.id;


--
-- Name: game_events; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.game_events (
    id uuid NOT NULL,
    round_id integer NOT NULL,
    source_type character varying(191) NOT NULL,
    source_id bigint NOT NULL,
    target_type character varying(191),
    target_id bigint,
    type character varying(191) NOT NULL,
    data text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: hero_battle_actions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_battle_actions (
    id integer NOT NULL,
    hero_battle_id integer NOT NULL,
    combatant_id integer NOT NULL,
    target_combatant_id integer,
    turn integer NOT NULL,
    action character varying(191) NOT NULL,
    damage integer NOT NULL,
    health integer NOT NULL,
    description character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: hero_battle_actions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_battle_actions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_battle_actions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_battle_actions_id_seq OWNED BY public.hero_battle_actions.id;


--
-- Name: hero_battle_queue; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_battle_queue (
    id integer NOT NULL,
    hero_id integer NOT NULL,
    level integer NOT NULL,
    rating integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: hero_battle_queue_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_battle_queue_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_battle_queue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_battle_queue_id_seq OWNED BY public.hero_battle_queue.id;


--
-- Name: hero_battles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_battles (
    id integer NOT NULL,
    round_id integer,
    current_turn integer DEFAULT 1 NOT NULL,
    winner_combatant_id integer,
    finished boolean DEFAULT false NOT NULL,
    last_processed_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    pvp boolean DEFAULT true NOT NULL
);


--
-- Name: hero_battles_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_battles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_battles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_battles_id_seq OWNED BY public.hero_battles.id;


--
-- Name: hero_combatants; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_combatants (
    id integer NOT NULL,
    hero_battle_id integer NOT NULL,
    hero_id integer,
    dominion_id integer,
    name character varying(191) NOT NULL,
    health integer NOT NULL,
    attack integer NOT NULL,
    defense integer NOT NULL,
    evasion integer NOT NULL,
    focus integer NOT NULL,
    counter integer NOT NULL,
    recover integer NOT NULL,
    current_health integer NOT NULL,
    has_focus boolean DEFAULT false NOT NULL,
    actions text,
    last_action character varying(191),
    time_bank integer DEFAULT 86400 NOT NULL,
    automated boolean,
    strategy character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    level integer DEFAULT 0 NOT NULL
);


--
-- Name: hero_combatants_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_combatants_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_combatants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_combatants_id_seq OWNED BY public.hero_combatants.id;


--
-- Name: hero_hero_upgrades; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_hero_upgrades (
    id integer NOT NULL,
    hero_id integer NOT NULL,
    hero_upgrade_id integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: hero_hero_upgrades_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_hero_upgrades_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_hero_upgrades_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_hero_upgrades_id_seq OWNED BY public.hero_hero_upgrades.id;


--
-- Name: hero_tournament_battles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_tournament_battles (
    id integer NOT NULL,
    hero_tournament_id integer NOT NULL,
    hero_battle_id integer NOT NULL,
    round_number integer DEFAULT 1 NOT NULL
);


--
-- Name: hero_tournament_battles_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_tournament_battles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_tournament_battles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_tournament_battles_id_seq OWNED BY public.hero_tournament_battles.id;


--
-- Name: hero_tournament_participants; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_tournament_participants (
    id integer NOT NULL,
    hero_tournament_id integer NOT NULL,
    hero_id integer NOT NULL,
    wins integer DEFAULT 0 NOT NULL,
    losses integer DEFAULT 0 NOT NULL,
    draws integer DEFAULT 0 NOT NULL,
    standing integer,
    eliminated boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: hero_tournament_participants_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_tournament_participants_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_tournament_participants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_tournament_participants_id_seq OWNED BY public.hero_tournament_participants.id;


--
-- Name: hero_tournaments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_tournaments (
    id integer NOT NULL,
    round_id integer,
    name character varying(191) NOT NULL,
    current_round_number integer DEFAULT 1 NOT NULL,
    finished boolean DEFAULT false NOT NULL,
    winner_dominion_id integer,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    start_date timestamp(0) without time zone
);


--
-- Name: hero_tournaments_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_tournaments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_tournaments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_tournaments_id_seq OWNED BY public.hero_tournaments.id;


--
-- Name: hero_upgrade_perks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_upgrade_perks (
    id integer NOT NULL,
    hero_upgrade_id integer NOT NULL,
    key character varying(191) NOT NULL,
    value character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: hero_upgrade_perks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_upgrade_perks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_upgrade_perks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_upgrade_perks_id_seq OWNED BY public.hero_upgrade_perks.id;


--
-- Name: hero_upgrades; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hero_upgrades (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    name character varying(191) NOT NULL,
    level integer NOT NULL,
    type character varying(191) NOT NULL,
    icon character varying(191) NOT NULL,
    classes text,
    active boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: hero_upgrades_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hero_upgrades_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hero_upgrades_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hero_upgrades_id_seq OWNED BY public.hero_upgrades.id;


--
-- Name: heroes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.heroes (
    id integer NOT NULL,
    dominion_id integer NOT NULL,
    name character varying(191),
    class character varying(191),
    experience double precision DEFAULT '0'::double precision NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    stat_combat_wins integer DEFAULT 0 NOT NULL,
    stat_combat_losses integer DEFAULT 0 NOT NULL,
    stat_combat_draws integer DEFAULT 0 NOT NULL,
    combat_rating integer DEFAULT 1000 NOT NULL
);


--
-- Name: heroes_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.heroes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: heroes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.heroes_id_seq OWNED BY public.heroes.id;


--
-- Name: info_ops; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.info_ops (
    id integer NOT NULL,
    source_realm_id integer NOT NULL,
    source_dominion_id integer NOT NULL,
    target_dominion_id integer NOT NULL,
    type character varying(191) NOT NULL,
    data text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    target_realm_id integer,
    latest boolean DEFAULT true NOT NULL
);


--
-- Name: info_ops_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.info_ops_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: info_ops_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.info_ops_id_seq OWNED BY public.info_ops.id;


--
-- Name: jobs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.jobs (
    id bigint NOT NULL,
    queue character varying(191) NOT NULL,
    payload text NOT NULL,
    attempts smallint NOT NULL,
    reserved_at integer,
    available_at integer NOT NULL,
    created_at integer NOT NULL
);


--
-- Name: jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.jobs_id_seq OWNED BY public.jobs.id;


--
-- Name: message_board_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.message_board_categories (
    id integer NOT NULL,
    name character varying(191) NOT NULL,
    slug character varying(191) NOT NULL,
    role_required character varying(191)
);


--
-- Name: message_board_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.message_board_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: message_board_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.message_board_categories_id_seq OWNED BY public.message_board_categories.id;


--
-- Name: message_board_posts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.message_board_posts (
    id integer NOT NULL,
    message_board_thread_id integer NOT NULL,
    user_id integer NOT NULL,
    body text NOT NULL,
    flagged_for_removal boolean DEFAULT false NOT NULL,
    flagged_by text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


--
-- Name: message_board_posts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.message_board_posts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: message_board_posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.message_board_posts_id_seq OWNED BY public.message_board_posts.id;


--
-- Name: message_board_threads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.message_board_threads (
    id integer NOT NULL,
    message_board_category_id integer NOT NULL,
    user_id integer NOT NULL,
    title character varying(191) NOT NULL,
    body text NOT NULL,
    flagged_for_removal boolean DEFAULT false NOT NULL,
    flagged_by text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone,
    last_activity timestamp(0) without time zone
);


--
-- Name: message_board_threads_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.message_board_threads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: message_board_threads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.message_board_threads_id_seq OWNED BY public.message_board_threads.id;


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.migrations (
    id integer NOT NULL,
    migration character varying(191) NOT NULL,
    batch integer NOT NULL
);


--
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.migrations_id_seq OWNED BY public.migrations.id;


--
-- Name: model_has_permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.model_has_permissions (
    permission_id integer NOT NULL,
    model_type character varying(191) NOT NULL,
    model_id bigint NOT NULL
);


--
-- Name: model_has_roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.model_has_roles (
    role_id integer NOT NULL,
    model_type character varying(191) NOT NULL,
    model_id bigint NOT NULL
);


--
-- Name: notifications; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.notifications (
    id uuid NOT NULL,
    type character varying(191) NOT NULL,
    notifiable_type character varying(191) NOT NULL,
    notifiable_id bigint NOT NULL,
    data text NOT NULL,
    read_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: packs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.packs (
    id integer NOT NULL,
    round_id integer NOT NULL,
    realm_id integer,
    name character varying(191) NOT NULL,
    password character varying(191) NOT NULL,
    size integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    closed_at timestamp(0) without time zone,
    creator_dominion_id integer DEFAULT 0 NOT NULL,
    rating integer DEFAULT 0 NOT NULL
);


--
-- Name: packs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.packs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: packs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.packs_id_seq OWNED BY public.packs.id;


--
-- Name: password_resets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.password_resets (
    email character varying(191) NOT NULL,
    token character varying(191) NOT NULL,
    created_at timestamp(0) without time zone NOT NULL
);


--
-- Name: permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.permissions (
    id integer NOT NULL,
    name character varying(191) NOT NULL,
    guard_name character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: race_perk_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.race_perk_types (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: race_perk_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.race_perk_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: race_perk_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.race_perk_types_id_seq OWNED BY public.race_perk_types.id;


--
-- Name: race_perks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.race_perks (
    id integer NOT NULL,
    race_id integer NOT NULL,
    race_perk_type_id integer NOT NULL,
    value double precision NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: race_perks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.race_perks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: race_perks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.race_perks_id_seq OWNED BY public.race_perks.id;


--
-- Name: races; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.races (
    id integer NOT NULL,
    name character varying(191) NOT NULL,
    alignment character varying(255) NOT NULL,
    home_land_type character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    description text,
    playable boolean DEFAULT true NOT NULL,
    attacker_difficulty integer DEFAULT 0 NOT NULL,
    explorer_difficulty integer DEFAULT 0 NOT NULL,
    converter_difficulty integer DEFAULT 0 NOT NULL,
    overall_difficulty integer DEFAULT 0 NOT NULL,
    key character varying(191),
    CONSTRAINT races_alignment_check CHECK (((alignment)::text = ANY ((ARRAY['good'::character varying, 'neutral'::character varying, 'evil'::character varying, 'other'::character varying])::text[]))),
    CONSTRAINT races_home_land_type_check CHECK (((home_land_type)::text = ANY ((ARRAY['plain'::character varying, 'mountain'::character varying, 'swamp'::character varying, 'cavern'::character varying, 'forest'::character varying, 'hill'::character varying, 'water'::character varying])::text[])))
);


--
-- Name: races_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.races_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: races_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.races_id_seq OWNED BY public.races.id;


--
-- Name: realm_history; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.realm_history (
    id integer NOT NULL,
    realm_id integer NOT NULL,
    dominion_id integer NOT NULL,
    event character varying(191) NOT NULL,
    delta text NOT NULL,
    created_at timestamp(0) without time zone
);


--
-- Name: realm_history_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.realm_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: realm_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.realm_history_id_seq OWNED BY public.realm_history.id;


--
-- Name: realm_wars; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.realm_wars (
    id integer NOT NULL,
    source_realm_id integer NOT NULL,
    source_realm_name_start character varying(191),
    target_realm_id integer NOT NULL,
    target_realm_name_start character varying(191),
    active_at timestamp(0) without time zone,
    inactive_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    source_realm_name_end character varying(191),
    target_realm_name_end character varying(191)
);


--
-- Name: realm_wars_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.realm_wars_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: realm_wars_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.realm_wars_id_seq OWNED BY public.realm_wars.id;


--
-- Name: realms; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.realms (
    id integer NOT NULL,
    round_id integer NOT NULL,
    monarch_dominion_id integer,
    alignment character varying(255) NOT NULL,
    number integer NOT NULL,
    name character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    motd text,
    motd_updated_at timestamp(0) without time zone,
    discord_role_id character varying(191),
    rating integer DEFAULT 0 NOT NULL,
    discord_category_id character varying(191),
    general_dominion_id integer,
    magister_dominion_id integer,
    mage_dominion_id integer,
    jester_dominion_id integer,
    spymaster_dominion_id integer,
    settings text,
    valor integer DEFAULT 0 NOT NULL,
    CONSTRAINT realms_alignment_check CHECK (((alignment)::text = ANY ((ARRAY['good'::character varying, 'neutral'::character varying, 'evil'::character varying])::text[])))
);


--
-- Name: realms_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.realms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: realms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.realms_id_seq OWNED BY public.realms.id;


--
-- Name: role_has_permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.role_has_permissions (
    permission_id integer NOT NULL,
    role_id integer NOT NULL
);


--
-- Name: roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    name character varying(191) NOT NULL,
    guard_name character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: round_leagues; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.round_leagues (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    description character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: round_leagues_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.round_leagues_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: round_leagues_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.round_leagues_id_seq OWNED BY public.round_leagues.id;


--
-- Name: round_wonder_damage; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.round_wonder_damage (
    id integer NOT NULL,
    round_wonder_id integer NOT NULL,
    realm_id integer NOT NULL,
    dominion_id integer NOT NULL,
    damage integer DEFAULT 0 NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    source character varying(191)
);


--
-- Name: round_wonder_damage_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.round_wonder_damage_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: round_wonder_damage_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.round_wonder_damage_id_seq OWNED BY public.round_wonder_damage.id;


--
-- Name: round_wonders; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.round_wonders (
    id integer NOT NULL,
    round_id integer NOT NULL,
    realm_id integer,
    wonder_id integer NOT NULL,
    power integer DEFAULT 0 NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: round_wonders_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.round_wonders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: round_wonders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.round_wonders_id_seq OWNED BY public.round_wonders.id;


--
-- Name: rounds; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rounds (
    id integer NOT NULL,
    round_league_id integer NOT NULL,
    number integer NOT NULL,
    name character varying(191) NOT NULL,
    start_date timestamp(0) without time zone NOT NULL,
    end_date timestamp(0) without time zone NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    realm_size integer DEFAULT 12 NOT NULL,
    pack_size integer DEFAULT 0 NOT NULL,
    players_per_race integer DEFAULT 0 NOT NULL,
    mixed_alignment boolean DEFAULT false NOT NULL,
    offensive_actions_prohibited_at timestamp(0) without time zone,
    discord_guild_id character varying(191),
    tech_version integer DEFAULT 1 NOT NULL,
    largest_hit integer DEFAULT 0 NOT NULL,
    assignment_complete boolean DEFAULT false NOT NULL
);


--
-- Name: rounds_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.rounds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: rounds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.rounds_id_seq OWNED BY public.rounds.id;


--
-- Name: spell_perk_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.spell_perk_types (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: spell_perk_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.spell_perk_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: spell_perk_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.spell_perk_types_id_seq OWNED BY public.spell_perk_types.id;


--
-- Name: spell_perks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.spell_perks (
    id integer NOT NULL,
    spell_id integer NOT NULL,
    spell_perk_type_id integer NOT NULL,
    value character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: spell_perks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.spell_perks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: spell_perks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.spell_perks_id_seq OWNED BY public.spell_perks.id;


--
-- Name: spells; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.spells (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    name character varying(191) NOT NULL,
    category character varying(191) NOT NULL,
    cost_mana double precision NOT NULL,
    cost_strength double precision NOT NULL,
    duration integer DEFAULT 0 NOT NULL,
    cooldown integer DEFAULT 0 NOT NULL,
    races text,
    active boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: spells_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.spells_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: spells_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.spells_id_seq OWNED BY public.spells.id;


--
-- Name: tech_perk_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tech_perk_types (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: tech_perk_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.tech_perk_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tech_perk_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.tech_perk_types_id_seq OWNED BY public.tech_perk_types.id;


--
-- Name: tech_perks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tech_perks (
    id integer NOT NULL,
    tech_id integer NOT NULL,
    tech_perk_type_id integer NOT NULL,
    value character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: tech_perks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.tech_perks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tech_perks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.tech_perks_id_seq OWNED BY public.tech_perks.id;


--
-- Name: techs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.techs (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    name character varying(191) NOT NULL,
    prerequisites text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    active boolean DEFAULT true NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    x integer DEFAULT 0 NOT NULL,
    y integer DEFAULT 0 NOT NULL
);


--
-- Name: techs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.techs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: techs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.techs_id_seq OWNED BY public.techs.id;


--
-- Name: telescope_entries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.telescope_entries (
    sequence bigint NOT NULL,
    uuid uuid NOT NULL,
    batch_id uuid NOT NULL,
    family_hash character varying(191),
    should_display_on_index boolean DEFAULT true NOT NULL,
    type character varying(20) NOT NULL,
    content text NOT NULL,
    created_at timestamp(0) without time zone
);


--
-- Name: telescope_entries_sequence_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.telescope_entries_sequence_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: telescope_entries_sequence_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.telescope_entries_sequence_seq OWNED BY public.telescope_entries.sequence;


--
-- Name: telescope_entries_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.telescope_entries_tags (
    entry_uuid uuid NOT NULL,
    tag character varying(191) NOT NULL
);


--
-- Name: telescope_monitoring; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.telescope_monitoring (
    tag character varying(191) NOT NULL
);


--
-- Name: unit_perk_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.unit_perk_types (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: unit_perk_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.unit_perk_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: unit_perk_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.unit_perk_types_id_seq OWNED BY public.unit_perk_types.id;


--
-- Name: unit_perks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.unit_perks (
    id integer NOT NULL,
    unit_id integer NOT NULL,
    unit_perk_type_id integer NOT NULL,
    value character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: unit_perks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.unit_perks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: unit_perks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.unit_perks_id_seq OWNED BY public.unit_perks.id;


--
-- Name: units; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.units (
    id integer NOT NULL,
    race_id integer NOT NULL,
    slot character varying(255) NOT NULL,
    name character varying(191) NOT NULL,
    cost_platinum integer NOT NULL,
    cost_ore integer NOT NULL,
    power_offense double precision NOT NULL,
    power_defense double precision NOT NULL,
    need_boat boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    type text,
    cost_mana integer NOT NULL,
    cost_lumber integer NOT NULL,
    cost_gems integer NOT NULL,
    CONSTRAINT units_slot_check CHECK (((slot)::text = ANY ((ARRAY['1'::character varying, '2'::character varying, '3'::character varying, '4'::character varying])::text[])))
);


--
-- Name: units_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.units_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: units_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.units_id_seq OWNED BY public.units.id;


--
-- Name: user_achievements; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_achievements (
    id integer NOT NULL,
    user_id integer NOT NULL,
    achievement_id integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: user_achievements_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_achievements_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_achievements_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_achievements_id_seq OWNED BY public.user_achievements.id;


--
-- Name: user_activities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_activities (
    id integer NOT NULL,
    user_id integer NOT NULL,
    ip character varying(191) NOT NULL,
    key character varying(191) NOT NULL,
    context text,
    created_at timestamp(0) without time zone,
    status character varying(191),
    device character varying(191)
);


--
-- Name: user_activities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_activities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_activities_id_seq OWNED BY public.user_activities.id;


--
-- Name: user_discord_users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_discord_users (
    id integer NOT NULL,
    user_id integer NOT NULL,
    discord_user_id bigint NOT NULL,
    username character varying(191) NOT NULL,
    discriminator integer NOT NULL,
    email character varying(191),
    refresh_token character varying(191) NOT NULL,
    expires_at timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: user_discord_users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_discord_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_discord_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_discord_users_id_seq OWNED BY public.user_discord_users.id;


--
-- Name: user_feedback; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_feedback (
    id integer NOT NULL,
    source_id integer NOT NULL,
    target_id integer NOT NULL,
    endorsed boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: user_feedback_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_feedback_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_feedback_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_feedback_id_seq OWNED BY public.user_feedback.id;


--
-- Name: user_identities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_identities (
    id integer NOT NULL,
    user_id integer NOT NULL,
    fingerprint character varying(191),
    user_agent character varying(191),
    count integer DEFAULT 1 NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: user_identities_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_identities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_identities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_identities_id_seq OWNED BY public.user_identities.id;


--
-- Name: user_origin_lookups; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_origin_lookups (
    id integer NOT NULL,
    ip_address character varying(191) NOT NULL,
    isp character varying(191),
    organization character varying(191),
    country character varying(191),
    region character varying(191),
    city character varying(191),
    vpn boolean,
    score double precision,
    data text,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: user_origin_lookups_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_origin_lookups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_origin_lookups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_origin_lookups_id_seq OWNED BY public.user_origin_lookups.id;


--
-- Name: user_origins; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_origins (
    id integer NOT NULL,
    user_id integer NOT NULL,
    dominion_id integer,
    ip_address character varying(191) NOT NULL,
    count integer DEFAULT 1 NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: user_origins_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_origins_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_origins_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_origins_id_seq OWNED BY public.user_origins.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(191) NOT NULL,
    password character varying(191) NOT NULL,
    display_name character varying(191) NOT NULL,
    remember_token character varying(100),
    activated boolean DEFAULT false NOT NULL,
    activation_code character varying(191) NOT NULL,
    last_online timestamp(0) without time zone,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    avatar character varying(191),
    settings text,
    skin text,
    message_board_last_read timestamp(0) without time zone,
    rating integer DEFAULT 0 NOT NULL
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: valor; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.valor (
    id integer NOT NULL,
    round_id integer NOT NULL,
    realm_id integer NOT NULL,
    dominion_id integer NOT NULL,
    source character varying(191) NOT NULL,
    amount double precision NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: valor_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.valor_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: valor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.valor_id_seq OWNED BY public.valor.id;


--
-- Name: wonder_perk_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.wonder_perk_types (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: wonder_perk_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.wonder_perk_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: wonder_perk_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.wonder_perk_types_id_seq OWNED BY public.wonder_perk_types.id;


--
-- Name: wonder_perks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.wonder_perks (
    id integer NOT NULL,
    wonder_id integer NOT NULL,
    wonder_perk_type_id integer NOT NULL,
    value character varying(191),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: wonder_perks_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.wonder_perks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: wonder_perks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.wonder_perks_id_seq OWNED BY public.wonder_perks.id;


--
-- Name: wonders; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.wonders (
    id integer NOT NULL,
    key character varying(191) NOT NULL,
    name character varying(191) NOT NULL,
    power integer NOT NULL,
    active boolean DEFAULT true NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


--
-- Name: wonders_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.wonders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: wonders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.wonders_id_seq OWNED BY public.wonders.id;


--
-- Name: achievements id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.achievements ALTER COLUMN id SET DEFAULT nextval('public.achievements_id_seq'::regclass);


--
-- Name: bounties id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bounties ALTER COLUMN id SET DEFAULT nextval('public.bounties_id_seq'::regclass);


--
-- Name: council_posts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_posts ALTER COLUMN id SET DEFAULT nextval('public.council_posts_id_seq'::regclass);


--
-- Name: council_threads id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_threads ALTER COLUMN id SET DEFAULT nextval('public.council_threads_id_seq'::regclass);


--
-- Name: daily_rankings id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.daily_rankings ALTER COLUMN id SET DEFAULT nextval('public.daily_rankings_id_seq'::regclass);


--
-- Name: dominion_history id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_history ALTER COLUMN id SET DEFAULT nextval('public.dominion_history_id_seq'::regclass);


--
-- Name: dominion_journals id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_journals ALTER COLUMN id SET DEFAULT nextval('public.dominion_journals_id_seq'::regclass);


--
-- Name: dominion_techs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_techs ALTER COLUMN id SET DEFAULT nextval('public.dominion_techs_id_seq'::regclass);


--
-- Name: dominion_tick id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_tick ALTER COLUMN id SET DEFAULT nextval('public.dominion_tick_id_seq'::regclass);


--
-- Name: dominions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions ALTER COLUMN id SET DEFAULT nextval('public.dominions_id_seq'::regclass);


--
-- Name: failed_jobs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.failed_jobs ALTER COLUMN id SET DEFAULT nextval('public.failed_jobs_id_seq'::regclass);


--
-- Name: forum_posts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_posts ALTER COLUMN id SET DEFAULT nextval('public.forum_posts_id_seq'::regclass);


--
-- Name: forum_threads id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_threads ALTER COLUMN id SET DEFAULT nextval('public.forum_threads_id_seq'::regclass);


--
-- Name: hero_battle_actions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_actions ALTER COLUMN id SET DEFAULT nextval('public.hero_battle_actions_id_seq'::regclass);


--
-- Name: hero_battle_queue id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_queue ALTER COLUMN id SET DEFAULT nextval('public.hero_battle_queue_id_seq'::regclass);


--
-- Name: hero_battles id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battles ALTER COLUMN id SET DEFAULT nextval('public.hero_battles_id_seq'::regclass);


--
-- Name: hero_combatants id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_combatants ALTER COLUMN id SET DEFAULT nextval('public.hero_combatants_id_seq'::regclass);


--
-- Name: hero_hero_upgrades id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_hero_upgrades ALTER COLUMN id SET DEFAULT nextval('public.hero_hero_upgrades_id_seq'::regclass);


--
-- Name: hero_tournament_battles id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_battles ALTER COLUMN id SET DEFAULT nextval('public.hero_tournament_battles_id_seq'::regclass);


--
-- Name: hero_tournament_participants id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_participants ALTER COLUMN id SET DEFAULT nextval('public.hero_tournament_participants_id_seq'::regclass);


--
-- Name: hero_tournaments id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournaments ALTER COLUMN id SET DEFAULT nextval('public.hero_tournaments_id_seq'::regclass);


--
-- Name: hero_upgrade_perks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_upgrade_perks ALTER COLUMN id SET DEFAULT nextval('public.hero_upgrade_perks_id_seq'::regclass);


--
-- Name: hero_upgrades id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_upgrades ALTER COLUMN id SET DEFAULT nextval('public.hero_upgrades_id_seq'::regclass);


--
-- Name: heroes id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.heroes ALTER COLUMN id SET DEFAULT nextval('public.heroes_id_seq'::regclass);


--
-- Name: info_ops id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.info_ops ALTER COLUMN id SET DEFAULT nextval('public.info_ops_id_seq'::regclass);


--
-- Name: jobs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.jobs ALTER COLUMN id SET DEFAULT nextval('public.jobs_id_seq'::regclass);


--
-- Name: message_board_categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_categories ALTER COLUMN id SET DEFAULT nextval('public.message_board_categories_id_seq'::regclass);


--
-- Name: message_board_posts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_posts ALTER COLUMN id SET DEFAULT nextval('public.message_board_posts_id_seq'::regclass);


--
-- Name: message_board_threads id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_threads ALTER COLUMN id SET DEFAULT nextval('public.message_board_threads_id_seq'::regclass);


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- Name: packs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packs ALTER COLUMN id SET DEFAULT nextval('public.packs_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: race_perk_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.race_perk_types ALTER COLUMN id SET DEFAULT nextval('public.race_perk_types_id_seq'::regclass);


--
-- Name: race_perks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.race_perks ALTER COLUMN id SET DEFAULT nextval('public.race_perks_id_seq'::regclass);


--
-- Name: races id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.races ALTER COLUMN id SET DEFAULT nextval('public.races_id_seq'::regclass);


--
-- Name: realm_history id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_history ALTER COLUMN id SET DEFAULT nextval('public.realm_history_id_seq'::regclass);


--
-- Name: realm_wars id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_wars ALTER COLUMN id SET DEFAULT nextval('public.realm_wars_id_seq'::regclass);


--
-- Name: realms id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms ALTER COLUMN id SET DEFAULT nextval('public.realms_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: round_leagues id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_leagues ALTER COLUMN id SET DEFAULT nextval('public.round_leagues_id_seq'::regclass);


--
-- Name: round_wonder_damage id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonder_damage ALTER COLUMN id SET DEFAULT nextval('public.round_wonder_damage_id_seq'::regclass);


--
-- Name: round_wonders id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonders ALTER COLUMN id SET DEFAULT nextval('public.round_wonders_id_seq'::regclass);


--
-- Name: rounds id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rounds ALTER COLUMN id SET DEFAULT nextval('public.rounds_id_seq'::regclass);


--
-- Name: spell_perk_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spell_perk_types ALTER COLUMN id SET DEFAULT nextval('public.spell_perk_types_id_seq'::regclass);


--
-- Name: spell_perks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spell_perks ALTER COLUMN id SET DEFAULT nextval('public.spell_perks_id_seq'::regclass);


--
-- Name: spells id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spells ALTER COLUMN id SET DEFAULT nextval('public.spells_id_seq'::regclass);


--
-- Name: tech_perk_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tech_perk_types ALTER COLUMN id SET DEFAULT nextval('public.tech_perk_types_id_seq'::regclass);


--
-- Name: tech_perks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tech_perks ALTER COLUMN id SET DEFAULT nextval('public.tech_perks_id_seq'::regclass);


--
-- Name: techs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.techs ALTER COLUMN id SET DEFAULT nextval('public.techs_id_seq'::regclass);


--
-- Name: telescope_entries sequence; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.telescope_entries ALTER COLUMN sequence SET DEFAULT nextval('public.telescope_entries_sequence_seq'::regclass);


--
-- Name: unit_perk_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.unit_perk_types ALTER COLUMN id SET DEFAULT nextval('public.unit_perk_types_id_seq'::regclass);


--
-- Name: unit_perks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.unit_perks ALTER COLUMN id SET DEFAULT nextval('public.unit_perks_id_seq'::regclass);


--
-- Name: units id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.units ALTER COLUMN id SET DEFAULT nextval('public.units_id_seq'::regclass);


--
-- Name: user_achievements id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements ALTER COLUMN id SET DEFAULT nextval('public.user_achievements_id_seq'::regclass);


--
-- Name: user_activities id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_activities ALTER COLUMN id SET DEFAULT nextval('public.user_activities_id_seq'::regclass);


--
-- Name: user_discord_users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_discord_users ALTER COLUMN id SET DEFAULT nextval('public.user_discord_users_id_seq'::regclass);


--
-- Name: user_feedback id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_feedback ALTER COLUMN id SET DEFAULT nextval('public.user_feedback_id_seq'::regclass);


--
-- Name: user_identities id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_identities ALTER COLUMN id SET DEFAULT nextval('public.user_identities_id_seq'::regclass);


--
-- Name: user_origin_lookups id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origin_lookups ALTER COLUMN id SET DEFAULT nextval('public.user_origin_lookups_id_seq'::regclass);


--
-- Name: user_origins id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origins ALTER COLUMN id SET DEFAULT nextval('public.user_origins_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: valor id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.valor ALTER COLUMN id SET DEFAULT nextval('public.valor_id_seq'::regclass);


--
-- Name: wonder_perk_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonder_perk_types ALTER COLUMN id SET DEFAULT nextval('public.wonder_perk_types_id_seq'::regclass);


--
-- Name: wonder_perks id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonder_perks ALTER COLUMN id SET DEFAULT nextval('public.wonder_perks_id_seq'::regclass);


--
-- Name: wonders id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonders ALTER COLUMN id SET DEFAULT nextval('public.wonders_id_seq'::regclass);


--
-- Name: achievements achievements_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.achievements
    ADD CONSTRAINT achievements_pkey PRIMARY KEY (id);


--
-- Name: bounties bounties_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bounties
    ADD CONSTRAINT bounties_pkey PRIMARY KEY (id);


--
-- Name: council_posts council_posts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_posts
    ADD CONSTRAINT council_posts_pkey PRIMARY KEY (id);


--
-- Name: council_threads council_threads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_threads
    ADD CONSTRAINT council_threads_pkey PRIMARY KEY (id);


--
-- Name: daily_rankings daily_rankings_dominion_id_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.daily_rankings
    ADD CONSTRAINT daily_rankings_dominion_id_key_unique UNIQUE (dominion_id, key);


--
-- Name: daily_rankings daily_rankings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.daily_rankings
    ADD CONSTRAINT daily_rankings_pkey PRIMARY KEY (id);


--
-- Name: dominion_history dominion_history_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_history
    ADD CONSTRAINT dominion_history_pkey PRIMARY KEY (id);


--
-- Name: dominion_journals dominion_journals_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_journals
    ADD CONSTRAINT dominion_journals_pkey PRIMARY KEY (id);


--
-- Name: dominion_queue dominion_queue_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_queue
    ADD CONSTRAINT dominion_queue_pkey PRIMARY KEY (dominion_id, source, resource, hours);


--
-- Name: dominion_spells dominion_spells_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_spells
    ADD CONSTRAINT dominion_spells_pkey PRIMARY KEY (dominion_id, spell_id);


--
-- Name: dominion_techs dominion_techs_dominion_id_tech_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_techs
    ADD CONSTRAINT dominion_techs_dominion_id_tech_id_unique UNIQUE (dominion_id, tech_id);


--
-- Name: dominion_techs dominion_techs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_techs
    ADD CONSTRAINT dominion_techs_pkey PRIMARY KEY (id);


--
-- Name: dominion_tick dominion_tick_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_tick
    ADD CONSTRAINT dominion_tick_pkey PRIMARY KEY (id);


--
-- Name: dominions dominions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_pkey PRIMARY KEY (id);


--
-- Name: dominions dominions_round_id_name_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_round_id_name_unique UNIQUE (round_id, name);


--
-- Name: dominions dominions_round_id_realm_id_ruler_name_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_round_id_realm_id_ruler_name_unique UNIQUE (round_id, realm_id, ruler_name);


--
-- Name: dominions dominions_user_id_round_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_user_id_round_id_unique UNIQUE (user_id, round_id);


--
-- Name: failed_jobs failed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.failed_jobs
    ADD CONSTRAINT failed_jobs_pkey PRIMARY KEY (id);


--
-- Name: forum_posts forum_posts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_posts
    ADD CONSTRAINT forum_posts_pkey PRIMARY KEY (id);


--
-- Name: forum_threads forum_threads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_threads
    ADD CONSTRAINT forum_threads_pkey PRIMARY KEY (id);


--
-- Name: game_events game_events_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.game_events
    ADD CONSTRAINT game_events_pkey PRIMARY KEY (id);


--
-- Name: hero_battle_actions hero_battle_actions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_actions
    ADD CONSTRAINT hero_battle_actions_pkey PRIMARY KEY (id);


--
-- Name: hero_battle_queue hero_battle_queue_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_queue
    ADD CONSTRAINT hero_battle_queue_pkey PRIMARY KEY (id);


--
-- Name: hero_battles hero_battles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battles
    ADD CONSTRAINT hero_battles_pkey PRIMARY KEY (id);


--
-- Name: hero_combatants hero_combatants_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_combatants
    ADD CONSTRAINT hero_combatants_pkey PRIMARY KEY (id);


--
-- Name: hero_hero_upgrades hero_hero_upgrades_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_hero_upgrades
    ADD CONSTRAINT hero_hero_upgrades_pkey PRIMARY KEY (id);


--
-- Name: hero_tournament_battles hero_tournament_battles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_battles
    ADD CONSTRAINT hero_tournament_battles_pkey PRIMARY KEY (id);


--
-- Name: hero_tournament_participants hero_tournament_participants_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_participants
    ADD CONSTRAINT hero_tournament_participants_pkey PRIMARY KEY (id);


--
-- Name: hero_tournaments hero_tournaments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournaments
    ADD CONSTRAINT hero_tournaments_pkey PRIMARY KEY (id);


--
-- Name: hero_upgrade_perks hero_upgrade_perks_hero_upgrade_id_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_upgrade_perks
    ADD CONSTRAINT hero_upgrade_perks_hero_upgrade_id_key_unique UNIQUE (hero_upgrade_id, key);


--
-- Name: hero_upgrade_perks hero_upgrade_perks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_upgrade_perks
    ADD CONSTRAINT hero_upgrade_perks_pkey PRIMARY KEY (id);


--
-- Name: hero_upgrades hero_upgrades_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_upgrades
    ADD CONSTRAINT hero_upgrades_pkey PRIMARY KEY (id);


--
-- Name: heroes heroes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.heroes
    ADD CONSTRAINT heroes_pkey PRIMARY KEY (id);


--
-- Name: info_ops info_ops_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.info_ops
    ADD CONSTRAINT info_ops_pkey PRIMARY KEY (id);


--
-- Name: jobs jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.jobs
    ADD CONSTRAINT jobs_pkey PRIMARY KEY (id);


--
-- Name: message_board_categories message_board_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_categories
    ADD CONSTRAINT message_board_categories_pkey PRIMARY KEY (id);


--
-- Name: message_board_posts message_board_posts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_posts
    ADD CONSTRAINT message_board_posts_pkey PRIMARY KEY (id);


--
-- Name: message_board_threads message_board_threads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_threads
    ADD CONSTRAINT message_board_threads_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- Name: model_has_permissions model_has_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.model_has_permissions
    ADD CONSTRAINT model_has_permissions_pkey PRIMARY KEY (permission_id, model_id, model_type);


--
-- Name: model_has_roles model_has_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.model_has_roles
    ADD CONSTRAINT model_has_roles_pkey PRIMARY KEY (role_id, model_id, model_type);


--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: packs packs_creator_dominion_id_round_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packs
    ADD CONSTRAINT packs_creator_dominion_id_round_id_unique UNIQUE (creator_dominion_id, round_id);


--
-- Name: packs packs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packs
    ADD CONSTRAINT packs_pkey PRIMARY KEY (id);


--
-- Name: packs packs_round_id_name_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packs
    ADD CONSTRAINT packs_round_id_name_unique UNIQUE (round_id, name);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: race_perk_types race_perk_types_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.race_perk_types
    ADD CONSTRAINT race_perk_types_key_unique UNIQUE (key);


--
-- Name: race_perk_types race_perk_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.race_perk_types
    ADD CONSTRAINT race_perk_types_pkey PRIMARY KEY (id);


--
-- Name: race_perks race_perks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.race_perks
    ADD CONSTRAINT race_perks_pkey PRIMARY KEY (id);


--
-- Name: races races_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.races
    ADD CONSTRAINT races_pkey PRIMARY KEY (id);


--
-- Name: realm_history realm_history_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_history
    ADD CONSTRAINT realm_history_pkey PRIMARY KEY (id);


--
-- Name: realm_wars realm_wars_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_wars
    ADD CONSTRAINT realm_wars_pkey PRIMARY KEY (id);


--
-- Name: realms realms_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_pkey PRIMARY KEY (id);


--
-- Name: role_has_permissions role_has_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_has_permissions
    ADD CONSTRAINT role_has_permissions_pkey PRIMARY KEY (permission_id, role_id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: round_leagues round_leagues_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_leagues
    ADD CONSTRAINT round_leagues_key_unique UNIQUE (key);


--
-- Name: round_leagues round_leagues_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_leagues
    ADD CONSTRAINT round_leagues_pkey PRIMARY KEY (id);


--
-- Name: round_wonder_damage round_wonder_damage_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonder_damage
    ADD CONSTRAINT round_wonder_damage_pkey PRIMARY KEY (id);


--
-- Name: round_wonders round_wonders_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonders
    ADD CONSTRAINT round_wonders_pkey PRIMARY KEY (id);


--
-- Name: round_wonders round_wonders_round_id_realm_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonders
    ADD CONSTRAINT round_wonders_round_id_realm_id_unique UNIQUE (round_id, realm_id);


--
-- Name: round_wonders round_wonders_round_id_wonder_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonders
    ADD CONSTRAINT round_wonders_round_id_wonder_id_unique UNIQUE (round_id, wonder_id);


--
-- Name: rounds rounds_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rounds
    ADD CONSTRAINT rounds_pkey PRIMARY KEY (id);


--
-- Name: spell_perk_types spell_perk_types_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spell_perk_types
    ADD CONSTRAINT spell_perk_types_key_unique UNIQUE (key);


--
-- Name: spell_perk_types spell_perk_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spell_perk_types
    ADD CONSTRAINT spell_perk_types_pkey PRIMARY KEY (id);


--
-- Name: spell_perks spell_perks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spell_perks
    ADD CONSTRAINT spell_perks_pkey PRIMARY KEY (id);


--
-- Name: spells spells_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spells
    ADD CONSTRAINT spells_pkey PRIMARY KEY (id);


--
-- Name: tech_perk_types tech_perk_types_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tech_perk_types
    ADD CONSTRAINT tech_perk_types_key_unique UNIQUE (key);


--
-- Name: tech_perk_types tech_perk_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tech_perk_types
    ADD CONSTRAINT tech_perk_types_pkey PRIMARY KEY (id);


--
-- Name: tech_perks tech_perks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tech_perks
    ADD CONSTRAINT tech_perks_pkey PRIMARY KEY (id);


--
-- Name: techs techs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.techs
    ADD CONSTRAINT techs_pkey PRIMARY KEY (id);


--
-- Name: telescope_entries telescope_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.telescope_entries
    ADD CONSTRAINT telescope_entries_pkey PRIMARY KEY (sequence);


--
-- Name: telescope_entries telescope_entries_uuid_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.telescope_entries
    ADD CONSTRAINT telescope_entries_uuid_unique UNIQUE (uuid);


--
-- Name: unit_perk_types unit_perk_types_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.unit_perk_types
    ADD CONSTRAINT unit_perk_types_key_unique UNIQUE (key);


--
-- Name: unit_perk_types unit_perk_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.unit_perk_types
    ADD CONSTRAINT unit_perk_types_pkey PRIMARY KEY (id);


--
-- Name: unit_perks unit_perks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.unit_perks
    ADD CONSTRAINT unit_perks_pkey PRIMARY KEY (id);


--
-- Name: units units_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_pkey PRIMARY KEY (id);


--
-- Name: units units_race_id_slot_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_race_id_slot_unique UNIQUE (race_id, slot);


--
-- Name: user_achievements user_achievements_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements
    ADD CONSTRAINT user_achievements_pkey PRIMARY KEY (id);


--
-- Name: user_achievements user_achievements_user_id_achievement_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements
    ADD CONSTRAINT user_achievements_user_id_achievement_id_unique UNIQUE (user_id, achievement_id);


--
-- Name: user_activities user_activities_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_activities
    ADD CONSTRAINT user_activities_pkey PRIMARY KEY (id);


--
-- Name: user_discord_users user_discord_users_discord_user_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_discord_users
    ADD CONSTRAINT user_discord_users_discord_user_id_unique UNIQUE (discord_user_id);


--
-- Name: user_discord_users user_discord_users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_discord_users
    ADD CONSTRAINT user_discord_users_pkey PRIMARY KEY (id);


--
-- Name: user_discord_users user_discord_users_user_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_discord_users
    ADD CONSTRAINT user_discord_users_user_id_unique UNIQUE (user_id);


--
-- Name: user_feedback user_feedback_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_feedback
    ADD CONSTRAINT user_feedback_pkey PRIMARY KEY (id);


--
-- Name: user_identities user_identities_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_identities
    ADD CONSTRAINT user_identities_pkey PRIMARY KEY (id);


--
-- Name: user_identities user_identities_user_id_fingerprint_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_identities
    ADD CONSTRAINT user_identities_user_id_fingerprint_unique UNIQUE (user_id, fingerprint);


--
-- Name: user_origin_lookups user_origin_lookups_ip_address_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origin_lookups
    ADD CONSTRAINT user_origin_lookups_ip_address_unique UNIQUE (ip_address);


--
-- Name: user_origin_lookups user_origin_lookups_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origin_lookups
    ADD CONSTRAINT user_origin_lookups_pkey PRIMARY KEY (id);


--
-- Name: user_origins user_origins_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origins
    ADD CONSTRAINT user_origins_pkey PRIMARY KEY (id);


--
-- Name: user_origins user_origins_user_id_dominion_id_ip_address_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origins
    ADD CONSTRAINT user_origins_user_id_dominion_id_ip_address_unique UNIQUE (user_id, dominion_id, ip_address);


--
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: valor valor_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.valor
    ADD CONSTRAINT valor_pkey PRIMARY KEY (id);


--
-- Name: wonder_perk_types wonder_perk_types_key_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonder_perk_types
    ADD CONSTRAINT wonder_perk_types_key_unique UNIQUE (key);


--
-- Name: wonder_perk_types wonder_perk_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonder_perk_types
    ADD CONSTRAINT wonder_perk_types_pkey PRIMARY KEY (id);


--
-- Name: wonder_perks wonder_perks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonder_perks
    ADD CONSTRAINT wonder_perks_pkey PRIMARY KEY (id);


--
-- Name: wonders wonders_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonders
    ADD CONSTRAINT wonders_pkey PRIMARY KEY (id);


--
-- Name: dominion_history_ip_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX dominion_history_ip_index ON public.dominion_history USING btree (ip);


--
-- Name: game_events_source_type_source_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX game_events_source_type_source_id_index ON public.game_events USING btree (source_type, source_id);


--
-- Name: game_events_target_type_target_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX game_events_target_type_target_id_index ON public.game_events USING btree (target_type, target_id);


--
-- Name: game_events_type_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX game_events_type_index ON public.game_events USING btree (type);


--
-- Name: jobs_queue_reserved_at_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX jobs_queue_reserved_at_index ON public.jobs USING btree (queue, reserved_at);


--
-- Name: model_has_permissions_model_type_model_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX model_has_permissions_model_type_model_id_index ON public.model_has_permissions USING btree (model_type, model_id);


--
-- Name: model_has_roles_model_type_model_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX model_has_roles_model_type_model_id_index ON public.model_has_roles USING btree (model_type, model_id);


--
-- Name: notifications_notifiable_type_notifiable_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX notifications_notifiable_type_notifiable_id_index ON public.notifications USING btree (notifiable_type, notifiable_id);


--
-- Name: password_resets_email_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX password_resets_email_index ON public.password_resets USING btree (email);


--
-- Name: password_resets_token_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX password_resets_token_index ON public.password_resets USING btree (token);


--
-- Name: telescope_entries_batch_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX telescope_entries_batch_id_index ON public.telescope_entries USING btree (batch_id);


--
-- Name: telescope_entries_family_hash_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX telescope_entries_family_hash_index ON public.telescope_entries USING btree (family_hash);


--
-- Name: telescope_entries_tags_entry_uuid_tag_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX telescope_entries_tags_entry_uuid_tag_index ON public.telescope_entries_tags USING btree (entry_uuid, tag);


--
-- Name: telescope_entries_tags_tag_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX telescope_entries_tags_tag_index ON public.telescope_entries_tags USING btree (tag);


--
-- Name: telescope_entries_type_should_display_on_index_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX telescope_entries_type_should_display_on_index_index ON public.telescope_entries USING btree (type, should_display_on_index);


--
-- Name: bounties bounties_collected_by_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bounties
    ADD CONSTRAINT bounties_collected_by_dominion_id_foreign FOREIGN KEY (collected_by_dominion_id) REFERENCES public.dominions(id);


--
-- Name: bounties bounties_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bounties
    ADD CONSTRAINT bounties_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: bounties bounties_source_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bounties
    ADD CONSTRAINT bounties_source_dominion_id_foreign FOREIGN KEY (source_dominion_id) REFERENCES public.dominions(id);


--
-- Name: bounties bounties_source_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bounties
    ADD CONSTRAINT bounties_source_realm_id_foreign FOREIGN KEY (source_realm_id) REFERENCES public.realms(id);


--
-- Name: bounties bounties_target_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.bounties
    ADD CONSTRAINT bounties_target_dominion_id_foreign FOREIGN KEY (target_dominion_id) REFERENCES public.dominions(id);


--
-- Name: council_posts council_posts_council_thread_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_posts
    ADD CONSTRAINT council_posts_council_thread_id_foreign FOREIGN KEY (council_thread_id) REFERENCES public.council_threads(id);


--
-- Name: council_posts council_posts_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_posts
    ADD CONSTRAINT council_posts_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: council_threads council_threads_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_threads
    ADD CONSTRAINT council_threads_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: council_threads council_threads_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.council_threads
    ADD CONSTRAINT council_threads_realm_id_foreign FOREIGN KEY (realm_id) REFERENCES public.realms(id);


--
-- Name: daily_rankings daily_rankings_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.daily_rankings
    ADD CONSTRAINT daily_rankings_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: daily_rankings daily_rankings_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.daily_rankings
    ADD CONSTRAINT daily_rankings_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: dominion_history dominion_history_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_history
    ADD CONSTRAINT dominion_history_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: dominion_journals dominion_journals_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_journals
    ADD CONSTRAINT dominion_journals_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: dominion_queue dominion_queue_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_queue
    ADD CONSTRAINT dominion_queue_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: dominion_spells dominion_spells_cast_by_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_spells
    ADD CONSTRAINT dominion_spells_cast_by_dominion_id_foreign FOREIGN KEY (cast_by_dominion_id) REFERENCES public.dominions(id);


--
-- Name: dominion_spells dominion_spells_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_spells
    ADD CONSTRAINT dominion_spells_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: dominion_spells dominion_spells_spell_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_spells
    ADD CONSTRAINT dominion_spells_spell_id_foreign FOREIGN KEY (spell_id) REFERENCES public.spells(id);


--
-- Name: dominion_techs dominion_techs_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_techs
    ADD CONSTRAINT dominion_techs_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: dominion_techs dominion_techs_tech_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_techs
    ADD CONSTRAINT dominion_techs_tech_id_foreign FOREIGN KEY (tech_id) REFERENCES public.techs(id);


--
-- Name: dominion_tick dominion_tick_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominion_tick
    ADD CONSTRAINT dominion_tick_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: dominions dominions_pack_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_pack_id_foreign FOREIGN KEY (pack_id) REFERENCES public.packs(id);


--
-- Name: dominions dominions_race_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_race_id_foreign FOREIGN KEY (race_id) REFERENCES public.races(id);


--
-- Name: dominions dominions_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_realm_id_foreign FOREIGN KEY (realm_id) REFERENCES public.realms(id);


--
-- Name: dominions dominions_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: dominions dominions_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.dominions
    ADD CONSTRAINT dominions_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: forum_posts forum_posts_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_posts
    ADD CONSTRAINT forum_posts_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: forum_posts forum_posts_forum_thread_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_posts
    ADD CONSTRAINT forum_posts_forum_thread_id_foreign FOREIGN KEY (forum_thread_id) REFERENCES public.forum_threads(id);


--
-- Name: forum_threads forum_threads_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_threads
    ADD CONSTRAINT forum_threads_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: forum_threads forum_threads_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.forum_threads
    ADD CONSTRAINT forum_threads_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: game_events game_events_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.game_events
    ADD CONSTRAINT game_events_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: hero_battle_actions hero_battle_actions_combatant_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_actions
    ADD CONSTRAINT hero_battle_actions_combatant_id_foreign FOREIGN KEY (combatant_id) REFERENCES public.hero_combatants(id);


--
-- Name: hero_battle_actions hero_battle_actions_hero_battle_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_actions
    ADD CONSTRAINT hero_battle_actions_hero_battle_id_foreign FOREIGN KEY (hero_battle_id) REFERENCES public.hero_battles(id);


--
-- Name: hero_battle_actions hero_battle_actions_target_combatant_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_actions
    ADD CONSTRAINT hero_battle_actions_target_combatant_id_foreign FOREIGN KEY (target_combatant_id) REFERENCES public.hero_combatants(id);


--
-- Name: hero_battle_queue hero_battle_queue_hero_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battle_queue
    ADD CONSTRAINT hero_battle_queue_hero_id_foreign FOREIGN KEY (hero_id) REFERENCES public.heroes(id);


--
-- Name: hero_battles hero_battles_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battles
    ADD CONSTRAINT hero_battles_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: hero_battles hero_battles_winner_combatant_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_battles
    ADD CONSTRAINT hero_battles_winner_combatant_id_foreign FOREIGN KEY (winner_combatant_id) REFERENCES public.hero_combatants(id);


--
-- Name: hero_combatants hero_combatants_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_combatants
    ADD CONSTRAINT hero_combatants_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: hero_combatants hero_combatants_hero_battle_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_combatants
    ADD CONSTRAINT hero_combatants_hero_battle_id_foreign FOREIGN KEY (hero_battle_id) REFERENCES public.hero_battles(id);


--
-- Name: hero_combatants hero_combatants_hero_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_combatants
    ADD CONSTRAINT hero_combatants_hero_id_foreign FOREIGN KEY (hero_id) REFERENCES public.heroes(id);


--
-- Name: hero_hero_upgrades hero_hero_upgrades_hero_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_hero_upgrades
    ADD CONSTRAINT hero_hero_upgrades_hero_id_foreign FOREIGN KEY (hero_id) REFERENCES public.heroes(id);


--
-- Name: hero_hero_upgrades hero_hero_upgrades_hero_upgrade_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_hero_upgrades
    ADD CONSTRAINT hero_hero_upgrades_hero_upgrade_id_foreign FOREIGN KEY (hero_upgrade_id) REFERENCES public.hero_upgrades(id);


--
-- Name: hero_tournament_battles hero_tournament_battles_hero_battle_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_battles
    ADD CONSTRAINT hero_tournament_battles_hero_battle_id_foreign FOREIGN KEY (hero_battle_id) REFERENCES public.hero_battles(id);


--
-- Name: hero_tournament_battles hero_tournament_battles_hero_tournament_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_battles
    ADD CONSTRAINT hero_tournament_battles_hero_tournament_id_foreign FOREIGN KEY (hero_tournament_id) REFERENCES public.hero_tournaments(id);


--
-- Name: hero_tournament_participants hero_tournament_participants_hero_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_participants
    ADD CONSTRAINT hero_tournament_participants_hero_id_foreign FOREIGN KEY (hero_id) REFERENCES public.heroes(id);


--
-- Name: hero_tournament_participants hero_tournament_participants_hero_tournament_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournament_participants
    ADD CONSTRAINT hero_tournament_participants_hero_tournament_id_foreign FOREIGN KEY (hero_tournament_id) REFERENCES public.hero_tournaments(id);


--
-- Name: hero_tournaments hero_tournaments_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournaments
    ADD CONSTRAINT hero_tournaments_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: hero_tournaments hero_tournaments_winner_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_tournaments
    ADD CONSTRAINT hero_tournaments_winner_dominion_id_foreign FOREIGN KEY (winner_dominion_id) REFERENCES public.dominions(id);


--
-- Name: hero_upgrade_perks hero_upgrade_perks_hero_upgrade_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hero_upgrade_perks
    ADD CONSTRAINT hero_upgrade_perks_hero_upgrade_id_foreign FOREIGN KEY (hero_upgrade_id) REFERENCES public.hero_upgrades(id);


--
-- Name: heroes heroes_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.heroes
    ADD CONSTRAINT heroes_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: info_ops info_ops_source_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.info_ops
    ADD CONSTRAINT info_ops_source_dominion_id_foreign FOREIGN KEY (source_dominion_id) REFERENCES public.dominions(id);


--
-- Name: info_ops info_ops_source_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.info_ops
    ADD CONSTRAINT info_ops_source_realm_id_foreign FOREIGN KEY (source_realm_id) REFERENCES public.realms(id);


--
-- Name: info_ops info_ops_target_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.info_ops
    ADD CONSTRAINT info_ops_target_dominion_id_foreign FOREIGN KEY (target_dominion_id) REFERENCES public.dominions(id);


--
-- Name: info_ops info_ops_target_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.info_ops
    ADD CONSTRAINT info_ops_target_realm_id_foreign FOREIGN KEY (target_realm_id) REFERENCES public.realms(id);


--
-- Name: message_board_posts message_board_posts_message_board_thread_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_posts
    ADD CONSTRAINT message_board_posts_message_board_thread_id_foreign FOREIGN KEY (message_board_thread_id) REFERENCES public.message_board_threads(id);


--
-- Name: message_board_posts message_board_posts_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_posts
    ADD CONSTRAINT message_board_posts_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: message_board_threads message_board_threads_message_board_category_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_threads
    ADD CONSTRAINT message_board_threads_message_board_category_id_foreign FOREIGN KEY (message_board_category_id) REFERENCES public.message_board_categories(id);


--
-- Name: message_board_threads message_board_threads_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message_board_threads
    ADD CONSTRAINT message_board_threads_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: model_has_permissions model_has_permissions_permission_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.model_has_permissions
    ADD CONSTRAINT model_has_permissions_permission_id_foreign FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: model_has_roles model_has_roles_role_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.model_has_roles
    ADD CONSTRAINT model_has_roles_role_id_foreign FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;


--
-- Name: packs packs_creator_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packs
    ADD CONSTRAINT packs_creator_dominion_id_foreign FOREIGN KEY (creator_dominion_id) REFERENCES public.dominions(id);


--
-- Name: packs packs_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packs
    ADD CONSTRAINT packs_realm_id_foreign FOREIGN KEY (realm_id) REFERENCES public.realms(id);


--
-- Name: packs packs_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.packs
    ADD CONSTRAINT packs_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: race_perks race_perks_race_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.race_perks
    ADD CONSTRAINT race_perks_race_id_foreign FOREIGN KEY (race_id) REFERENCES public.races(id);


--
-- Name: race_perks race_perks_race_perk_type_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.race_perks
    ADD CONSTRAINT race_perks_race_perk_type_id_foreign FOREIGN KEY (race_perk_type_id) REFERENCES public.race_perk_types(id);


--
-- Name: realm_history realm_history_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_history
    ADD CONSTRAINT realm_history_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: realm_history realm_history_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_history
    ADD CONSTRAINT realm_history_realm_id_foreign FOREIGN KEY (realm_id) REFERENCES public.realms(id);


--
-- Name: realm_wars realm_wars_source_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_wars
    ADD CONSTRAINT realm_wars_source_realm_id_foreign FOREIGN KEY (source_realm_id) REFERENCES public.realms(id);


--
-- Name: realm_wars realm_wars_target_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realm_wars
    ADD CONSTRAINT realm_wars_target_realm_id_foreign FOREIGN KEY (target_realm_id) REFERENCES public.realms(id);


--
-- Name: realms realms_general_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_general_dominion_id_foreign FOREIGN KEY (general_dominion_id) REFERENCES public.dominions(id);


--
-- Name: realms realms_jester_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_jester_dominion_id_foreign FOREIGN KEY (jester_dominion_id) REFERENCES public.dominions(id);


--
-- Name: realms realms_mage_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_mage_dominion_id_foreign FOREIGN KEY (mage_dominion_id) REFERENCES public.dominions(id);


--
-- Name: realms realms_magister_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_magister_dominion_id_foreign FOREIGN KEY (magister_dominion_id) REFERENCES public.dominions(id);


--
-- Name: realms realms_monarch_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_monarch_dominion_id_foreign FOREIGN KEY (monarch_dominion_id) REFERENCES public.dominions(id);


--
-- Name: realms realms_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: realms realms_spymaster_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.realms
    ADD CONSTRAINT realms_spymaster_dominion_id_foreign FOREIGN KEY (spymaster_dominion_id) REFERENCES public.dominions(id);


--
-- Name: role_has_permissions role_has_permissions_permission_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_has_permissions
    ADD CONSTRAINT role_has_permissions_permission_id_foreign FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: role_has_permissions role_has_permissions_role_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.role_has_permissions
    ADD CONSTRAINT role_has_permissions_role_id_foreign FOREIGN KEY (role_id) REFERENCES public.roles(id) ON DELETE CASCADE;


--
-- Name: round_wonder_damage round_wonder_damage_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonder_damage
    ADD CONSTRAINT round_wonder_damage_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: round_wonder_damage round_wonder_damage_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonder_damage
    ADD CONSTRAINT round_wonder_damage_realm_id_foreign FOREIGN KEY (realm_id) REFERENCES public.realms(id);


--
-- Name: round_wonder_damage round_wonder_damage_round_wonder_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonder_damage
    ADD CONSTRAINT round_wonder_damage_round_wonder_id_foreign FOREIGN KEY (round_wonder_id) REFERENCES public.round_wonders(id);


--
-- Name: round_wonders round_wonders_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonders
    ADD CONSTRAINT round_wonders_realm_id_foreign FOREIGN KEY (realm_id) REFERENCES public.realms(id);


--
-- Name: round_wonders round_wonders_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonders
    ADD CONSTRAINT round_wonders_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: round_wonders round_wonders_wonder_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.round_wonders
    ADD CONSTRAINT round_wonders_wonder_id_foreign FOREIGN KEY (wonder_id) REFERENCES public.wonders(id);


--
-- Name: rounds rounds_round_league_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rounds
    ADD CONSTRAINT rounds_round_league_id_foreign FOREIGN KEY (round_league_id) REFERENCES public.round_leagues(id);


--
-- Name: spell_perks spell_perks_spell_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spell_perks
    ADD CONSTRAINT spell_perks_spell_id_foreign FOREIGN KEY (spell_id) REFERENCES public.spells(id);


--
-- Name: spell_perks spell_perks_spell_perk_type_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.spell_perks
    ADD CONSTRAINT spell_perks_spell_perk_type_id_foreign FOREIGN KEY (spell_perk_type_id) REFERENCES public.spell_perk_types(id);


--
-- Name: tech_perks tech_perks_tech_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tech_perks
    ADD CONSTRAINT tech_perks_tech_id_foreign FOREIGN KEY (tech_id) REFERENCES public.techs(id);


--
-- Name: tech_perks tech_perks_tech_perk_type_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tech_perks
    ADD CONSTRAINT tech_perks_tech_perk_type_id_foreign FOREIGN KEY (tech_perk_type_id) REFERENCES public.tech_perk_types(id);


--
-- Name: telescope_entries_tags telescope_entries_tags_entry_uuid_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.telescope_entries_tags
    ADD CONSTRAINT telescope_entries_tags_entry_uuid_foreign FOREIGN KEY (entry_uuid) REFERENCES public.telescope_entries(uuid) ON DELETE CASCADE;


--
-- Name: unit_perks unit_perks_unit_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.unit_perks
    ADD CONSTRAINT unit_perks_unit_id_foreign FOREIGN KEY (unit_id) REFERENCES public.units(id);


--
-- Name: unit_perks unit_perks_unit_perk_type_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.unit_perks
    ADD CONSTRAINT unit_perks_unit_perk_type_id_foreign FOREIGN KEY (unit_perk_type_id) REFERENCES public.unit_perk_types(id);


--
-- Name: units units_race_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_race_id_foreign FOREIGN KEY (race_id) REFERENCES public.races(id);


--
-- Name: user_achievements user_achievements_achievement_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements
    ADD CONSTRAINT user_achievements_achievement_id_foreign FOREIGN KEY (achievement_id) REFERENCES public.achievements(id);


--
-- Name: user_achievements user_achievements_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements
    ADD CONSTRAINT user_achievements_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: user_activities user_activities_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_activities
    ADD CONSTRAINT user_activities_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: user_discord_users user_discord_users_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_discord_users
    ADD CONSTRAINT user_discord_users_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: user_feedback user_feedback_source_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_feedback
    ADD CONSTRAINT user_feedback_source_id_foreign FOREIGN KEY (source_id) REFERENCES public.users(id);


--
-- Name: user_feedback user_feedback_target_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_feedback
    ADD CONSTRAINT user_feedback_target_id_foreign FOREIGN KEY (target_id) REFERENCES public.users(id);


--
-- Name: user_identities user_identities_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_identities
    ADD CONSTRAINT user_identities_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: user_origins user_origins_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origins
    ADD CONSTRAINT user_origins_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: user_origins user_origins_ip_address_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origins
    ADD CONSTRAINT user_origins_ip_address_foreign FOREIGN KEY (ip_address) REFERENCES public.user_origin_lookups(ip_address);


--
-- Name: user_origins user_origins_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_origins
    ADD CONSTRAINT user_origins_user_id_foreign FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: valor valor_dominion_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.valor
    ADD CONSTRAINT valor_dominion_id_foreign FOREIGN KEY (dominion_id) REFERENCES public.dominions(id);


--
-- Name: valor valor_realm_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.valor
    ADD CONSTRAINT valor_realm_id_foreign FOREIGN KEY (realm_id) REFERENCES public.realms(id);


--
-- Name: valor valor_round_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.valor
    ADD CONSTRAINT valor_round_id_foreign FOREIGN KEY (round_id) REFERENCES public.rounds(id);


--
-- Name: wonder_perks wonder_perks_wonder_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonder_perks
    ADD CONSTRAINT wonder_perks_wonder_id_foreign FOREIGN KEY (wonder_id) REFERENCES public.wonders(id);


--
-- Name: wonder_perks wonder_perks_wonder_perk_type_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.wonder_perks
    ADD CONSTRAINT wonder_perks_wonder_perk_type_id_foreign FOREIGN KEY (wonder_perk_type_id) REFERENCES public.wonder_perk_types(id);
