class MedicationOrder < ApplicationRecord
  enum Route: [:Oral, :Intravenous, :Subcutaneous, :Intramuscular, :Topical, :Inhalation]


  composed_of :dose,
    class_name: "Dose",
    mapping: [
      ${$mapping}, 
      %w[dose_unit unit]
    ]

  has_many :Order, class_name: 'ClinicalOrder'
  has_many :Pharmacy, class_name: 'Pharmacy'
  has_many :Dispenses, class_name: 'MedicationDispense'

end
