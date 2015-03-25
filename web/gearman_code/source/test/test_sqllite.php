<?php
//     $db = new SQLite3('/data/DockerManager/db/test.db');
//     $db->exec('CREATE TABLE files (id INTEGER PRIMARY KEY, filename TEXT, content BLOB);');
    
//     $statement = $db->prepare('INSERT INTO files (filename, content) VALUES (?, ?);');
//     $statement->bindValue('filename', 'Archive.zip');
//     $statement->bindValue('content', file_get_contents('Archive.zip'));
//     $statement->execute();
    
//     $fp = $db->openBlob('files', 'content', $id);
    
//     while (!feof($fp))
//     {
//         echo fgets($fp);
//     }
    
//     fclose($fp);

$db = new SQLite3('/data/DockerManager/db/DockerManager.db');

$results = $db->query('SELECT *  FROM images');
while ($row = $results->fetchArray()) {
    var_dump($row);
}
?>
