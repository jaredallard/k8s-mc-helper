package minecraft

// Version is a rep of a minecraft version
type Version struct {
	Major int64
	Minor int64
	Patch int64
}

// Server is a minecraft server
type Server struct {
	PID int
	Dir string
}

// Mod is a minecraft mod
type Mod struct {
	Name             string
	Path             string
	Version          string
	MinecraftVersion string
	URL              string
}
