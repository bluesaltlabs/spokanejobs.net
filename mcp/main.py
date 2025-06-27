#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys
import json
import os
import logging

# Basic logging setup (optional but helpful for debugging)
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def initialize(params):
    """
    This function is called by LM Studio when the plugin starts.
    It MUST return a dictionary with a specific structure.
    """
    print("Plugin 'initialize' method called.")

    # This is the data structure LM Studio is expecting.
    response = {
        "protocolVersion": "1.0",
        "serverInfo": {
            "name": "Bedrock MCP Server Prototype",
            "version": "0.1.0"
        },
        "capabilities": {
            "get_cwd": True
        }
    }

    print(f"Returning initialization response: {response}")
    # Make sure you are returning this dictionary!
    return response

def send_response(response_dict):
    """Sends a JSON response object to stdout."""
    response_json = json.dumps(response_dict)
    logging.info(f"Sending response: {response_json}")
    print(response_json, flush=True) # flush=True is important!

def send_error(request_id, code, message):
    """Sends a JSON-RPC error response."""
    error_response = {
        "jsonrpc": "2.0",
        "id": request_id,
        "error": {
            "code": code,
            "message": message
        }
    }
    send_response(error_response)

# MCP Error Codes (subset)
ERROR_METHOD_NOT_FOUND = -32601
ERROR_INVALID_PARAMS = -32602
ERROR_INTERNAL = -32603

def handle_list_tools(request_id):
    """Handles the ListTools request."""
    tools = [
        {
            "name": "get_cwd",
            "description": "Gets the server's current working directory.",
            # No parameters needed for this tool
            "inputSchema": { "type": "object", "properties": {} }
        }
    ]
    response = {
        "jsonrpc": "2.0",
        "id": request_id,
        "result": {
            "tools": tools
        }
    }
    send_response(response)

def handle_call_tool(request_id, params):
    """Handles the CallTool request."""
    tool_name = params.get("name")

    if tool_name == "get_cwd":
        try:
            cwd = os.getcwd()
            response = {
                "jsonrpc": "2.0",
                "id": request_id,
                "result": {
                    # Content should be an array of ContentPart objects
                    "content": [
                        {
                            "type": "text",
                            "text": f"Current Working Directory: {cwd}"
                        }
                    ]
                }
            }
            send_response(response)
        except Exception as e:
            logging.error(f"Error getting CWD: {e}")
            send_error(request_id, ERROR_INTERNAL, f"Failed to get current directory: {e}")
    else:
        # Tool not found
        send_error(request_id, ERROR_METHOD_NOT_FOUND, f"Tool '{tool_name}' not found.")

def main():
    """Main request handling loop."""
    logging.info("Python MCP Server started. Waiting for requests on stdin...")
    for line in sys.stdin:
        logging.info(f"Received line: {line.strip()}")
        try:
            request = json.loads(line)
            request_id = request.get("id")
            method = request.get("method")
            params = request.get("params", {})

            if not all([request.get("jsonrpc") == "2.0", method, request_id is not None]):
                 # Send a general error if basic JSON-RPC structure is wrong
                 # Note: A robust server might handle batch requests etc.
                 if request_id is not None:
                     send_error(request_id, -32600, "Invalid Request structure")
                 else:
                     # Cannot reply if ID is missing
                     logging.error("Received invalid request with no ID.")
                 continue # Skip processing this line

            if method == "initialize":
                initialize(request_id)
            elif method == "ListTools":
                handle_list_tools(request_id)
            elif method == "CallTool":
                handle_call_tool(request_id, params)
            else:
                # Method not supported by this simple server
                send_error(request_id, ERROR_METHOD_NOT_FOUND, f"Method '{method}' not supported.")

        except json.JSONDecodeError:
            logging.error(f"Failed to decode JSON: {line.strip()}")
            # Cannot reply if JSON is invalid and we can't get an ID
            continue
        except Exception as e:
            logging.exception("An unexpected error occurred during request handling.")
            # Try to send an internal error if we have an ID
            request_id = request.get("id") if 'request' in locals() else None
            if request_id is not None:
                send_error(request_id, ERROR_INTERNAL, f"An internal server error occurred: {e}")

if __name__ == "__main__":
    main()
