#!/bin/bash
rover dev --url http://localhost:3000/graphql --name=user & \
sleep 10 && \
rover dev --url http://localhost:3001/graphql --name=inventory & \
sleep 10 && \
rover dev --url http://localhost:3002/graphql --name=order