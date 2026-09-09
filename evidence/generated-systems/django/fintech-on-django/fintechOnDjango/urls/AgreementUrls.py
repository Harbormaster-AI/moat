from django.urls import path
from fintechOnDjango.views import AgreementView

urlpatterns = [
    path('', AgreementView.index, name='index'),
	path('create', AgreementView.get, name='create'),
	path('get/<int:agreementId>/', AgreementView.get, name='get'),
	path('save', AgreementView.save, name='save'),
	path('getAll', AgreementView.getAll, name='getAll'),
	path('delete/<int:agreementId>/', AgreementView.delete, name='delete'),
	path('assignCustomer/<int:agreementId>/<int:CustomerId>/', AgreementView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:agreementId>/', AgreementView.unassignCustomer, name='unassignCustomer'),
	path('assignProductOffering/<int:agreementId>/<int:ProductOfferingId>/', AgreementView.assignProductOffering, name='assignProductOffering'),
	path('unassignProductOffering/<int:agreementId>/', AgreementView.unassignProductOffering, name='unassignProductOffering'),
]
