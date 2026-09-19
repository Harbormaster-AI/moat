Rails.application.routes.draw do
  root "application#health"
  resources :enterprises do
    resources :businessunits
    resources :plants
    resources :suppliers
    resources :customers
  end
  resources :businessunits do
    resource :enterprise
    resources :items
    resources :plants
  end
  resources :plants do
    resource :enterprise
    resources :productionlines
    resources :workcenters
    resources :warehouses
    resources :assets
    resources :productionschedules
  end
  resources :productionlines do
    resource :plant
    resources :workcenters
  end
  resources :workcenters do
    resource :productionline
    resources :assets
    resources :maintenanceorders
  end
  resources :items do
    resource :businessunit
    resources :boms
    resources :routings
    resources :suppliers
    resources :qualityspecifications
    resources :inventoryitems
  end
  resources :boms do
    resource :parentitem
    resources :bomitems
  end
  resources :bomitems do
    resource :bom
    resource :component
  end
  resources :routings do
    resource :item
    resources :operations
  end
  resources :operations do
    resource :routing
    resource :workcenter
    resource :inspectionplan
  end
  resources :workorders do
    resource :item
    resource :plant
    resource :routing
    resource :bom
    resource :productionschedule
    resource :salesorder
  end
  resources :productionschedules do
    resource :plant
    resources :workorders
  end
  resources :suppliers do
    resources :enterprises
    resources :items
    resources :purchaseorders
  end
  resources :purchaseorders do
    resource :supplier
    resource :plant
    resources :lines
    resources :goodsreceipts
  end
  resources :purchaseorderlines do
    resource :purchaseorder
    resource :item
  end
  resources :goodsreceipts do
    resource :purchaseorder
    resource :warehouse
    resources :lines
  end
  resources :goodsreceiptlines do
    resource :goodsreceipt
    resource :item
    resource :inventorytransaction
  end
  resources :warehouses do
    resource :plant
    resources :locations
    resources :inventoryitems
  end
  resources :locations do
    resource :warehouse
    resources :inventoryitems
  end
  resources :inventoryitems do
    resource :item
    resource :location
  end
  resources :inventorytransactions do
    resource :item
    resource :location
    resource :workorder
    resource :purchaseorder
    resource :salesorder
  end
  resources :customers do
    resources :enterprises
    resources :salesorders
  end
  resources :salesorders do
    resource :customer
    resource :plant
    resources :lines
    resources :workorders
  end
  resources :salesorderlines do
    resource :salesorder
    resource :item
  end
  resources :qualityspecifications do
    resource :item
  end
  resources :inspectionplans do
    resource :item
    resources :characteristics
  end
  resources :inspectioncharacteristics do
    resource :inspectionplan
  end
  resources :inspectionlots do
    resource :item
    resource :workorder
    resource :goodsreceipt
    resources :results
  end
  resources :inspectionresults do
    resource :inspectionlot
    resource :characteristic
  end
  resources :nonconformances do
    resource :item
    resource :workorder
    resource :inspectionlot
    resource :correctiveaction
  end
  resources :correctiveactions do
    resource :nonconformance
    resource :owner
  end
  resources :assets do
    resource :plant
    resource :workcenter
    resources :maintenanceorders
    resources :maintenanceplans
  end
  resources :maintenanceplans do
    resource :asset
    resources :maintenanceorders
  end
  resources :maintenanceorders do
    resource :asset
    resource :plan
    resource :workcenter
  end
  resources :employees do
    resource :workcenter
    resources :shiftassignments
    resources :correctiveactions
  end
  resources :shifts do
    resource :plant
    resources :assignments
  end
  resources :shiftassignments do
    resource :shift
    resource :employee
    resource :workcenter
  end
  resources :forecasts do
    resources :lines
  end
  resources :forecastlines do
    resource :forecast
    resource :item
  end
  resources :mrpruns do
    resource :plant
    resources :plannedorders
  end
  resources :plannedorders do
    resource :mrprun
    resource :item
    resource :plant
  end
end
