# tagitup
webapp where users can add and tag links, and search for links their friends have tagged

Current prototype is an example of a full-stack CRUD app in React.js and Golang.

Users can add and tag links, then search intelligently for the tags they are interested in.

It's currently hosted on Heroku (https://secure-meadow-75884.herokuapp.com), with a SQL database (wrapped with in-memory caching). It also uses OAuth to let users log in with their favourite service (Facebook, Twitter, etc), though for now I've only integrated with Github.

There's still a lot of functionality to add, but I'm most excited about extending the architecture to a micro-services-ready implementation, then letting users follow each other and search based on the tags of users they follow
