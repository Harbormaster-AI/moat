class AuditWorkpaper < ApplicationRecord


  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :Engagement, class_name: 'AuditEngagement'
  has_many :Evidence, class_name: 'Evidence'
  has_many :Findings, class_name: 'AuditFinding'

end
