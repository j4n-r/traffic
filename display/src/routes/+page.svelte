<script>
    import "../app.css";
    let status = $state("idle")
    let packets = $state([])

    const socket = new WebSocket("ws://localhost:8999");
    socket.addEventListener("open", () => {
       status = "open";
    });
    socket.addEventListener("error", (event) => {
        status = event;
    })
    socket.addEventListener("message", (event) => {
        packets.push(event.data);
    })
</script>

<div class="table-container">
    <table class="packet-table">
        <thead>
            <tr>
                <th>Timestamp</th>
                <th>Source IP</th>
                <th>Destination IP</th>
                <th>Protocol</th>
                <th>Size</th>
            </tr>
        </thead>
        <tbody>
            {#each packets as packet}
                <tr>
                    <td>{packet.timestamp}</td>
                    <td>{packet.sourceIP}</td>
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
        background-color: var(--rp-surface);
        border-radius: 10px;
        padding: 2em;
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
        color: var(--rp-text);
    }
</style>
