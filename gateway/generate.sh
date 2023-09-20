#!/bin/bash
rover supergraph compose --config  gateway/router.yaml > gateway/schema/supergraph.graphql  && \
sed 's/localhost:3000/user\-api:3000/' gateway/schema/supergraph.graphql > gateway/schema/supergraph.graphql.tmp && \
mv gateway/schema/supergraph.graphql.tmp gateway/schema/supergraph.graphql && \
sed 's/localhost:3001/inventory\-api:3000/' gateway/schema/supergraph.graphql > gateway/schema/supergraph.graphql.tmp && \
mv gateway/schema/supergraph.graphql.tmp gateway/schema/supergraph.graphql && \
sed 's/localhost:3002/order\-api:3000/' gateway/schema/supergraph.graphql > gateway/schema/supergraph.graphql.tmp && \
mv gateway/schema/supergraph.graphql.tmp gateway/schema/supergraph.graphql