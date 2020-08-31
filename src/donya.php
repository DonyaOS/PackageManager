<?php
/*
 * @Name: DonyaOS Package Manager (draft in PHP 7.4)
 * @Author: Max Base
 * @Repository: https://github.com/DonyaOS/PackageManager
 * @Date: August 30, 2020
 */
$flags=[];
if($argc === 1) {
	help();
}
else {
	if(startsWith($argv[1], "-")) {
		help();
	}
	else {
		$command=$argv[1];
		callCommand($command);
	}
}

function startsWith($haystack, $needle ) {
	$length = strlen($needle);
	return substr($haystack, 0, $length) === $needle;
}

function endsWith($haystack, $needle) {
	$length = strlen( $needle );
	if(!$length ) {
		return true;
	}
	return substr($haystack, -$length) === $needle;
}

function help($command="") {
  // show special help for commands: if($command == "install") {}
	print "Usage:
  donya [command]

Available Commands:
  help        Help about any command
  install     Installing package(s) in DonyaOS
  remove      Removing package(s) in DonyaOS
  list        Listing package(s) in DonyaOS
  search      Search package(s) in the repository of DonyaOS

Flags:
  -h, --help   help for donya

Use \"donya [command] --help\" for more information about a command.
";
}

function callCommand($command) {
	global $argv;
	$args=$argv;
	unset($args[0]); // software name. e.g: donya
	unset($args[1]); // command name
	$args=array_values($args); // start index of items from 0 in Array
	switch ($command) {
		case "i":
		case "install":
			commandInstall($args);
			break;
		case "r":
		case "remove":
			commandRemove($args);
			break;
		case "u":
		case "update":
			commandUpdate($args);
			break;
		case "l":
		case "list":
			commandList($args);
			break;
		case "s":
		case "search":
			commandSearch($args);
			break;
		case "h":
		case "help":
		default:
			help($command);
			break;
	}
}

function commandInstall($args) {
	$flags = flags("install", $args);
	$args = $flags[1];
	$flags = $flags[0];
	print_r($args);
	print_r($flags);
	foreach($args as $arg) {
		// Read list of package from `cache` file and detect category/group name of packages!
		$group="core"; // WE need to set this from `cache` file
		$data = file_get_contents("https://donyaos.github.io/Packages/$group/$arg/package.donya");
		if(($data !== null || $data !== "") and strlen($data) > 0) {
			print $data;
		}
	}
}

function commandRemove($args) {
	$flags = flags("remove", $args);
	$args = $flags[1];
	$flags = $flags[0];
	print_r($args);
	print_r($flags);
}

function commandList($args) {
	$flags = flags("list", $args);
	$args = $flags[1];
	$flags = $flags[0];
	print_r($args);
	print_r($flags);
}

function commandSearch($args) {
	$flags = flags("search", $args);
	$args = $flags[1];
	$flags = $flags[0];
	print_r($args);
	print_r($flags);
}

function flags($command, $args) {
	$flags=[];
	foreach($args as $i=>$arg) {
		if(startsWith($arg, "-")) {
			if($arg === "-h" || $arg === "--help") {
				unset($args[$i]);
				help();
			}
			else if($command == "install" && ($arg == "-r" || $arg == '--rewrite')) {
				$flags[]="rewrite";
			}
		}
	}
	return [$flags, $args];
}
