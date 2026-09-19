class Application < ApplicationRecord
  enum Status: [:Draft, :Submitted, :UnderReview, :Quoted, :Declined, :Withdrawn, :Bound]


  has_many :Customer, class_name: 'Customer'
  has_many :Product, class_name: 'InsuranceProduct'
  has_many :Distributor, class_name: 'Distributor'
  has_many :Quotes, class_name: 'Quote'
  has_many :SelectedQuote, class_name: 'Quote'

end
