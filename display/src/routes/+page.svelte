<script>
    import "../app.css";

    // Sample data for the packet table
    let packets = [
        {
            timestamp: "2024-11-01 12:34:56",
            sourceIP: "192.168.1.1",
            destinationIP: "192.168.1.2",
            protocol: "TCP",
            size: "1500 bytes",
        },
        {
            timestamp: "2024-11-01 12:35:01",
            sourceIP: "192.168.1.3",
            destinationIP: "192.168.1.4",
            protocol: "UDP",
            size: "850 bytes",
        },
        // Add more packets as needed
    ];
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
th, td {
    padding: 12px 15px;
    color: var(--rp-text);
}



</style>
