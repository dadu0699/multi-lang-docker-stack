from os import getenv

from redis import Redis

HOST = getenv("REDIS_HOST", "localhost")
PORT = getenv("REDIS_PORT", 6379)

cache = Redis(host=HOST, port=int(PORT))


async def incr_cache_count():
    return cache.incr("cache_count")
