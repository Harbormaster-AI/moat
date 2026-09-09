from django.urls import path
from hrOnDjango.views import DependentView

urlpatterns = [
    path('', DependentView.index, name='index'),
	path('create', DependentView.get, name='create'),
	path('get/<int:dependentId>/', DependentView.get, name='get'),
	path('save', DependentView.save, name='save'),
	path('getAll', DependentView.getAll, name='getAll'),
	path('delete/<int:dependentId>/', DependentView.delete, name='delete'),
	path('assignBenefitEnrollment/<int:dependentId>/<int:BenefitEnrollmentId>/', DependentView.assignBenefitEnrollment, name='assignBenefitEnrollment'),
	path('unassignBenefitEnrollment/<int:dependentId>/', DependentView.unassignBenefitEnrollment, name='unassignBenefitEnrollment'),
	path('assignEmployee/<int:dependentId>/<int:EmployeeId>/', DependentView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:dependentId>/', DependentView.unassignEmployee, name='unassignEmployee'),
]
