import { Pool, PoolClient } from "pg";

export const pool = new Pool({
  host: process.env.DB_HOST,
  port: Number(process.env.DB_PORT),
  database: process.env.DB_DATABASE,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  max: 20,
  idleTimeoutMillis: 30000,
  connectionTimeoutMillis: 2000,
});

interface DbWrapper {
  read: <T>(query: string, bindings?: Array<unknown>) => Promise<T[]>;
  readFirst: <T>(query: string, bindings?: Array<unknown>) => Promise<T | undefined>;
  execute: <T>(query: string, bindings?: Array<unknown>, client?: PoolClient) => Promise<T>;
}

export const db = {
  read: async <T>(query: string, bindings?: Array<unknown>) => {
    const client = await pool.connect();

    try {
      const result = await client.query(query, bindings);
      return result.rows as T[];
    } catch (error) {
      console.log(error);
      throw error;
    } finally {
      await client.release();
    }
  },

  readFirst: async <T>(query: string, bindings?: Array<unknown>) => {
    const client = await pool.connect();

    try {
      const result = await client.query(`${query} LIMIT 1`, bindings);
      return result.rows[0] as T | undefined;
    } catch (error) {
      console.log(error);
      throw error;
    } finally {
      await client.release();
    }
  },

  execute: async <T>(query: string, bindings?: Array<unknown>, client?: PoolClient) => {
    let singleQuery = false;
    if (!client) {
      client = await pool.connect();
      singleQuery = true;
    }

    try {
      const result = await client.query(query, bindings);
      return result.rows[0] as T;
    } catch (error) {
      console.log(error);
      throw error;
    } finally {
      if (singleQuery) await client.release();
    }
  },

  transaction: async <T>(callback: (db: DbWrapper, client: PoolClient) => Promise<T>) => {
    const client = await pool.connect();
    try {
      await client.query("BEGIN");
      const result = await callback(db, client);
      await client.query("COMMIT");
      return result;
    } catch (error) {
      await client.query("ROLLBACK");
      throw error;
    } finally {
      await client.release();
    }
  },
};

export const inQuery = <T extends string>(array: T[]) => {
  const joinedArray = array.map((str) => `'${str}'`).join(", ");
  return `( ${joinedArray} )`;
};
