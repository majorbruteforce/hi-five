Services: 
- Socket connection manager: 
    - Creates new socket connections
    - Keeps track of open connections (status and number)
    - Cleans up stale connections
- Match maker: 
    - Takes in peer with an array of keywords
    - Finds best matches between the nodes using Blossom's Algo
    - Uses an optimistic wait time to find higher weighted pairs
    - Queues match making requests and clears them up
- Engagement status manager:
    - Keeps a live count of online and matched participants and broadcasts them constantly
    - Keeps track of keyword frequency and broadcasts them in realtime to users



Middlewares: 
- Auth checker: 
    - Checks for existing sessions in every request, sessions have an expiration time, session only stores name and ASL
    - If not exists, creates a local session upon hitting the app and stores name, age, sex and location in db for sometime
    - Get user details from sessions table
