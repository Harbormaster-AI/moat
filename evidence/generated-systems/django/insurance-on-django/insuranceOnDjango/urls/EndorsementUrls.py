from django.urls import path
from insuranceOnDjango.views import EndorsementView

urlpatterns = [
    path('', EndorsementView.index, name='index'),
	path('create', EndorsementView.get, name='create'),
	path('get/<int:endorsementId>/', EndorsementView.get, name='get'),
	path('save', EndorsementView.save, name='save'),
	path('getAll', EndorsementView.getAll, name='getAll'),
	path('delete/<int:endorsementId>/', EndorsementView.delete, name='delete'),
	path('assignPolicy/<int:endorsementId>/<int:PolicyId>/', EndorsementView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:endorsementId>/', EndorsementView.unassignPolicy, name='unassignPolicy'),
]
