from fastapi import FastAPI, Request
from pydantic import BaseModel

app = FastAPI()

class LogEntry(BaseModel):
    timestamp: str
    level: str
    msg: str

@app.post("/analyze")
async def analyze_log(log: LogEntry):
    # Simple anomaly rule: flag any 'error' level logs
    is_anomaly = log.level.lower() == "error"
    return {"anomaly": is_anomaly}
