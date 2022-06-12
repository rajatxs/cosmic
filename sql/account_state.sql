
CREATE TABLE IF NOT EXISTS `account_state`(

   -- Account address
   addr BLOB(32) PRIMARY KEY,

   -- Account round
   rnd UNSIGNED INTEGER DEFAULT 0,

   -- Account balance
   bal UNSIGNED BIGINT DEFAULT 0
);
