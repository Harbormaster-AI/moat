class DataSource < ApplicationRecord
  enum SourceType: [:Database, :File, :Stream, :API, :DataWarehouse, :DataLake]
  enum Format: [:CSV, :JSON, :Parquet, :Avro, :ORC, :XML]


  composed_of :connectionInfo,
    class_name: "ConnectionInfo",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[connectionInfo_username username]
    ]

  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :ProducedDatasets, class_name: 'DataSet'
  has_many :Pipelines, class_name: 'DataPipeline'

end
