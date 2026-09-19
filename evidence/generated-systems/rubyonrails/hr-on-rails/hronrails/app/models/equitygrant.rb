class EquityGrant < ApplicationRecord
  enum GrantType: [:RSU, :StockOption, :ESPP]


  has_many :CompensationPackage, class_name: 'CompensationPackage'

end
