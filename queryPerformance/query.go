package queryperformance

import (
    "fmt"
    "time"
    "github.com/gocql/gocql"

)


func CompareQueryPerformance(session *gocql.Session)(float64,float64){
	timeDuration2 := batchInsertion(session);
	fmt.Println(timeDuration2);
	timeDuration1 := insertOneByOne(session);
	fmt.Println(timeDuration1);
	return timeDuration2,timeDuration1;
}

func insertOneByOne(session *gocql.Session) float64{
	fmt.Println("In insert one by one operation");
    start := time.Now();
    UploadCsvDataOneByOne(session);
    totalDuration := time.Since(start).Seconds();
	return totalDuration;
}


func batchInsertion(session *gocql.Session) float64{
	fmt.Println("In batch insertion operation");
	start := time.Now();
	BatchUploadData(session);
	totalDuration := time.Since(start).Seconds();
	return totalDuration;
}