// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    chatPostMessage, err := UnmarshalChatPostMessage(bytes)
//    bytes, err = chatPostMessage.Marshal()

package chat_postMessage

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalChatPostMessage(data []byte) (ChatPostMessage, error) {
	var r ChatPostMessage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ChatPostMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ChatPostMessage struct {
	Ok                 *bool             `json:"ok,omitempty"`                 
	Channel            *string           `json:"channel,omitempty"`            
	Ts                 *string           `json:"ts,omitempty"`                 
	Message            *Message          `json:"message,omitempty"`            
	Error              *string           `json:"error,omitempty"`              
	DeprecatedArgument *string           `json:"deprecated_argument,omitempty"`
	ResponseMetadata   *ResponseMetadata `json:"response_metadata,omitempty"`  
	Needed             *string           `json:"needed,omitempty"`             
	Provided           *string           `json:"provided,omitempty"`           
}

type Message struct {
	BotID        *string        `json:"bot_id,omitempty"`        
	Type         *string        `json:"type,omitempty"`          
	Text         *string        `json:"text,omitempty"`          
	User         *string        `json:"user,omitempty"`          
	Ts           *string        `json:"ts,omitempty"`            
	Team         *string        `json:"team,omitempty"`          
	BotProfile   *BotProfile    `json:"bot_profile,omitempty"`   
	ThreadTs     *string        `json:"thread_ts,omitempty"`     
	ParentUserID *string        `json:"parent_user_id,omitempty"`
	Subtype      *string        `json:"subtype,omitempty"`       
	Username     *string        `json:"username,omitempty"`      
	Icons        *MessageIcons  `json:"icons,omitempty"`         
	Root         *Root          `json:"root,omitempty"`          
	Blocks       []MessageBlock `json:"blocks,omitempty"`        
	Attachments  []Attachment   `json:"attachments,omitempty"`   
}

type Attachment struct {
	Text               *string           `json:"text,omitempty"`                 
	ID                 *int64            `json:"id,omitempty"`                   
	Fallback           *string           `json:"fallback,omitempty"`             
	Blocks             []AttachmentBlock `json:"blocks,omitempty"`               
	Color              *string           `json:"color,omitempty"`                
	MsgSubtype         *string           `json:"msg_subtype,omitempty"`          
	CallbackID         *string           `json:"callback_id,omitempty"`          
	Pretext            *string           `json:"pretext,omitempty"`              
	ServiceURL         *string           `json:"service_url,omitempty"`          
	ServiceName        *string           `json:"service_name,omitempty"`         
	ServiceIcon        *string           `json:"service_icon,omitempty"`         
	AuthorID           *string           `json:"author_id,omitempty"`            
	AuthorName         *string           `json:"author_name,omitempty"`          
	AuthorLink         *string           `json:"author_link,omitempty"`          
	AuthorIcon         *string           `json:"author_icon,omitempty"`          
	FromURL            *string           `json:"from_url,omitempty"`             
	OriginalURL        *string           `json:"original_url,omitempty"`         
	AuthorSubname      *string           `json:"author_subname,omitempty"`       
	ChannelID          *string           `json:"channel_id,omitempty"`           
	ChannelName        *string           `json:"channel_name,omitempty"`         
	BotID              *string           `json:"bot_id,omitempty"`               
	Indent             *bool             `json:"indent,omitempty"`               
	IsMsgUnfurl        *bool             `json:"is_msg_unfurl,omitempty"`        
	IsReplyUnfurl      *bool             `json:"is_reply_unfurl,omitempty"`      
	IsThreadRootUnfurl *bool             `json:"is_thread_root_unfurl,omitempty"`
	IsAppUnfurl        *bool             `json:"is_app_unfurl,omitempty"`        
	AppUnfurlURL       *string           `json:"app_unfurl_url,omitempty"`       
	Title              *string           `json:"title,omitempty"`                
	TitleLink          *string           `json:"title_link,omitempty"`           
	Fields             []Field           `json:"fields,omitempty"`               
	ImageURL           *string           `json:"image_url,omitempty"`            
	ImageWidth         *int64            `json:"image_width,omitempty"`          
	ImageHeight        *int64            `json:"image_height,omitempty"`         
	ImageBytes         *int64            `json:"image_bytes,omitempty"`          
	ThumbURL           *string           `json:"thumb_url,omitempty"`            
	ThumbWidth         *int64            `json:"thumb_width,omitempty"`          
	ThumbHeight        *int64            `json:"thumb_height,omitempty"`         
	VideoHTML          *string           `json:"video_html,omitempty"`           
	VideoHTMLWidth     *int64            `json:"video_html_width,omitempty"`     
	VideoHTMLHeight    *int64            `json:"video_html_height,omitempty"`    
	Footer             *string           `json:"footer,omitempty"`               
	FooterIcon         *string           `json:"footer_icon,omitempty"`          
	Ts                 *string           `json:"ts,omitempty"`                   
	MrkdwnIn           []string          `json:"mrkdwn_in,omitempty"`            
	Actions            []Action          `json:"actions,omitempty"`              
	Filename           *string           `json:"filename,omitempty"`             
	Size               *int64            `json:"size,omitempty"`                 
	Mimetype           *string           `json:"mimetype,omitempty"`             
	URL                *string           `json:"url,omitempty"`                  
	Metadata           *Metadata         `json:"metadata,omitempty"`             
}

type Action struct {
	ID              *string        `json:"id,omitempty"`              
	Name            *string        `json:"name,omitempty"`            
	Text            *string        `json:"text,omitempty"`            
	Style           *string        `json:"style,omitempty"`           
	Type            *string        `json:"type,omitempty"`            
	Value           *string        `json:"value,omitempty"`           
	Confirm         *ActionConfirm `json:"confirm,omitempty"`         
	Options         []Option       `json:"options,omitempty"`         
	SelectedOptions []Option       `json:"selected_options,omitempty"`
	DataSource      *string        `json:"data_source,omitempty"`     
	MinQueryLength  *int64         `json:"min_query_length,omitempty"`
	OptionGroups    []OptionGroup  `json:"option_groups,omitempty"`   
	URL             *string        `json:"url,omitempty"`             
}

type ActionConfirm struct {
	Title       *string `json:"title,omitempty"`       
	Text        *string `json:"text,omitempty"`        
	OkText      *string `json:"ok_text,omitempty"`     
	DismissText *string `json:"dismiss_text,omitempty"`
}

type OptionGroup struct {
	Text *string `json:"text,omitempty"`
}

type Option struct {
	Text  *string `json:"text,omitempty"` 
	Value *string `json:"value,omitempty"`
}

type AttachmentBlock struct {
	Type        *string         `json:"type,omitempty"`        
	Elements    []PurpleElement `json:"elements,omitempty"`    
	BlockID     *string         `json:"block_id,omitempty"`    
	Fallback    *string         `json:"fallback,omitempty"`    
	ImageURL    *string         `json:"image_url,omitempty"`   
	ImageWidth  *int64          `json:"image_width,omitempty"` 
	ImageHeight *int64          `json:"image_height,omitempty"`
	ImageBytes  *int64          `json:"image_bytes,omitempty"` 
	AltText     *string         `json:"alt_text,omitempty"`    
	Title       *TextElement    `json:"title,omitempty"`       
	Text        *TextElement    `json:"text,omitempty"`        
	Fields      []TextElement   `json:"fields,omitempty"`      
	Accessory   *Accessory      `json:"accessory,omitempty"`   
}

type Accessory struct {
	Type        *string `json:"type,omitempty"`        
	ImageURL    *string `json:"image_url,omitempty"`   
	AltText     *string `json:"alt_text,omitempty"`    
	Fallback    *string `json:"fallback,omitempty"`    
	ImageWidth  *int64  `json:"image_width,omitempty"` 
	ImageHeight *int64  `json:"image_height,omitempty"`
	ImageBytes  *int64  `json:"image_bytes,omitempty"` 
}

type PurpleElement struct {
	Type                         *string         `json:"type,omitempty"`                           
	Text                         *TextElement    `json:"text,omitempty"`                           
	ActionID                     *string         `json:"action_id,omitempty"`                      
	URL                          *string         `json:"url,omitempty"`                            
	Value                        *string         `json:"value,omitempty"`                          
	Style                        *string         `json:"style,omitempty"`                          
	Confirm                      *ElementConfirm `json:"confirm,omitempty"`                        
	Placeholder                  *TextElement    `json:"placeholder,omitempty"`                    
	InitialChannel               *string         `json:"initial_channel,omitempty"`                
	ResponseURLEnabled           *bool           `json:"response_url_enabled,omitempty"`           
	InitialConversation          *string         `json:"initial_conversation,omitempty"`           
	DefaultToCurrentConversation *bool           `json:"default_to_current_conversation,omitempty"`
	Filter                       *Filter         `json:"filter,omitempty"`                         
	InitialDate                  *string         `json:"initial_date,omitempty"`                   
	InitialOption                *InitialOption  `json:"initial_option,omitempty"`                 
	MinQueryLength               *int64          `json:"min_query_length,omitempty"`               
	ImageURL                     *string         `json:"image_url,omitempty"`                      
	AltText                      *string         `json:"alt_text,omitempty"`                       
	Fallback                     *string         `json:"fallback,omitempty"`                       
	ImageWidth                   *int64          `json:"image_width,omitempty"`                    
	ImageHeight                  *int64          `json:"image_height,omitempty"`                   
	ImageBytes                   *int64          `json:"image_bytes,omitempty"`                    
	InitialUser                  *string         `json:"initial_user,omitempty"`                   
}

type ElementConfirm struct {
	Title   *TextElement `json:"title,omitempty"`  
	Text    *TextElement `json:"text,omitempty"`   
	Confirm *TextElement `json:"confirm,omitempty"`
	Deny    *TextElement `json:"deny,omitempty"`   
	Style   *string      `json:"style,omitempty"`  
}

type TextElement struct {
	Type     *string `json:"type,omitempty"`    
	Text     *string `json:"text,omitempty"`    
	Emoji    *bool   `json:"emoji,omitempty"`   
	Verbatim *bool   `json:"verbatim,omitempty"`
}

type Filter struct {
	ExcludeExternalSharedChannels *bool `json:"exclude_external_shared_channels,omitempty"`
	ExcludeBotUsers               *bool `json:"exclude_bot_users,omitempty"`               
}

type InitialOption struct {
	Text        *TextElement `json:"text,omitempty"`       
	Value       *string      `json:"value,omitempty"`      
	Description *TextElement `json:"description,omitempty"`
	URL         *string      `json:"url,omitempty"`        
}

type Field struct {
	Title *string `json:"title,omitempty"`
	Value *string `json:"value,omitempty"`
	Short *bool   `json:"short,omitempty"`
}

type Metadata struct {
	Thumb64    *bool   `json:"thumb_64,omitempty"`   
	Thumb80    *bool   `json:"thumb_80,omitempty"`   
	Thumb160   *bool   `json:"thumb_160,omitempty"`  
	OriginalW  *int64  `json:"original_w,omitempty"` 
	OriginalH  *int64  `json:"original_h,omitempty"` 
	Thumb360_W *int64  `json:"thumb_360_w,omitempty"`
	Thumb360_H *int64  `json:"thumb_360_h,omitempty"`
	Format     *string `json:"format,omitempty"`     
	Extension  *string `json:"extension,omitempty"`  
	Rotation   *int64  `json:"rotation,omitempty"`   
	ThumbTiny  *string `json:"thumb_tiny,omitempty"` 
}

type MessageBlock struct {
	Type        *string         `json:"type,omitempty"`        
	BlockID     *string         `json:"block_id,omitempty"`    
	Text        *TextElement    `json:"text,omitempty"`        
	Accessory   *Accessory      `json:"accessory,omitempty"`   
	Elements    []FluffyElement `json:"elements,omitempty"`    
	Fallback    *string         `json:"fallback,omitempty"`    
	ImageURL    *string         `json:"image_url,omitempty"`   
	ImageWidth  *int64          `json:"image_width,omitempty"` 
	ImageHeight *int64          `json:"image_height,omitempty"`
	ImageBytes  *int64          `json:"image_bytes,omitempty"` 
	AltText     *string         `json:"alt_text,omitempty"`    
	Title       *TextElement    `json:"title,omitempty"`       
	Fields      []TextElement   `json:"fields,omitempty"`      
}

type FluffyElement struct {
	Type                         *string         `json:"type,omitempty"`                           
	ActionID                     *string         `json:"action_id,omitempty"`                      
	Text                         *TextUnion      `json:"text"`                                     
	Value                        *string         `json:"value,omitempty"`                          
	Verbatim                     *bool           `json:"verbatim,omitempty"`                       
	Emoji                        *bool           `json:"emoji,omitempty"`                          
	URL                          *string         `json:"url,omitempty"`                            
	Style                        *string         `json:"style,omitempty"`                          
	Confirm                      *ElementConfirm `json:"confirm,omitempty"`                        
	Placeholder                  *TextElement    `json:"placeholder,omitempty"`                    
	InitialChannel               *string         `json:"initial_channel,omitempty"`                
	ResponseURLEnabled           *bool           `json:"response_url_enabled,omitempty"`           
	InitialConversation          *string         `json:"initial_conversation,omitempty"`           
	DefaultToCurrentConversation *bool           `json:"default_to_current_conversation,omitempty"`
	Filter                       *Filter         `json:"filter,omitempty"`                         
	InitialDate                  *string         `json:"initial_date,omitempty"`                   
	InitialOption                *InitialOption  `json:"initial_option,omitempty"`                 
	MinQueryLength               *int64          `json:"min_query_length,omitempty"`               
	ImageURL                     *string         `json:"image_url,omitempty"`                      
	AltText                      *string         `json:"alt_text,omitempty"`                       
	Fallback                     *string         `json:"fallback,omitempty"`                       
	ImageWidth                   *int64          `json:"image_width,omitempty"`                    
	ImageHeight                  *int64          `json:"image_height,omitempty"`                   
	ImageBytes                   *int64          `json:"image_bytes,omitempty"`                    
	InitialUser                  *string         `json:"initial_user,omitempty"`                   
}

type BotProfile struct {
	ID      *string          `json:"id,omitempty"`     
	Deleted *bool            `json:"deleted,omitempty"`
	Name    *string          `json:"name,omitempty"`   
	Updated *int64           `json:"updated,omitempty"`
	AppID   *string          `json:"app_id,omitempty"` 
	Icons   *BotProfileIcons `json:"icons,omitempty"`  
	TeamID  *string          `json:"team_id,omitempty"`
}

type BotProfileIcons struct {
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
}

type MessageIcons struct {
	Emoji   *string `json:"emoji,omitempty"`   
	Image64 *string `json:"image_64,omitempty"`
}

type Root struct {
	Type            *string       `json:"type,omitempty"`             
	Subtype         *string       `json:"subtype,omitempty"`          
	Text            *string       `json:"text,omitempty"`             
	Ts              *string       `json:"ts,omitempty"`               
	Username        *string       `json:"username,omitempty"`         
	Icons           *MessageIcons `json:"icons,omitempty"`            
	BotID           *string       `json:"bot_id,omitempty"`           
	ThreadTs        *string       `json:"thread_ts,omitempty"`        
	ParentUserID    *string       `json:"parent_user_id,omitempty"`   
	ReplyCount      *int64        `json:"reply_count,omitempty"`      
	ReplyUsersCount *int64        `json:"reply_users_count,omitempty"`
	LatestReply     *string       `json:"latest_reply,omitempty"`     
	ReplyUsers      []string      `json:"reply_users,omitempty"`      
	Subscribed      *bool         `json:"subscribed,omitempty"`       
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
}

type TextUnion struct {
	String      *string
	TextElement *TextElement
}

func (x *TextUnion) UnmarshalJSON(data []byte) error {
	x.TextElement = nil
	var c TextElement
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TextElement = &c
	}
	return nil
}

func (x *TextUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TextElement != nil, x.TextElement, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
