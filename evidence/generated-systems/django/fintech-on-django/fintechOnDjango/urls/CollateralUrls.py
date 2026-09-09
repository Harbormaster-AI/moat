from django.urls import path
from fintechOnDjango.views import CollateralView

urlpatterns = [
    path('', CollateralView.index, name='index'),
	path('create', CollateralView.get, name='create'),
	path('get/<int:collateralId>/', CollateralView.get, name='get'),
	path('save', CollateralView.save, name='save'),
	path('getAll', CollateralView.getAll, name='getAll'),
	path('delete/<int:collateralId>/', CollateralView.delete, name='delete'),
	path('assignLoan/<int:collateralId>/<int:LoanId>/', CollateralView.assignLoan, name='assignLoan'),
	path('unassignLoan/<int:collateralId>/', CollateralView.unassignLoan, name='unassignLoan'),
]
