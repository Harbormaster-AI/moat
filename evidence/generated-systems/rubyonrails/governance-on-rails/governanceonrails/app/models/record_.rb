class Record_ < ApplicationRecord
  enum RecordType: [:PolicyRecord, :ContractRecord, :FinancialRecord, :HRRecord, :CustomerRecord, :TechnicalRecord, :AuditRecord, :LegalRecord]
  enum Classification: [:Public, :Internal, :Confidential, :Restricted, :HighlyRestricted]
  enum Status: [:Active, :Archived, :PendingDisposition, :Disposed, :OnHold]


  has_many :Repository, class_name: 'RecordsRepository'
  has_many :RetentionSchedule, class_name: 'RetentionSchedule'
  has_many :ProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :DataCategories, class_name: 'DataCategory'
  has_many :LegalHolds, class_name: 'LegalHold'
  has_many :DataSubjectRequests, class_name: 'DataSubjectRequest'

end
