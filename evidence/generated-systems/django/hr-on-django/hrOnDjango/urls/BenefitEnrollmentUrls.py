from django.urls import path
from hrOnDjango.views import BenefitEnrollmentView

urlpatterns = [
    path('', BenefitEnrollmentView.index, name='index'),
	path('create', BenefitEnrollmentView.get, name='create'),
	path('get/<int:benefitEnrollmentId>/', BenefitEnrollmentView.get, name='get'),
	path('save', BenefitEnrollmentView.save, name='save'),
	path('getAll', BenefitEnrollmentView.getAll, name='getAll'),
	path('delete/<int:benefitEnrollmentId>/', BenefitEnrollmentView.delete, name='delete'),
	path('assignBenefitPlan/<int:benefitEnrollmentId>/<int:BenefitPlanId>/', BenefitEnrollmentView.assignBenefitPlan, name='assignBenefitPlan'),
	path('unassignBenefitPlan/<int:benefitEnrollmentId>/', BenefitEnrollmentView.unassignBenefitPlan, name='unassignBenefitPlan'),
	path('assignEmployee/<int:benefitEnrollmentId>/<int:EmployeeId>/', BenefitEnrollmentView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:benefitEnrollmentId>/', BenefitEnrollmentView.unassignEmployee, name='unassignEmployee'),
	path('addDependents/<int:benefitEnrollmentId>/<DependentsIds>/', BenefitEnrollmentView.addDependents, name='addDependents'),
	path('removeDependents/<int:benefitEnrollmentId>/<DependentsIds>/', BenefitEnrollmentView.removeDependents, name='removeDependents'),
]
