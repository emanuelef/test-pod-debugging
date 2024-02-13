import os
import asyncio
import logging
from aiohttp import web, ClientSession


async def handleRoot(request):
    text = "Hello!"
    return web.Response(text=text)


async def handle(request):
    name = request.match_info.get("name", "Anonymous")
    text = await callExternalApi()
    return web.Response(text=f"Hello, {name}! External API Response: {text}")


async def health(request):
    return web.Response(text="OK")


async def callExternalApi():
    api_host = os.getenv("API_HOST", "localhost")
    url = f"http://{api_host}:8080"
    async with ClientSession() as session:
        async with session.get(url) as response:
            response_text = await response.text()
            return response_text


async def init():
    app = web.Application()
    app.router.add_get("/", handleRoot)
    app.router.add_get("/{name}", handle)
    app.router.add_get("/health", health)

    # Setup the aiohttp server
    runner = web.AppRunner(app)
    await runner.setup()

    # Create the aiohttp server and start it
    site = web.TCPSite(runner, "0.0.0.0", 8080)
    await site.start()

    logging.warning("Server started on http://0.0.0.0:8080")


# Run the asyncio event loop
if __name__ == "__main__":
    os.environ["PYTHONUNBUFFERED"] = "1"
    loop = asyncio.get_event_loop()
    loop.run_until_complete(init())
    try:
        loop.run_forever()
    except KeyboardInterrupt:
        pass
    finally:
        loop.run_until_complete(loop.shutdown_asyncgens())
        loop.close()
