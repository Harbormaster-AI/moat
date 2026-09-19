class RetentionSchedule < ApplicationRecord
  enum RetentionTrigger: [:CreationDate, :LastModified, :Termination, :ContractEnd, :EventCompletion, :FiscalYearEnd]
  enum DispositionAction: [:Destroy, :TransferToArchive, :Review, :SecureDelete, :ReturnToOwner]
  enum Status: [:Draft, :Approved, :InEffect, :Suspended, :Retired]


  has_many :Repositories, class_name: 'RecordsRepository'
  has_many :Records, class_name: 'Record_'
  has_many :Exceptions, class_name: 'Exception_'
  has_many :DispositionReviews, class_name: 'DispositionReview'

end
