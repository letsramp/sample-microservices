#!/bin/sh
sudo dockerd &
exec "$@"