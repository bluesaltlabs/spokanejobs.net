# bedrock

Bedrock project monorepo. Source code for spokanejobs.net


---


## Frontend

> `/frontend`


The Bedrock Frontend is a vue PWA app, deployed to [SpokaneJobs.Net](https://spokanejobs.net).



## Infra

> `/infra`


Originally the API for this project was going to reside on Supabase.
Currently the data is retrieved from [bluesaltlabs/spokanejobs_data](https://github.com/bluesaltlabs/spokanejobs_data) for simplicity and cost.



## Scraper


> `/scraper`


The scraper is a go app deployed as a docker container.
This container is initiated via [this scheduler script](https://github.com/bluesaltlabs/bluesaltlabs.local/blob/main/scripts/bedrock-scheduler.sh).


## MCP

> `/mcp`


This is an experiement to see how one would implement an MCP server using python for Bedrock.
See sub-directory for more information.
