CREATE TABLE "gid_wasm_exec" (
    one_row_id boolean DEFAULT true CHECK (one_row_id) PRIMARY KEY,
    height bigint NOT NULL,
    contract_address text NOT NULL,
    reg_name text,
    duration bigint NOT NULL,
    sender text,
    coins      COIN[]  NOT NULL
);

CREATE UNIQUE INDEX "wasm_exec_pkey" ON "wasm_exec"(one_row_id bool_ops);
