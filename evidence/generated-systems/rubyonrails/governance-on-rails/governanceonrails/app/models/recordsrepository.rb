class RecordsRepository < ApplicationRecord
  enum RepositoryType: [:DocumentManagement, :RecordsArchive, :EmailArchive, :FileShare, :ContentServices, :DataLake]


  has_many :Organization, class_name: 'Organization'
  has_many :Records, class_name: 'Record_'
  has_many :Systems, class_name: 'System_'
  has_many :RetentionSchedules, class_name: 'RetentionSchedule'
  has_many :LegalHolds, class_name: 'LegalHold'

end
