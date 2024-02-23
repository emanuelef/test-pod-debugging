import os
import asyncio
import logging
from aiohttp import web, ClientSession


async def handleRoot(request):
    logging.warning("Called root API")
    text = "Hello!"
    return web.Response(text=text)


async def handle(request):
    logging.warning("Called Hello API")
    name = request.match_info.get("name", "Anonymous")
    text = await callExternalApi()
    return web.Response(text=f"Hello, {name}! External API Response: {text}")


async def callExternalApi():
    api_host = os.getenv(
        "SECOND_API_HOST",
        "localhost",
    )
    url = f"http://{api_host}:8080"
    logging.warning(f"Calling {api_host}")
    async with ClientSession() as session:
        async with session.get(url) as response:
            response_text = await response.text()
            return response_text


async def health(request):
    return web.Response(text="OK")


async def init():
    logging.basicConfig(
        format="%(asctime)s %(message)s",
        datefmt="%m/%d/%Y %H:%M:%S",
        level=logging.WARNING,
    )
    app = web.Application()
    app.router.add_get("/", handleRoot)
    app.router.add_get("/hello/{name}", handle)
    app.router.add_get("/health", health)  # Register the health endpoint

    # Setup the aiohttp server
    runner = web.AppRunner(app)
    await runner.setup()

    # Create the aiohttp server and start it
    site = web.TCPSite(runner, "0.0.0.0", 8080)
    await site.start()

    logging.warning("Server started on http://0.0.0.0:8080")

    # Keep the application running indefinitely
    while True:
        await asyncio.sleep(60)


if __name__ == "__main__":
    os.environ["PYTHONUNBUFFERED"] = "1"
    asyncio.run(init())
