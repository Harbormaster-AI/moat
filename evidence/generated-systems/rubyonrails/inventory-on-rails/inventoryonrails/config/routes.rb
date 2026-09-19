Rails.application.routes.draw do
  root "application#health"
  resources :stockkeepingunits do
    resources :inventoryitems
    resources :uomconversions
    resources :replenishmentpolicies
    resources :lots
    resources :serialnumbers
  end
  resources :warehouses do
    resources :storagelocations
    resources :inventoryitems
    resources :inboundshipments
    resources :outboundallocations
    resources :origintransfers
    resources :destinationtransfers
    resources :cyclecounts
  end
  resources :storagelocations do
    resource :warehouse
    resource :parentlocation
    resources :childlocations
    resources :inventoryitems
  end
  resources :inventoryitems do
    resource :sku
    resource :warehouse
    resource :location
    resource :lot
    resources :serialnumbers
    resources :transactions
    resources :reservations
  end
  resources :lots do
    resource :sku
    resources :inventoryitems
  end
  resources :serialnumbers do
    resource :sku
    resource :currentinventoryitem
    resource :lot
  end
  resources :reservations do
    resource :sku
    resource :warehouse
    resource :location
    resource :inventoryitem
    resource :lot
    resources :serialnumbers
    resource :demandsignal
  end
  resources :demandsignals do
    resource :sku
    resources :reservations
  end
  resources :inventorytransactions do
    resource :sku
    resource :warehouse
    resource :location
    resource :lot
    resources :serialnumbers
    resource :relatedreservation
    resource :transferorder
    resource :adjustment
    resource :cyclecount
  end
  resources :transferorders do
    resource :originwarehouse
    resource :destinationwarehouse
    resources :lines
    resources :transactions
  end
  resources :transferorderlines do
    resource :transferorder
    resource :sku
    resource :lot
    resources :serialnumbers
    resource :fromlocation
    resource :tolocation
  end
  resources :stockadjustments do
    resource :warehouse
    resources :lines
    resources :transactions
  end
  resources :stockadjustmentlines do
    resource :adjustment
    resource :sku
    resource :lot
    resource :location
    resources :serialnumbers
  end
  resources :cyclecounts do
    resource :warehouse
    resources :locations
    resources :entries
    resources :transactions
  end
  resources :cyclecountentrys do
    resource :cyclecount
    resource :sku
    resource :lot
    resource :location
    resources :serialnumbers
  end
  resources :replenishmentpolicys do
    resource :sku
    resource :warehouse
    resource :location
  end
  resources :uomconversions do
    resource :sku
  end
  resources :inventorythresholdalerts do
    resource :sku
    resource :warehouse
    resource :location
    resource :relatedpolicy
  end
  resources :quarantines do
    resource :warehouse
    resources :items
    resource :lot
    resources :serialnumbers
  end
  resources :expirationpolicys do
    resource :sku
    resource :warehouse
  end
  resources :inboundshipments do
    resource :warehouse
    resources :lines
    resources :transactions
  end
  resources :inboundshipmentlines do
    resource :inboundshipment
    resource :sku
    resource :lot
    resources :serialnumbers
    resource :destinationlocation
  end
  resources :outboundallocations do
    resource :warehouse
    resource :sku
    resource :inventoryitem
    resource :reservation
    resource :lot
    resources :serialnumbers
    resource :sourcelocation
  end
end
