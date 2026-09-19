Rails.application.routes.draw do
  root "application#health"
  resources :aerospacemanufacturers do
    resources :programs
    resources :plants
    resources :suppliers
    resources :productioncertificates
  end
  resources :aircraftprograms do
    resource :manufacturer
    resources :aircraftfamilies
    resource :typecertificate
    resources :keysuppliers
  end
  resources :aircraftfamilys do
    resource :program
    resources :aircraftmodels
  end
  resources :aircraftmodels do
    resource :family
    resources :variants
    resources :enginetypes
  end
  resources :enginetypes do
    resource :supplier
    resources :compatiblemodels
  end
  resources :aircraftvariants do
    resource :model
    resource :enginetype
    resource :avionicssuite
    resource :apu
    resource :landinggear
    resources :cabinlayouts
    resources :options
    resources :packages
  end
  resources :avionicssuites do
    resource :supplier
    resources :variants
    resources :softwareloads
  end
  resources :apus do
    resource :supplier
    resources :variants
  end
  resources :landinggears do
    resource :supplier
    resources :variants
  end
  resources :aircraftoptions do
    resources :variants
    resources :packages
  end
  resources :aircraftpackages do
    resources :options
    resources :variants
  end
  resources :suppliers do
    resources :manufacturers
    resources :components
    resources :enginetypes
    resources :avionicssuites
    resources :apus
    resources :landinggears
  end
  resources :component_s do
    resource :supplier
  end
  resources :plants do
    resource :manufacturer
    resources :productionlines
    resources :warehouses
  end
  resources :productionlines do
    resource :plant
    resources :workcenters
  end
  resources :workcenters do
    resource :productionline
  end
  resources :productionorders do
    resource :variant
    resource :plant
    resource :aircraftorder
  end
  resources :buildschedules do
    resources :productionorders
  end
  resources :warehouses do
    resources :inventoryitems
  end
  resources :inventoryitems do
    resource :component
    resource :warehouse
  end
  resources :operators do
    resources :aircraftorders
    resources :operatedaircraft
    resource :salesregion
  end
  resources :aircraftorders do
    resource :operator
    resource :variant
    resource :quote
    resource :purchaseagreement
  end
  resources :quotes do
    resource :aircraftorder
  end
  resources :purchaseagreements do
    resource :aircraftorder
  end
  resources :aircrafts do
    resource :variant
    resource :operator
    resource :registration
    resource :warranty
    resources :maintenancerecords
    resource :connectedaircraft
    resource :cabinlayout
  end
  resources :registrations do
    resource :aircraft
  end
  resources :warrantys do
    resource :aircraft
  end
  resources :cabinlayouts do
    resource :variant
    resources :aircraft
    resources :options
  end
  resources :mrofacilitys do
    resources :appointments
    resources :workorders
  end
  resources :maintenanceappointments do
    resource :aircraft
    resource :mrofacility
    resource :workorder
  end
  resources :maintenanceworkorders do
    resource :aircraft
    resource :airworthinessdirective
    resource :servicebulletin
  end
  resources :airworthinessdirectives do
    resources :workorders
  end
  resources :servicebulletins do
    resources :workorders
    resources :variants
  end
  resources :connectedaircrafts do
    resource :aircraft
    resources :flighthealthevents
    resources :softwareloads
  end
  resources :flighthealthevents do
    resource :connectedaircraft
  end
  resources :softwareloads do
    resource :connectedaircraft
    resource :avionicssuite
  end
  resources :typecertificates do
    resource :program
  end
  resources :productioncertificates do
    resource :manufacturer
  end
  resources :salesregions do
    resources :operators
    resources :salescampaigns
  end
  resources :salescampaigns do
    resource :region
    resource :operator
    resources :quotes
  end
end
