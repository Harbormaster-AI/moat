from django.urls import path
from fintechOnDjango.views import PaymentProcessorView

urlpatterns = [
    path('', PaymentProcessorView.index, name='index'),
	path('create', PaymentProcessorView.get, name='create'),
	path('get/<int:paymentProcessorId>/', PaymentProcessorView.get, name='get'),
	path('save', PaymentProcessorView.save, name='save'),
	path('getAll', PaymentProcessorView.getAll, name='getAll'),
	path('delete/<int:paymentProcessorId>/', PaymentProcessorView.delete, name='delete'),
	path('addInstitutions/<int:paymentProcessorId>/<InstitutionsIds>/', PaymentProcessorView.addInstitutions, name='addInstitutions'),
	path('removeInstitutions/<int:paymentProcessorId>/<InstitutionsIds>/', PaymentProcessorView.removeInstitutions, name='removeInstitutions'),
	path('addContracts/<int:paymentProcessorId>/<ContractsIds>/', PaymentProcessorView.addContracts, name='addContracts'),
	path('removeContracts/<int:paymentProcessorId>/<ContractsIds>/', PaymentProcessorView.removeContracts, name='removeContracts'),
	path('addSettlements/<int:paymentProcessorId>/<SettlementsIds>/', PaymentProcessorView.addSettlements, name='addSettlements'),
	path('removeSettlements/<int:paymentProcessorId>/<SettlementsIds>/', PaymentProcessorView.removeSettlements, name='removeSettlements'),
]
