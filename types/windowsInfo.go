package types

import (
	"image"
	"log"
	"os"
	"strings"

	"github.com/tc-hib/winres"
	"github.com/tc-hib/winres/version"
)

type WindowsInfo struct {
	Name        string
	Version     string //v0.0.0.0
	IconPath    string
	Description string
	SysPath     string // ../rsrc_windows_amd64.syso
	CopyRight   string
	CompanyName string
	RS          winres.ResourceSet
}

func (wi *WindowsInfo) checkVersion() {
	if !strings.HasPrefix(wi.Version, "v") {
		wi.Version = "v" + wi.Version
	}
	if len(wi.Version) != 8 {
		wi.Version = wi.Version + ".0"
	}
}

func (wi *WindowsInfo) createVersion() [4]uint16 {
	wi.checkVersion()
	v := [4]uint16{}
	ver := wi.Version
	for i, n := range ver {
		if n == '.' || n == 'v' {
			continue
		}
		if i > 3 {
			break
		}

		v[i] = uint16(n)
	}
	return v
}

func (wi *WindowsInfo) setIcon() {
	f, err := os.Open(wi.IconPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
	icon, er := winres.NewIconFromResizedImage(img, nil)
	if er != nil {
		log.Fatalln(er)
	}
	wi.RS.SetIcon(winres.Name("APPICON"), icon)
}

func (wi *WindowsInfo) setVersionInfo() {
	vi := version.Info{
		FileVersion:    [4]uint16{1, 0, 0, 0},
		ProductVersion: wi.createVersion(),
		Type:           version.App,
	}
	vi.Set(0, version.ProductName, wi.Name)
	vi.Set(0, version.ProductVersion, wi.Version)
	vi.Set(0, version.LegalCopyright, wi.CopyRight)
	vi.Set(0, version.FileDescription, wi.Description)
	vi.Set(0, version.InternalName, wi.Name)
	vi.Set(0, version.CompanyName, wi.CompanyName)
	// Add the VersionInfo to the resource set
	wi.RS.SetVersionInfo(vi)
}

func (wi *WindowsInfo) setManifest() {
	wi.RS.SetManifest(winres.AppManifest{
		Description:                  wi.Description,
		Compatibility:                winres.Win10AndAbove,
		DisableWindowFiltering:       true,
		UseCommonControlsV6:          true,
		UIAccess:                     false,
		HighResolutionScrollingAware: true,
		LongPathAware:                true,
		ExecutionLevel:               winres.AsInvoker,
		Identity:                     winres.AssemblyIdentity{Name: wi.Name, Version: wi.createVersion()},
	})
}

func (wi *WindowsInfo) writeObject() error {
	out, err := os.Create(wi.SysPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	return wi.RS.WriteObject(out, winres.ArchAMD64)
}

func (wi *WindowsInfo) Write() error {
	wi.setIcon()
	wi.setVersionInfo()
	wi.setManifest()
	return wi.writeObject()
}
