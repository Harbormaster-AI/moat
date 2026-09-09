from django.urls import path
from insuranceOnDjango.views import BeneficiaryView

urlpatterns = [
    path('', BeneficiaryView.index, name='index'),
	path('create', BeneficiaryView.get, name='create'),
	path('get/<int:beneficiaryId>/', BeneficiaryView.get, name='get'),
	path('save', BeneficiaryView.save, name='save'),
	path('getAll', BeneficiaryView.getAll, name='getAll'),
	path('delete/<int:beneficiaryId>/', BeneficiaryView.delete, name='delete'),
	path('assignPolicy/<int:beneficiaryId>/<int:PolicyId>/', BeneficiaryView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:beneficiaryId>/', BeneficiaryView.unassignPolicy, name='unassignPolicy'),
	path('assignCustomer/<int:beneficiaryId>/<int:CustomerId>/', BeneficiaryView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:beneficiaryId>/', BeneficiaryView.unassignCustomer, name='unassignCustomer'),
]
