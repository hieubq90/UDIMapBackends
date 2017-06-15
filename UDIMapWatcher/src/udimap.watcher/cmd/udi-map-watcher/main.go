package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"udimap.watcher/config"
	"udimap.watcher/models"
	"udimap.watcher/utils/sliceutils"

	"time"

	soap "github.com/achiku/soapc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mkideal/cli"
)

func main() {
	if config.InitFromYAML() {
		if err := cli.Root(root,
			cli.Tree(help),
			cli.Tree(migrate),
			cli.Tree(crawl),
		).Run(os.Args[1:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

var help = cli.HelpCommand("UDIMap Watcher")

// root command
type rootT struct {
	cli.Helper
	Version string `cli:"v,version" usage:"version" dft:"0.0.1"`
	Name    string `cli:"name" usage:"your name" dft:"root"`
}

var root = &cli.Command{
	Desc: "this is root command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} {
		return new(rootT)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		ctx.String("UDIMap Watcher - version %s\n", argv.Version)
		return nil
	},
}

// migrate database
type migrateT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name" dft:"migrate"`
}

var migrate = &cli.Command{
	Name: "migrate",
	Desc: "UDIMap Watcher - Migrate Database",
	Argv: func() interface{} {
		return new(migrateT)
	},
	Fn: func(ctx *cli.Context) error {
		ctx.String("UDIMap Watcher - Start Migrate Database\n")

		db, err := gorm.Open("postgres", config.AppConfig.DBConnection)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Automatically create the "accounts" table based on the Account model.
		db.CreateTable(&models.Camera{})
		db.CreateTable(&models.TramDoMua{})
		db.CreateTable(&models.TramDoTrieu{})

		return nil
	},
}

// migrate database
type crawlT struct {
	cli.Helper
	Name string `cli:"name" usage:"your name" dft:"crawl"`
}

var crawl = &cli.Command{
	Name: "crawl",
	Desc: "UDIMap Watcher - Migrate Database",
	Argv: func() interface{} {
		return new(crawlT)
	},
	Fn: func(ctx *cli.Context) error {
		ctx.String("UDIMap Watcher - Start Migrate Database\n")

		db, err := gorm.Open("postgres", config.AppConfig.DBConnection)
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(10)
		db.DB().SetConnMaxLifetime(time.Duration(10) * time.Minute)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		db.AutoMigrate(&models.Camera{})
		db.AutoMigrate(&models.TramDoMua{})
		db.AutoMigrate(&models.TramDoTrieu{})
		db.AutoMigrate(&models.TramQuanTracNgap{})
		db.AutoMigrate(&models.FloodPoint{})

		for {
			fmt.Println(time.Now().Format("02-01-2006 15:04"), "Start Crawling...")
			CrawlCamera(db)
			CrawlDSTramDoMua(db)
			CrawlDSTramDoTrieu(db)
			CrawlQuanTracNgap(db)
			CrawlDiemNgap(db)
			fmt.Println(time.Now().Format("02-01-2006 15:04"), "Done Crawling...")
			time.Sleep(time.Duration(config.AppConfig.IntervalTime) * time.Minute)
		}

		return nil
	},
}

func CrawlCamera(db *gorm.DB) {
	fmt.Println("CrawlCamera")
	client := soap.NewClient(config.AppConfig.UDIMapEndpoint, false, nil)
	body := models.GetCameraBodyContent{Key: config.AppConfig.UDIMapKey}
	resp := models.GetCameraResponse{}
	crawErr := client.Call(config.AppConfig.GetCameraSoapAction, body, &resp, nil)
	if crawErr != nil {
		fmt.Println("err: ", crawErr)
	} else {
		listCameras := make([]*models.CameraNgap, 0)
		if parseErr := json.Unmarshal([]byte(resp.Result), &listCameras); parseErr != nil {
			fmt.Println(time.Now().Unix(), "CrawlCamera | parse json error", parseErr)
		} else {
			existedCameras := make([]*models.Camera, 0)
			currentCameras := make([]*models.Camera, 0)
			db.Find(&existedCameras)
			if len(listCameras) > 0 {
				for _, c := range listCameras {
					lat, _ := strconv.ParseFloat(c.Lat, 64)
					lng, _ := strconv.ParseFloat(c.Lng, 64)
					camera := &models.Camera{
						ID:       fmt.Sprintf("%s_%s", c.Lat, c.Lng),
						Address:  c.Address,
						ImageUrl: fmt.Sprintf("%s/CGIProxy.fcgi?cmd=snapPicture&usr=tho&pwd=SETh93jp", c.Domain),
						Lat:      lat,
						Lng:      lng,
					}
					currentCameras = append(currentCameras, camera)
				}
			}
			if len(existedCameras) > 0 && len(currentCameras) > 0 {
				// remove if camera is not contained in current cameras
				for _, c := range existedCameras {
					if !sliceutils.ContainCamera(currentCameras, c) {
						db.Delete(c)
					}
				}
			}
			if len(currentCameras) > 0 && len(existedCameras) > 0 {
				for _, c := range currentCameras {
					if sliceutils.ContainCamera(existedCameras, c) {
						db.Save(c)
					} else {
						db.Create(c)
					}
				}
			} else {
				for _, c := range currentCameras {
					db.Create(c)
				}
			}
		}
	}
}

func CrawlDSTramDoMua(db *gorm.DB) {
	fmt.Println("CrawlDSTramDoMua")
	client := soap.NewClient(config.AppConfig.UDIMapEndpoint, false, nil)
	body := models.GetTramDoMuaBodyContent{Key: config.AppConfig.UDIMapKey}
	resp := models.GetTramDoMuaResponse{}
	crawErr := client.Call(config.AppConfig.GetTramDoMuaSoapAction, body, &resp, nil)
	if crawErr != nil {
		fmt.Println("err: ", crawErr)
	} else {
		currentTramDoMua := make([]*models.TramDoMua, 0)
		if parseErr := json.Unmarshal([]byte(resp.Result), &currentTramDoMua); parseErr != nil {
			fmt.Println(time.Now().Unix(), "CrawlDSTramDoMua | parse json error", parseErr)
		} else {
			existedTramDoMua := make([]*models.TramDoMua, 0)
			db.Find(&existedTramDoMua)

			if len(existedTramDoMua) > 0 && len(currentTramDoMua) > 0 {
				// remove if camera is not contained in current cameras
				for _, c := range existedTramDoMua {
					if !sliceutils.ContainTramDoMua(currentTramDoMua, c) {
						db.Delete(c)
					}
				}
			}
			if len(currentTramDoMua) > 0 && len(existedTramDoMua) > 0 {
				for _, c := range currentTramDoMua {
					if sliceutils.ContainTramDoMua(existedTramDoMua, c) {
						db.Save(c)
					} else {
						db.Create(c)
					}
				}
			} else {
				for _, c := range currentTramDoMua {
					db.Create(c)
				}
			}
		}
	}
}

func CrawlDSTramDoTrieu(db *gorm.DB) {
	fmt.Println("CrawlDSTramDoTrieu")
	client := soap.NewClient(config.AppConfig.UDIMapEndpoint, false, nil)
	body := models.GetTramDoTrieuBodyContent{Key: config.AppConfig.UDIMapKey}
	resp := models.GetTramDoTrieuResponse{}
	crawErr := client.Call(config.AppConfig.GetTramDoTrieuSoapAction, body, &resp, nil)
	if crawErr != nil {
		fmt.Println("err: ", crawErr)
	} else {
		currentTramDoTrieu := make([]*models.TramDoTrieu, 0)
		if parseErr := json.Unmarshal([]byte(resp.Result), &currentTramDoTrieu); parseErr != nil {
			fmt.Println(time.Now().Unix(), "CrawlDSTramDoTrieu | parse json error", parseErr)
		} else {
			existedTramDoTrieu := make([]*models.TramDoTrieu, 0)
			db.Find(&existedTramDoTrieu)

			if len(existedTramDoTrieu) > 0 && len(currentTramDoTrieu) > 0 {
				// remove if camera is not contained in current cameras
				for _, c := range existedTramDoTrieu {
					if !sliceutils.ContainTramDoTrieu(currentTramDoTrieu, c) {
						db.Delete(c)
					}
				}
			}
			if len(currentTramDoTrieu) > 0 && len(existedTramDoTrieu) > 0 {
				for _, c := range currentTramDoTrieu {
					if sliceutils.ContainTramDoTrieu(existedTramDoTrieu, c) {
						id, _ := strconv.ParseInt(c.IDText, 10, 64)
						c.ID = id
						db.Save(c)
					} else {
						db.Create(c)
					}
				}
			} else {
				for _, c := range currentTramDoTrieu {
					db.Create(c)
				}
			}
		}
	}
}

func CrawlQuanTracNgap(db *gorm.DB) {
	fmt.Println("CrawlQuanTracNgap")
	client := soap.NewClient(config.AppConfig.UDIMapEndpoint, false, nil)
	body := models.GetDSQuanTracNgapBodyContent{Key: config.AppConfig.UDIMapKey}
	resp := models.GetDSQuanTracNgapResponse{}
	crawErr := client.Call(config.AppConfig.GetDSQuanTracNgapSoapAction, body, &resp, nil)
	if crawErr != nil {
		fmt.Println("err: ", crawErr)
	} else {
		listQuanTracNgap := make([]*models.QuanTracNgap, 0)
		if parseErr := json.Unmarshal([]byte(resp.Result), &listQuanTracNgap); parseErr != nil {
			fmt.Println(time.Now().Unix(), "CrawlQuanTracNgap | parse json error", parseErr)
		} else {
			existedQuanTracNgap := make([]*models.TramQuanTracNgap, 0)
			currentQuanTracNgap := make([]*models.TramQuanTracNgap, 0)
			db.Find(&existedQuanTracNgap)
			if len(listQuanTracNgap) > 0 {
				for _, c := range listQuanTracNgap {
					lat, _ := strconv.ParseFloat(c.Lat, 64)
					lng, _ := strconv.ParseFloat(c.Lng, 64)
					id, _ := strconv.ParseInt(c.IDText, 10, 64)
					tramQuanTracNgap := &models.TramQuanTracNgap{
						ID:         id,
						Name:       c.Name,
						Address:    c.Address,
						Lat:        lat,
						Lng:        lng,
						FloodDeep:  c.FloodDeep,
						StatusID:   c.StatusID,
						Status:     c.Status,
						StatusText: c.StatusText,
						LastUpdate: c.LastUpdate,
					}
					currentQuanTracNgap = append(currentQuanTracNgap, tramQuanTracNgap)
				}
			}
			if len(existedQuanTracNgap) > 0 && len(currentQuanTracNgap) > 0 {
				// remove if camera is not contained in current cameras
				for _, c := range existedQuanTracNgap {
					if !sliceutils.ContainTramQuanTracNgap(currentQuanTracNgap, c) {
						db.Delete(c)
					}
				}
			}
			if len(currentQuanTracNgap) > 0 && len(existedQuanTracNgap) > 0 {
				for _, c := range currentQuanTracNgap {
					if sliceutils.ContainTramQuanTracNgap(existedQuanTracNgap, c) {
						db.Save(c)
					} else {
						db.Create(c)
					}
				}
			} else {
				for _, c := range currentQuanTracNgap {
					db.Create(c)
				}
			}

		}
	}
}

func CrawlDiemNgap(db *gorm.DB) {
	fmt.Println("CrawlDiemNgap")
	client := soap.NewClient(config.AppConfig.UDIMapEndpoint, false, nil)
	body := models.GetDSDiemNgapBodyContent{Key: config.AppConfig.UDIMapKey}
	resp := models.GetDSDiemNgapResponse{}
	crawErr := client.Call(config.AppConfig.GetDSDiemNgapSoapAction, body, &resp, nil)
	if crawErr != nil {
		fmt.Println("err: ", crawErr)
	} else {
		listDiemNgap := make([]*models.DiemNgap, 0)
		if parseErr := json.Unmarshal([]byte(resp.Result), &listDiemNgap); parseErr != nil {
			fmt.Println(time.Now().Unix(), "CrawlDiemNgap | parse json error", parseErr)
		} else {
			existedFloodPoints := make([]*models.FloodPoint, 0)
			currentFloodPoints := make([]*models.FloodPoint, 0)
			db.Find(&existedFloodPoints)
			if len(listDiemNgap) > 0 {
				for _, c := range listDiemNgap {
					lat, _ := strconv.ParseFloat(c.Lat, 64)
					lng, _ := strconv.ParseFloat(c.Lng, 64)
					floodPoint := &models.FloodPoint{
						ID:           c.ID,
						RoadName:     c.RoadName,
						DistrictName: c.DistrictName,
						From:         c.From,
						To:           c.To,
						Lat:          lat,
						Lng:          lng,
						FloodDeep:    c.FloodDeep,
						Status:       c.Status,
						Expected:     c.Expected,
						Warning:      c.Warning,
						LastUpdate:   c.LastUpdate,
					}
					currentFloodPoints = append(currentFloodPoints, floodPoint)
				}
			}
			if len(existedFloodPoints) > 0 && len(currentFloodPoints) > 0 {
				// remove if camera is not contained in current cameras
				for _, c := range existedFloodPoints {
					if !sliceutils.ContainFloodPoint(currentFloodPoints, c) {
						db.Delete(c)
					}
				}
			}
			if len(currentFloodPoints) > 0 && len(existedFloodPoints) > 0 {
				for _, c := range currentFloodPoints {
					if sliceutils.ContainFloodPoint(existedFloodPoints, c) {
						db.Save(c)
					} else {
						db.Create(c)
					}
				}
			} else {
				for _, c := range currentFloodPoints {
					db.Create(c)
				}
			}

		}
	}
}
