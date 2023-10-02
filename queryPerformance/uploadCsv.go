package queryperformance

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/gocql/gocql"
)

func UploadCsvDataOneByOne(session *gocql.Session){
	f, err := os.Open("title.episode_processed.csv")
    if err != nil {
        log.Fatal(err)
    }
    i := 1;
    defer f.Close()
    csvReader := csv.NewReader(f)
    for {
        rec, err := csvReader.Read()
        if err != nil {
            break;
        }
        if i== 1000000{
            break;
        }
        col1 := rec[0]
        col2 := rec[1]
        col3 := rec[2]
        col4 := rec[3];

          if err := session.Query(
            "INSERT INTO sampleData1(tconst,parentTconst,seasonNumber,episodeNumber) VALUES (?, ?, ?, ?)",
            col1, col2,col3,col4).Exec(); err != nil {
            fmt.Println("Error in insertion")
			fmt.Println(err.Error());
			break;
        }
        i++;
    }
}


func BatchUploadData(session *gocql.Session){
    f, err := os.Open("title.episode_processed.csv")
    if err != nil {
        log.Fatal(err)
    }
    i := 1;
    batch := session.NewBatch(gocql.UnloggedBatch);
    defer f.Close()
    csvReader := csv.NewReader(f)
    for {
        rec, err := csvReader.Read()
        if err != nil {
            break;
        }
        col1 := rec[0]
        col2 := rec[1]
        col3 := rec[2]
        col4 := rec[3];
         batch.Query("INSERT INTO sampleData2 (tconst,parentTconst,seasonNumber,episodeNumber) VALUES (?, ?,?,?)", col1, col2,col3,col4);
		 if i%1000 == 0{
			if err := session.ExecuteBatch(batch); err != nil{
				fmt.Println("Error");
				fmt.Print(err.Error());
                break;
			}
			batch = session.NewBatch(gocql.UnloggedBatch);
		 }
        //  if i==1000000{
        //     break;
        //  }
         i++;
    }
}