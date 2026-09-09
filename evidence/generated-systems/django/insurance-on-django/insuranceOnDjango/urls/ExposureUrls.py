from django.urls import path
from insuranceOnDjango.views import ExposureView

urlpatterns = [
    path('', ExposureView.index, name='index'),
	path('create', ExposureView.get, name='create'),
	path('get/<int:exposureId>/', ExposureView.get, name='get'),
	path('save', ExposureView.save, name='save'),
	path('getAll', ExposureView.getAll, name='getAll'),
	path('delete/<int:exposureId>/', ExposureView.delete, name='delete'),
	path('assignClaim/<int:exposureId>/<int:ClaimId>/', ExposureView.assignClaim, name='assignClaim'),
	path('unassignClaim/<int:exposureId>/', ExposureView.unassignClaim, name='unassignClaim'),
	path('assignPolicyCoverage/<int:exposureId>/<int:PolicyCoverageId>/', ExposureView.assignPolicyCoverage, name='assignPolicyCoverage'),
	path('unassignPolicyCoverage/<int:exposureId>/', ExposureView.unassignPolicyCoverage, name='unassignPolicyCoverage'),
	path('assignInsuredObject/<int:exposureId>/<int:InsuredObjectId>/', ExposureView.assignInsuredObject, name='assignInsuredObject'),
	path('unassignInsuredObject/<int:exposureId>/', ExposureView.unassignInsuredObject, name='unassignInsuredObject'),
	path('addReserves/<int:exposureId>/<ReservesIds>/', ExposureView.addReserves, name='addReserves'),
	path('removeReserves/<int:exposureId>/<ReservesIds>/', ExposureView.removeReserves, name='removeReserves'),
	path('addPayments/<int:exposureId>/<PaymentsIds>/', ExposureView.addPayments, name='addPayments'),
	path('removePayments/<int:exposureId>/<PaymentsIds>/', ExposureView.removePayments, name='removePayments'),
]
