1) Запустите следующую команду в корне вашего проекта, там где лежит go.mod:
```
go list -f "{{.Name}}__{{.ImportPath}}__{{.Imports}}" ./... > pkg_list.txt
```

2) Установка:
```
go install github.com/iziCode/gopkgplantuml/cmd/gopkgplantuml@latest
```

3) Использование:
```
gopkgplantuml /path/to/your/pkg_list.txt
```

4) Открыть сгенерированный файл можно в:
   1) Jetbrains IDE установив плагин PlantUML Integration
   2) Онлайн, например, http://www.plantuml.com/plantuml/uml/


