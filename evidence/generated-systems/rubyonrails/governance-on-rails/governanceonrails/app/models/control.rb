class Control < ApplicationRecord
  enum ControlType: [:Preventive, :Detective, :Corrective, :Directive]
  enum Frequency: [:Continuous, :Daily, :Weekly, :Monthly, :Quarterly, :Annually, :AdHoc]
  enum Status: [:Designed, :Implemented, :Operating, :Retired]


  has_many :Policy, class_name: 'Policy'
  has_many :ControlTests, class_name: 'ControlTest_'
  has_many :Evidence, class_name: 'Evidence'
  has_many :Risks, class_name: 'Risk'
  has_many :Obligations, class_name: 'Obligation'
  has_many :Procedures, class_name: 'Procedure'
  has_many :Issues, class_name: 'Issue'

end
