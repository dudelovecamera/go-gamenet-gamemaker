//global.players = ds_map_create(); // Create a map to store player positions

position = "0,0";

///@description Say hi to server that will broadcast to every client
connected = false
client = network_create_socket(network_socket_tcp)
network_connect_raw_async(client, "185.94.99.107", 8080)
//network_connect_raw_async(client, "127.0.0.1", 8080)
msg = ""
num = ""
c_buffer = buffer_create(1, buffer_grow, 1)
xx=x
yy=y