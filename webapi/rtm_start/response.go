// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    rtmStart, err := UnmarshalRtmStart(bytes)
//    bytes, err = rtmStart.Marshal()

package rtm_start

import "encoding/json"

func UnmarshalRtmStart(data []byte) (RtmStart, error) {
	var r RtmStart
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RtmStart) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RtmStart struct {
	Ok                       *bool     `json:"ok,omitempty"`                         
	Self                     *Self     `json:"self,omitempty"`                       
	Team                     *Team     `json:"team,omitempty"`                       
	AcceptTosURL             *string   `json:"accept_tos_url,omitempty"`             
	LatestEventTs            *string   `json:"latest_event_ts,omitempty"`            
	Channels                 []Channel `json:"channels,omitempty"`                   
	Groups                   []Group   `json:"groups,omitempty"`                     
	Ims                      []IM      `json:"ims,omitempty"`                        
	CacheTs                  *int64    `json:"cache_ts,omitempty"`                   
	MobileAppRequiresUpgrade *bool     `json:"mobile_app_requires_upgrade,omitempty"`
	ReadOnlyChannels         []string  `json:"read_only_channels,omitempty"`         
	NonThreadableChannels    []string  `json:"non_threadable_channels,omitempty"`    
	ThreadOnlyChannels       []string  `json:"thread_only_channels,omitempty"`       
	CanManageSharedChannels  *bool     `json:"can_manage_shared_channels,omitempty"` 
	Subteams                 *Subteams `json:"subteams,omitempty"`                   
	DND                      *DND      `json:"dnd,omitempty"`                        
	Users                    []User    `json:"users,omitempty"`                      
	CacheVersion             *string   `json:"cache_version,omitempty"`              
	CacheTsVersion           *string   `json:"cache_ts_version,omitempty"`           
	Bots                     []Bot     `json:"bots,omitempty"`                       
	URL                      *string   `json:"url,omitempty"`                        
	IsEurope                 *bool     `json:"is_europe,omitempty"`                  
	Error                    *string   `json:"error,omitempty"`                      
	Needed                   *string   `json:"needed,omitempty"`                     
	Provided                 *string   `json:"provided,omitempty"`                   
}

type Bot struct {
	ID            *string   `json:"id,omitempty"`             
	Deleted       *bool     `json:"deleted,omitempty"`        
	Name          *string   `json:"name,omitempty"`           
	Updated       *int64    `json:"updated,omitempty"`        
	AppID         *string   `json:"app_id,omitempty"`         
	Icons         *BotIcons `json:"icons,omitempty"`          
	IsWorkflowBot *bool     `json:"is_workflow_bot,omitempty"`
	TeamID        *string   `json:"team_id,omitempty"`        
}

type BotIcons struct {
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
}

type Channel struct {
	ID                      *string  `json:"id,omitempty"`                        
	Name                    *string  `json:"name,omitempty"`                      
	IsChannel               *bool    `json:"is_channel,omitempty"`                
	Created                 *int64   `json:"created,omitempty"`                   
	IsArchived              *bool    `json:"is_archived,omitempty"`               
	IsGeneral               *bool    `json:"is_general,omitempty"`                
	Unlinked                *int64   `json:"unlinked,omitempty"`                  
	Creator                 *string  `json:"creator,omitempty"`                   
	NameNormalized          *string  `json:"name_normalized,omitempty"`           
	IsShared                *bool    `json:"is_shared,omitempty"`                 
	IsOrgShared             *bool    `json:"is_org_shared,omitempty"`             
	HasPins                 *bool    `json:"has_pins,omitempty"`                  
	IsMember                *bool    `json:"is_member,omitempty"`                 
	IsPrivate               *bool    `json:"is_private,omitempty"`                
	IsMpim                  *bool    `json:"is_mpim,omitempty"`                   
	PreviousNames           []string `json:"previous_names,omitempty"`            
	Priority                *int64   `json:"priority,omitempty"`                  
	LastRead                *string  `json:"last_read,omitempty"`                 
	Members                 []string `json:"members,omitempty"`                   
	Topic                   *Purpose `json:"topic,omitempty"`                     
	Purpose                 *Purpose `json:"purpose,omitempty"`                   
	IsGroup                 *bool    `json:"is_group,omitempty"`                  
	IsIM                    *bool    `json:"is_im,omitempty"`                     
	IsEXTShared             *bool    `json:"is_ext_shared,omitempty"`             
	SharedTeamIDS           []string `json:"shared_team_ids,omitempty"`           
	InternalTeamIDS         []string `json:"internal_team_ids,omitempty"`         
	ConnectedTeamIDS        []string `json:"connected_team_ids,omitempty"`        
	PendingShared           []string `json:"pending_shared,omitempty"`            
	PendingConnectedTeamIDS []string `json:"pending_connected_team_ids,omitempty"`
	IsPendingEXTShared      *bool    `json:"is_pending_ext_shared,omitempty"`     
	ConversationHostID      *string  `json:"conversation_host_id,omitempty"`      
}

type Purpose struct {
	Value   *string `json:"value,omitempty"`   
	Creator *string `json:"creator,omitempty"` 
	LastSet *int64  `json:"last_set,omitempty"`
}

type DND struct {
	DNDEnabled     *bool  `json:"dnd_enabled,omitempty"`      
	NextDNDStartTs *int64 `json:"next_dnd_start_ts,omitempty"`
	NextDNDEndTs   *int64 `json:"next_dnd_end_ts,omitempty"`  
	SnoozeEnabled  *bool  `json:"snooze_enabled,omitempty"`   
}

type Group struct {
	ID                 *string  `json:"id,omitempty"`                  
	Name               *string  `json:"name,omitempty"`                
	NameNormalized     *string  `json:"name_normalized,omitempty"`     
	IsGroup            *bool    `json:"is_group,omitempty"`            
	Created            *int64   `json:"created,omitempty"`             
	Creator            *string  `json:"creator,omitempty"`             
	IsArchived         *bool    `json:"is_archived,omitempty"`         
	IsMpim             *bool    `json:"is_mpim,omitempty"`             
	IsOpen             *bool    `json:"is_open,omitempty"`             
	IsReadOnly         *bool    `json:"is_read_only,omitempty"`        
	IsThreadOnly       *bool    `json:"is_thread_only,omitempty"`      
	ParentGroup        *string  `json:"parent_group,omitempty"`        
	Topic              *Purpose `json:"topic,omitempty"`               
	Purpose            *Purpose `json:"purpose,omitempty"`             
	LastRead           *string  `json:"last_read,omitempty"`           
	Latest             *Latest  `json:"latest,omitempty"`              
	UnreadCount        *int64   `json:"unread_count,omitempty"`        
	UnreadCountDisplay *int64   `json:"unread_count_display,omitempty"`
	Priority           *float64 `json:"priority,omitempty"`            
}

type Latest struct {
	ClientMsgID  *string      `json:"client_msg_id,omitempty"` 
	Type         *string      `json:"type,omitempty"`          
	Subtype      *string      `json:"subtype,omitempty"`       
	Team         *string      `json:"team,omitempty"`          
	User         *string      `json:"user,omitempty"`          
	Username     *string      `json:"username,omitempty"`      
	ParentUserID *string      `json:"parent_user_id,omitempty"`
	Text         *string      `json:"text,omitempty"`          
	Topic        *string      `json:"topic,omitempty"`         
	Root         *Root        `json:"root,omitempty"`          
	Upload       *bool        `json:"upload,omitempty"`        
	DisplayAsBot *bool        `json:"display_as_bot,omitempty"`
	BotID        *string      `json:"bot_id,omitempty"`        
	BotLink      *string      `json:"bot_link,omitempty"`      
	BotProfile   *Bot         `json:"bot_profile,omitempty"`   
	ThreadTs     *string      `json:"thread_ts,omitempty"`     
	Ts           *string      `json:"ts,omitempty"`            
	Icons        *LatestIcons `json:"icons,omitempty"`         
	Edited       *Edited      `json:"edited,omitempty"`        
}

type Edited struct {
	User *string `json:"user,omitempty"`
	Ts   *string `json:"ts,omitempty"`  
}

type LatestIcons struct {
	Emoji   *string `json:"emoji,omitempty"`   
	Image36 *string `json:"image_36,omitempty"`
	Image48 *string `json:"image_48,omitempty"`
	Image64 *string `json:"image_64,omitempty"`
	Image72 *string `json:"image_72,omitempty"`
}

type Root struct {
	Text            *string      `json:"text,omitempty"`             
	User            *string      `json:"user,omitempty"`             
	ParentUserID    *string      `json:"parent_user_id,omitempty"`   
	Username        *string      `json:"username,omitempty"`         
	Team            *string      `json:"team,omitempty"`             
	BotID           *string      `json:"bot_id,omitempty"`           
	Mrkdwn          *bool        `json:"mrkdwn,omitempty"`           
	Type            *string      `json:"type,omitempty"`             
	Subtype         *string      `json:"subtype,omitempty"`          
	ThreadTs        *string      `json:"thread_ts,omitempty"`        
	Icons           *LatestIcons `json:"icons,omitempty"`            
	BotProfile      *Bot         `json:"bot_profile,omitempty"`      
	ReplyCount      *int64       `json:"reply_count,omitempty"`      
	ReplyUsersCount *int64       `json:"reply_users_count,omitempty"`
	LatestReply     *string      `json:"latest_reply,omitempty"`     
	Subscribed      *bool        `json:"subscribed,omitempty"`       
	LastRead        *string      `json:"last_read,omitempty"`        
	UnreadCount     *int64       `json:"unread_count,omitempty"`     
	Ts              *string      `json:"ts,omitempty"`               
}

type IM struct {
	ID          *string `json:"id,omitempty"`           
	Created     *int64  `json:"created,omitempty"`      
	IsArchived  *bool   `json:"is_archived,omitempty"`  
	IsIM        *bool   `json:"is_im,omitempty"`        
	IsOrgShared *bool   `json:"is_org_shared,omitempty"`
	User        *string `json:"user,omitempty"`         
	HasPins     *bool   `json:"has_pins,omitempty"`     
	LastRead    *string `json:"last_read,omitempty"`    
	IsOpen      *bool   `json:"is_open,omitempty"`      
	Priority    *int64  `json:"priority,omitempty"`     
}

type Self struct {
	ID             *string    `json:"id,omitempty"`             
	Name           *string    `json:"name,omitempty"`           
	Prefs          *SelfPrefs `json:"prefs,omitempty"`          
	Created        *int64     `json:"created,omitempty"`        
	FirstLogin     *int64     `json:"first_login,omitempty"`    
	ManualPresence *string    `json:"manual_presence,omitempty"`
}

type SelfPrefs struct {
	UserColors                                     *string         `json:"user_colors,omitempty"`                                         
	ColorNamesInList                               *bool           `json:"color_names_in_list,omitempty"`                                 
	EmailAlerts                                    *string         `json:"email_alerts,omitempty"`                                        
	EmailAlertsSleepUntil                          *int64          `json:"email_alerts_sleep_until,omitempty"`                            
	EmailTips                                      *bool           `json:"email_tips,omitempty"`                                          
	EmailWeekly                                    *bool           `json:"email_weekly,omitempty"`                                        
	EmailOffers                                    *bool           `json:"email_offers,omitempty"`                                        
	EmailResearch                                  *bool           `json:"email_research,omitempty"`                                      
	EmailDeveloper                                 *bool           `json:"email_developer,omitempty"`                                     
	WelcomeMessageHidden                           *bool           `json:"welcome_message_hidden,omitempty"`                              
	SearchSort                                     *string         `json:"search_sort,omitempty"`                                         
	SearchFileSort                                 *string         `json:"search_file_sort,omitempty"`                                    
	SearchChannelSort                              *string         `json:"search_channel_sort,omitempty"`                                 
	SearchPeopleSort                               *string         `json:"search_people_sort,omitempty"`                                  
	ExpandInlineImgs                               *bool           `json:"expand_inline_imgs,omitempty"`                                  
	ExpandInternalInlineImgs                       *bool           `json:"expand_internal_inline_imgs,omitempty"`                         
	ExpandSnippets                                 *bool           `json:"expand_snippets,omitempty"`                                     
	PostsFormattingGuide                           *bool           `json:"posts_formatting_guide,omitempty"`                              
	SeenWelcome2                                   *bool           `json:"seen_welcome_2,omitempty"`                                      
	SeenSsbPrompt                                  *bool           `json:"seen_ssb_prompt,omitempty"`                                     
	SpacesNewXPBannerDismissed                     *bool           `json:"spaces_new_xp_banner_dismissed,omitempty"`                      
	SearchOnlyMyChannels                           *bool           `json:"search_only_my_channels,omitempty"`                             
	SearchOnlyCurrentTeam                          *bool           `json:"search_only_current_team,omitempty"`                            
	SearchHideMyChannels                           *bool           `json:"search_hide_my_channels,omitempty"`                             
	SearchOnlyShowOnline                           *bool           `json:"search_only_show_online,omitempty"`                             
	SearchHideDeactivatedUsers                     *bool           `json:"search_hide_deactivated_users,omitempty"`                       
	EmojiMode                                      *string         `json:"emoji_mode,omitempty"`                                          
	EmojiUse                                       *string         `json:"emoji_use,omitempty"`                                           
	EmojiUseOrg                                    *string         `json:"emoji_use_org,omitempty"`                                       
	HasInvited                                     *bool           `json:"has_invited,omitempty"`                                         
	HasUploaded                                    *bool           `json:"has_uploaded,omitempty"`                                        
	HasCreatedChannel                              *bool           `json:"has_created_channel,omitempty"`                                 
	HasCreatedChannelSection                       *bool           `json:"has_created_channel_section,omitempty"`                         
	HasSearched                                    *bool           `json:"has_searched,omitempty"`                                        
	SearchExcludeChannels                          *string         `json:"search_exclude_channels,omitempty"`                             
	MessagesTheme                                  *string         `json:"messages_theme,omitempty"`                                      
	WebappSpellcheck                               *bool           `json:"webapp_spellcheck,omitempty"`                                   
	NoJoinedOverlays                               *bool           `json:"no_joined_overlays,omitempty"`                                  
	NoCreatedOverlays                              *bool           `json:"no_created_overlays,omitempty"`                                 
	DropboxEnabled                                 *bool           `json:"dropbox_enabled,omitempty"`                                     
	SeenDomainInviteReminder                       *bool           `json:"seen_domain_invite_reminder,omitempty"`                         
	SeenMemberInviteReminder                       *bool           `json:"seen_member_invite_reminder,omitempty"`                         
	MuteSounds                                     *bool           `json:"mute_sounds,omitempty"`                                         
	ArrowHistory                                   *bool           `json:"arrow_history,omitempty"`                                       
	TabUIReturnSelects                             *bool           `json:"tab_ui_return_selects,omitempty"`                               
	ObeyInlineImgLimit                             *bool           `json:"obey_inline_img_limit,omitempty"`                               
	RequireAt                                      *bool           `json:"require_at,omitempty"`                                          
	SsbSpaceWindow                                 *string         `json:"ssb_space_window,omitempty"`                                    
	MACSsbBounce                                   *string         `json:"mac_ssb_bounce,omitempty"`                                      
	MACSsbBullet                                   *bool           `json:"mac_ssb_bullet,omitempty"`                                      
	ExpandNonMediaAttachments                      *bool           `json:"expand_non_media_attachments,omitempty"`                        
	ShowTyping                                     *bool           `json:"show_typing,omitempty"`                                         
	PagekeysHandled                                *bool           `json:"pagekeys_handled,omitempty"`                                    
	LastSnippetType                                *string         `json:"last_snippet_type,omitempty"`                                   
	DisplayRealNamesOverride                       *int64          `json:"display_real_names_override,omitempty"`                         
	DisplayDisplayNames                            *bool           `json:"display_display_names,omitempty"`                               
	Time24                                         *bool           `json:"time24,omitempty"`                                              
	EnterIsSpecialInTbt                            *bool           `json:"enter_is_special_in_tbt,omitempty"`                             
	MsgInputSendBtn                                *bool           `json:"msg_input_send_btn,omitempty"`                                  
	MsgInputSendBtnAutoSet                         *bool           `json:"msg_input_send_btn_auto_set,omitempty"`                         
	MsgInputStickyComposer                         *bool           `json:"msg_input_sticky_composer,omitempty"`                           
	GraphicEmoticons                               *bool           `json:"graphic_emoticons,omitempty"`                                   
	ConvertEmoticons                               *bool           `json:"convert_emoticons,omitempty"`                                   
	SsEmojis                                       *bool           `json:"ss_emojis,omitempty"`                                           
	SeenOnboardingStart                            *bool           `json:"seen_onboarding_start,omitempty"`                               
	OnboardingCancelled                            *bool           `json:"onboarding_cancelled,omitempty"`                                
	SeenOnboardingSlackbotConversation             *bool           `json:"seen_onboarding_slackbot_conversation,omitempty"`               
	SeenOnboardingChannels                         *bool           `json:"seen_onboarding_channels,omitempty"`                            
	SeenOnboardingDirectMessages                   *bool           `json:"seen_onboarding_direct_messages,omitempty"`                     
	SeenOnboardingInvites                          *bool           `json:"seen_onboarding_invites,omitempty"`                             
	SeenOnboardingSearch                           *bool           `json:"seen_onboarding_search,omitempty"`                              
	SeenOnboardingRecentMentions                   *bool           `json:"seen_onboarding_recent_mentions,omitempty"`                     
	SeenOnboardingStarredItems                     *bool           `json:"seen_onboarding_starred_items,omitempty"`                       
	SeenOnboardingPrivateGroups                    *bool           `json:"seen_onboarding_private_groups,omitempty"`                      
	SeenOnboardingBanner                           *bool           `json:"seen_onboarding_banner,omitempty"`                              
	OnboardingSlackbotConversationStep             *int64          `json:"onboarding_slackbot_conversation_step,omitempty"`               
	SetTzAutomatically                             *bool           `json:"set_tz_automatically,omitempty"`                                
	SuppressLinkWarning                            *bool           `json:"suppress_link_warning,omitempty"`                               
	SeenEmojiPackCta                               *int64          `json:"seen_emoji_pack_cta,omitempty"`                                 
	SeenEmojiPackDialog                            *bool           `json:"seen_emoji_pack_dialog,omitempty"`                              
	DNDEnabled                                     *bool           `json:"dnd_enabled,omitempty"`                                         
	DNDStartHour                                   *string         `json:"dnd_start_hour,omitempty"`                                      
	DNDEndHour                                     *string         `json:"dnd_end_hour,omitempty"`                                        
	DNDBeforeMonday                                *string         `json:"dnd_before_monday,omitempty"`                                   
	DNDAfterMonday                                 *string         `json:"dnd_after_monday,omitempty"`                                    
	DNDEnabledMonday                               *string         `json:"dnd_enabled_monday,omitempty"`                                  
	DNDBeforeTuesday                               *string         `json:"dnd_before_tuesday,omitempty"`                                  
	DNDAfterTuesday                                *string         `json:"dnd_after_tuesday,omitempty"`                                   
	DNDEnabledTuesday                              *string         `json:"dnd_enabled_tuesday,omitempty"`                                 
	DNDBeforeWednesday                             *string         `json:"dnd_before_wednesday,omitempty"`                                
	DNDAfterWednesday                              *string         `json:"dnd_after_wednesday,omitempty"`                                 
	DNDEnabledWednesday                            *string         `json:"dnd_enabled_wednesday,omitempty"`                               
	DNDBeforeThursday                              *string         `json:"dnd_before_thursday,omitempty"`                                 
	DNDAfterThursday                               *string         `json:"dnd_after_thursday,omitempty"`                                  
	DNDEnabledThursday                             *string         `json:"dnd_enabled_thursday,omitempty"`                                
	DNDBeforeFriday                                *string         `json:"dnd_before_friday,omitempty"`                                   
	DNDAfterFriday                                 *string         `json:"dnd_after_friday,omitempty"`                                    
	DNDEnabledFriday                               *string         `json:"dnd_enabled_friday,omitempty"`                                  
	DNDBeforeSaturday                              *string         `json:"dnd_before_saturday,omitempty"`                                 
	DNDAfterSaturday                               *string         `json:"dnd_after_saturday,omitempty"`                                  
	DNDEnabledSaturday                             *string         `json:"dnd_enabled_saturday,omitempty"`                                
	DNDBeforeSunday                                *string         `json:"dnd_before_sunday,omitempty"`                                   
	DNDAfterSunday                                 *string         `json:"dnd_after_sunday,omitempty"`                                    
	DNDEnabledSunday                               *string         `json:"dnd_enabled_sunday,omitempty"`                                  
	DNDDays                                        *string         `json:"dnd_days,omitempty"`                                            
	DNDWeekdaysOffAllday                           *bool           `json:"dnd_weekdays_off_allday,omitempty"`                             
	DNDCustomNewBadgeSeen                          *bool           `json:"dnd_custom_new_badge_seen,omitempty"`                           
	DNDNotificationScheduleNewBadgeSeen            *bool           `json:"dnd_notification_schedule_new_badge_seen,omitempty"`            
	CallsSurveyLastSeen                            *string         `json:"calls_survey_last_seen,omitempty"`                              
	SidebarBehavior                                *string         `json:"sidebar_behavior,omitempty"`                                    
	ChannelSort                                    *string         `json:"channel_sort,omitempty"`                                        
	SeparatePrivateChannels                        *bool           `json:"separate_private_channels,omitempty"`                           
	SeparateSharedChannels                         *bool           `json:"separate_shared_channels,omitempty"`                            
	SidebarTheme                                   *string         `json:"sidebar_theme,omitempty"`                                       
	SidebarThemeCustomValues                       *string         `json:"sidebar_theme_custom_values,omitempty"`                         
	NoInvitesWidgetInSidebar                       *bool           `json:"no_invites_widget_in_sidebar,omitempty"`                        
	NoOmniboxInChannels                            *bool           `json:"no_omnibox_in_channels,omitempty"`                              
	KKeyOmniboxAutoHideCount                       *int64          `json:"k_key_omnibox_auto_hide_count,omitempty"`                       
	ShowSidebarQuickswitcherButton                 *bool           `json:"show_sidebar_quickswitcher_button,omitempty"`                   
	EntOrgWideChannelsSidebar                      *bool           `json:"ent_org_wide_channels_sidebar,omitempty"`                       
	MarkMsgsReadImmediately                        *bool           `json:"mark_msgs_read_immediately,omitempty"`                          
	StartScrollAtOldest                            *bool           `json:"start_scroll_at_oldest,omitempty"`                              
	SnippetEditorWrapLongLines                     *bool           `json:"snippet_editor_wrap_long_lines,omitempty"`                      
	LsDisabled                                     *bool           `json:"ls_disabled,omitempty"`                                         
	FKeySearch                                     *bool           `json:"f_key_search,omitempty"`                                        
	KKeyOmnibox                                    *bool           `json:"k_key_omnibox,omitempty"`                                       
	PromptedForEmailDisabling                      *bool           `json:"prompted_for_email_disabling,omitempty"`                        
	NoMacelectronBanner                            *bool           `json:"no_macelectron_banner,omitempty"`                               
	NoMacssb1Banner                                *bool           `json:"no_macssb1_banner,omitempty"`                                   
	NoMacssb2Banner                                *bool           `json:"no_macssb2_banner,omitempty"`                                   
	NoWinssb1Banner                                *bool           `json:"no_winssb1_banner,omitempty"`                                   
	HideUserGroupInfoPane                          *bool           `json:"hide_user_group_info_pane,omitempty"`                           
	MentionsExcludeAtUserGroups                    *bool           `json:"mentions_exclude_at_user_groups,omitempty"`                     
	MentionsExcludeReactions                       *bool           `json:"mentions_exclude_reactions,omitempty"`                          
	PrivacyPolicySeen                              *bool           `json:"privacy_policy_seen,omitempty"`                                 
	EnterpriseMigrationSeen                        *bool           `json:"enterprise_migration_seen,omitempty"`                           
	SearchExcludeBots                              *bool           `json:"search_exclude_bots,omitempty"`                                 
	LoadLato2                                      *bool           `json:"load_lato_2,omitempty"`                                         
	FullerTimestamps                               *bool           `json:"fuller_timestamps,omitempty"`                                   
	LastSeenAtChannelWarning                       *int64          `json:"last_seen_at_channel_warning,omitempty"`                        
	EmojiAutocompleteBig                           *bool           `json:"emoji_autocomplete_big,omitempty"`                              
	TwoFactorAuthEnabled                           *bool           `json:"two_factor_auth_enabled,omitempty"`                             
	HideHexSwatch                                  *bool           `json:"hide_hex_swatch,omitempty"`                                     
	ShowJumperScores                               *bool           `json:"show_jumper_scores,omitempty"`                                  
	EnterpriseMdmCustomMsg                         *string         `json:"enterprise_mdm_custom_msg,omitempty"`                           
	ClientLogsPri                                  *string         `json:"client_logs_pri,omitempty"`                                     
	FlannelServerPool                              *string         `json:"flannel_server_pool,omitempty"`                                 
	MentionsExcludeAtChannels                      *bool           `json:"mentions_exclude_at_channels,omitempty"`                        
	ConfirmClearAllUnreads                         *bool           `json:"confirm_clear_all_unreads,omitempty"`                           
	ConfirmUserMarkedAway                          *bool           `json:"confirm_user_marked_away,omitempty"`                            
	BoxEnabled                                     *bool           `json:"box_enabled,omitempty"`                                         
	SeenSingleEmojiMsg                             *bool           `json:"seen_single_emoji_msg,omitempty"`                               
	ConfirmShCallStart                             *bool           `json:"confirm_sh_call_start,omitempty"`                               
	PreferredSkinTone                              *string         `json:"preferred_skin_tone,omitempty"`                                 
	ShowAllSkinTones                               *bool           `json:"show_all_skin_tones,omitempty"`                                 
	WhatsNewRead                                   *int64          `json:"whats_new_read,omitempty"`                                      
	FrecencyJumper                                 *string         `json:"frecency_jumper,omitempty"`                                     
	FrecencyEntJumper                              *string         `json:"frecency_ent_jumper,omitempty"`                                 
	FrecencyEntJumperBackup                        *string         `json:"frecency_ent_jumper_backup,omitempty"`                          
	Jumbomoji                                      *bool           `json:"jumbomoji,omitempty"`                                           
	NewxpSeenLastMessage                           *int64          `json:"newxp_seen_last_message,omitempty"`                             
	ShowMemoryInstrument                           *bool           `json:"show_memory_instrument,omitempty"`                              
	EnableUnreadView                               *bool           `json:"enable_unread_view,omitempty"`                                  
	SeenUnreadViewCoachmark                        *bool           `json:"seen_unread_view_coachmark,omitempty"`                          
	EnableReactEmojiPicker                         *bool           `json:"enable_react_emoji_picker,omitempty"`                           
	SeenCustomStatusBadge                          *bool           `json:"seen_custom_status_badge,omitempty"`                            
	SeenCustomStatusCallout                        *bool           `json:"seen_custom_status_callout,omitempty"`                          
	SeenCustomStatusExpirationBadge                *bool           `json:"seen_custom_status_expiration_badge,omitempty"`                 
	UsedCustomStatusKBShortcut                     *bool           `json:"used_custom_status_kb_shortcut,omitempty"`                      
	SeenGuestAdminSlackbotAnnouncement             *bool           `json:"seen_guest_admin_slackbot_announcement,omitempty"`              
	SeenThreadsNotificationBanner                  *bool           `json:"seen_threads_notification_banner,omitempty"`                    
	SeenNameTaggingCoachmark                       *bool           `json:"seen_name_tagging_coachmark,omitempty"`                         
	AllUnreadsSortOrder                            *string         `json:"all_unreads_sort_order,omitempty"`                              
	AllUnreadsSectionFilter                        *string         `json:"all_unreads_section_filter,omitempty"`                          
	Locale                                         *string         `json:"locale,omitempty"`                                              
	SeenIntlChannelNamesCoachmark                  *bool           `json:"seen_intl_channel_names_coachmark,omitempty"`                   
	SeenP2LocaleChangeMessage                      *int64          `json:"seen_p2_locale_change_message,omitempty"`                       
	SeenLocaleChangeMessage                        *int64          `json:"seen_locale_change_message,omitempty"`                          
	SeenJapaneseLocaleChangeMessage                *bool           `json:"seen_japanese_locale_change_message,omitempty"`                 
	SeenSharedChannelsCoachmark                    *bool           `json:"seen_shared_channels_coachmark,omitempty"`                      
	SeenSharedChannelsOptInChangeMessage           *bool           `json:"seen_shared_channels_opt_in_change_message,omitempty"`          
	HasRecentlySharedAChannel                      *bool           `json:"has_recently_shared_a_channel,omitempty"`                       
	SeenChannelBrowserAdminCoachmark               *bool           `json:"seen_channel_browser_admin_coachmark,omitempty"`                
	SeenAdministrationMenu                         *bool           `json:"seen_administration_menu,omitempty"`                            
	SeenDraftsSectionCoachmark                     *bool           `json:"seen_drafts_section_coachmark,omitempty"`                       
	SeenEmojiUpdateOverlayCoachmark                *bool           `json:"seen_emoji_update_overlay_coachmark,omitempty"`                 
	SeenSonicDeluxeToast                           *int64          `json:"seen_sonic_deluxe_toast,omitempty"`                             
	SeenWYSIWYGDeluxeToast                         *bool           `json:"seen_wysiwyg_deluxe_toast,omitempty"`                           
	SeenMarkdownPasteToast                         *int64          `json:"seen_markdown_paste_toast,omitempty"`                           
	SeenMarkdownPasteShortcut                      *int64          `json:"seen_markdown_paste_shortcut,omitempty"`                        
	SeenIaEducation                                *bool           `json:"seen_ia_education,omitempty"`                                   
	ShowIaTourRelaunch                             *int64          `json:"show_ia_tour_relaunch,omitempty"`                               
	PlainTextMode                                  *bool           `json:"plain_text_mode,omitempty"`                                     
	ShowSharedChannelsEducationBanner              *bool           `json:"show_shared_channels_education_banner,omitempty"`               
	IaSlackbotSurveyTimestamp48H                   *int64          `json:"ia_slackbot_survey_timestamp_48h,omitempty"`                    
	IaSlackbotSurveyTimestamp7D                    *int64          `json:"ia_slackbot_survey_timestamp_7d,omitempty"`                     
	AllowCallsToSetCurrentStatus                   *bool           `json:"allow_calls_to_set_current_status,omitempty"`                   
	InInteractiveMasMigrationFlow                  *bool           `json:"in_interactive_mas_migration_flow,omitempty"`                   
	SunsetInteractiveMessageViews                  *int64          `json:"sunset_interactive_message_views,omitempty"`                    
	ShdepPromoCodeSubmitted                        *bool           `json:"shdep_promo_code_submitted,omitempty"`                          
	SeenShdepSlackbotMessage                       *bool           `json:"seen_shdep_slackbot_message,omitempty"`                         
	SeenCallsInteractiveCoachmark                  *bool           `json:"seen_calls_interactive_coachmark,omitempty"`                    
	AllowCmdTabIss                                 *bool           `json:"allow_cmd_tab_iss,omitempty"`                                   
	SeenWorkflowBuilderDeluxeToast                 *bool           `json:"seen_workflow_builder_deluxe_toast,omitempty"`                  
	WorkflowBuilderIntroModalClickedThrough        *bool           `json:"workflow_builder_intro_modal_clicked_through,omitempty"`        
	WorkflowBuilderCoachmarks                      *string         `json:"workflow_builder_coachmarks,omitempty"`                         
	SeenGdriveCoachmark                            *bool           `json:"seen_gdrive_coachmark,omitempty"`                               
	SeenFirstInstallCoachmark                      *bool           `json:"seen_first_install_coachmark,omitempty"`                        
	SeenExistingInstallCoachmark                   *bool           `json:"seen_existing_install_coachmark,omitempty"`                     
	SeenLinkUnfurlCoachmark                        *bool           `json:"seen_link_unfurl_coachmark,omitempty"`                          
	OverloadedMessageEnabled                       *bool           `json:"overloaded_message_enabled,omitempty"`                          
	SeenHighlightsCoachmark                        *bool           `json:"seen_highlights_coachmark,omitempty"`                           
	SeenHighlightsArrowsCoachmark                  *bool           `json:"seen_highlights_arrows_coachmark,omitempty"`                    
	SeenHighlightsWarmWelcome                      *bool           `json:"seen_highlights_warm_welcome,omitempty"`                        
	SeenNewSearchUI                                *bool           `json:"seen_new_search_ui,omitempty"`                                  
	SeenChannelSearch                              *bool           `json:"seen_channel_search,omitempty"`                                 
	SeenPeopleSearch                               *bool           `json:"seen_people_search,omitempty"`                                  
	SeenPeopleSearchCount                          *int64          `json:"seen_people_search_count,omitempty"`                            
	DismissedScrollSearchTooltipCount              *int64          `json:"dismissed_scroll_search_tooltip_count,omitempty"`               
	LastDismissedScrollSearchTooltipTimestamp      *int64          `json:"last_dismissed_scroll_search_tooltip_timestamp,omitempty"`      
	HasUsedQuickswitcherShortcut                   *bool           `json:"has_used_quickswitcher_shortcut,omitempty"`                     
	SeenQuickswitcherShortcutTipCount              *int64          `json:"seen_quickswitcher_shortcut_tip_count,omitempty"`               
	BrowsersDismissedChannelsLowResultsEducation   *bool           `json:"browsers_dismissed_channels_low_results_education,omitempty"`   
	BrowsersSeenInitialChannelsEducation           *bool           `json:"browsers_seen_initial_channels_education,omitempty"`            
	BrowsersDismissedPeopleLowResultsEducation     *bool           `json:"browsers_dismissed_people_low_results_education,omitempty"`     
	BrowsersSeenInitialPeopleEducation             *bool           `json:"browsers_seen_initial_people_education,omitempty"`              
	BrowsersDismissedUserGroupsLowResultsEducation *bool           `json:"browsers_dismissed_user_groups_low_results_education,omitempty"`
	BrowsersSeenInitialUserGroupsEducation         *bool           `json:"browsers_seen_initial_user_groups_education,omitempty"`         
	BrowsersDismissedFilesLowResultsEducation      *bool           `json:"browsers_dismissed_files_low_results_education,omitempty"`      
	BrowsersSeenInitialFilesEducation              *bool           `json:"browsers_seen_initial_files_education,omitempty"`               
	BrowsersDismissedInitialDraftsEducation        *bool           `json:"browsers_dismissed_initial_drafts_education,omitempty"`         
	BrowsersSeenInitialDraftsEducation             *bool           `json:"browsers_seen_initial_drafts_education,omitempty"`              
	BrowsersDismissedInitialActivityEducation      *bool           `json:"browsers_dismissed_initial_activity_education,omitempty"`       
	BrowsersSeenInitialActivityEducation           *bool           `json:"browsers_seen_initial_activity_education,omitempty"`            
	BrowsersDismissedInitialSavedEducation         *bool           `json:"browsers_dismissed_initial_saved_education,omitempty"`          
	BrowsersSeenInitialSavedEducation              *bool           `json:"browsers_seen_initial_saved_education,omitempty"`               
	SeenEditMode                                   *bool           `json:"seen_edit_mode,omitempty"`                                      
	SeenEditModeEdu                                *bool           `json:"seen_edit_mode_edu,omitempty"`                                  
	A11YAnimations                                 *bool           `json:"a11y_animations,omitempty"`                                     
	SeenKeyboardShortcutsCoachmark                 *bool           `json:"seen_keyboard_shortcuts_coachmark,omitempty"`                   
	NeedsInitialPasswordSet                        *bool           `json:"needs_initial_password_set,omitempty"`                          
	LessonsEnabled                                 *bool           `json:"lessons_enabled,omitempty"`                                     
	TractorEnabled                                 *bool           `json:"tractor_enabled,omitempty"`                                     
	TractorExperimentGroup                         *string         `json:"tractor_experiment_group,omitempty"`                            
	OpenedSlackbotDm                               *bool           `json:"opened_slackbot_dm,omitempty"`                                  
	NewxpSeenHelpMessage                           *int64          `json:"newxp_seen_help_message,omitempty"`                             
	NewxpSuggestedChannels                         *string         `json:"newxp_suggested_channels,omitempty"`                            
	OnboardingComplete                             *bool           `json:"onboarding_complete,omitempty"`                                 
	WelcomePlaceState                              *string         `json:"welcome_place_state,omitempty"`                                 
	HasReceivedThreadedMessage                     *bool           `json:"has_received_threaded_message,omitempty"`                       
	OnboardingState                                *int64          `json:"onboarding_state,omitempty"`                                    
	WhocanseethisDmMpdmBadge                       *bool           `json:"whocanseethis_dm_mpdm_badge,omitempty"`                         
	HighlightWords                                 *string         `json:"highlight_words,omitempty"`                                     
	ThreadsEverything                              *bool           `json:"threads_everything,omitempty"`                                  
	NoTextInNotifications                          *bool           `json:"no_text_in_notifications,omitempty"`                            
	PushShowPreview                                *bool           `json:"push_show_preview,omitempty"`                                   
	GrowlsEnabled                                  *bool           `json:"growls_enabled,omitempty"`                                      
	AllChannelsLoud                                *bool           `json:"all_channels_loud,omitempty"`                                   
	PushDmAlert                                    *bool           `json:"push_dm_alert,omitempty"`                                       
	PushMentionAlert                               *bool           `json:"push_mention_alert,omitempty"`                                  
	PushEverything                                 *bool           `json:"push_everything,omitempty"`                                     
	PushIdleWait                                   *int64          `json:"push_idle_wait,omitempty"`                                      
	PushSound                                      *string         `json:"push_sound,omitempty"`                                          
	NewMsgSnd                                      *string         `json:"new_msg_snd,omitempty"`                                         
	PushLoudChannels                               *string         `json:"push_loud_channels,omitempty"`                                  
	PushMentionChannels                            *string         `json:"push_mention_channels,omitempty"`                               
	PushLoudChannelsSet                            *string         `json:"push_loud_channels_set,omitempty"`                              
	LoudChannels                                   *string         `json:"loud_channels,omitempty"`                                       
	NeverChannels                                  *string         `json:"never_channels,omitempty"`                                      
	LoudChannelsSet                                *string         `json:"loud_channels_set,omitempty"`                                   
	AtChannelSuppressedChannels                    *string         `json:"at_channel_suppressed_channels,omitempty"`                      
	PushAtChannelSuppressedChannels                *string         `json:"push_at_channel_suppressed_channels,omitempty"`                 
	MutedChannels                                  *string         `json:"muted_channels,omitempty"`                                      
	AllNotificationsPrefs                          *string         `json:"all_notifications_prefs,omitempty"`                             
	GrowthMsgLimitApproachingCtaCount              *int64          `json:"growth_msg_limit_approaching_cta_count,omitempty"`              
	GrowthMsgLimitApproachingCtaTs                 *int64          `json:"growth_msg_limit_approaching_cta_ts,omitempty"`                 
	GrowthMsgLimitReachedCtaCount                  *int64          `json:"growth_msg_limit_reached_cta_count,omitempty"`                  
	GrowthMsgLimitReachedCtaLastTs                 *int64          `json:"growth_msg_limit_reached_cta_last_ts,omitempty"`                
	GrowthMsgLimitLongReachedCtaCount              *int64          `json:"growth_msg_limit_long_reached_cta_count,omitempty"`             
	GrowthMsgLimitLongReachedCtaLastTs             *int64          `json:"growth_msg_limit_long_reached_cta_last_ts,omitempty"`           
	GrowthMsgLimitSixtyDayBannerCtaCount           *int64          `json:"growth_msg_limit_sixty_day_banner_cta_count,omitempty"`         
	GrowthMsgLimitSixtyDayBannerCtaLastTs          *int64          `json:"growth_msg_limit_sixty_day_banner_cta_last_ts,omitempty"`       
	GrowthAllBannersPrefs                          *string         `json:"growth_all_banners_prefs,omitempty"`                            
	AnalyticsUpsellCoachmarkSeen                   *bool           `json:"analytics_upsell_coachmark_seen,omitempty"`                     
	SeenAppSpaceCoachmark                          *bool           `json:"seen_app_space_coachmark,omitempty"`                            
	SeenAppSpaceTutorial                           *bool           `json:"seen_app_space_tutorial,omitempty"`                             
	DismissedAppLauncherWelcome                    *bool           `json:"dismissed_app_launcher_welcome,omitempty"`                      
	DismissedAppLauncherLimit                      *bool           `json:"dismissed_app_launcher_limit,omitempty"`                        
	Purchaser                                      *bool           `json:"purchaser,omitempty"`                                           
	ShowEntOnboarding                              *bool           `json:"show_ent_onboarding,omitempty"`                                 
	FoldersEnabled                                 *bool           `json:"folders_enabled,omitempty"`                                     
	FolderData                                     *string         `json:"folder_data,omitempty"`                                         
	SeenCorporateExportAlert                       *bool           `json:"seen_corporate_export_alert,omitempty"`                         
	ShowAutocompleteHelp                           *int64          `json:"show_autocomplete_help,omitempty"`                              
	DeprecationToastLastSeen                       *int64          `json:"deprecation_toast_last_seen,omitempty"`                         
	DeprecationModalLastSeen                       *int64          `json:"deprecation_modal_last_seen,omitempty"`                         
	Iap1Lab                                        *int64          `json:"iap1_lab,omitempty"`                                            
	IaTopNavTheme                                  *string         `json:"ia_top_nav_theme,omitempty"`                                    
	IaPlatformActionsLab                           *int64          `json:"ia_platform_actions_lab,omitempty"`                             
	ActivityView                                   *string         `json:"activity_view,omitempty"`                                       
	SavedView                                      *string         `json:"saved_view,omitempty"`                                          
	SeenFloatingSidebarCoachmark                   *bool           `json:"seen_floating_sidebar_coachmark,omitempty"`                     
	FailoverProxyCheckCompleted                    *int64          `json:"failover_proxy_check_completed,omitempty"`                      
	ChimeAccessCheckCompleted                      *int64          `json:"chime_access_check_completed,omitempty"`                        
	MXCalendarType                                 *string         `json:"mx_calendar_type,omitempty"`                                    
	EdgeUploadProxyCheckCompleted                  *int64          `json:"edge_upload_proxy_check_completed,omitempty"`                   
	AppSubdomainCheckCompleted                     *int64          `json:"app_subdomain_check_completed,omitempty"`                       
	AddPromptInteracted                            *bool           `json:"add_prompt_interacted,omitempty"`                               
	AddAppsPromptDismissed                         *bool           `json:"add_apps_prompt_dismissed,omitempty"`                           
	AddChannelPromptDismissed                      *bool           `json:"add_channel_prompt_dismissed,omitempty"`                        
	ChannelSidebarHideInvite                       *bool           `json:"channel_sidebar_hide_invite,omitempty"`                         
	ChannelSidebarHideBrowseDmsLink                *bool           `json:"channel_sidebar_hide_browse_dms_link,omitempty"`                
	InProdSurveysEnabled                           *bool           `json:"in_prod_surveys_enabled,omitempty"`                             
	DismissedInstalledAppDmSuggestions             *string         `json:"dismissed_installed_app_dm_suggestions,omitempty"`              
	SeenContextualMessageShortcutsModal            *bool           `json:"seen_contextual_message_shortcuts_modal,omitempty"`             
	SeenMessageNavigationEducationalToast          *bool           `json:"seen_message_navigation_educational_toast,omitempty"`           
	ContextualMessageShortcutsModalWasSeen         *bool           `json:"contextual_message_shortcuts_modal_was_seen,omitempty"`         
	MessageNavigationToastWasSeen                  *bool           `json:"message_navigation_toast_was_seen,omitempty"`                   
	UpToBrowseKBShortcut                           *bool           `json:"up_to_browse_kb_shortcut,omitempty"`                            
	ChannelSections                                *string         `json:"channel_sections,omitempty"`                                    
	ShowQuickReactions                             *bool           `json:"show_quick_reactions,omitempty"`                                
	HasReceivedMentionOrReaction                   *bool           `json:"has_received_mention_or_reaction,omitempty"`                    
	HasStarredItem                                 *bool           `json:"has_starred_item,omitempty"`                                    
	HasDraftedMessage                              *bool           `json:"has_drafted_message,omitempty"`                                 
	EnableMentionsAndReactionsView                 *bool           `json:"enable_mentions_and_reactions_view,omitempty"`                  
	EnableSavedItemsView                           *bool           `json:"enable_saved_items_view,omitempty"`                             
	EnableAllDmsView                               *bool           `json:"enable_all_dms_view,omitempty"`                                 
	EnableChannelBrowserView                       *bool           `json:"enable_channel_browser_view,omitempty"`                         
	EnableFileBrowserView                          *bool           `json:"enable_file_browser_view,omitempty"`                            
	EnablePeopleBrowserView                        *bool           `json:"enable_people_browser_view,omitempty"`                          
	EnableAppBrowserView                           *bool           `json:"enable_app_browser_view,omitempty"`                             
	ReachedAllDmsDisclosure                        *bool           `json:"reached_all_dms_disclosure,omitempty"`                          
	HasAcknowledgedShortcutSpeedbump               *bool           `json:"has_acknowledged_shortcut_speedbump,omitempty"`                 
	Tz                                             *string         `json:"tz,omitempty"`                                                  
	LocalesEnabled                                 *LocalesEnabled `json:"locales_enabled,omitempty"`                                     
	JoinerNotificationsMuted                       *bool           `json:"joiner_notifications_muted,omitempty"`                          
	InviteAcceptedNotificationsMuted               *bool           `json:"invite_accepted_notifications_muted,omitempty"`                 
	SuppressExternalInvitesFromComposeWarning      *bool           `json:"suppress_external_invites_from_compose_warning,omitempty"`      
	FilePickerVariant                              *int64          `json:"file_picker_variant,omitempty"`                                 
	HelpModalOpenTimestamp                         *int64          `json:"help_modal_open_timestamp,omitempty"`                           
	HelpModalConsultBannerDismissed                *bool           `json:"help_modal_consult_banner_dismissed,omitempty"`                 
	SeenChannelEmailTooltip                        *bool           `json:"seen_channel_email_tooltip,omitempty"`                          
	JoinCallsDeviceSettings                        *string         `json:"join_calls_device_settings,omitempty"`                          
	A11YDyslexic                                   *bool           `json:"a11y_dyslexic,omitempty"`                                       
	ConnectDmEarlyAccess                           *bool           `json:"connect_dm_early_access,omitempty"`                             
	SeenConnectDmCoachmark                         *bool           `json:"seen_connect_dm_coachmark,omitempty"`                           
	XwsSidebarVariant                              *int64          `json:"xws_sidebar_variant,omitempty"`                                 
}

type LocalesEnabled struct {
	DeDE *string `json:"de-DE,omitempty"`
	EnGB *string `json:"en-GB,omitempty"`
	EnUS *string `json:"en-US,omitempty"`
	EsES *string `json:"es-ES,omitempty"`
	EsLA *string `json:"es-LA,omitempty"`
	FrFR *string `json:"fr-FR,omitempty"`
	PtBR *string `json:"pt-BR,omitempty"`
	JaJP *string `json:"ja-JP,omitempty"`
}

type Subteams struct {
	Self []string `json:"self,omitempty"`
	All  []All    `json:"all,omitempty"` 
}

type All struct {
	ID                  *string   `json:"id,omitempty"`                   
	TeamID              *string   `json:"team_id,omitempty"`              
	IsUsergroup         *bool     `json:"is_usergroup,omitempty"`         
	IsSubteam           *bool     `json:"is_subteam,omitempty"`           
	Name                *string   `json:"name,omitempty"`                 
	Description         *string   `json:"description,omitempty"`          
	Handle              *string   `json:"handle,omitempty"`               
	IsExternal          *bool     `json:"is_external,omitempty"`          
	DateCreate          *int64    `json:"date_create,omitempty"`          
	DateUpdate          *int64    `json:"date_update,omitempty"`          
	DateDelete          *int64    `json:"date_delete,omitempty"`          
	AutoProvision       *bool     `json:"auto_provision,omitempty"`       
	EnterpriseSubteamID *string   `json:"enterprise_subteam_id,omitempty"`
	CreatedBy           *string   `json:"created_by,omitempty"`           
	UpdatedBy           *string   `json:"updated_by,omitempty"`           
	Prefs               *AllPrefs `json:"prefs,omitempty"`                
	UserCount           *int64    `json:"user_count,omitempty"`           
	ChannelCount        *int64    `json:"channel_count,omitempty"`        
}

type AllPrefs struct {
	Channels []string `json:"channels,omitempty"`
	Groups   []Group  `json:"groups,omitempty"`  
}

type Team struct {
	ID                  *string    `json:"id,omitempty"`                   
	Name                *string    `json:"name,omitempty"`                 
	EmailDomain         *string    `json:"email_domain,omitempty"`         
	Domain              *string    `json:"domain,omitempty"`               
	MsgEditWindowMins   *int64     `json:"msg_edit_window_mins,omitempty"` 
	Prefs               *TeamPrefs `json:"prefs,omitempty"`                
	Icon                *Icon      `json:"icon,omitempty"`                 
	OverStorageLimit    *bool      `json:"over_storage_limit,omitempty"`   
	MessagesCount       *int64     `json:"messages_count,omitempty"`       
	Plan                *string    `json:"plan,omitempty"`                 
	OnboardingChannelID *string    `json:"onboarding_channel_id,omitempty"`
	DateCreate          *int64     `json:"date_create,omitempty"`          
	LimitTs             *int64     `json:"limit_ts,omitempty"`             
	AvatarBaseURL       *string    `json:"avatar_base_url,omitempty"`      
}

type Icon struct {
	Image34       *string `json:"image_34,omitempty"`      
	Image44       *string `json:"image_44,omitempty"`      
	Image68       *string `json:"image_68,omitempty"`      
	Image88       *string `json:"image_88,omitempty"`      
	Image102      *string `json:"image_102,omitempty"`     
	Image132      *string `json:"image_132,omitempty"`     
	Image230      *string `json:"image_230,omitempty"`     
	ImageOriginal *string `json:"image_original,omitempty"`
}

type TeamPrefs struct {
	DefaultChannels                             []string       `json:"default_channels,omitempty"`                                  
	AllowCalls                                  *bool          `json:"allow_calls,omitempty"`                                       
	DisplayEmailAddresses                       *bool          `json:"display_email_addresses,omitempty"`                           
	GdriveEnabledTeam                           *bool          `json:"gdrive_enabled_team,omitempty"`                               
	AllUsersCanPurchase                         *bool          `json:"all_users_can_purchase,omitempty"`                            
	EnableSharedChannels                        *int64         `json:"enable_shared_channels,omitempty"`                            
	CanReceiveSharedChannelsInvites             *bool          `json:"can_receive_shared_channels_invites,omitempty"`               
	DropboxLegacyPicker                         *bool          `json:"dropbox_legacy_picker,omitempty"`                             
	AppWhitelistEnabled                         *bool          `json:"app_whitelist_enabled,omitempty"`                             
	WhoCanManageIntegrations                    *WhoCan        `json:"who_can_manage_integrations,omitempty"`                       
	WelcomePlaceEnabled                         *bool          `json:"welcome_place_enabled,omitempty"`                             
	MsgEditWindowMins                           *int64         `json:"msg_edit_window_mins,omitempty"`                              
	AllowMessageDeletion                        *bool          `json:"allow_message_deletion,omitempty"`                            
	Locale                                      *string        `json:"locale,omitempty"`                                            
	SlackbotResponsesDisabled                   *bool          `json:"slackbot_responses_disabled,omitempty"`                       
	HideReferers                                *bool          `json:"hide_referers,omitempty"`                                     
	CallingAppName                              *string        `json:"calling_app_name,omitempty"`                                  
	CallsApps                                   *CallsApps     `json:"calls_apps,omitempty"`                                        
	AllowCallsInteractiveScreenSharing          *bool          `json:"allow_calls_interactive_screen_sharing,omitempty"`            
	DisplayRealNames                            *bool          `json:"display_real_names,omitempty"`                                
	WhoCanAtEveryone                            *string        `json:"who_can_at_everyone,omitempty"`                               
	WhoCanAtChannel                             *string        `json:"who_can_at_channel,omitempty"`                                
	WhoCanCreateChannels                        *string        `json:"who_can_create_channels,omitempty"`                           
	WhoCanArchiveChannels                       *string        `json:"who_can_archive_channels,omitempty"`                          
	WhoCanCreateGroups                          *string        `json:"who_can_create_groups,omitempty"`                             
	WhoCanManageChannelPostingPrefs             *string        `json:"who_can_manage_channel_posting_prefs,omitempty"`              
	WhoCanPostGeneral                           *string        `json:"who_can_post_general,omitempty"`                              
	WhoCanKickChannels                          *string        `json:"who_can_kick_channels,omitempty"`                             
	WhoCanKickGroups                            *string        `json:"who_can_kick_groups,omitempty"`                               
	WorkflowBuilderEnabled                      *bool          `json:"workflow_builder_enabled,omitempty"`                          
	WhoCanViewMessageActivity                   *WhoCan        `json:"who_can_view_message_activity,omitempty"`                     
	WorkflowExtensionStepsBetaOptIn             *bool          `json:"workflow_extension_steps_beta_opt_in,omitempty"`              
	ChannelEmailAddressesEnabled                *bool          `json:"channel_email_addresses_enabled,omitempty"`                   
	RetentionType                               *int64         `json:"retention_type,omitempty"`                                    
	RetentionDuration                           *int64         `json:"retention_duration,omitempty"`                                
	GroupRetentionType                          *int64         `json:"group_retention_type,omitempty"`                              
	GroupRetentionDuration                      *int64         `json:"group_retention_duration,omitempty"`                          
	DmRetentionType                             *int64         `json:"dm_retention_type,omitempty"`                                 
	DmRetentionDuration                         *int64         `json:"dm_retention_duration,omitempty"`                             
	FileRetentionType                           *int64         `json:"file_retention_type,omitempty"`                               
	FileRetentionDuration                       *int64         `json:"file_retention_duration,omitempty"`                           
	AllowRetentionOverride                      *bool          `json:"allow_retention_override,omitempty"`                          
	AllowAdminRetentionOverride                 *int64         `json:"allow_admin_retention_override,omitempty"`                    
	DefaultRxns                                 []string       `json:"default_rxns,omitempty"`                                      
	ComplianceExportStart                       *int64         `json:"compliance_export_start,omitempty"`                           
	WarnBeforeAtChannel                         *string        `json:"warn_before_at_channel,omitempty"`                            
	DisallowPublicFileUrls                      *bool          `json:"disallow_public_file_urls,omitempty"`                         
	WhoCanCreateDeleteUserGroups                *string        `json:"who_can_create_delete_user_groups,omitempty"`                 
	WhoCanEditUserGroups                        *string        `json:"who_can_edit_user_groups,omitempty"`                          
	WhoCanChangeTeamProfile                     *string        `json:"who_can_change_team_profile,omitempty"`                       
	SubteamsAutoCreateOwner                     *bool          `json:"subteams_auto_create_owner,omitempty"`                        
	SubteamsAutoCreateAdmin                     *bool          `json:"subteams_auto_create_admin,omitempty"`                        
	Discoverable                                *string        `json:"discoverable,omitempty"`                                      
	DNDDays                                     *string        `json:"dnd_days,omitempty"`                                          
	InvitesOnlyAdmins                           *bool          `json:"invites_only_admins,omitempty"`                               
	InviteRequestsEnabled                       *bool          `json:"invite_requests_enabled,omitempty"`                           
	DisableFileUploads                          *string        `json:"disable_file_uploads,omitempty"`                              
	DisableFileEditing                          *bool          `json:"disable_file_editing,omitempty"`                              
	DisableFileDeleting                         *bool          `json:"disable_file_deleting,omitempty"`                             
	FileLimitWhitelisted                        *bool          `json:"file_limit_whitelisted,omitempty"`                            
	UsesCustomizedCustomStatusPresets           *bool          `json:"uses_customized_custom_status_presets,omitempty"`             
	DisableEmailIngestion                       *bool          `json:"disable_email_ingestion,omitempty"`                           
	WhoCanManageGuests                          *WhoCan        `json:"who_can_manage_guests,omitempty"`                             
	WhoCanCreateSharedChannels                  *string        `json:"who_can_create_shared_channels,omitempty"`                    
	WhoCanManageSharedChannels                  *WhoCan        `json:"who_can_manage_shared_channels,omitempty"`                    
	WhoCanPostInSharedChannels                  *WhoCan        `json:"who_can_post_in_shared_channels,omitempty"`                   
	WhoCanManageEXTSharedChannels               *WhoCan        `json:"who_can_manage_ext_shared_channels,omitempty"`                
	WhoCanDmAnyone                              *WhoCan        `json:"who_can_dm_anyone,omitempty"`                                 
	BoxAppInstalled                             *bool          `json:"box_app_installed,omitempty"`                                 
	OnedriveAppInstalled                        *bool          `json:"onedrive_app_installed,omitempty"`                            
	OnedriveEnabledTeam                         *bool          `json:"onedrive_enabled_team,omitempty"`                             
	FilepickerAppFirstInstall                   *bool          `json:"filepicker_app_first_install,omitempty"`                      
	UseBrowserPicker                            *bool          `json:"use_browser_picker,omitempty"`                                
	ReceivedEscRouteToChannelAwarenessMessage   *bool          `json:"received_esc_route_to_channel_awareness_message,omitempty"`   
	WhoCanApproveEXTSharedChannelInvites        *WhoCan        `json:"who_can_approve_ext_shared_channel_invites,omitempty"`        
	WhoCanCreateEXTSharedChannelInvites         *WhoCan        `json:"who_can_create_ext_shared_channel_invites,omitempty"`         
	EnterpriseDefaultChannels                   []string       `json:"enterprise_default_channels,omitempty"`                       
	EnterpriseHasCorporateExports               *bool          `json:"enterprise_has_corporate_exports,omitempty"`                  
	EnterpriseMandatoryChannels                 []string       `json:"enterprise_mandatory_channels,omitempty"`                     
	EnterpriseMdmDisableFileDownload            *bool          `json:"enterprise_mdm_disable_file_download,omitempty"`              
	MobilePasscodeTimeoutInSeconds              *int64         `json:"mobile_passcode_timeout_in_seconds,omitempty"`                
	NotificationRedactionType                   *string        `json:"notification_redaction_type,omitempty"`                       
	HasComplianceExport                         *bool          `json:"has_compliance_export,omitempty"`                             
	HasHipaaCompliance                          *bool          `json:"has_hipaa_compliance,omitempty"`                              
	SelfServeSelect                             *bool          `json:"self_serve_select,omitempty"`                                 
	LoudChannelMentionsLimit                    *int64         `json:"loud_channel_mentions_limit,omitempty"`                       
	ShowJoinLeave                               *bool          `json:"show_join_leave,omitempty"`                                   
	WhoCanManagePublicChannels                  *WhoCanManageP `json:"who_can_manage_public_channels,omitempty"`                    
	WhoCanManagePrivateChannels                 *WhoCanManageP `json:"who_can_manage_private_channels,omitempty"`                   
	WhoCanManagePrivateChannelsAtWorkspaceLevel *WhoCanManageP `json:"who_can_manage_private_channels_at_workspace_level,omitempty"`
	EnterpriseMobileDeviceCheck                 *bool          `json:"enterprise_mobile_device_check,omitempty"`                    
	DefaultChannelCreationEnabled               *bool          `json:"default_channel_creation_enabled,omitempty"`                  
	GgEnabled                                   *bool          `json:"gg_enabled,omitempty"`                                        
	CreatedWithGoogle                           *bool          `json:"created_with_google,omitempty"`                               
	HasSeenPartnerPromo                         *bool          `json:"has_seen_partner_promo,omitempty"`                            
	DisableSidebarConnectPrompts                []string       `json:"disable_sidebar_connect_prompts,omitempty"`                   
	DisableSidebarInstallPrompts                []string       `json:"disable_sidebar_install_prompts,omitempty"`                   
	BlockFileDownload                           *bool          `json:"block_file_download,omitempty"`                               
	SingleUserExports                           *bool          `json:"single_user_exports,omitempty"`                               
	AppManagementApps                           []string       `json:"app_management_apps,omitempty"`                               
	NTLMCredentialDomains                       *string        `json:"ntlm_credential_domains,omitempty"`                           
	DNDEnabled                                  *bool          `json:"dnd_enabled,omitempty"`                                       
	DNDStartHour                                *string        `json:"dnd_start_hour,omitempty"`                                    
	DNDEndHour                                  *string        `json:"dnd_end_hour,omitempty"`                                      
	DNDBeforeMonday                             *string        `json:"dnd_before_monday,omitempty"`                                 
	DNDAfterMonday                              *string        `json:"dnd_after_monday,omitempty"`                                  
	DNDBeforeTuesday                            *string        `json:"dnd_before_tuesday,omitempty"`                                
	DNDAfterTuesday                             *string        `json:"dnd_after_tuesday,omitempty"`                                 
	DNDBeforeWednesday                          *string        `json:"dnd_before_wednesday,omitempty"`                              
	DNDAfterWednesday                           *string        `json:"dnd_after_wednesday,omitempty"`                               
	DNDBeforeThursday                           *string        `json:"dnd_before_thursday,omitempty"`                               
	DNDAfterThursday                            *string        `json:"dnd_after_thursday,omitempty"`                                
	DNDBeforeFriday                             *string        `json:"dnd_before_friday,omitempty"`                                 
	DNDAfterFriday                              *string        `json:"dnd_after_friday,omitempty"`                                  
	DNDBeforeSaturday                           *string        `json:"dnd_before_saturday,omitempty"`                               
	DNDAfterSaturday                            *string        `json:"dnd_after_saturday,omitempty"`                                
	DNDBeforeSunday                             *string        `json:"dnd_before_sunday,omitempty"`                                 
	DNDAfterSunday                              *string        `json:"dnd_after_sunday,omitempty"`                                  
	DNDEnabledMonday                            *string        `json:"dnd_enabled_monday,omitempty"`                                
	DNDEnabledTuesday                           *string        `json:"dnd_enabled_tuesday,omitempty"`                               
	DNDEnabledWednesday                         *string        `json:"dnd_enabled_wednesday,omitempty"`                             
	DNDEnabledThursday                          *string        `json:"dnd_enabled_thursday,omitempty"`                              
	DNDEnabledFriday                            *string        `json:"dnd_enabled_friday,omitempty"`                                
	DNDEnabledSaturday                          *string        `json:"dnd_enabled_saturday,omitempty"`                              
	DNDEnabledSunday                            *string        `json:"dnd_enabled_sunday,omitempty"`                                
	DNDWeekdaysOffAllday                        *bool          `json:"dnd_weekdays_off_allday,omitempty"`                           
	CustomStatusPresets                         [][]string     `json:"custom_status_presets,omitempty"`                             
	CustomStatusDefaultEmoji                    *string        `json:"custom_status_default_emoji,omitempty"`                       
	AuthMode                                    *string        `json:"auth_mode,omitempty"`                                         
	WhoCanCreateWorkflows                       *WhoCan        `json:"who_can_create_workflows,omitempty"`                          
	WorkflowsWebhookTriggerEnabled              *bool          `json:"workflows_webhook_trigger_enabled,omitempty"`                 
	WorkflowsExportCSVEnabled                   *bool          `json:"workflows_export_csv_enabled,omitempty"`                      
	WhoCanCreateChannelEmailAddresses           *WhoCan        `json:"who_can_create_channel_email_addresses,omitempty"`            
	InvitesLimit                                *bool          `json:"invites_limit,omitempty"`                                     
	MemberAnalyticsDisabled                     *bool          `json:"member_analytics_disabled,omitempty"`                         
	CallsLocations                              []string       `json:"calls_locations,omitempty"`                                   
	WorkflowExtensionStepsEnabled               *bool          `json:"workflow_extension_steps_enabled,omitempty"`                  
}

type CallsApps struct {
	Audio               []string `json:"audio,omitempty"`                
	Video               []Video  `json:"video,omitempty"`                
	ProfileFieldOptions []string `json:"profile_field_options,omitempty"`
}

type Video struct {
	ID    *string `json:"id,omitempty"`   
	Name  *string `json:"name,omitempty"` 
	Image *string `json:"image,omitempty"`
}

type WhoCan struct {
	Type []string `json:"type,omitempty"`
}

type WhoCanManageP struct {
	User []string `json:"user,omitempty"`
	Type []string `json:"type,omitempty"`
}

type User struct {
	ID                *string  `json:"id,omitempty"`                 
	TeamID            *string  `json:"team_id,omitempty"`            
	Name              *string  `json:"name,omitempty"`               
	Deleted           *bool    `json:"deleted,omitempty"`            
	Color             *string  `json:"color,omitempty"`              
	RealName          *string  `json:"real_name,omitempty"`          
	Tz                *string  `json:"tz,omitempty"`                 
	TzLabel           *string  `json:"tz_label,omitempty"`           
	TzOffset          *int64   `json:"tz_offset,omitempty"`          
	Profile           *Profile `json:"profile,omitempty"`            
	IsAdmin           *bool    `json:"is_admin,omitempty"`           
	IsOwner           *bool    `json:"is_owner,omitempty"`           
	IsPrimaryOwner    *bool    `json:"is_primary_owner,omitempty"`   
	IsRestricted      *bool    `json:"is_restricted,omitempty"`      
	IsUltraRestricted *bool    `json:"is_ultra_restricted,omitempty"`
	IsBot             *bool    `json:"is_bot,omitempty"`             
	IsAppUser         *bool    `json:"is_app_user,omitempty"`        
	Updated           *int64   `json:"updated,omitempty"`            
	Presence          *string  `json:"presence,omitempty"`           
	IsWorkflowBot     *bool    `json:"is_workflow_bot,omitempty"`    
}

type Profile struct {
	Title                 *string          `json:"title,omitempty"`                  
	Phone                 *string          `json:"phone,omitempty"`                  
	Skype                 *string          `json:"skype,omitempty"`                  
	RealName              *string          `json:"real_name,omitempty"`              
	RealNameNormalized    *string          `json:"real_name_normalized,omitempty"`   
	DisplayName           *string          `json:"display_name,omitempty"`           
	DisplayNameNormalized *string          `json:"display_name_normalized,omitempty"`
	Fields                map[string]Field `json:"fields,omitempty"`                 
	StatusText            *string          `json:"status_text,omitempty"`            
	StatusEmoji           *string          `json:"status_emoji,omitempty"`           
	StatusExpiration      *int64           `json:"status_expiration,omitempty"`      
	AvatarHash            *string          `json:"avatar_hash,omitempty"`            
	ImageOriginal         *string          `json:"image_original,omitempty"`         
	IsCustomImage         *bool            `json:"is_custom_image,omitempty"`        
	Email                 *string          `json:"email,omitempty"`                  
	FirstName             *string          `json:"first_name,omitempty"`             
	LastName              *string          `json:"last_name,omitempty"`              
	Image24               *string          `json:"image_24,omitempty"`               
	Image32               *string          `json:"image_32,omitempty"`               
	Image48               *string          `json:"image_48,omitempty"`               
	Image72               *string          `json:"image_72,omitempty"`               
	Image192              *string          `json:"image_192,omitempty"`              
	Image512              *string          `json:"image_512,omitempty"`              
	Image1024             *string          `json:"image_1024,omitempty"`             
	StatusTextCanonical   *string          `json:"status_text_canonical,omitempty"`  
	Team                  *string          `json:"team,omitempty"`                   
	APIAppID              *string          `json:"api_app_id,omitempty"`             
	BotID                 *string          `json:"bot_id,omitempty"`                 
	AlwaysActive          *bool            `json:"always_active,omitempty"`          
	GuestInvitedBy        *string          `json:"guest_invited_by,omitempty"`       
}

type Field struct {
	Value *string `json:"value,omitempty"`
	Alt   *string `json:"alt,omitempty"`  
}
