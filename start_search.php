<?php
  first = $_POST["first"];
  second = $_POST["second"];
  exec("./chandler " . $first . " " . $second . " /var/www/result.php", $output);
  echo "Here is the map: " . $output;
?>
