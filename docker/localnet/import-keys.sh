#!/bin/bash

set -euox pipefail

KEYRING_BACKEND="test"

function import_key() {
    ALIAS=${1}
    MNEMONIC=${2}

    echo "Importing key: ${ALIAS}"

    if canow-chain keys show "${ALIAS}" --keyring-backend ${KEYRING_BACKEND}
    then
      echo "Key ${ALIAS} already exists"
      return 0
    fi

    echo "${MNEMONIC}" | canow-chain keys add "${ALIAS}" --keyring-backend ${KEYRING_BACKEND} --recover
}


import_key "base_account_1" "web region erupt kitchen ignore scout always cool advance tip window thank become liberty uncle reject powder task wheat industry blouse frozen trend two"
import_key "base_account_2" "field result budget animal friend solar update diesel sock almost casino play emotion pink honey conduct check witness copy eagle unlock genius brown dice"
import_key "base_vesting_account" "decide black crew connect cricket duck return finish piece license rough design lunch rude remember faculty shy cannon list toddler throw divide rent antique"
import_key "continuous_vesting_account" "pill soap false obvious echo still marine salute wheel patrol tourist sunset pizza destroy know alpha scare foot tragic lamp twin zero tonight defy"
import_key "delayed_vesting_account" "grant sample panda define master just pink mesh trash bulk north nominee avocado car banner wide hip amateur boost seek basic ribbon phrase day"
import_key "periodic_vesting_account" "cattle deliver practice infant clip want tag exercise inch guilt equal license connect shoe boat high garage people slim party display demise lesson curious"