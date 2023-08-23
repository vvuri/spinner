# spinner

- Slides from book and article.
- Markdown text with image convert to slide. 
- Search by tags and title.

### Run


#### Develop
Init DB and migration
```
$ spinner init
```
Run API server
```
$ spinner run -p 8080
```

### Plan Steps
1. Read dir './data' with *.md files
2. MD file format:
```
[!Title article]
[#tag1] 
[#tag2] 
[#tag3]
***
slide1 markdown
***
slide2 markdown
```
3. Connect to DB - SQLite
4. Migration
5. Save all data in DB
- title and tags
- text with md
6. API server - Get
- /list - all titles with id
- /tags - all tags with id
- /article/id - all slides from article
- /tag/id - all article title and id wit this tag
7. WEB UI - SPA page
- convert to html
- send to web
8. CSS style
9. Add images to article
10. Docker 
11. Docker compose
12. Add metrics - request per minute, get slides per minute
13. Kind deploy
14. Nginx links to web UI
15. Prometheus + Grafana
16. Linter
17. Tests
