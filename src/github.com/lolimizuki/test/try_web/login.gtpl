<html>
    <head>
        <title>
            "Helle Go Web"
        </title>
    </head>
    <body>
        <form action="/login" method="post">
            用戶名:<input type="text" name="username"><br />
            蜜子馬:<input type="password" name="password"><br />
            年齡?<input type="text" name="age"><br />
            <!-- drop menu -->
            <select name="fruit">
                <option value="apple">apple</option>
                <option value="pear">pear</option>
                <option value="banane">banane</option>
            </select>
            
            <!-- options -->
            <br />
            <input type="radio" name="girls" value="1">Yuyuko<br /> 
            <input type="radio" name="girls" value="2">Mizuki<br />
            <input type="radio" name="girls" value="3">Ayaya<br />
            
            <!-- check box(multi selections) -->
            <br />
            <input type="checkbox" name="sleeplist" value="remi">Remiria<br />
            <input type="checkbox" name="sleeplist" value="pachi">Pachi<br />
            <input type="checkbox" name="sleeplist" value="meilin">Meilin<br />
            
            <input type="submit" value="Login">
        </form>
    </body>
</html>