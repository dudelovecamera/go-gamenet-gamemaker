
if (keyboard_check(vk_left)) {
    xx = x - 5; // Move left
}
if (keyboard_check(vk_right)) {
    xx = x+ 5; // Move right
}
if (keyboard_check(vk_up)) {
    yy = y - 5; // Move up
}
if (keyboard_check(vk_down)) {
 yy = y + 5
 
}

if(xprevious!=xx or yprevious!=yy){
   
   
var jsonData = ds_map_create();

// Add some key-value pairs to the map

ds_map_add(jsonData, "x", xx);
ds_map_add(jsonData, "y", yy);

// Convert the map to a JSON string
var jsonString = json_encode(jsonData)+"\n";
ds_map_destroy(jsonData);	
// Output the JSON string (for example, to the console)
//show_message(jsonString);

// Clean up the map


///@description Say hi to server that will broadcast to every client
buffer_seek(c_buffer, buffer_seek_start, 0)

buffer_write(c_buffer, buffer_string, jsonString)
   
 ;
	network_send_udp_raw(udp_socket, server_ip, server_port, c_buffer, buffer_tell(c_buffer));
}