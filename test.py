from fastapi import FastAPI, Request, Depends, HTTPException, Query
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, Session

app = FastAPI()

class Item(BaseModel):
    name: str
    description : str

@app.get("/")
def read_root():
    return {"message":"root"}

@app.get("/items")  # ✅ Add this GET route
def get_items():
    print("in items")
    return {"items": ["item1", "item2"]}

@app.post("/items")
def add_data(item: Item):
    return {
        "name": item.name,
        "description": item.description
    }

@app.get("/items/{item_id}")
def search_id(item_id: int): 
    return {"item_id": item_id}

@app.get("/items/query")
def query_items(q: str = Query(..., description="desc"), limit: int = 10):
    return {"query": q, "limit": limit}
