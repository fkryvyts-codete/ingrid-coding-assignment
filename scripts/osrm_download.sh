#!/bin/bash

wget http://download.geofabrik.de/europe/germany/berlin-latest.osm.pbf

docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-extract -p /opt/car.lua /data/berlin-latest.osm.pbf
docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-partition /data/berlin-latest.osrm
docker run -t -v "${PWD}:/data" osrm/osrm-backend osrm-customize /data/berlin-latest.osrm
