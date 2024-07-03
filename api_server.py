from fastapi import FastAPI, Request
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from app.config import routes as config_routes
from app.alpha import routes as alpha_routes

app = FastAPI()

# Mount static files (CSS, JS) to be served
app.mount("/static", StaticFiles(directory="app/static"), name="static")

# Jinja2 templates for rendering HTML
templates = Jinja2Templates(directory="app")

# Include routes from config and alpha modules
app.include_router(config_routes.router, prefix="/config", tags=["config"])
app.include_router(alpha_routes.router, prefix="/alpha", tags=["alpha"])

# Example route to render HTML template
@app.get("/config/view", response_class=HTMLResponse)
async def view_config(request: Request):
    config_data = {"account_id": "12345", "target_position": 100.0, "target_offset": 10.0}
    return templates.TemplateResponse("config_view.html", {"request": request, "config": config_data})
