<?php
include("flag.php");
if (isset($_GET['source'])) {
    show_source(__FILE__);
    die();
}
if (!isset($_GET['password'])) {
    echo <<<EOT
    <h2>Can you guess the magic password?</h2>
    <form>
        <input name="password" type="text" placeholder="password"/>
        <input type="submit" value="Login"/>
    </form>
    <a href="/?source">Source Code</a>
    EOT;
} else {
    $password = $_GET['password'];
    $correctPassword = "T0uUfK2Q8cik";

    # TODO: Remove blocking of admin password
    if ($password === $correctPassword) {
        die("This password is currently disabled");
    }
    # It's secure to compare hashes of passwords!
    if (md5($password) == md5($correctPassword)) {
        echo "<h1>Welcome admin, here's your secret flag: $flag</h1>";
    } else {
        echo "Incorrect password!";
    }
}
?>
