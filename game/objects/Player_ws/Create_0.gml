connected=false
xx=x
yy=y
c_buffer = buffer_create(1, buffer_grow, 1)
server_ip = "185.94.99.107"; // Replace with your server's IP address
//server_ip="127.0.0.1"
server_port = 8080; // Replace with your server's port
ws_socket = network_create_socket(network_socket_ws);
network_connect_raw_async(ws_socket,server_ip, server_port)
