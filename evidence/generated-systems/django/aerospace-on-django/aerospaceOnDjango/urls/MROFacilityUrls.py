from django.urls import path
from aerospaceOnDjango.views import MROFacilityView

urlpatterns = [
    path('', MROFacilityView.index, name='index'),
	path('create', MROFacilityView.get, name='create'),
	path('get/<int:mROFacilityId>/', MROFacilityView.get, name='get'),
	path('save', MROFacilityView.save, name='save'),
	path('getAll', MROFacilityView.getAll, name='getAll'),
	path('delete/<int:mROFacilityId>/', MROFacilityView.delete, name='delete'),
	path('addAppointments/<int:mROFacilityId>/<AppointmentsIds>/', MROFacilityView.addAppointments, name='addAppointments'),
	path('removeAppointments/<int:mROFacilityId>/<AppointmentsIds>/', MROFacilityView.removeAppointments, name='removeAppointments'),
	path('addWorkOrders/<int:mROFacilityId>/<WorkOrdersIds>/', MROFacilityView.addWorkOrders, name='addWorkOrders'),
	path('removeWorkOrders/<int:mROFacilityId>/<WorkOrdersIds>/', MROFacilityView.removeWorkOrders, name='removeWorkOrders'),
]
