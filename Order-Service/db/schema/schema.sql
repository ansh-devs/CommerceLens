CREATE TABLE "users" (
  "id" varchar PRIMARY KEY NOT NULL,
  "email" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "password" varchar NOT NULL,
  "address" varchar NOT NULL,
  "created_at" timestamptx NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" varchar PRIMARY KEY NOT NULL,
  "product_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "total_cost" varchar NOT NULL,
  "status" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "address" varchar NOT NULL,
  "product_name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" varchar NOT NULL,
  "created_at" timestamptx NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" varchar PRIMARY KEY NOT NULL,
  "product_name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" varchar NOT NULL,
  "created_at" timestamptx NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "orders" ("id");

CREATE INDEX ON "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_name") REFERENCES "products" ("product_name");

ALTER TABLE "orders" ADD FOREIGN KEY ("description") REFERENCES "products" ("description");

ALTER TABLE "orders" ADD FOREIGN KEY ("price") REFERENCES "products" ("price");

ALTER TABLE "orders" ADD FOREIGN KEY ("fullname") REFERENCES "users" ("fullname");

ALTER TABLE "orders" ADD FOREIGN KEY ("address") REFERENCES "users" ("address");
