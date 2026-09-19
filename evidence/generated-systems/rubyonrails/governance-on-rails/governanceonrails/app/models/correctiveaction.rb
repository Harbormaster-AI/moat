class CorrectiveAction < ApplicationRecord
  enum Status: [:NotStarted, :InProgress, :Deferred, :Completed, :Cancelled]


  has_many :Finding, class_name: 'AuditFinding'
  has_many :Issue, class_name: 'Issue'

end
