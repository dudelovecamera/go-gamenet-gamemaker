/*

position = "0,0";

///@description Say hi to server that will broadcast to every client

client = network_create_socket(network_socket_tcp)
network_connect_raw_async(client, "185.94.99.107", 8080)
//network_connect_raw_async(client, "127.0.0.1", 8080)
msg = ""
num = ""




*/
connected = false
xx=x
yy=y
c_buffer = buffer_create(1, buffer_grow, 1)
server_ip = "185.94.99.107"; // Replace with your server's IP address
//server_ip="127.0.0.1"
server_port = 8090; // Replace with your server's port
server_port = irandom_range(8080, 8090);
udp_socket = network_create_socket(network_socket_udp);
network_connect_raw_async(udp_socket,server_ip, server_port)

/*
if (udp_socket == -1) {
    show_error("Failed to create UDP socket.", true);
} else {
    var result = network_connect(udp_socket, server_ip, server_port);
    if (result != 0) {
        show_error("Failed to connect to server: " + string(result), true);
    }
}