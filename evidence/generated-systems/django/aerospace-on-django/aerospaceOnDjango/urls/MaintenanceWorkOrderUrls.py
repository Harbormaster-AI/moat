from django.urls import path
from aerospaceOnDjango.views import MaintenanceWorkOrderView

urlpatterns = [
    path('', MaintenanceWorkOrderView.index, name='index'),
	path('create', MaintenanceWorkOrderView.get, name='create'),
	path('get/<int:maintenanceWorkOrderId>/', MaintenanceWorkOrderView.get, name='get'),
	path('save', MaintenanceWorkOrderView.save, name='save'),
	path('getAll', MaintenanceWorkOrderView.getAll, name='getAll'),
	path('delete/<int:maintenanceWorkOrderId>/', MaintenanceWorkOrderView.delete, name='delete'),
	path('assignAircraft/<int:maintenanceWorkOrderId>/<int:AircraftId>/', MaintenanceWorkOrderView.assignAircraft, name='assignAircraft'),
	path('unassignAircraft/<int:maintenanceWorkOrderId>/', MaintenanceWorkOrderView.unassignAircraft, name='unassignAircraft'),
	path('assignAirworthinessDirective/<int:maintenanceWorkOrderId>/<int:AirworthinessDirectiveId>/', MaintenanceWorkOrderView.assignAirworthinessDirective, name='assignAirworthinessDirective'),
	path('unassignAirworthinessDirective/<int:maintenanceWorkOrderId>/', MaintenanceWorkOrderView.unassignAirworthinessDirective, name='unassignAirworthinessDirective'),
	path('assignServiceBulletin/<int:maintenanceWorkOrderId>/<int:ServiceBulletinId>/', MaintenanceWorkOrderView.assignServiceBulletin, name='assignServiceBulletin'),
	path('unassignServiceBulletin/<int:maintenanceWorkOrderId>/', MaintenanceWorkOrderView.unassignServiceBulletin, name='unassignServiceBulletin'),
]
