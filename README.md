# [Ten2](https://ten2.will-mccall.com)
## I created this website, mostly for myself, to be able to see religious data about every country in the world
### I did this because I think the Joshua Project website is hard to work with (especially on mobile)

* I chose to do almost everything on the server because my server doesn't have the memory for most popular front-end frameworks and why overcomplicate things?
* I chose Golang because I like how simple and easy it is and because it's pretty memory efficient for small applications like mine
* I chose SQLite because it's fast and lightweight (especially for read operations which is most of what this website needs to do)

I'm particularly proud of all the [maps](helpers/maps), I used Plotly and Google Maps for all the geolocation stuff and [Joshua Project](https://joshuaproject.net) to get all my data
