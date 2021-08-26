#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    local PP1=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        -e "s#\${PEERPEM1}#$PP1#" \
        -e "s#\${P0PORT1}#$7#" \
        ./ccp-template.json
}

{{range .Organizations}}
ORG={{.OrgNum}}
P0PORT={{.P0Port}}
CAPORT={{.CAPort}}
P0PORT1={{.P1Port}}
PEERPEM=../../artifacts/channel/crypto-config/peerOrganizations/{{.OrgName}}.example.com/peers/peer0.{{.OrgName}}.example.com/msp/tlscacerts/tlsca.{{.OrgName}}.example.com-cert.pem
PEERPEM1=../../artifacts/channel/crypto-config/peerOrganizations/{{.OrgName}}.example.com/peers/peer1.{{.OrgName}}.example.com/msp/tlscacerts/tlsca.{{.OrgName}}.example.com-cert.pem
CAPEM=../../artifacts/channel/crypto-config/peerOrganizations/{{.OrgName}}.example.com/msp/tlscacerts/tlsca.{{.OrgName}}.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $PEERPEM1 $P0PORT1)" > connection-{{.OrgName}}.json
{{end}}
