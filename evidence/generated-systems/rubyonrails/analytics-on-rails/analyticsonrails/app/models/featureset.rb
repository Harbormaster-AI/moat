class FeatureSet < ApplicationRecord
  enum StoreType: [:Online, :Offline, :Hybrid]


  composed_of :cronSchedule,
    class_name: "CronSchedule",
    mapping: [
      ${$mapping}, 
      %w[cronSchedule_timezone timezone]
    ]

  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :Features, class_name: 'Feature'
  has_many :Datasets, class_name: 'DataSet'
  has_many :Models, class_name: 'Model'
  has_many :ModelVersions, class_name: 'ModelVersion'
  has_many :Tags, class_name: 'Tag'

end
