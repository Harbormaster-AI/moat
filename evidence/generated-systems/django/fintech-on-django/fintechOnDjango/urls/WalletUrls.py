from django.urls import path
from fintechOnDjango.views import WalletView

urlpatterns = [
    path('', WalletView.index, name='index'),
	path('create', WalletView.get, name='create'),
	path('get/<int:walletId>/', WalletView.get, name='get'),
	path('save', WalletView.save, name='save'),
	path('getAll', WalletView.getAll, name='getAll'),
	path('delete/<int:walletId>/', WalletView.delete, name='delete'),
	path('assignCustomer/<int:walletId>/<int:CustomerId>/', WalletView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:walletId>/', WalletView.unassignCustomer, name='unassignCustomer'),
	path('addTransactions/<int:walletId>/<TransactionsIds>/', WalletView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:walletId>/<TransactionsIds>/', WalletView.removeTransactions, name='removeTransactions'),
]
