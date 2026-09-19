class Notebook < ApplicationRecord
  enum Language: [:Python, :R, :SQL, :Julia]


  composed_of :repositoryRef,
    class_name: "RepositoryRef",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[repositoryRef_path path]
    ]

  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :Datasets, class_name: 'DataSet'
  has_many :Experiments, class_name: 'Experiment'
  has_many :Queries, class_name: 'BIQuery'

end
