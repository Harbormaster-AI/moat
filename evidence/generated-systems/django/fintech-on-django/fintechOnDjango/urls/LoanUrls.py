from django.urls import path
from fintechOnDjango.views import LoanView

urlpatterns = [
    path('', LoanView.index, name='index'),
	path('create', LoanView.get, name='create'),
	path('get/<int:loanId>/', LoanView.get, name='get'),
	path('save', LoanView.save, name='save'),
	path('getAll', LoanView.getAll, name='getAll'),
	path('delete/<int:loanId>/', LoanView.delete, name='delete'),
	path('assignCustomer/<int:loanId>/<int:CustomerId>/', LoanView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:loanId>/', LoanView.unassignCustomer, name='unassignCustomer'),
	path('addSchedule/<int:loanId>/<ScheduleIds>/', LoanView.addSchedule, name='addSchedule'),
	path('removeSchedule/<int:loanId>/<ScheduleIds>/', LoanView.removeSchedule, name='removeSchedule'),
	path('addCollateral/<int:loanId>/<CollateralIds>/', LoanView.addCollateral, name='addCollateral'),
	path('removeCollateral/<int:loanId>/<CollateralIds>/', LoanView.removeCollateral, name='removeCollateral'),
	path('addTransactions/<int:loanId>/<TransactionsIds>/', LoanView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:loanId>/<TransactionsIds>/', LoanView.removeTransactions, name='removeTransactions'),
]
