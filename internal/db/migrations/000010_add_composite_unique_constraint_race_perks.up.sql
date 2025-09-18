ALTER TABLE race_perks
    ADD CONSTRAINT foreign_ids_unique UNIQUE (race_id, race_perk_type_id);

ALTER TABLE unit_perks
    ADD CONSTRAINT unit_id_unit_perk_type_id_unique UNIQUE (unit_id, unit_perk_type_id);

ALTER TABLE units
    ADD CONSTRAINT race_id_slot_unique UNIQUE (race_id, slot);

ALTER TABLE spell_perks
    ADD CONSTRAINT spell_id_spell_perk_type_id_unique UNIQUE (spell_id, spell_perk_type_id);

ALTER TABLE tech_perks
    ADD CONSTRAINT tech_id_tech_perk_type_id_unique UNIQUE (tech_id, tech_perk_type_id);

ALTER TABLE wonder_perks
    ADD CONSTRAINT wonder_id_wonder_perk_type_id_unique UNIQUE (wonder_id, wonder_perk_type_id);

ALTER TABLE hero_upgrade_perks
    DROP CONSTRAINT IF EXISTS hero_upgrade_perks_key_unique