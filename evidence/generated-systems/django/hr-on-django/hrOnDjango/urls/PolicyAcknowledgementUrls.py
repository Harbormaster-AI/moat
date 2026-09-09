from django.urls import path
from hrOnDjango.views import PolicyAcknowledgementView

urlpatterns = [
    path('', PolicyAcknowledgementView.index, name='index'),
	path('create', PolicyAcknowledgementView.get, name='create'),
	path('get/<int:policyAcknowledgementId>/', PolicyAcknowledgementView.get, name='get'),
	path('save', PolicyAcknowledgementView.save, name='save'),
	path('getAll', PolicyAcknowledgementView.getAll, name='getAll'),
	path('delete/<int:policyAcknowledgementId>/', PolicyAcknowledgementView.delete, name='delete'),
	path('assignPolicy/<int:policyAcknowledgementId>/<int:PolicyId>/', PolicyAcknowledgementView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:policyAcknowledgementId>/', PolicyAcknowledgementView.unassignPolicy, name='unassignPolicy'),
	path('assignEmployee/<int:policyAcknowledgementId>/<int:EmployeeId>/', PolicyAcknowledgementView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:policyAcknowledgementId>/', PolicyAcknowledgementView.unassignEmployee, name='unassignEmployee'),
]
