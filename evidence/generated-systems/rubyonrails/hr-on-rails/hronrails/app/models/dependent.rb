class Dependent < ApplicationRecord
  enum Relationship: [:Spouse, :DomesticPartner, :Child, :Other]


  has_many :BenefitEnrollment, class_name: 'BenefitEnrollment'
  has_many :Employee, class_name: 'Employee'

end
