package jq

import "fmt"

// This is an example. Replace thie following with whatever steps are needed to
// install required components into
// const dockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli
// `

/*
ENV JQ_VERSION='1.6'

RUN wget --no-check-certificate https://raw.githubusercontent.com/stedolan/jq/master/sig/jq-release.key -O /tmp/jq-release.key && \
    wget --no-check-certificate https://raw.githubusercontent.com/stedolan/jq/master/sig/v${JQ_VERSION}/jq-linux64.asc -O /tmp/jq-linux64.asc && \
    wget --no-check-certificate https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64 -O /tmp/jq-linux64 && \
    gpg --import /tmp/jq-release.key && \
    gpg --verify /tmp/jq-linux64.asc /tmp/jq-linux64 && \
    cp /tmp/jq-linux64 /usr/bin/jq && \
    chmod +x /usr/bin/jq && \
    rm -f /tmp/jq-release.key && \
    rm -f /tmp/jq-linux64.asc && \
    rm -f /tmp/jq-linux64
*/

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	fmt.Fprintln(m.Out, `RUN apt-get update && apt-get install -y wget && rm -rf /var/lib/apt/lists/*`)
	fmt.Fprintln(m.Out, `RUN apt-get update && apt-get install gnupg2 -y`)
	fmt.Fprintln(m.Out, `ENV JQ_VERSION='1.6'`)
	fmt.Fprintln(m.Out, `RUN wget --no-check-certificate https://raw.githubusercontent.com/stedolan/jq/master/sig/jq-release.key -O /tmp/jq-release.key && \
    wget --no-check-certificate https://raw.githubusercontent.com/stedolan/jq/master/sig/v${JQ_VERSION}/jq-linux64.asc -O /tmp/jq-linux64.asc && \
    wget --no-check-certificate https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64 -O /tmp/jq-linux64 && \
    gpg --import /tmp/jq-release.key && \
    gpg --verify /tmp/jq-linux64.asc /tmp/jq-linux64 && \
    cp /tmp/jq-linux64 /usr/bin/jq && \
    chmod +x /usr/bin/jq && \
    rm -f /tmp/jq-release.key && \
    rm -f /tmp/jq-linux64.asc && \
    rm -f /tmp/jq-linux64`)
	return nil
}
