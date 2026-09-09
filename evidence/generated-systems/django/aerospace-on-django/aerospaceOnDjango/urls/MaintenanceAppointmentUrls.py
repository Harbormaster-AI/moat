from django.urls import path
from aerospaceOnDjango.views import MaintenanceAppointmentView

urlpatterns = [
    path('', MaintenanceAppointmentView.index, name='index'),
	path('create', MaintenanceAppointmentView.get, name='create'),
	path('get/<int:maintenanceAppointmentId>/', MaintenanceAppointmentView.get, name='get'),
	path('save', MaintenanceAppointmentView.save, name='save'),
	path('getAll', MaintenanceAppointmentView.getAll, name='getAll'),
	path('delete/<int:maintenanceAppointmentId>/', MaintenanceAppointmentView.delete, name='delete'),
	path('assignAircraft/<int:maintenanceAppointmentId>/<int:AircraftId>/', MaintenanceAppointmentView.assignAircraft, name='assignAircraft'),
	path('unassignAircraft/<int:maintenanceAppointmentId>/', MaintenanceAppointmentView.unassignAircraft, name='unassignAircraft'),
	path('assignMroFacility/<int:maintenanceAppointmentId>/<int:MroFacilityId>/', MaintenanceAppointmentView.assignMroFacility, name='assignMroFacility'),
	path('unassignMroFacility/<int:maintenanceAppointmentId>/', MaintenanceAppointmentView.unassignMroFacility, name='unassignMroFacility'),
	path('assignWorkOrder/<int:maintenanceAppointmentId>/<int:WorkOrderId>/', MaintenanceAppointmentView.assignWorkOrder, name='assignWorkOrder'),
	path('unassignWorkOrder/<int:maintenanceAppointmentId>/', MaintenanceAppointmentView.unassignWorkOrder, name='unassignWorkOrder'),
]
