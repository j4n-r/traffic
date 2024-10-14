#include "capture.h"
#include <pcap/pcap.h>
#include <stdlib.h>

void packet_handler(u_char* args, const struct pcap_pkthdr* hdr,
                    const u_char* pkt_body);
void pcap_err(int err, char* errbuf);

int main(int argc, char* argv[]) {
    char errbuf[PCAP_ERRBUF_SIZE] = {};
    pcap_err(pcap_init(PCAP_CHAR_ENC_UTF_8, errbuf), errbuf);

    pcap_t* handle = pcap_open_live("any", 262144, 1, 1000, errbuf);
    if (handle == NULL) {
        printf("%s\n", errbuf);
        exit(1);
    }

    int err = pcap_loop(handle, -1, packet_handler, NULL);
    if (err == PCAP_ERROR) {
        printf("PCAP_ERROR");
        exit(1);
    }
    pcap_close(handle);
}

void packet_handler(u_char* args, const struct pcap_pkthdr* header,
                    const u_char* packet) {
    printf("Packet capture length: %d\n", header->caplen);
    printf("Packet total length %d\n", header->len);
    printf("Packet body %s\n", packet);
    return;
}

void pcap_err(int err, char* errbuf) {
    if (err == PCAP_ERROR) {
        printf("%s", errbuf);
        exit(1);
    }
}
