package main

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/qedus/osmpbf"
	"github.com/gaffo/goosm"
	"io"
	"log"
	"os"
	_ "runtime"
)

func main() {
	f, err := os.Open("washington-latest.osm.pbf")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	o, err := os.Create("washington-latest-highway.osm")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	osm := goosm.NewOsm()

	stat, _ := f.Stat()
	size := stat.Size()
	bar := pb.New64(size).SetUnits(pb.U_BYTES)
	bar.ShowSpeed = true
	bar.Start()

	d := osmpbf.NewDecoder(bar.NewProxyReader(f))
	err = d.Start(1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var nc, wc, rc uint64
	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch t := v.(type) {
			case *osmpbf.Node:
				n := v.(*osmpbf.Node)
				for k, _ := range n.Tags {
					if k == "highway" {
						osm.AppendNode(goosm.Node{
							Id: fmt.Sprintf("%d", n.ID),
							Lat: n.Lat,
							Lon: n.Lon,
							Visible: n.Info.Visible})
					}
				}
				// Process Node v.
				nc++
			case *osmpbf.Way:
				n := v.(*osmpbf.Way)
				for k, _ := range n.Tags {
					if k == "highway" {
						way := goosm.NewWay()
						for _, nd := range n.NodeIDs {
							way.AppendNd(string(nd))
						}
						osm.AppendWay(*way)
						wc++
					}
				}
			case *osmpbf.Relation:
				n := v.(*osmpbf.Relation)
				for k, _ := range n.Tags {
					if k == "highway" {
						rc++
					}
				}
			default:
				log.Fatalf("unknown type %T\n", t)
			}
		}
	}
	fmt.Println()
	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
	osm.Write(o)
}
