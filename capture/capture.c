#include "capture.h"
#include <pcap/pcap.h>
#include <stdlib.h>

void packet_handler(u_char *args, const struct pcap_pkthdr *hdr,
                    const u_char *pkt_body);
void pcap_err(int err, char *errbuf);

int main(int argc, char *argv[]) {
	pcap_if_t *dev;
	char errbuf[PCAP_ERRBUF_SIZE] = {};

	pcap_findalldevs(&dev, errbuf);
	if (dev == NULL) {
		fprintf(stderr, "Couldn't find default device: %s\n", errbuf);
		return (2);
	}
	printf("Device: %s\n", dev->name);
	pcap_t *handle = pcap_open_live(dev->name, 262144, 1, 1000, errbuf);
	if (handle == NULL) {
		printf("%s\n", errbuf);
		exit(1);
	}
	if (pcap_datalink(handle) != DLT_EN10MB) {
		fprintf(stderr,
		        "Device %s doesn't provide Ethernet headers - not supported\n",
		        dev->name);
		return (2);
	}

	int err = pcap_loop(handle, -1, packet_handler, NULL);
	if (err == PCAP_ERROR) {
		printf("PCAP_ERROR");
		exit(1);
	}
	pcap_close(handle);
}

void packet_handler(u_char *args, const struct pcap_pkthdr *header,
                    const u_char *packet) {
	printf("Packet capture length: %d\n", header->caplen);
	printf("Packet total length %d\n", header->len);
	printf("Packet body %s\n", packet);
	return;
}

void pcap_err(int err, char *errbuf) {
	if (err == PCAP_ERROR) {
		printf("%s", errbuf);
		exit(1);
	}
}
