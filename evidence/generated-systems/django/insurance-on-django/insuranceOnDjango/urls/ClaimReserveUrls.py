from django.urls import path
from insuranceOnDjango.views import ClaimReserveView

urlpatterns = [
    path('', ClaimReserveView.index, name='index'),
	path('create', ClaimReserveView.get, name='create'),
	path('get/<int:claimReserveId>/', ClaimReserveView.get, name='get'),
	path('save', ClaimReserveView.save, name='save'),
	path('getAll', ClaimReserveView.getAll, name='getAll'),
	path('delete/<int:claimReserveId>/', ClaimReserveView.delete, name='delete'),
	path('assignClaim/<int:claimReserveId>/<int:ClaimId>/', ClaimReserveView.assignClaim, name='assignClaim'),
	path('unassignClaim/<int:claimReserveId>/', ClaimReserveView.unassignClaim, name='unassignClaim'),
	path('assignExposure/<int:claimReserveId>/<int:ExposureId>/', ClaimReserveView.assignExposure, name='assignExposure'),
	path('unassignExposure/<int:claimReserveId>/', ClaimReserveView.unassignExposure, name='unassignExposure'),
]
