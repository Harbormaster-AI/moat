from django.urls import path
from fintechOnDjango.views import PaymentContractView

urlpatterns = [
    path('', PaymentContractView.index, name='index'),
	path('create', PaymentContractView.get, name='create'),
	path('get/<int:paymentContractId>/', PaymentContractView.get, name='get'),
	path('save', PaymentContractView.save, name='save'),
	path('getAll', PaymentContractView.getAll, name='getAll'),
	path('delete/<int:paymentContractId>/', PaymentContractView.delete, name='delete'),
	path('assignMerchant/<int:paymentContractId>/<int:MerchantId>/', PaymentContractView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:paymentContractId>/', PaymentContractView.unassignMerchant, name='unassignMerchant'),
	path('assignAcquirer/<int:paymentContractId>/<int:AcquirerId>/', PaymentContractView.assignAcquirer, name='assignAcquirer'),
	path('unassignAcquirer/<int:paymentContractId>/', PaymentContractView.unassignAcquirer, name='unassignAcquirer'),
]
