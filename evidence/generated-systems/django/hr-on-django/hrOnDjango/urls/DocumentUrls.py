from django.urls import path
from hrOnDjango.views import DocumentView

urlpatterns = [
    path('', DocumentView.index, name='index'),
	path('create', DocumentView.get, name='create'),
	path('get/<int:documentId>/', DocumentView.get, name='get'),
	path('save', DocumentView.save, name='save'),
	path('getAll', DocumentView.getAll, name='getAll'),
	path('delete/<int:documentId>/', DocumentView.delete, name='delete'),
	path('assignCandidate/<int:documentId>/<int:CandidateId>/', DocumentView.assignCandidate, name='assignCandidate'),
	path('unassignCandidate/<int:documentId>/', DocumentView.unassignCandidate, name='unassignCandidate'),
	path('assignEmployee/<int:documentId>/<int:EmployeeId>/', DocumentView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:documentId>/', DocumentView.unassignEmployee, name='unassignEmployee'),
]
