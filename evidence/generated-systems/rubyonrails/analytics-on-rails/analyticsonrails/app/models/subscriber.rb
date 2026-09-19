class Subscriber < ApplicationRecord
  enum Channel: [:Email, :SMS, :Webhook, :Chat]


  has_many :Alerts, class_name: 'Alert'

end
