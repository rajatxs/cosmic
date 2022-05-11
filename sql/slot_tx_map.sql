
CREATE TABLE IF NOT EXISTS `slot_tx_map` (
   slot_seq INTEGER REFERENCES slot_headers(seq),
   tx_seq INTEGER DEFAULT 0, 
   tx_sig VARCHAR(66) NOT NULL
);
