package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DataSource Declaration
//==============================================================
type DataSource struct {
    gorm.Model
     Name                                    string
    Connection                                                            string
    Streaming                                    bool
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     ProducedDatasets           []DataSet `gorm:"foreignKey:ProducedDatasetsFromDataSourceId"`
     Pipelines           []DataPipeline `gorm:"foreignKey:PipelinesFromDataSourceId"`
    SourceType                      DataSourceType
    Format                      DataFormat

// parent associations as their child

}

