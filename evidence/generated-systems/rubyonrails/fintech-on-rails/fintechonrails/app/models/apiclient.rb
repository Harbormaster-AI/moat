class APIClient < ApplicationRecord
  enum ClientType: [:Confidential, :Public]


  has_many :Consents, class_name: 'Consent'

end
