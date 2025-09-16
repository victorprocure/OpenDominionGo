ALTER TABLE race_perks
    DELETE CONSTRAINT foreign_ids_unique;

ALTER TABLE unit_perks
    DELETE CONSTRAINT unit_id_unit_perk_type_id_unique;

ALTER TABLE units
    DELETE CONSTRAINT race_id_slot_unique;