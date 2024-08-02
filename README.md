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

### Step by step
1. main
   - TODO: init config: cleanenv
   - TODO: init logger: slog
   - TODO: init storage: sqlite
   - TODO: init router: chi, chi render
   - TODO: init server
2. ```go get -u github.com/ilyakaznacheev/cleanenv```
3. Must в имени функции означает, что можно паниковать если не сработает
4. ```$env:CGO_ENABLED=1``` - для работы с БД SQLITE
   - Open File -> Settings -> Go -> Vendoring & Build Tags
   - Set CGO Support to Disabled
   - метод не рабочий
5. Альтернатива без CGO это ```go get modernc.org/sqlite``` но методы другие - отказался
6. Скачал и истановил [x86_64-13.1.0-release-posix-seh-msvcrt-rt_v11-rev1.7z](https://github.com/niXman/mingw-builds-binaries/releases)
   антивирус помещает в карантин - добавлять исключения 
   ```bash 
   $ gcc --version
   gcc.exe (x86_64-posix-seh-rev1, Built by MinGW-Builds project) 13.1.0 
   ```
7. Add router ```go get github.com/go-chi/chi/v5```
8. Add render chi ```go get github.com/go-chi/render```
9. Add validator 
   - ```go get github.com/go-playground/validator/v10```
   - in struct ```validate:"required,url"```
   - in code ```validator.New().Struct(req)```
10. Add assert for testing ```go get github.com/stretchr/testify/assert```
11. Mock generator:
    - ```go install github.com/vektra/mockery/v2@v2.34.2```
    - add to .go file ```//go:generate ...``` for interface
    - ```//go:generate go run github.com/vektra/mockery/v2@v2.34.2 --name=URLSaver```
12. Add testify:
    - ```go get github.com/stretchr/testify```
    - ```go get github.com/stretchr/testify/mock```
13. Chi
    - ```go get -u github.com/go-chi/chi/v5```
14. Fake data lib
    - ```go get github.com/brianvoe/gofakeit/v7```
15. 
 

