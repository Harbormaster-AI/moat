from django.urls import path
from manufacturingOnDjango.views import MaintenancePlanView

urlpatterns = [
    path('', MaintenancePlanView.index, name='index'),
	path('create', MaintenancePlanView.get, name='create'),
	path('get/<int:maintenancePlanId>/', MaintenancePlanView.get, name='get'),
	path('save', MaintenancePlanView.save, name='save'),
	path('getAll', MaintenancePlanView.getAll, name='getAll'),
	path('delete/<int:maintenancePlanId>/', MaintenancePlanView.delete, name='delete'),
	path('assignAsset/<int:maintenancePlanId>/<int:AssetId>/', MaintenancePlanView.assignAsset, name='assignAsset'),
	path('unassignAsset/<int:maintenancePlanId>/', MaintenancePlanView.unassignAsset, name='unassignAsset'),
	path('addMaintenanceOrders/<int:maintenancePlanId>/<MaintenanceOrdersIds>/', MaintenancePlanView.addMaintenanceOrders, name='addMaintenanceOrders'),
	path('removeMaintenanceOrders/<int:maintenancePlanId>/<MaintenanceOrdersIds>/', MaintenancePlanView.removeMaintenanceOrders, name='removeMaintenanceOrders'),
]
