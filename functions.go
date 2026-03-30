package main

import "golang.org/x/sys/windows/registry"

func isDarkModeEnabled() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	val, _, err := k.GetIntegerValue("AppsUseLightTheme")
	if err != nil {
		return false
	}
	return val == 0
}

func isOldMenuEnabled() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}\InprocServer32`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	return true
}

func isArrowEnabled() bool {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Shell Icons`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	return true
}

func isRecycleBinHidden() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\HideDesktopIcons\NewStartPanel`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	val, _, err := k.GetIntegerValue("{645FF040-5081-101B-9F08-00AA002F954E}")
	if err != nil {
		return false
	}
	return val == 1
}

func isHiddenFilesShown() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	val, _, err := k.GetIntegerValue("Hidden")
	if err != nil {
		return false
	}

	return val == 1
}

func isExtensionsShown() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.QUERY_VALUE)
	if err != nil {
		return true
	}
	defer k.Close()
	val, _, err := k.GetIntegerValue("HideFileExt")
	if err != nil {
		return true
	}
	return val == 0
}

func getSearchMode() string {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Search`, registry.QUERY_VALUE)
	if err != nil {
		return "Icon Only"
	}
	defer k.Close()

	val, _, err := k.GetIntegerValue("SearchboxTaskbarMode")
	if err != nil {
		return "Icon Only"
	}

	switch val {
	case 0:
		return "Hidden"
	case 1:
		return "Icon Only"
	case 3:
		return "Search Box"
	default:
		return "Icon Only"
	}
}

func getTaskbarAlignment() string {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.QUERY_VALUE)
	if err != nil {
		return "Center"
	}
	defer k.Close()
	val, _, err := k.GetIntegerValue("TaskbarAl")
	if err != nil || val == 1 {
		return "Center"
	}
	return "Left"
}

func isTaskViewHidden() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	val, _, err := k.GetIntegerValue("ShowTaskViewButton")
	if err != nil {
		return false
	}
	return val == 0
}

func isSecondsEnabled() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()
	val, _, err := k.GetIntegerValue("ShowSecondsInSystemClock")
	if err != nil {
		return false
	}
	return val == 1
}
