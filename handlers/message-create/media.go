package messageCreate

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

var mediaFolderPath string = "d:/Meme"

func HandleMedia(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid arguments for media command.")
		return
	}

	// Send "typing..." status
	err := s.ChannelTyping(m.ChannelID)
	if err != nil {
		log.Println("Failed to send typing indicator:", err)
	}

	// for now, use the "d:/Meme" Folder to look for the media title
	mediaFileName := args[1]

	// Check if the media file exists
	filePath, err := findMediaFile(mediaFolderPath, mediaFileName)
	if err != nil {
		log.Println(err)
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %v", err))
		return
	}

	// Open the file for sending
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Failed to open file:", err)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Failed to load the requested media file.")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Failed to close file:", err)
		}
	}(file)

	// Send the media file
	_, err = s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: "Here is the media you requested:",
		Files: []*discordgo.File{
			{
				Name:   mediaFileName,
				Reader: file,
			},
		},
	})
	if err != nil {
		log.Println("Failed to send media file:", err)
	}
}

func findMediaFile(folderPath, fileName string) (string, error) {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return "", fmt.Errorf("failed to read media folder: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && entry.Name() == fileName {
			return folderPath + "/" + fileName, nil
		}
	}
	return "", fmt.Errorf("file %q not found in folder %q", fileName, folderPath)
}
