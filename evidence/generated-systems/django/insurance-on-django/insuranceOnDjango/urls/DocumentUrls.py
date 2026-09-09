from django.urls import path
from insuranceOnDjango.views import DocumentView

urlpatterns = [
    path('', DocumentView.index, name='index'),
	path('create', DocumentView.get, name='create'),
	path('get/<int:documentId>/', DocumentView.get, name='get'),
	path('save', DocumentView.save, name='save'),
	path('getAll', DocumentView.getAll, name='getAll'),
	path('delete/<int:documentId>/', DocumentView.delete, name='delete'),
	path('assignPolicy/<int:documentId>/<int:PolicyId>/', DocumentView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:documentId>/', DocumentView.unassignPolicy, name='unassignPolicy'),
	path('assignClaim/<int:documentId>/<int:ClaimId>/', DocumentView.assignClaim, name='assignClaim'),
	path('unassignClaim/<int:documentId>/', DocumentView.unassignClaim, name='unassignClaim'),
	path('assignApplication/<int:documentId>/<int:ApplicationId>/', DocumentView.assignApplication, name='assignApplication'),
	path('unassignApplication/<int:documentId>/', DocumentView.unassignApplication, name='unassignApplication'),
	path('assignCustomer/<int:documentId>/<int:CustomerId>/', DocumentView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:documentId>/', DocumentView.unassignCustomer, name='unassignCustomer'),
]
