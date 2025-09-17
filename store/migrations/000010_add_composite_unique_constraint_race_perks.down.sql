ALTER TABLE race_perks
    DELETE CONSTRAINT foreign_ids_unique;

ALTER TABLE unit_perks
    DELETE CONSTRAINT unit_id_unit_perk_type_id_unique;

ALTER TABLE units
    DELETE CONSTRAINT race_id_slot_unique;

ALTER TABLE spell_perks
    DELETE CONSTRAINT spell_id_spell_perk_type_id_unique

ALTER TABLE tech_perks
    DELETE CONSTRAINT tech_id_tech_perk_type_id_unique

ALTER TABLE wonder_perks
    DELETE CONSTRAINT wonder_id_wonder_perk_type_id_unique

ALTER TABLE hero_upgrade_perks
    ADD CONSTRAINT hero_upgrade_perks_key_unique UNIQUE (key);