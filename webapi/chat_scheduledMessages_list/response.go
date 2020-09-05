// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatScheduledMessagesList, err := UnmarshalChatScheduledMessagesList(bytes)
//    bytes, err = chatScheduledMessagesList.Marshal()

package chat_scheduledMessages_list

import "encoding/json"

func UnmarshalChatScheduledMessagesList(data []byte) (ChatScheduledMessagesList, error) {
	var r ChatScheduledMessagesList
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatScheduledMessagesList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatScheduledMessagesList struct {
	Ok                *bool              `json:"ok,omitempty"`                
	ScheduledMessages []ScheduledMessage `json:"scheduled_messages,omitempty"`
	ResponseMetadata  *ResponseMetadata  `json:"response_metadata,omitempty"` 
	Error             *string            `json:"error,omitempty"`             
	Needed            *string            `json:"needed,omitempty"`            
	Provided          *string            `json:"provided,omitempty"`          
}

type ResponseMetadata struct {
	NextCursor *string `json:"next_cursor,omitempty"`
}

type ScheduledMessage struct {
	ID          *string `json:"id,omitempty"`          
	ChannelID   *string `json:"channel_id,omitempty"`  
	Text        *string `json:"text,omitempty"`        
	PostAt      *int64  `json:"post_at,omitempty"`     
	DateCreated *int64  `json:"date_created,omitempty"`
}
