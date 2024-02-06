import os
import json
import openai

CONFIG_PATH = "../deploy/config/dev.json"

# env variables config
with open(CONFIG_PATH) as json_file:
    config_data = json.load(json_file)
for key, value in config_data.items():
    os.environ[key] = value

# open ai key
openai.api_key = os.getenv("OPENAI_API_KEY")

# basic example

## set open ai params
params = {
    "model": "gpt-3.5-turbo",
    "temperature": 0.7,
    "max_tokens": 256,
    "top_p": 1,
    "frequency_penalty": 0,
    "presence_penalty": 0
}

## create the content
content = "The sky is"

## create the messages
messages = [
    {
        "role": "user",
        "content": content
    }
]

## get the response
response = openai.chat.completions.create(
    model = params["model"],
    messages = messages,
    temperature = params["temperature"],
    max_tokens = params["max_tokens"],
    top_p = params["top_p"],
    frequency_penalty = params["frequency_penalty"],
    presence_penalty = params["presence_penalty"]
)

print(response)
