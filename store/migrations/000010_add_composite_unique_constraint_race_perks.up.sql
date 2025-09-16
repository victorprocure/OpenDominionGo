ALTER TABLE race_perks
    ADD CONSTRAINT foreign_ids_unique UNIQUE (race_id, race_perk_type_id);

ALTER TABLE unit_perks
    ADD CONSTRAINT unit_id_unit_perk_type_id_unique UNIQUE (unit_id, unit_perk_type_id);

ALTER TABLE units
    ADD CONSTRAINT race_id_slot_unique UNIQUE (race_id, slot);