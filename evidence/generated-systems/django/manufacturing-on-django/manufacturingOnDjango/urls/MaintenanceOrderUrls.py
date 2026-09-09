from django.urls import path
from manufacturingOnDjango.views import MaintenanceOrderView

urlpatterns = [
    path('', MaintenanceOrderView.index, name='index'),
	path('create', MaintenanceOrderView.get, name='create'),
	path('get/<int:maintenanceOrderId>/', MaintenanceOrderView.get, name='get'),
	path('save', MaintenanceOrderView.save, name='save'),
	path('getAll', MaintenanceOrderView.getAll, name='getAll'),
	path('delete/<int:maintenanceOrderId>/', MaintenanceOrderView.delete, name='delete'),
	path('assignAsset/<int:maintenanceOrderId>/<int:AssetId>/', MaintenanceOrderView.assignAsset, name='assignAsset'),
	path('unassignAsset/<int:maintenanceOrderId>/', MaintenanceOrderView.unassignAsset, name='unassignAsset'),
	path('assignPlan/<int:maintenanceOrderId>/<int:PlanId>/', MaintenanceOrderView.assignPlan, name='assignPlan'),
	path('unassignPlan/<int:maintenanceOrderId>/', MaintenanceOrderView.unassignPlan, name='unassignPlan'),
	path('assignWorkCenter/<int:maintenanceOrderId>/<int:WorkCenterId>/', MaintenanceOrderView.assignWorkCenter, name='assignWorkCenter'),
	path('unassignWorkCenter/<int:maintenanceOrderId>/', MaintenanceOrderView.unassignWorkCenter, name='unassignWorkCenter'),
]
