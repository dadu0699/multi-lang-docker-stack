from os import getenv

from fastapi import FastAPI
from redis import Redis

HOST = getenv("REDIS_HOST", "localhost")
PORT = getenv("REDIS_PORT", "6379")
USER = getenv("REDIS_USER", "")
PASSWORD = getenv("REDIS_PASSWORD", "")

app = FastAPI()

cache = Redis(
    host=HOST,
    port=int(PORT),
    username=USER if USER else None,
    password=PASSWORD if PASSWORD else None
)


@app.get("/greet")
async def read_root():
    count = cache.incr("cache_count")
    return {"message": f"Welcome, you are the visitor No. {count}"}
