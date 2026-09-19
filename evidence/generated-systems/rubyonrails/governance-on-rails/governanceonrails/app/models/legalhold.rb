class LegalHold < ApplicationRecord
  enum HoldStatus: [:Active, :Released, :Superseded]


  has_many :Repositories, class_name: 'RecordsRepository'
  has_many :Records, class_name: 'Record_'
  has_many :Matter, class_name: 'Matter'

end
