package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func getImg() string {
	resp, err := http.Get("http://cn.bing.com/")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	s := string(body)

	reg := regexp.MustCompile("http[^,\"]+jpg")
	arr := reg.FindAllString(s, -1)
	if len(arr) > 0 {
		return strings.Replace(arr[0], "\\", "", -1)
	} else {
		return ""
	}
}

func downLoadImg(imgurl string) string {
	path := "/home/fighter/Pictures/bing_wallpaper/"

	filepath := path + imgurl[strings.LastIndex(imgurl, "/")+1:]
	imgrsp, err := http.Get(imgurl)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer imgrsp.Body.Close()

	body, err := ioutil.ReadAll(imgrsp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	ioutil.WriteFile(filepath, body, 0600)

	return filepath

}

func setWallPaper(filepath string) {
	out, err := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", "file://"+filepath).Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("in all caps: %q\n", out)
}

func getWallPaper() string {
	out, err := exec.Command("gsettings", "get", "org.gnome.desktop.background", "picture-uri").Output()
	if err != nil {
		fmt.Println(err)
	}
	filepath := string(out)
	return filepath
}

func main() {
	//exec.Command("sleep", "20").Output()
	imgurl := getImg()
	fmt.Println(imgurl)
	if len(imgurl) > 0 {
		filepath := downLoadImg(imgurl)
		fmt.Println(filepath)
		setWallPaper(filepath)
	}

	fmt.Printf("wallpaper is %s\n", getWallPaper())
}
