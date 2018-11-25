package minecraft

import (
	"io/ioutil"
	"path/filepath"
)

// ListMods returns mods
func (s *Server) ListMods() ([]*Mod, error) {
	modsPath := filepath.Join(s.Dir, "/mods")
	files, err := ioutil.ReadDir(modsPath)
	if err != nil {
		return nil, err
	}

	mods := make([]*Mod, 0)
	for _, f := range files {
		name := f.Name()
		mods = append(mods, &Mod{
			Name: name,
			Path: filepath.Join(modsPath, name),
		})
	}

	return mods, nil
}

// InstallMod installs a mod
func (s *Server) InstallMod() (*Mod, error) {
	return nil, nil
}

// RemoveMod removes a mod
func (s *Server) RemoveMod(m *Mod) error {
	return nil
}
