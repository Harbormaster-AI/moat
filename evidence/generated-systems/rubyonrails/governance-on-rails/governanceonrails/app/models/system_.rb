class System_ < ApplicationRecord
  enum SystemType: [:Application, :Database, :DataWarehouse, :SaaS, :Infrastructure, :Endpoint]


  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :RecordsRepositories, class_name: 'RecordsRepository'

end
