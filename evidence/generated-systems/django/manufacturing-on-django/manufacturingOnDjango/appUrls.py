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
    path('Enterprise/', include('manufacturingOnDjango.urls.EnterpriseUrls')),
    path('BusinessUnit/', include('manufacturingOnDjango.urls.BusinessUnitUrls')),
    path('Plant/', include('manufacturingOnDjango.urls.PlantUrls')),
    path('ProductionLine/', include('manufacturingOnDjango.urls.ProductionLineUrls')),
    path('WorkCenter/', include('manufacturingOnDjango.urls.WorkCenterUrls')),
    path('Item/', include('manufacturingOnDjango.urls.ItemUrls')),
    path('BOM/', include('manufacturingOnDjango.urls.BOMUrls')),
    path('BOMItem/', include('manufacturingOnDjango.urls.BOMItemUrls')),
    path('Routing/', include('manufacturingOnDjango.urls.RoutingUrls')),
    path('Operation/', include('manufacturingOnDjango.urls.OperationUrls')),
    path('WorkOrder/', include('manufacturingOnDjango.urls.WorkOrderUrls')),
    path('ProductionSchedule/', include('manufacturingOnDjango.urls.ProductionScheduleUrls')),
    path('Supplier/', include('manufacturingOnDjango.urls.SupplierUrls')),
    path('PurchaseOrder/', include('manufacturingOnDjango.urls.PurchaseOrderUrls')),
    path('PurchaseOrderLine/', include('manufacturingOnDjango.urls.PurchaseOrderLineUrls')),
    path('GoodsReceipt/', include('manufacturingOnDjango.urls.GoodsReceiptUrls')),
    path('GoodsReceiptLine/', include('manufacturingOnDjango.urls.GoodsReceiptLineUrls')),
    path('Warehouse/', include('manufacturingOnDjango.urls.WarehouseUrls')),
    path('Location/', include('manufacturingOnDjango.urls.LocationUrls')),
    path('InventoryItem/', include('manufacturingOnDjango.urls.InventoryItemUrls')),
    path('InventoryTransaction/', include('manufacturingOnDjango.urls.InventoryTransactionUrls')),
    path('Customer/', include('manufacturingOnDjango.urls.CustomerUrls')),
    path('SalesOrder/', include('manufacturingOnDjango.urls.SalesOrderUrls')),
    path('SalesOrderLine/', include('manufacturingOnDjango.urls.SalesOrderLineUrls')),
    path('QualitySpecification/', include('manufacturingOnDjango.urls.QualitySpecificationUrls')),
    path('InspectionPlan/', include('manufacturingOnDjango.urls.InspectionPlanUrls')),
    path('InspectionCharacteristic/', include('manufacturingOnDjango.urls.InspectionCharacteristicUrls')),
    path('InspectionLot/', include('manufacturingOnDjango.urls.InspectionLotUrls')),
    path('InspectionResult/', include('manufacturingOnDjango.urls.InspectionResultUrls')),
    path('Nonconformance/', include('manufacturingOnDjango.urls.NonconformanceUrls')),
    path('CorrectiveAction/', include('manufacturingOnDjango.urls.CorrectiveActionUrls')),
    path('Asset/', include('manufacturingOnDjango.urls.AssetUrls')),
    path('MaintenancePlan/', include('manufacturingOnDjango.urls.MaintenancePlanUrls')),
    path('MaintenanceOrder/', include('manufacturingOnDjango.urls.MaintenanceOrderUrls')),
    path('Employee/', include('manufacturingOnDjango.urls.EmployeeUrls')),
    path('Shift/', include('manufacturingOnDjango.urls.ShiftUrls')),
    path('ShiftAssignment/', include('manufacturingOnDjango.urls.ShiftAssignmentUrls')),
    path('Forecast/', include('manufacturingOnDjango.urls.ForecastUrls')),
    path('ForecastLine/', include('manufacturingOnDjango.urls.ForecastLineUrls')),
    path('MRPRun/', include('manufacturingOnDjango.urls.MRPRunUrls')),
    path('PlannedOrder/', include('manufacturingOnDjango.urls.PlannedOrderUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]