class WorkAuthorization < ApplicationRecord
  enum Status: [:NotRequired, :Pending, :Authorized, :Expired]


  has_many :Employee, class_name: 'Employee'
  has_many :Documents, class_name: 'Document'

end
