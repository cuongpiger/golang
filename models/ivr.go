package models

import "time"

/*
Key represents a key on mobile keyboard.
*/
type Key string

const (
	MainKey  Key = "main" // `main` key, used for the main menu in which users can select options or navigate through the IVR controller
	StarKey  Key = "*"    // `*` key, used for the star key in the IVR controller
	HashKey  Key = "#"    // `#` key, used for the hash key in the IVR controller
	ZeroKey  Key = "0"    // `0` key, used for the zero key in the IVR controller
	OneKey   Key = "1"    // `1` key, used for the one key in the IVR controller
	TwoKey   Key = "2"    // `2` key, used for the two key in the IVR controller
	ThreeKey Key = "3"    // `3` key, used for the three key in the IVR controller
	FourKey  Key = "4"    // `4` key, used for the four key in the IVR controller
	FiveKey  Key = "5"    // `5` key, used for the five key in the IVR controller
	SixKey   Key = "6"    // `6` key, used for the six key in the IVR controller
	SevenKey Key = "7"    // `7` key, used for the seven key in the IVR controller
	EightKey Key = "8"    // `8` key, used for the eight key in the IVR controller
	NineKey  Key = "9"    // `9` key, used for the nine key in the IVR controller
)

/*
Action represents an action that can be performed in the IVR controller.
*/
type Action string

const (
	RouteAction       Action = "route"        // `route` action, used to route the call to a specific destination in the IVR controller
	PlaySoundAction   Action = "play_sound"   // `play_sound` action, used to play a sound file in the IVR controller
	ReturnAction      Action = "return"       // `return` action, used to return to the previous menu in the IVR controller
	RepeatSoundAction Action = "repeat_sound" // `repeat_announcement` action, used to repeat the current announcement in the IVR controller
	EndCallAction     Action = "end_call"     // `end_call` action, used to end the call in the IVR controller
)

/*
Priority represents the priority of an IVR node.
*/
type Priority string

const (
	LowPriority    Priority = "low"    // `low` priority, used for low priority IVR nodes
	MediumPriority Priority = "medium" // `medium` priority, used for medium priority IVR nodes
	HighPriority   Priority = "high"   // `high` priority, used for high priority IVR nodes
)

type Sound struct {
	SoundPath      string `json:"sound_path,omitempty"` // Path to the sound file
	Duration       int    `json:"duration,omitempty"`
	RepeatCount    int    `json:"repeat_count,omitempty"`           // Number of times to repeat the sound
	RepeatInterval int    `json:"repeat_interval,omitempty"` //  -1 is infinite, 0 is no repeat, >0 is number of repeats
}

type ListSounds []Sound

type IVRContent struct {
	Key         Key          `json:"key,omitempty"`
	Action      Action       `json:"action,omitempty"`
	Description string       `json:"description,omitempty"`
	Value       string       `json:"value,omitempty"`
	Priority    Priority     `json:"priority,omitempty"`
	Sounds      ListSounds   `json:"sounds,omitempty"`
	Options     []IVRContent `json:"options,omitempty"`
}

//go:generate goqueryset -in ivr.go
//gen:qs
type IVR struct {
	Id            int64       `json:"id"`
	TenantId      string      `json:"tenant_id"`
	PhoneNumber   string      `json:"phone_number" gorm:"uniqueIndex"`
	Configuration *IVRContent `json:"configuration,omitempty" gorm:"type:JSONB;serializer:json"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
