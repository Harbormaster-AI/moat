class MROFacility < ApplicationRecord


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Appointments, class_name: 'MaintenanceAppointment'
  has_many :WorkOrders, class_name: 'MaintenanceWorkOrder'

end
