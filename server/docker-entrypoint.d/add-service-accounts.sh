#!/usr/bin/env bash

echo "-- Creating service-accounts (admin, viewer, and agent)"
htpasswd -b -B -c /go-working-dir/config/static-passwords admin "${ADMIN_PASSWORD}"
htpasswd -b -B /go-working-dir/config/static-passwords viewer "${VIEWER_PASSWORD}"
htpasswd -b -B /go-working-dir/config/static-passwords agent "${AGENT_PASSWORD}"
