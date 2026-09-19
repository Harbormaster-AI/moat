class Organization < ApplicationRecord


  has_many :Departments, class_name: 'Department'
  has_many :Locations, class_name: 'Location'
  has_many :JobFamilies, class_name: 'JobFamily'
  has_many :BenefitPlans, class_name: 'BenefitPlan'
  has_many :CostCenters, class_name: 'CostCenter'
  has_many :PayrollCalendars, class_name: 'PayrollCalendar'

end
