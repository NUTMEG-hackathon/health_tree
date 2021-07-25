package model

import (
  "fmt"
)

type Answer struct {
  ID        int    `json:"id" gorm:"praimaly_key"`
  Title      string `json:"title"`
  Point     float64    `json:"point"`
  QuestionID    int    `json:"question_id"`
}

type Answers []Answer

func AnswerIndex() Answers {
  db := Init()
  defer db.Close()

  var answers Answers
  db.Find(&answers)
  return answers
}

func AnswerSeed(){
  db := Init()
  defer db.Close()

  answers := Answers{
    { ID:1,  Title:"ans1-1",  Point:-10, QuestionID:1  },
    { ID:2,  Title:"ans1-2",  Point:20,  QuestionID:1  },
    { ID:3,  Title:"ans1-3",  Point:30,  QuestionID:1  },

    { ID:4,  Title:"ans2-1",  Point:-40, QuestionID:2  },
    { ID:5,  Title:"ans2-2",  Point:50,  QuestionID:2  },
    { ID:6,  Title:"ans2-3",  Point:60,  QuestionID:2  },

    { ID:7,  Title:"ans3-1",  Point:-70, QuestionID:3  },
    { ID:8,  Title:"ans3-2",  Point:80,  QuestionID:3  },
    { ID:9,  Title:"ans3-3",  Point:90,  QuestionID:3  },

    { ID:10, Title:"ans4-1",  Point:-10, QuestionID:4  },
    { ID:11, Title:"ans4-2",  Point:20,  QuestionID:4  },
    { ID:12, Title:"ans4-3",  Point:30,  QuestionID:4  },

    { ID:13, Title:"ans5-1",  Point:-40, QuestionID:5  },
    { ID:14, Title:"ans5-2",  Point:50,  QuestionID:5  },
    { ID:15, Title:"ans5-3",  Point:60,  QuestionID:5  },

    { ID:16, Title:"ans6-1",  Point:-70, QuestionID:6  },
    { ID:17, Title:"ans6-2",  Point:80,  QuestionID:6  },
    { ID:18, Title:"ans6-3",  Point:90,  QuestionID:6  },

    { ID:19, Title:"ans7-1",  Point:-10, QuestionID:7  },
    { ID:20, Title:"ans7-2",  Point:20,  QuestionID:7  },
    { ID:21, Title:"ans7-3",  Point:30,  QuestionID:7  },

    { ID:22, Title:"ans8-1",  Point:-40, QuestionID:8  },
    { ID:23, Title:"ans8-2",  Point:50,  QuestionID:8  },
    { ID:24, Title:"ans8-3",  Point:60,  QuestionID:8  },

    { ID:25, Title:"ans9-1",  Point:-70, QuestionID:9  },
    { ID:26, Title:"ans9-2",  Point:80,  QuestionID:9  },
    { ID:27, Title:"ans9-3",  Point:90,  QuestionID:9  },

    { ID:28, Title:"ans10-1", Point:-10, QuestionID:10 },
    { ID:29, Title:"ans10-2", Point:20,  QuestionID:10 },
    { ID:30, Title:"ans10-3", Point:30,  QuestionID:10 },

    { ID:31, Title:"ans11-1", Point:-40, QuestionID:11 },
    { ID:32, Title:"ans11-2", Point:50,  QuestionID:11 },
    { ID:33, Title:"ans11-3", Point:60,  QuestionID:11 },

    { ID:34, Title:"ans12-1", Point:-70, QuestionID:12 },
    { ID:35, Title:"ans12-2", Point:80,  QuestionID:12 },
    { ID:36, Title:"ans12-3", Point:90,  QuestionID:12 },

    { ID:37, Title:"ans13-1", Point:-10, QuestionID:13 },
    { ID:38, Title:"ans13-2", Point:20,  QuestionID:13 },
    { ID:39, Title:"ans13-3", Point:30,  QuestionID:13 },

    { ID:40, Title:"ans14-1", Point:-40, QuestionID:14 },
    { ID:41, Title:"ans14-2", Point:50,  QuestionID:14 },
    { ID:42, Title:"ans14-3", Point:60,  QuestionID:14 },

    { ID:43, Title:"ans15-1", Point:-70, QuestionID:15 },
    { ID:44, Title:"ans15-2", Point:80,  QuestionID:15 },
    { ID:45, Title:"ans15-3", Point:90,  QuestionID:15 },
  }

  for i,answer := range answers{
    fmt.Println(i,answer)
    db.Create(answer)
  }

}

