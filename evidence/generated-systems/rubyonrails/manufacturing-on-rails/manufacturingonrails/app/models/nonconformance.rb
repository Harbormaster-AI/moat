class Nonconformance < ApplicationRecord
  enum NcType: [:Dimension, :Functional, :Cosmetic, :Documentation, :Supplier, :Process]
  enum Severity: [:Minor, :Major, :Critical]
  enum Status: [:Open, :Contained, :UnderInvestigation, :Dispositioned, :Closed]


  has_many :Item, class_name: 'Item'
  has_many :WorkOrder, class_name: 'WorkOrder'
  has_many :InspectionLot, class_name: 'InspectionLot'
  has_many :CorrectiveAction, class_name: 'CorrectiveAction'

end
