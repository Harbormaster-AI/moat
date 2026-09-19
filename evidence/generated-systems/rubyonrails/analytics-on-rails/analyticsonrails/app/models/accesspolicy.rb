class AccessPolicy < ApplicationRecord
  enum AccessLevel: [:View, :Query, :Modify, :Admin]
  enum SubjectType: [:User, :Group, :Service]


  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :Datasets, class_name: 'DataSet'
  has_many :Dashboards, class_name: 'Dashboard'
  has_many :Reports, class_name: 'Report'
  has_many :Models, class_name: 'Model'
  has_many :FeatureSets, class_name: 'FeatureSet'

end
