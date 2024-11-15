<script>
    import "../app.css";
    let status = $state("idle");
    let packets = $state([]);

    let packet = {
        version: "",
        protocol: "",
        timestamp: "",
        addr_type: "",
        src_addr: "",
        src_port: "",
        dst_addr: "",
        dst_port: "",
    };

    function construct_packet(bytes) {
        packet.version = bytes.slice(0, 1).toString();
        packet.protocol = bytes.slice(1, 2).toString();
        packet.timestamp = bytes.slice(2, 10).toString();
        if (bytes.slice(10, 11) != 0x00) {
            // check addr type for IP version
            packet.src_addr = bytes.slice(11, 15).toString();
            packet.src_port = bytes.slice(15, 17).toString();
            packet.dst_addr = bytes.slice(17, 21).toString();
            packet.dst_port = bytes.slice(21, 23).toString();
        }
        return packet;
    }

    const socket = new WebSocket("ws://localhost:8999");
    socket.addEventListener("open", () => {
        status = "open";
    });
    socket.addEventListener("error", (event) => {
        status = event;
    });
    socket.addEventListener("message", (event) => {
        packet =  construct_packet(Uint8Array(event.data))
        packets.push(packet);
    });
</script>

<div class="table-container">
    <table class="packet-table">
        <thead>
            <tr>
                <th></th>
                <th>Source IP</th>
                <th>Destination IP</th>
                <th>Protocol</th>
                <th>Size</th>
            </tr>
        </thead>
        <tbody>
            {#each packets as packet}
                <tr>
                    <td>{packet}</td>
                    <td>{packet.}</td>
                    <td>{packet.}</td>
                    <td>{packet.destinationIP}</td>
                    <td>{packet.protocol}</td>
                    <td>{packet.size}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>

<style>
    table {
        border-collapse: collapse;
    }

    .table-container {
        background-color: white;
        border-radius: 10px;
        padding-left: 2rem;
        padding-right: 2rem;
        overflow-y: auto;
        height: 90vh;
    }
    table {
        width: 100%;
    }

    thead tr {
        color: #ffffff;
        text-align: left;
    }
    th,
    td {
        padding: 12px 15px;
        color: var(--rp-dawn-text);
    }
    th {
        padding: 2rem;
        background-color: white;
        position: sticky;
        z-index: 2;
        top: 0;
    }
</style>
