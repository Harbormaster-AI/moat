from django.urls import path
from fintechOnDjango.views import ChargebackView

urlpatterns = [
    path('', ChargebackView.index, name='index'),
	path('create', ChargebackView.get, name='create'),
	path('get/<int:chargebackId>/', ChargebackView.get, name='get'),
	path('save', ChargebackView.save, name='save'),
	path('getAll', ChargebackView.getAll, name='getAll'),
	path('delete/<int:chargebackId>/', ChargebackView.delete, name='delete'),
	path('assignDispute/<int:chargebackId>/<int:DisputeId>/', ChargebackView.assignDispute, name='assignDispute'),
	path('unassignDispute/<int:chargebackId>/', ChargebackView.unassignDispute, name='unassignDispute'),
	path('assignTransaction/<int:chargebackId>/<int:TransactionId>/', ChargebackView.assignTransaction, name='assignTransaction'),
	path('unassignTransaction/<int:chargebackId>/', ChargebackView.unassignTransaction, name='unassignTransaction'),
]
