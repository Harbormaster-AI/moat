from django.urls import path
from fintechOnDjango.views import LoanApplicationView

urlpatterns = [
    path('', LoanApplicationView.index, name='index'),
	path('create', LoanApplicationView.get, name='create'),
	path('get/<int:loanApplicationId>/', LoanApplicationView.get, name='get'),
	path('save', LoanApplicationView.save, name='save'),
	path('getAll', LoanApplicationView.getAll, name='getAll'),
	path('delete/<int:loanApplicationId>/', LoanApplicationView.delete, name='delete'),
	path('assignCustomer/<int:loanApplicationId>/<int:CustomerId>/', LoanApplicationView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:loanApplicationId>/', LoanApplicationView.unassignCustomer, name='unassignCustomer'),
	path('assignRiskAssessment/<int:loanApplicationId>/<int:RiskAssessmentId>/', LoanApplicationView.assignRiskAssessment, name='assignRiskAssessment'),
	path('unassignRiskAssessment/<int:loanApplicationId>/', LoanApplicationView.unassignRiskAssessment, name='unassignRiskAssessment'),
	path('assignLoan/<int:loanApplicationId>/<int:LoanId>/', LoanApplicationView.assignLoan, name='assignLoan'),
	path('unassignLoan/<int:loanApplicationId>/', LoanApplicationView.unassignLoan, name='unassignLoan'),
]
