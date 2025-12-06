from os import getenv

from redis.asyncio import Redis

HOST = getenv("REDIS_HOST", "localhost")
PORT = getenv("REDIS_PORT", 6379)
PASSWORD = getenv("REDIS_PASSWORD", None)

cache = Redis(host=HOST, port=int(PORT), password=PASSWORD)


async def incr_cache_count():
    return await cache.incr("cache_count")
