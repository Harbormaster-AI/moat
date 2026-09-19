class BonusPlan < ApplicationRecord


  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :CompensationPackages, class_name: 'CompensationPackage'

end
