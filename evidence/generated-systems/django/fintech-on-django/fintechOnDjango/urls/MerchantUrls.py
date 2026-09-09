from django.urls import path
from fintechOnDjango.views import MerchantView

urlpatterns = [
    path('', MerchantView.index, name='index'),
	path('create', MerchantView.get, name='create'),
	path('get/<int:merchantId>/', MerchantView.get, name='get'),
	path('save', MerchantView.save, name='save'),
	path('getAll', MerchantView.getAll, name='getAll'),
	path('delete/<int:merchantId>/', MerchantView.delete, name='delete'),
	path('addTerminals/<int:merchantId>/<TerminalsIds>/', MerchantView.addTerminals, name='addTerminals'),
	path('removeTerminals/<int:merchantId>/<TerminalsIds>/', MerchantView.removeTerminals, name='removeTerminals'),
	path('addPaymentContracts/<int:merchantId>/<PaymentContractsIds>/', MerchantView.addPaymentContracts, name='addPaymentContracts'),
	path('removePaymentContracts/<int:merchantId>/<PaymentContractsIds>/', MerchantView.removePaymentContracts, name='removePaymentContracts'),
	path('addPayouts/<int:merchantId>/<PayoutsIds>/', MerchantView.addPayouts, name='addPayouts'),
	path('removePayouts/<int:merchantId>/<PayoutsIds>/', MerchantView.removePayouts, name='removePayouts'),
	path('addSettlements/<int:merchantId>/<SettlementsIds>/', MerchantView.addSettlements, name='addSettlements'),
	path('removeSettlements/<int:merchantId>/<SettlementsIds>/', MerchantView.removeSettlements, name='removeSettlements'),
	path('addDisputes/<int:merchantId>/<DisputesIds>/', MerchantView.addDisputes, name='addDisputes'),
	path('removeDisputes/<int:merchantId>/<DisputesIds>/', MerchantView.removeDisputes, name='removeDisputes'),
	path('addInvoices/<int:merchantId>/<InvoicesIds>/', MerchantView.addInvoices, name='addInvoices'),
	path('removeInvoices/<int:merchantId>/<InvoicesIds>/', MerchantView.removeInvoices, name='removeInvoices'),
]
