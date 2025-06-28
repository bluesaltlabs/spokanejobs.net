from fastmcp import FastMCP
import httpx
from typing import List, Dict, Any

# Create a server instance
mcp = FastMCP(name="Bedrock MCP Server prototype")

JOBS_URL = "https://raw.githubusercontent.com/bluesaltlabs/spokanejobs_data/refs/heads/main/api/jobs.json"
COMPANIES_URL = "https://raw.githubusercontent.com/bluesaltlabs/spokanejobs_data/refs/heads/main/api/companies.json"

def fetch_json_sync(url: str) -> List[Dict[str, Any]]:
    """Synchronous function to fetch JSON data with error handling."""
    try:
        with httpx.Client(timeout=30.0) as client:
            response = client.get(url)
            response.raise_for_status()
            return response.json()
    except Exception as e:
        print(f"Error fetching data from {url}: {e}")
        return []

@mcp.tool
def get_jobs() -> List[Dict[str, Any]]:
    """Returns a list of all jobs from the remote API."""
    return fetch_json_sync(JOBS_URL)

@mcp.tool
def get_companies() -> List[Dict[str, Any]]:
    """Returns a list of all companies from the remote API."""
    return fetch_json_sync(COMPANIES_URL)

if __name__ == "__main__":
    mcp.run()

# todo:
# count_companies()
# count_jobs()
# get_company()
# get_job()
