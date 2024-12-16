from fastapi import APIRouter

from app.services.cache_service import incr_cache_count

router = APIRouter()


@router.get("/count")
async def cache_count():
    count = await incr_cache_count()
    return {"greet": f"Hello, the cache count is {count}"}
