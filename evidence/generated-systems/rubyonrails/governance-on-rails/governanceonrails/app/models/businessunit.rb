class BusinessUnit < ApplicationRecord


  has_many :Organization, class_name: 'Organization'
  has_many :Audits, class_name: 'AuditEngagement'

end
