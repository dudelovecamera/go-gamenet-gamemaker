
var type = ds_map_find_value(async_load, "type"); 
//show_debug_message("this is the number I got: " + string(type))
if (type == network_type_non_blocking_connect) {
	connected = true
}
else if (type == network_type_data) {
	var t_buffer = ds_map_find_value(async_load, "buffer"); 
	//var cmd_type = buffer_read(t_buffer, buffer_u8);

//show_debug_message(cmd_type)

var message = buffer_read(t_buffer, buffer_string)
			//show_debug_message("Reason: " + reason)
			var parts = string_split(message, "\n");
   // for (var i = 0; i < array_length(parts); i++) {
		//joined player1 0,0
      //  var reason =parts[i]// parts[i];
		//var cmd = string_split(reason, " ");
		//switch(cmd[0]) {
		//case "0":
		//joined
		//break;
		//case "1":
		//moved
		//show_message(array_length(cmd))
		//show_message(reason )
		// if ( cmd[2] != "") {
			 
			var pos_tokens = string_split(message, ",");
			 if ( array_length(pos_tokens) == 3) {
			show_debug_message(pos_tokens[0])
			show_debug_message(pos_tokens[1])
			
			//var mx =real(string_trim(pos_tokens[0]));
			//var my= real(string_trim(pos_tokens[1]));
			//if(xprevious!=mx){
				//show_message(xprevious)
				x= real(string_trim(pos_tokens[0]))
			//}
			//if(yprevious!=my){
				//show_message(yprevious)
				y=  real(string_trim(pos_tokens[1]))
			//}
			
			 }else{
				// show_debug_message("here" + reason)
			 }
			//show_message(pos_tokens[1])
			//x =real(string_trim(string_delete(pos_tokens[0],0,1)));
			
			
			
		//}
		
		//break;
		//}
       
	//	}
			
	
	//show_debug_message("other_x:" + other_x)
	//show_debug_message("other_y:" + other_y)
	/*switch(cmd_type) {
		case NET_GET_REQUESTED_NUM:
			var s = buffer_read(t_buffer, buffer_u16);
			num = string(s)
			show_debug_message("this is the number I got: " + num)
			break;
		case NET_GET_KICKED:
			var reason = buffer_read(t_buffer, buffer_string)
			show_message_async("You Have been kicked from the server.\nReason: " + reason)
			network_destroy(client)
			break;
		case NET_GET_MSG:
			msg = buffer_read(t_buffer, buffer_string)
			show_debug_message("this is the number I got: " + msg)
			break;
	}*/
}
/*

    // Process the message
    var parts = string_split(message, "\n");
    for (var i = 0; i < array_length(parts); i++) {
        var msg = parts[i];
        if (msg != "") {
            // Example: "Player 1 moved to position 100,200"
            var tokens = string_split(msg, " ");
            if (array_length(tokens) >= 5) {
                var player_id = tokens[1]; // Player ID
                var player_position = tokens[4]; // Position (e.g., "100,200")
                
                // Update other players' positions
                // You will need to implement logic to update the position of other players
                // For example, you could store player positions in a data structure
                // and draw them in the Draw Event.
            }
        }
    }
