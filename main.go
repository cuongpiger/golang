package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Calculate duration
	duration, err := GetAudioDurationByFFprobe("./gato.ogg")
	if err != nil {
		fmt.Printf("Error calculating duration: %v\n", err)
		return
	}

	fmt.Printf("Duration of gato.ogg: %d seconds\n", int(duration + 1))
}

func GetAudioDurationByFFprobe(filename string) (float64, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1", filename)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	durationStr := strings.TrimSpace(string(out))
	return strconv.ParseFloat(durationStr, 64)
}
