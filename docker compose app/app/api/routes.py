from fastapi import APIRouter, status

from app.services.cache_service import incr_cache_count

router = APIRouter()

@router.get("/favicon.ico",status_code=status.HTTP_204_NO_CONTENT)
async def favicon():
    return None

@router.get("/count")
async def cache_count():
    count = await incr_cache_count()
    return {"greet": f"Hello, the cache count is {count}"}
