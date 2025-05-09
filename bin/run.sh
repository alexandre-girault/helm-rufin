#!/bin/bash

PLUGIN_BIN="${HELM_PLUGIN_DIR}/bin/rufin"

HELM_ARGS_WITH_SECRETS="$($PLUGIN_BIN $@)"


helm ${HELM_ARGS_WITH_SECRETS}

#rm -rf with-secrets-*.yaml

