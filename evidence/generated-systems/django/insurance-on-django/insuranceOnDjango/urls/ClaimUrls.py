from django.urls import path
from insuranceOnDjango.views import ClaimView

urlpatterns = [
    path('', ClaimView.index, name='index'),
	path('create', ClaimView.get, name='create'),
	path('get/<int:claimId>/', ClaimView.get, name='get'),
	path('save', ClaimView.save, name='save'),
	path('getAll', ClaimView.getAll, name='getAll'),
	path('delete/<int:claimId>/', ClaimView.delete, name='delete'),
	path('assignPolicy/<int:claimId>/<int:PolicyId>/', ClaimView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:claimId>/', ClaimView.unassignPolicy, name='unassignPolicy'),
	path('assignCustomer/<int:claimId>/<int:CustomerId>/', ClaimView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:claimId>/', ClaimView.unassignCustomer, name='unassignCustomer'),
	path('assignAdjuster/<int:claimId>/<int:AdjusterId>/', ClaimView.assignAdjuster, name='assignAdjuster'),
	path('unassignAdjuster/<int:claimId>/', ClaimView.unassignAdjuster, name='unassignAdjuster'),
	path('assignIncident/<int:claimId>/<int:IncidentId>/', ClaimView.assignIncident, name='assignIncident'),
	path('unassignIncident/<int:claimId>/', ClaimView.unassignIncident, name='unassignIncident'),
	path('addExposures/<int:claimId>/<ExposuresIds>/', ClaimView.addExposures, name='addExposures'),
	path('removeExposures/<int:claimId>/<ExposuresIds>/', ClaimView.removeExposures, name='removeExposures'),
	path('addReserves/<int:claimId>/<ReservesIds>/', ClaimView.addReserves, name='addReserves'),
	path('removeReserves/<int:claimId>/<ReservesIds>/', ClaimView.removeReserves, name='removeReserves'),
	path('addClaimPayments/<int:claimId>/<ClaimPaymentsIds>/', ClaimView.addClaimPayments, name='addClaimPayments'),
	path('removeClaimPayments/<int:claimId>/<ClaimPaymentsIds>/', ClaimView.removeClaimPayments, name='removeClaimPayments'),
	path('addServiceProviders/<int:claimId>/<ServiceProvidersIds>/', ClaimView.addServiceProviders, name='addServiceProviders'),
	path('removeServiceProviders/<int:claimId>/<ServiceProvidersIds>/', ClaimView.removeServiceProviders, name='removeServiceProviders'),
	path('addSubrogations/<int:claimId>/<SubrogationsIds>/', ClaimView.addSubrogations, name='addSubrogations'),
	path('removeSubrogations/<int:claimId>/<SubrogationsIds>/', ClaimView.removeSubrogations, name='removeSubrogations'),
]
