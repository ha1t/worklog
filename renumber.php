<?php
/**
 *
 *
 */

$file_list = array();
foreach (glob('./screencapture/*') as $filename) {
    $file_list[] = basename($filename);
}

sort($file_list);

$number = 1;
foreach ($file_list as $filename) {
    $new_filename = sprintf("%08d", $number) . '.png';
    $number++;
    if ($filename != $new_filename) {
        echo "{$filename} -> {$new_filename}" . PHP_EOL;
        rename('./screencapture/' . $filename, './screencapture/' . $new_filename);
    }
}
