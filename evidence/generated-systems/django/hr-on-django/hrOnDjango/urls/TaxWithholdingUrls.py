from django.urls import path
from hrOnDjango.views import TaxWithholdingView

urlpatterns = [
    path('', TaxWithholdingView.index, name='index'),
	path('create', TaxWithholdingView.get, name='create'),
	path('get/<int:taxWithholdingId>/', TaxWithholdingView.get, name='get'),
	path('save', TaxWithholdingView.save, name='save'),
	path('getAll', TaxWithholdingView.getAll, name='getAll'),
	path('delete/<int:taxWithholdingId>/', TaxWithholdingView.delete, name='delete'),
	path('assignEmployee/<int:taxWithholdingId>/<int:EmployeeId>/', TaxWithholdingView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:taxWithholdingId>/', TaxWithholdingView.unassignEmployee, name='unassignEmployee'),
]
