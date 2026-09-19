class Registration < ApplicationRecord


  composed_of :tailNumber,
    class_name: "TailNumber",
    mapping: [
      ${$mapping}, 
      %w[tailNumber_country country]
    ]

  has_many :Aircraft, class_name: 'Aircraft'

end
