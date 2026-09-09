from django.urls import path
from fintechOnDjango.views import SettlementBatchView

urlpatterns = [
    path('', SettlementBatchView.index, name='index'),
	path('create', SettlementBatchView.get, name='create'),
	path('get/<int:settlementBatchId>/', SettlementBatchView.get, name='get'),
	path('save', SettlementBatchView.save, name='save'),
	path('getAll', SettlementBatchView.getAll, name='getAll'),
	path('delete/<int:settlementBatchId>/', SettlementBatchView.delete, name='delete'),
	path('assignProcessor/<int:settlementBatchId>/<int:ProcessorId>/', SettlementBatchView.assignProcessor, name='assignProcessor'),
	path('unassignProcessor/<int:settlementBatchId>/', SettlementBatchView.unassignProcessor, name='unassignProcessor'),
	path('assignMerchant/<int:settlementBatchId>/<int:MerchantId>/', SettlementBatchView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:settlementBatchId>/', SettlementBatchView.unassignMerchant, name='unassignMerchant'),
	path('addPayouts/<int:settlementBatchId>/<PayoutsIds>/', SettlementBatchView.addPayouts, name='addPayouts'),
	path('removePayouts/<int:settlementBatchId>/<PayoutsIds>/', SettlementBatchView.removePayouts, name='removePayouts'),
	path('addTransactions/<int:settlementBatchId>/<TransactionsIds>/', SettlementBatchView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:settlementBatchId>/<TransactionsIds>/', SettlementBatchView.removeTransactions, name='removeTransactions'),
]
