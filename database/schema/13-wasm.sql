CREATE TABLE "wasm-exec" (
    id integer DEFAULT nextval('"wasm-exec_id_seq"'::regclass) PRIMARY KEY,
    height bigint NOT NULL,
    contract_address text NOT NULL,
    params jsonb,
    sender text
);

CREATE UNIQUE INDEX "wasm-exec_pkey" ON "wasm-exec"(id int4_ops);
