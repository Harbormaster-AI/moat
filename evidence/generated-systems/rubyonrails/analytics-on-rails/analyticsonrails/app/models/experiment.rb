class Experiment < ApplicationRecord
  enum Status: [:Planned, :Running, :Completed, :Failed, :Stopped]


  has_many :Workspace, class_name: 'AnalyticsWorkspace'
  has_many :TrainingRuns, class_name: 'TrainingRun'
  has_many :Models, class_name: 'Model'
  has_many :Notebooks, class_name: 'Notebook'

end
