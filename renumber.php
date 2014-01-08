<?php
$file_list = array();
$target_dir = $argv[1] . '/';
$target_date = $argv[2];

if (!is_numeric($target_date)) {
    echo "invalid date";
    exit;
}

if (!file_exists($target_dir)) {
    echo "directory: {$target_dir} not found.";
    exit;
}

foreach (glob($target_dir . $target_date . '_*.png') as $filename) {
    $file_list[] = basename($filename);
}

sort($file_list);

$number = 1;
foreach ($file_list as $filename) {
    $new_filename = $target_date . '_' . sprintf("%05d", $number) . '.png';
    $number++;
    if ($filename != $new_filename) {
        echo "{$filename} -> {$new_filename}" . PHP_EOL;
        rename($target_dir . $filename, $target_dir . $new_filename);
    }
}
