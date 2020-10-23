package gobat2exe

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Bat2Exe converts Windows batch files into executable files.
func Bat2Exe(batPath string) error {
	sedFile, err := ioutil.TempFile(os.TempDir(), "sed-")
	if err != nil {
		return err
	}
	log.Println("sedFile.Name", sedFile.Name())

	batName := filepath.Base(batPath)
	log.Println("batName", batName)

	batDirectoryPath, err := filepath.Abs(filepath.Dir(batPath))
	if err != nil {
		return err
	}
	log.Println("batDirectoryPath", batDirectoryPath)

	sedFileContents := getSEDFileContents(batName, batDirectoryPath)
	log.Println("sedFileContents", sedFileContents)

	text := []byte(sedFileContents)
	if _, err = sedFile.Write(text); err != nil {
		return err
	}

	if err := sedFile.Close(); err != nil {
		return err
	}

	s := []string{"cmd.exe", "/C", "start", "iexpress", "/n", "/q", "/m", sedFile.Name()}
	log.Println("cmd", s)

	cmd := exec.Command(s[0], s[1:]...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// getSEDFileContents return string representation of Self Extraction Directive https://ss64.com/nt/iexpress-sed.html
func getSEDFileContents(batName string, batDirectoryPath string) string {
	return `
		[Version]
		Class=IEXPRESS
		SEDVersion=3
		[Options]
		PackagePurpose=InstallApp
		ShowInstallProgramWindow=0
		HideExtractAnimation=1
		UseLongFileName=1
		InsideCompressed=0
		CAB_FixedSize=0
		CAB_ResvCodeSigning=0
		RebootMode=N
		InstallPrompt=%InstallPrompt%
		DisplayLicense=%DisplayLicense%
		FinishMessage=%FinishMessage%
		TargetName=%TargetName%
		FriendlyName=%FriendlyName%
		AppLaunched=%AppLaunched%
		PostInstallCmd=%PostInstallCmd%
		AdminQuietInstCmd=%AdminQuietInstCmd%
		UserQuietInstCmd=%UserQuietInstCmd%
		SourceFiles=SourceFiles
		[Strings]
		InstallPrompt=
		DisplayLicense=
		FinishMessage=
		TargetName=` + batDirectoryPath + `\` + batName + `.exe
		FriendlyName=` + batName + `
		AppLaunched=cmd /c ` + batName + `
		PostInstallCmd=<None>
		AdminQuietInstCmd=
		UserQuietInstCmd=%UserQuietInstCmd%
		FILE0="` + batName + `"
		[SourceFiles]
		SourceFiles0=` + batDirectoryPath + `
		[SourceFiles0]
		%FILE0%=
		`
}
