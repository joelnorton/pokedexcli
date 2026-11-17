package main

// Pokemon represents a Pokémon from the PokeAPI.
// Source: https://pokeapi.co/api/v2/pokemon/{id}/
type Pokemon struct {
	ID                     int              `json:"id"`
	Name                   string           `json:"name"`
	BaseExperience         int              `json:"base_experience"`
	Height                 int              `json:"height"`
	Weight                 int              `json:"weight"`
	IsDefault              bool             `json:"is_default"`
	Order                  int              `json:"order"`
	Abilities              []Ability        `json:"abilities"`
	PastAbilities          []PastAbility    `json:"past_abilities"`
	Forms                  []Form           `json:"forms"`
	GameIndices            []GameIndex      `json:"game_indices"`
	HeldItems              []HeldItem       `json:"held_items"`
	LocationAreaEncounters string           `json:"location_area_encounters"`
	Moves                  []Move           `json:"moves"`
	PastStats              []PastStat       `json:"past_stats"`
	PastTypes              []PastType       `json:"past_types"`
	Sprites                Sprites          `json:"sprites"`
	Stats                  []Stat           `json:"stats"`
	Types                  []Type           `json:"types"`
	Cries                  Cries            `json:"cries"`
	Species                NamedAPIResource `json:"species"`
}

// Ability represents a Pokémon's ability.
type Ability struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

// PastAbility represents past generations' abilities for a Pokémon.
type PastAbility struct {
	Generation Generation `json:"generation"`
	Abilities  []Ability  `json:"abilities"`
}

// Form represents a Pokémon's form variation.
type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GameIndex represents a Pokémon's appearance in a game version.
type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

// HeldItem represents an item held by a Pokémon.
type HeldItem struct {
	Item           NamedAPIResource `json:"item"`
	VersionDetails []VersionDetail  `json:"version_details"`
}

// VersionDetail holds rarity and version info for a held item.
type VersionDetail struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

// Move represents a Pokémon's move entry.
type Move struct {
	Move                NamedAPIResource         `json:"move"`
	VersionGroupDetails []MoveVersionGroupDetail `json:"version_group_details"`
}

// MoveVersionGroupDetail describes how a move is learned in a version group.
type MoveVersionGroupDetail struct {
	LevelLearnedAt   int              `json:"level_learned_at"`
	VersionGroup     NamedAPIResource `json:"version_group"`
	MoveLearnMethod  NamedAPIResource `json:"move_learn_method"`
	Order            *int             `json:"order"`
	EnhancementLevel *int             `json:"enhancement_level"`
}

// Stat represents a Pokémon's base stat entry.
type Stat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

// PastStat represents past generations' stats for a Pokémon.
type PastStat struct {
	Generation Generation `json:"generation"`
	Stats      []Stat     `json:"stats"`
}

// Type represents a Pokémon's type entry.
type Type struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

// PastType represents past generations' types for a Pokémon.
type PastType struct {
	Generation Generation `json:"generation"`
	Types      []Type     `json:"types"`
}

// Sprites holds URLs for a Pokémon's artwork and sprites.
type Sprites struct {
	BackDefault      string         `json:"back_default"`
	BackFemale       *string        `json:"back_female"`
	BackShiny        string         `json:"back_shiny"`
	BackShinyFemale  *string        `json:"back_shiny_female"`
	FrontDefault     string         `json:"front_default"`
	FrontFemale      *string        `json:"front_female"`
	FrontShiny       string         `json:"front_shiny"`
	FrontShinyFemale *string        `json:"front_shiny_female"`
	Other            SpriteOther    `json:"other"`
	Versions         SpriteVersions `json:"versions"`
}

// SpriteOther holds non-version-specific sprite sets.
type SpriteOther struct {
	DreamWorld      *SpriteHome            `json:"dream_world"`
	Home            *SpriteHome            `json:"home"`
	OfficialArtwork *SpriteOfficialArtwork `json:"official-artwork"`
	Showdown        *SpriteHome            `json:"showdown"`
}

// SpriteHome holds sprite URLs used across several "other" categories.
type SpriteHome struct {
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
}

// SpriteOfficialArtwork holds official artwork URLs.
type SpriteOfficialArtwork struct {
	FrontDefault     string  `json:"front_default"`
	FrontShiny       string  `json:"front_shiny"`
	FrontFemale      *string `json:"front_female"`
	FrontShinyFemale *string `json:"front_shiny_female"`
	Versions         map[string]map[string]struct {
		FrontDefault *string `json:"front_default"`
	} `json:"versions"`
}

// SpriteVersions holds generation-specific sprites.
type SpriteVersions map[string]interface{}

// Cries holds URLs for a Pokémon's cry sounds.
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

// NamedAPIResource is a minimal reference object with name and URL.
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Generation represents a game generation reference.
type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
