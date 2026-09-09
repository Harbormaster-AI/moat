from django.contrib import admin

# Register your models here.
from .models.StockKeepingUnit import StockKeepingUnit
from .models.Warehouse import Warehouse
from .models.StorageLocation import StorageLocation
from .models.InventoryItem import InventoryItem
from .models.Lot import Lot
from .models.SerialNumber import SerialNumber
from .models.Reservation import Reservation
from .models.DemandSignal import DemandSignal
from .models.InventoryTransaction import InventoryTransaction
from .models.TransferOrder import TransferOrder
from .models.TransferOrderLine import TransferOrderLine
from .models.StockAdjustment import StockAdjustment
from .models.StockAdjustmentLine import StockAdjustmentLine
from .models.CycleCount import CycleCount
from .models.CycleCountEntry import CycleCountEntry
from .models.ReplenishmentPolicy import ReplenishmentPolicy
from .models.UoMConversion import UoMConversion
from .models.InventoryThresholdAlert import InventoryThresholdAlert
from .models.Quarantine import Quarantine
from .models.ExpirationPolicy import ExpirationPolicy
from .models.InboundShipment import InboundShipment
from .models.InboundShipmentLine import InboundShipmentLine
from .models.OutboundAllocation import OutboundAllocation

# Need to add this for each model that requires managing

admin.site.register(StockKeepingUnit)
admin.site.register(Warehouse)
admin.site.register(StorageLocation)
admin.site.register(InventoryItem)
admin.site.register(Lot)
admin.site.register(SerialNumber)
admin.site.register(Reservation)
admin.site.register(DemandSignal)
admin.site.register(InventoryTransaction)
admin.site.register(TransferOrder)
admin.site.register(TransferOrderLine)
admin.site.register(StockAdjustment)
admin.site.register(StockAdjustmentLine)
admin.site.register(CycleCount)
admin.site.register(CycleCountEntry)
admin.site.register(ReplenishmentPolicy)
admin.site.register(UoMConversion)
admin.site.register(InventoryThresholdAlert)
admin.site.register(Quarantine)
admin.site.register(ExpirationPolicy)
admin.site.register(InboundShipment)
admin.site.register(InboundShipmentLine)
admin.site.register(OutboundAllocation)
