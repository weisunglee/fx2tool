package pgjson

/* sample
{
    "LivePresets": [
        {
            "display_order": 0,
            "is_favorite": false,
            "preset_name": "Crunchy Cookies",
            "preset_uuid": "06977430-37AA-41D2-BD69-B593AD81EB3B"
        },
        {
            "display_order": 1,
            "is_favorite": false,
            "preset_name": "British Love",
            "preset_uuid": "073EF34E-4526-4D44-9B02-2B3CDBFD8576"
        }
    ]
}
*/

// LivePresets : json struct
type LivePresets struct {
	DisplayOrder int    `json:"display_order"`
	IsFavorite   bool   `json:"is_favorite"`
	PresetName   string `json:"preset_name"`
	PresetUUID   string `json:"preset_uuid"`
}

// Preset : json struct
type Preset struct {
	LivePresets []LivePresets `json:"LivePresets"`
}
