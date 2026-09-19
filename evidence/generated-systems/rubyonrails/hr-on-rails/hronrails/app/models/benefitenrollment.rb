class BenefitEnrollment < ApplicationRecord
  enum Status: [:Pending, :Active, :Waived, :Cancelled, :Terminated]
  enum CoverageLevel: [:EmployeeOnly, :EmployeeSpouse, :EmployeeChildren, :Family]


  has_many :BenefitPlan, class_name: 'BenefitPlan'
  has_many :Employee, class_name: 'Employee'
  has_many :Dependents, class_name: 'Dependent'

end
