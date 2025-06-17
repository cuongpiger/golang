package main

import (
	"log"
	"os"

	mdl "github.com/cuongpiger/golang/models"
	db "github.com/cuongpiger/golang/postgre"
	ivrRepo "github.com/cuongpiger/golang/repo"
)

func main() {
	// Get the username and password via environment variables or configuration
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	config, err := db.NewPostgresClient(
		"localhost",
		"5432",
		user,
		password,
		dbname,
	)

	if err != nil {
		panic(err)
	}

	ivrRepository := ivrRepo.NewIVRRepo(config)

	ivrSample := &mdl.IVR{
		Id:          4,
		PhoneNumber: "+18573824222",
		TenantId:    "tenant-123",
		Configuration: &mdl.IVRContent{
			Key:         mdl.OneKey,
			Action:      mdl.RouteAction,
			Description: "Main Menu",
			Value:       "main_menu",
			Priority:    mdl.MediumPriority,
			Sounds: mdl.ListSounds{
				{
					SoundPath:      "welcome.wav",
					Duration:       5,
					RepeatCount:    1,
					RepeatInterval: 0,
				},
			},
			Options: []mdl.IVRContent{
				{
					Key:         mdl.TwoKey,
					Action:      mdl.PlaySoundAction,
					Description: "Press 2 for Sales",
					Value:       "sales",
					Priority:    mdl.LowPriority,
					Sounds: mdl.ListSounds{
						{
							SoundPath:      "sales.wav",
							Duration:       5,
							RepeatCount:    1,
							RepeatInterval: 0,
						},
					},
				},
			},
		},
	}

	err = ivrRepository.CreateIVR(ivrSample)
	if err != nil {
		log.Fatal(err, "failed to create IVR")
	}
	log.Println("IVR created successfully:", ivrSample.PhoneNumber)

	ivr, err := ivrRepository.GetIVRByPhoneNumber("+18573824222")
	if err != nil {
		log.Fatal(err, "failed to get IVR by phone number")
	}

	log.Printf("IVR: %+v\n", ivr)
}
