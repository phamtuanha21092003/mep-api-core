#!/bin/sh
tusd -behind-proxy -hooks-enabled-events pre-create -hooks-http http://gateway:8080/api/v1/hooks/tus -base-path /api/files -s3-bucket my-bucket -s3-endpoint http://minio:9000 -host 0.0.0.0 -port 8000
