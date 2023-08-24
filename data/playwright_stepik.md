[!Автоматизация тестирования с помощью Playwright Python]
[#playwright]
[#autotest]
[#python]
***
[Stepik](https://stepik.org/lesson/826357)
```python
from playwright.sync_api import Playwright, sync_playwright, expect
```
Следующие три строки отвечают за запуск браузера и создание в нем контекста
```
browser = playwright.chromium.launch(headless=False)
context = browser.new_context()
page = context.new_page()
````

playwright.chromium.launch -  запуск браузера chromium

headless=False - дает команду, чтобы браузер chromium отображался и был видимым при запуске кода. Ели вы установите True, то браузер не будет отображаться. Но при этом все записанные в коде действия сценария будут выполнены.

new_context() - создает изолированный сеанс браузера.

new_page()  - открывает новую страницу(tab) в браузере

После того как браузер открыт и подготовлен, можно начать с ним работать.

page.goto("https://playwright-todomvc.antonzimaiev.repl.co/#/")
Метод page.goto() необходим, чтобы открыть веб-сайт.

click()- эмулирует клик левой кнопкой мышки по веб-элементу

fill() - этот метод вводит значения, переданные ему в качестве аргумента в веб-элемент. В нашем случае это текст  - "Создать первый сценарий playwright"

press("Enter") - эмулирует нажатие клавиши Enter на клавиатуре.

***
### Browsers
Для работы каждой версии  Playwright требуются определенные версии браузеров. В отличие от Selenium, Playwright не использует webdriver. С каждой новым выпуском, Playwright обновляет версии браузеров которые он поддерживает. Это означает, что при каждом обновлении Playwright вам, возможно придется заново запускать команду  playwright install

### BrowserContext
Playwright использует контексты браузера для достижения изоляции тестов. Страницы в двух отдельных контекстах не имеют общих cookie, настроек профиля.  Контекстом можно назвать независимую сессию браузера, схожую с режимом инкогнито.

### Pages
Page содержит содержимое загруженного веб-сайта. Каждый Context может иметь несколько страниц.

***
### Генератор
```
$ playwright codegen --viewport-size=800,600 https://playwright-todomvc.antonzimaiev.repl.co

$ playwright codegen -o lesson.py https://playwright-todomvc.antonzimaiev.repl.co
            -o lesson.py сразу сохранянть в файл
```

