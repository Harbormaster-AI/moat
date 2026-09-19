Rails.application.routes.draw do
  root "application#health"
  resources :healthsystems do
    resources :facilities
    resources :suppliers
  end
  resources :facilitys do
    resource :healthsystem
    resources :departments
    resources :careteams
    resources :laboratories
    resources :imagingcenters
    resources :pharmacies
    resources :inventoryitems
  end
  resources :departments do
    resource :facility
    resources :careteams
  end
  resources :careteams do
    resource :department
    resources :clinicians
    resources :patients
  end
  resources :clinicians do
    resources :careteams
    resources :appointments
    resources :encounters
    resources :procedures
    resources :imagingreports
  end
  resources :patients do
    resources :appointments
    resources :encounters
    resources :careplans
    resources :allergies
    resources :conditions
    resources :medicationorders
    resources :laborders
    resources :imagingorders
    resources :coverages
    resources :claims
    resources :devices
    resources :observations
  end
  resources :appointments do
    resource :patient
    resource :clinician
    resource :facility
    resource :encounter
  end
  resources :encounters do
    resource :patient
    resource :clinician
    resource :facility
    resource :appointment
    resources :diagnoses
    resources :procedures
    resources :observations
    resources :orders
    resource :admission
    resource :discharge
  end
  resources :admissions do
    resource :encounter
    resource :facility
  end
  resources :discharges do
    resource :encounter
  end
  resources :clinicalorders do
    resource :patient
    resource :encounter
    resource :orderingclinician
    resources :medicationorders
    resources :laboratoryorders
    resources :imagingorders
    resources :procedureorders
    resources :authorizations
  end
  resources :medicationorders do
    resource :order
    resource :pharmacy
    resources :dispenses
  end
  resources :laboratorys do
    resource :facility
    resources :laboratoryorders
    resources :labresults
  end
  resources :laboratoryorders do
    resource :order
    resource :laboratory
    resources :results
  end
  resources :labresults do
    resource :laboratoryorder
    resources :observations
    resource :laboratory
  end
  resources :imagingcenters do
    resource :facility
    resources :imagingorders
    resources :imagingreports
  end
  resources :imagingorders do
    resource :order
    resource :imagingcenter
    resources :reports
  end
  resources :imagingreports do
    resource :imagingorder
    resource :clinician
    resource :encounter
    resource :imagingcenter
  end
  resources :procedureorders do
    resource :order
    resource :facility
    resource :procedure
  end
  resources :procedures do
    resource :encounter
    resource :performer
    resource :procedureorder
  end
  resources :pharmacys do
    resource :facility
    resources :medicationdispenses
    resources :medicationorders
  end
  resources :medicationdispenses do
    resource :medicationorder
    resource :pharmacy
    resource :patient
  end
  resources :diagnosiss do
    resource :encounter
    resource :patient
  end
  resources :observations do
    resource :encounter
    resource :patient
    resource :device
    resource :labresult
  end
  resources :careplans do
    resource :patient
    resources :encounters
    resources :tasks
    resource :careteam
  end
  resources :caretasks do
    resource :careplan
    resource :assignedto
    resource :encounter
  end
  resources :allergys do
    resource :patient
  end
  resources :conditions do
    resource :patient
  end
  resources :insurancepayers do
    resources :plans
    resources :claims
  end
  resources :insuranceplans do
    resource :payer
    resources :coverages
  end
  resources :coverages do
    resource :patient
    resource :plan
    resources :claims
    resources :authorizations
  end
  resources :claims do
    resource :patient
    resource :coverage
    resource :encounter
    resources :invoices
    resource :payer
  end
  resources :authorizations do
    resource :coverage
    resource :order
  end
  resources :invoices do
    resource :patient
    resource :claim
    resources :payments
  end
  resources :payments do
    resource :invoice
    resource :payer
  end
  resources :medicaldevices do
    resource :patient
    resources :observations
    resources :softwareupdates
  end
  resources :softwareupdates do
    resource :device
  end
  resources :medicalsuppliers do
    resources :facilities
    resources :inventoryitems
  end
  resources :inventoryitems do
    resource :facility
    resource :supplier
  end
end
