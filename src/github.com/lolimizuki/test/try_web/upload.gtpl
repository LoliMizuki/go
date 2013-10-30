<html>
    <head>
        <title>
            Upload your secret
        </title>
    </head>
    
    <body>
        <form enctype="multipart/form-data" action="/upload_file" method="post">
            <input type="file" name="uploadfile" /><br />
            <input type="hidden" name="token" value="{{.}}" /><br />
            <input type="submit" name="Upload Your Secret" />
        </form>
    </body>
</html>