package main

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	MOVE_DOWN = iota
	MOVE_UP
	MISC
)
const (
	TOTAL_DISK_SIZE  int64 = 70000000
	NEEDED_DISK_SIZE int64 = 30000000
)

func UpdateSize(newSize int64, DirMap map[string]DirectoryNode, DirList []string) map[string]DirectoryNode {
	CurrDirName := fmt.Sprintf("/%v", strings.Join(DirList, "/"))
	result, _ := DirMap[CurrDirName]
	result.Size += newSize
	DirMap[CurrDirName] = result

	if len(DirList) == 0 {
		return DirMap
	}
	//*originalMap = DirMap
	return UpdateSize(newSize, DirMap, DirList[:len(DirList)-1])
}

func parseCommand(commArr []string, DirListPoint *[]string) int {
	Command := commArr[1]
	DirList := *DirListPoint
	defer func() {
		*DirListPoint = DirList
	}()

	if Command == "cd" {
		DestDir := commArr[2]
		if DestDir == ".." {
			DirList = DirList[:len(DirList)-1]
			return MOVE_UP
		}
		if DestDir == "/" {
			DirList = []string{}
			return MOVE_UP
		}
		DirList = append(DirList, DestDir)
		return MOVE_DOWN
	}
	return MISC
}

type DirectoryNode struct {
	Name string
	Size int64
}

func main() {
	var (
		CurrDir     DirectoryNode
		CurrDirSize int64
		DirList     []string
	)
	DirMap := make(map[string]DirectoryNode)
	CurrDir = DirectoryNode{Name: "/"}
	DirMap["/"] = CurrDir

	utils.ReadFile("Day07/input.txt", func(text string) {
		text = strings.TrimSpace(text)

		textArr := strings.Split(text, " ")
		if textArr[0] == "$" {
			DirMap = UpdateSize(CurrDirSize, DirMap, DirList)
			defer func() { CurrDirSize = 0 }()

			CommandType := parseCommand(textArr, &DirList)
			CurrDirName := fmt.Sprintf("/%v", strings.Join(DirList, "/"))

			switch CommandType {
			case MOVE_DOWN:
				if oldDir, ok := DirMap[CurrDirName]; ok {
					CurrDir = oldDir
				} else {
					CurrDir = DirectoryNode{
						Name: CurrDirName,
						Size: 0,
					}
					DirMap[CurrDirName] = CurrDir
				}
			case MOVE_UP:
				CurrDir = DirMap[CurrDirName]
			}
		} else {
			if textArr[0] == "dir" {
				return
			}
			itemSize, _ := strconv.ParseInt(textArr[0], 10, 64)
			CurrDirSize += itemSize
		}
	})

	var totalSize int64
	for _, node := range DirMap {
		if node.Size <= 100000 {
			totalSize += node.Size
		}
	}
	fmt.Println("First Task: Total Size:", utils.PrettyPrintNum(totalSize))

	rootSize := DirMap["/"].Size
	emptyDiskSpace := TOTAL_DISK_SIZE - rootSize
	neededEmptySpace := NEEDED_DISK_SIZE - emptyDiskSpace

	fmt.Println()
	fmt.Println("Root Size:", utils.PrettyPrintNum(rootSize))
	fmt.Println("Empty Size:", utils.PrettyPrintNum(emptyDiskSpace))
	fmt.Println("Needed Empty Size:", utils.PrettyPrintNum(neededEmptySpace))

	currentSmallestDir := TOTAL_DISK_SIZE
	DirToBeDeleted := DirectoryNode{}
	for _, directory := range DirMap {
		if directory.Size > neededEmptySpace {
			if directory.Size < currentSmallestDir {
				currentSmallestDir = directory.Size
				DirToBeDeleted = directory
			}
		}
	}

	fmt.Println()
	fmt.Println("Smallest Directory to be deleted:", fmt.Sprintf("%+v", DirToBeDeleted))
	fmt.Println("Smallest Directory to be deleted size:", utils.PrettyPrintNum(DirToBeDeleted.Size))
}
