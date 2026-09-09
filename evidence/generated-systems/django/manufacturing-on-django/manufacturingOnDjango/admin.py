from django.contrib import admin

# Register your models here.
from .models.Enterprise import Enterprise
from .models.BusinessUnit import BusinessUnit
from .models.Plant import Plant
from .models.ProductionLine import ProductionLine
from .models.WorkCenter import WorkCenter
from .models.Item import Item
from .models.BOM import BOM
from .models.BOMItem import BOMItem
from .models.Routing import Routing
from .models.Operation import Operation
from .models.WorkOrder import WorkOrder
from .models.ProductionSchedule import ProductionSchedule
from .models.Supplier import Supplier
from .models.PurchaseOrder import PurchaseOrder
from .models.PurchaseOrderLine import PurchaseOrderLine
from .models.GoodsReceipt import GoodsReceipt
from .models.GoodsReceiptLine import GoodsReceiptLine
from .models.Warehouse import Warehouse
from .models.Location import Location
from .models.InventoryItem import InventoryItem
from .models.InventoryTransaction import InventoryTransaction
from .models.Customer import Customer
from .models.SalesOrder import SalesOrder
from .models.SalesOrderLine import SalesOrderLine
from .models.QualitySpecification import QualitySpecification
from .models.InspectionPlan import InspectionPlan
from .models.InspectionCharacteristic import InspectionCharacteristic
from .models.InspectionLot import InspectionLot
from .models.InspectionResult import InspectionResult
from .models.Nonconformance import Nonconformance
from .models.CorrectiveAction import CorrectiveAction
from .models.Asset import Asset
from .models.MaintenancePlan import MaintenancePlan
from .models.MaintenanceOrder import MaintenanceOrder
from .models.Employee import Employee
from .models.Shift import Shift
from .models.ShiftAssignment import ShiftAssignment
from .models.Forecast import Forecast
from .models.ForecastLine import ForecastLine
from .models.MRPRun import MRPRun
from .models.PlannedOrder import PlannedOrder

# Need to add this for each model that requires managing

admin.site.register(Enterprise)
admin.site.register(BusinessUnit)
admin.site.register(Plant)
admin.site.register(ProductionLine)
admin.site.register(WorkCenter)
admin.site.register(Item)
admin.site.register(BOM)
admin.site.register(BOMItem)
admin.site.register(Routing)
admin.site.register(Operation)
admin.site.register(WorkOrder)
admin.site.register(ProductionSchedule)
admin.site.register(Supplier)
admin.site.register(PurchaseOrder)
admin.site.register(PurchaseOrderLine)
admin.site.register(GoodsReceipt)
admin.site.register(GoodsReceiptLine)
admin.site.register(Warehouse)
admin.site.register(Location)
admin.site.register(InventoryItem)
admin.site.register(InventoryTransaction)
admin.site.register(Customer)
admin.site.register(SalesOrder)
admin.site.register(SalesOrderLine)
admin.site.register(QualitySpecification)
admin.site.register(InspectionPlan)
admin.site.register(InspectionCharacteristic)
admin.site.register(InspectionLot)
admin.site.register(InspectionResult)
admin.site.register(Nonconformance)
admin.site.register(CorrectiveAction)
admin.site.register(Asset)
admin.site.register(MaintenancePlan)
admin.site.register(MaintenanceOrder)
admin.site.register(Employee)
admin.site.register(Shift)
admin.site.register(ShiftAssignment)
admin.site.register(Forecast)
admin.site.register(ForecastLine)
admin.site.register(MRPRun)
admin.site.register(PlannedOrder)
