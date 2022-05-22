
CREATE TABLE IF NOT EXISTS `block_tx_map` (
   block_id INTEGER REFERENCES block_headers(id),
   tx_seq INTEGER DEFAULT 0, 
   tx_sig VARCHAR(66) NOT NULL
);
