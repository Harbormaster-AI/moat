from django.urls import path
from fintechOnDjango.views import RepaymentScheduleView

urlpatterns = [
    path('', RepaymentScheduleView.index, name='index'),
	path('create', RepaymentScheduleView.get, name='create'),
	path('get/<int:repaymentScheduleId>/', RepaymentScheduleView.get, name='get'),
	path('save', RepaymentScheduleView.save, name='save'),
	path('getAll', RepaymentScheduleView.getAll, name='getAll'),
	path('delete/<int:repaymentScheduleId>/', RepaymentScheduleView.delete, name='delete'),
	path('assignLoan/<int:repaymentScheduleId>/<int:LoanId>/', RepaymentScheduleView.assignLoan, name='assignLoan'),
	path('unassignLoan/<int:repaymentScheduleId>/', RepaymentScheduleView.unassignLoan, name='unassignLoan'),
	path('addPayments/<int:repaymentScheduleId>/<PaymentsIds>/', RepaymentScheduleView.addPayments, name='addPayments'),
	path('removePayments/<int:repaymentScheduleId>/<PaymentsIds>/', RepaymentScheduleView.removePayments, name='removePayments'),
]
