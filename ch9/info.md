Обновить список зависимостей в файле go.mod и загрузить используемые в проекте модули:
```bash
go mod tidy
```
Посмотреть список пакетов в проекте:
```bash
go list
```
Посмотреть список модулей:
```bash
go list -m
```
Посмотреть, какие версии модуля доступны:
```bash
go list -m -versions github.com/nickname/reponame
```
Откатится или обновить версию модуля
```bash
go get github.com/nickname/reponame@v1.0.0
```
Обновиться до версии патча, исправляющей ошибки текущей второстепенной версии:
```bash
go get -u=patch github.com/nickname/reponame
```
Получить самую свежую версию модуля
```bash
go get -u github.com/nickname/reponame
```
Сохранить все копии зависимостей в текущем модуле, в папке vendor (вендоринг):
```bash
go mod vendor
```