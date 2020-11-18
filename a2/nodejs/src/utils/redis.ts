import redis, { RedisClient } from "redis";

const client = redis.createClient("http://localhost:6379");

export function getRedisClient(): RedisClient {
  return client;
}
