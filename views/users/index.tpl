<!DOCTYPE html>
<html>
    <body>
        <td><a href="/users/create">Добавить</a></td>
        <table border = "1" style ="width:80%;">
            {{range.index}}
            <tr>
                <td>{{.Id}}</td>
                <td>{{.Name}}</td>
                <td>{{.Email}}</td>
                <td>{{.Phone}}</td>
                <td>{{.Password}}</td>
                <!-- <td><a href="/groups/view?Id={{.Id}}">Просмотр</a></td>   -->
                <td><a href="/users/delete?Id={{.Id}}">Удалить</a></td> 
                <!-- <td><a href="/groups/update?Id={{.Id}}&flag=1">Обновить</a></td> -->
            </tr>
            {{end}}
        </table>
    </body>
</html>
