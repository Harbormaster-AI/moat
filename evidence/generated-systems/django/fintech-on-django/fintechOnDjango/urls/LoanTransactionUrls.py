from django.urls import path
from fintechOnDjango.views import LoanTransactionView

urlpatterns = [
    path('', LoanTransactionView.index, name='index'),
	path('create', LoanTransactionView.get, name='create'),
	path('get/<int:loanTransactionId>/', LoanTransactionView.get, name='get'),
	path('save', LoanTransactionView.save, name='save'),
	path('getAll', LoanTransactionView.getAll, name='getAll'),
	path('delete/<int:loanTransactionId>/', LoanTransactionView.delete, name='delete'),
	path('assignLoan/<int:loanTransactionId>/<int:LoanId>/', LoanTransactionView.assignLoan, name='assignLoan'),
	path('unassignLoan/<int:loanTransactionId>/', LoanTransactionView.unassignLoan, name='unassignLoan'),
]
