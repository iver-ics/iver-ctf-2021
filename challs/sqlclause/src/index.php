<?php
if (!isset($_GET['username']) || !isset($_GET['password'])) {
    echo <<<EOT
    <h2>Can you guess the magic password?</h2>
    <form>
        <input name="username" type="text" placeholder="username"/>
        <input name="password" type="text" placeholder="password"/>
        <input type="submit" value="Login"/>
    </form>
    EOT;
} else {
    $username = $_GET['username'];
    $password = $_GET['password'];

    $db = new SQLite3("1jkn24n1j2n4k1j2nkj1n2fkjn12kjdn1kn.db");
    $db->query("
    CREATE TABLE IF NOT EXISTS users (
        username TEXT UNIQUE,
        password TEXT
    );
    INSERT OR REPLACE INTO users(username, password) VALUES('admin', 'admin1337!');
    ");
    $res = $db->query("SELECT * FROM users WHERE username='$username' AND password='$password'");
    if (!$res) {
        die("Internal SQL error!");
    }
    $row = $res->fetchArray();
    if ($row) {
        echo "<h1>Welcome admin! Here's your flag: haxmas{Pr3par3d_St4teMeNt5_4Re_A_MuS7}</h1>";
    } else {
        echo "Incorrect password!";
    }
}
?>
