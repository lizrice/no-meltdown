#!/bin/sh
set -x
curl "127.0.0.1/?name=Robert'%3B%20drop%20table%20Students%3B--"
