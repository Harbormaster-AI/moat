class Procedure < ApplicationRecord
  enum Status: [:Draft, :InReview, :Approved, :Retired]


  has_many :Policy, class_name: 'Policy'
  has_many :Controls, class_name: 'Control'

end
