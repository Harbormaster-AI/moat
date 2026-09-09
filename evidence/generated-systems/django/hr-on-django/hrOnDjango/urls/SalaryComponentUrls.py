from django.urls import path
from hrOnDjango.views import SalaryComponentView

urlpatterns = [
    path('', SalaryComponentView.index, name='index'),
	path('create', SalaryComponentView.get, name='create'),
	path('get/<int:salaryComponentId>/', SalaryComponentView.get, name='get'),
	path('save', SalaryComponentView.save, name='save'),
	path('getAll', SalaryComponentView.getAll, name='getAll'),
	path('delete/<int:salaryComponentId>/', SalaryComponentView.delete, name='delete'),
	path('assignCompensationPackage/<int:salaryComponentId>/<int:CompensationPackageId>/', SalaryComponentView.assignCompensationPackage, name='assignCompensationPackage'),
	path('unassignCompensationPackage/<int:salaryComponentId>/', SalaryComponentView.unassignCompensationPackage, name='unassignCompensationPackage'),
]
