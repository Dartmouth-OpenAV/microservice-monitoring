<?php
header('Content-Type: application/json');
define("NAMESERVER", "129.170.17.4");
define("EXTERNAL_HOST", "google.com");

$request_method = $_SERVER['REQUEST_METHOD'];

///////////////////////////////////////////////////
//
//  Endpoints:
//  external - ping EXTERNAL_HOST - usage: /pingexternal
//  router - ping router_host - usage: /pingrouter/{router_host}
//  dns - ping NAMESERVER - usage: /pingdns
//
///////////////////////////////////////////////////

$router_host = "";
$parts = explode('/', $_GET['command'], 2);
$endpoint = $parts[0];
if (isset($parts[1])) {
    $router_host = $parts[1];
}
$nameserver = "129.170.17.4";
$external_host = "google.com";
$output = '"false"';

if ($request_method == "GET") {
    if ($endpoint === 'pingexternal') {
        if (get_ping($external_host)) {
            $output = '"true"';
        }
    } elseif ($endpoint === 'pingrouter') {
        if (get_ping($router_host)) {
            $output = '"true"';
        }
    } elseif ($endpoint === 'pingdns') {
        if (get_ping($nameserver)) {
            $output = '"true"';
        }
    }
}

print $output;


///////////////////////////////////////////////////
//
//  Functions
//  get_ping($host, $timeout)
//
///////////////////////////////////////////////////

function get_ping($host, $timeout = 1) {
    $command = "ping -c 1 -W " . $timeout . " " . escapeshellarg($host) . " 2>&1";

    $output = [];
    $result = 0;
    exec($command, $output, $result);

    return $result === 0;
}
