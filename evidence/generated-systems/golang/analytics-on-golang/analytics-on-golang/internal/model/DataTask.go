package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DataTask Declaration
//==============================================================
type DataTask struct {
    gorm.Model
     Name                                    string
    Command                                    string
    Retries                                                            string
    PipelineId         *uint
    Pipeline           *DataPipeline `gorm:"foreignKey:PipelineId"`
     InputDatasets           []DataSet `gorm:"foreignKey:InputDatasetsFromDataTaskId"`
     OutputDatasets           []DataSet `gorm:"foreignKey:OutputDatasetsFromDataTaskId"`
    TaskType                      DataTaskType

// parent associations as their child

}

