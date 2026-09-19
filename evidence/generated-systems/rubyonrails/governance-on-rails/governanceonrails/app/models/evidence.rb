class Evidence < ApplicationRecord
  enum EvidenceType: [:Document, :Screenshot, :LogExport, :SystemReport, :Ticket, :Attestation, :Configuration, :Dataset]


  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :ControlTest, class_name: 'ControlTest_'
  has_many :Control, class_name: 'Control'
  has_many :Obligation, class_name: 'Obligation'
  has_many :Workpaper, class_name: 'AuditWorkpaper'

end
