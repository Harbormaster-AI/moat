from django.urls import path
from insuranceOnDjango.views import SubrogationRecoveryView

urlpatterns = [
    path('', SubrogationRecoveryView.index, name='index'),
	path('create', SubrogationRecoveryView.get, name='create'),
	path('get/<int:subrogationRecoveryId>/', SubrogationRecoveryView.get, name='get'),
	path('save', SubrogationRecoveryView.save, name='save'),
	path('getAll', SubrogationRecoveryView.getAll, name='getAll'),
	path('delete/<int:subrogationRecoveryId>/', SubrogationRecoveryView.delete, name='delete'),
	path('assignClaim/<int:subrogationRecoveryId>/<int:ClaimId>/', SubrogationRecoveryView.assignClaim, name='assignClaim'),
	path('unassignClaim/<int:subrogationRecoveryId>/', SubrogationRecoveryView.unassignClaim, name='unassignClaim'),
	path('assignExposure/<int:subrogationRecoveryId>/<int:ExposureId>/', SubrogationRecoveryView.assignExposure, name='assignExposure'),
	path('unassignExposure/<int:subrogationRecoveryId>/', SubrogationRecoveryView.unassignExposure, name='unassignExposure'),
	path('assignCounterparty/<int:subrogationRecoveryId>/<int:CounterpartyId>/', SubrogationRecoveryView.assignCounterparty, name='assignCounterparty'),
	path('unassignCounterparty/<int:subrogationRecoveryId>/', SubrogationRecoveryView.unassignCounterparty, name='unassignCounterparty'),
]
