from django.urls import path
from fintechOnDjango.views import TransactionView

urlpatterns = [
    path('', TransactionView.index, name='index'),
	path('create', TransactionView.get, name='create'),
	path('get/<int:transactionId>/', TransactionView.get, name='get'),
	path('save', TransactionView.save, name='save'),
	path('getAll', TransactionView.getAll, name='getAll'),
	path('delete/<int:transactionId>/', TransactionView.delete, name='delete'),
	path('assignAccount/<int:transactionId>/<int:AccountId>/', TransactionView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:transactionId>/', TransactionView.unassignAccount, name='unassignAccount'),
	path('assignWallet/<int:transactionId>/<int:WalletId>/', TransactionView.assignWallet, name='assignWallet'),
	path('unassignWallet/<int:transactionId>/', TransactionView.unassignWallet, name='unassignWallet'),
	path('assignPaymentOrder/<int:transactionId>/<int:PaymentOrderId>/', TransactionView.assignPaymentOrder, name='assignPaymentOrder'),
	path('unassignPaymentOrder/<int:transactionId>/', TransactionView.unassignPaymentOrder, name='unassignPaymentOrder'),
	path('assignMerchant/<int:transactionId>/<int:MerchantId>/', TransactionView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:transactionId>/', TransactionView.unassignMerchant, name='unassignMerchant'),
	path('assignCard/<int:transactionId>/<int:CardId>/', TransactionView.assignCard, name='assignCard'),
	path('unassignCard/<int:transactionId>/', TransactionView.unassignCard, name='unassignCard'),
	path('addRelatedTransactions/<int:transactionId>/<RelatedTransactionsIds>/', TransactionView.addRelatedTransactions, name='addRelatedTransactions'),
	path('removeRelatedTransactions/<int:transactionId>/<RelatedTransactionsIds>/', TransactionView.removeRelatedTransactions, name='removeRelatedTransactions'),
	path('addAlerts/<int:transactionId>/<AlertsIds>/', TransactionView.addAlerts, name='addAlerts'),
	path('removeAlerts/<int:transactionId>/<AlertsIds>/', TransactionView.removeAlerts, name='removeAlerts'),
]
