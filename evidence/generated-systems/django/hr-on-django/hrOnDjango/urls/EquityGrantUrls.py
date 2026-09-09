from django.urls import path
from hrOnDjango.views import EquityGrantView

urlpatterns = [
    path('', EquityGrantView.index, name='index'),
	path('create', EquityGrantView.get, name='create'),
	path('get/<int:equityGrantId>/', EquityGrantView.get, name='get'),
	path('save', EquityGrantView.save, name='save'),
	path('getAll', EquityGrantView.getAll, name='getAll'),
	path('delete/<int:equityGrantId>/', EquityGrantView.delete, name='delete'),
	path('assignCompensationPackage/<int:equityGrantId>/<int:CompensationPackageId>/', EquityGrantView.assignCompensationPackage, name='assignCompensationPackage'),
	path('unassignCompensationPackage/<int:equityGrantId>/', EquityGrantView.unassignCompensationPackage, name='unassignCompensationPackage'),
]
