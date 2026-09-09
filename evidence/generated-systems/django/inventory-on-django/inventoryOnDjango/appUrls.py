"""mainsite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
urlpatterns = [
    path('StockKeepingUnit/', include('inventoryOnDjango.urls.StockKeepingUnitUrls')),
    path('Warehouse/', include('inventoryOnDjango.urls.WarehouseUrls')),
    path('StorageLocation/', include('inventoryOnDjango.urls.StorageLocationUrls')),
    path('InventoryItem/', include('inventoryOnDjango.urls.InventoryItemUrls')),
    path('Lot/', include('inventoryOnDjango.urls.LotUrls')),
    path('SerialNumber/', include('inventoryOnDjango.urls.SerialNumberUrls')),
    path('Reservation/', include('inventoryOnDjango.urls.ReservationUrls')),
    path('DemandSignal/', include('inventoryOnDjango.urls.DemandSignalUrls')),
    path('InventoryTransaction/', include('inventoryOnDjango.urls.InventoryTransactionUrls')),
    path('TransferOrder/', include('inventoryOnDjango.urls.TransferOrderUrls')),
    path('TransferOrderLine/', include('inventoryOnDjango.urls.TransferOrderLineUrls')),
    path('StockAdjustment/', include('inventoryOnDjango.urls.StockAdjustmentUrls')),
    path('StockAdjustmentLine/', include('inventoryOnDjango.urls.StockAdjustmentLineUrls')),
    path('CycleCount/', include('inventoryOnDjango.urls.CycleCountUrls')),
    path('CycleCountEntry/', include('inventoryOnDjango.urls.CycleCountEntryUrls')),
    path('ReplenishmentPolicy/', include('inventoryOnDjango.urls.ReplenishmentPolicyUrls')),
    path('UoMConversion/', include('inventoryOnDjango.urls.UoMConversionUrls')),
    path('InventoryThresholdAlert/', include('inventoryOnDjango.urls.InventoryThresholdAlertUrls')),
    path('Quarantine/', include('inventoryOnDjango.urls.QuarantineUrls')),
    path('ExpirationPolicy/', include('inventoryOnDjango.urls.ExpirationPolicyUrls')),
    path('InboundShipment/', include('inventoryOnDjango.urls.InboundShipmentUrls')),
    path('InboundShipmentLine/', include('inventoryOnDjango.urls.InboundShipmentLineUrls')),
    path('OutboundAllocation/', include('inventoryOnDjango.urls.OutboundAllocationUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]