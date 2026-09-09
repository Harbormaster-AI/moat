from django.urls import path
from hrOnDjango.views import WorkAuthorizationView

urlpatterns = [
    path('', WorkAuthorizationView.index, name='index'),
	path('create', WorkAuthorizationView.get, name='create'),
	path('get/<int:workAuthorizationId>/', WorkAuthorizationView.get, name='get'),
	path('save', WorkAuthorizationView.save, name='save'),
	path('getAll', WorkAuthorizationView.getAll, name='getAll'),
	path('delete/<int:workAuthorizationId>/', WorkAuthorizationView.delete, name='delete'),
	path('assignEmployee/<int:workAuthorizationId>/<int:EmployeeId>/', WorkAuthorizationView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:workAuthorizationId>/', WorkAuthorizationView.unassignEmployee, name='unassignEmployee'),
	path('addDocuments/<int:workAuthorizationId>/<DocumentsIds>/', WorkAuthorizationView.addDocuments, name='addDocuments'),
	path('removeDocuments/<int:workAuthorizationId>/<DocumentsIds>/', WorkAuthorizationView.removeDocuments, name='removeDocuments'),
]
