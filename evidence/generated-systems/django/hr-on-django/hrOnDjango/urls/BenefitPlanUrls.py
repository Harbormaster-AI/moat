from django.urls import path
from hrOnDjango.views import BenefitPlanView

urlpatterns = [
    path('', BenefitPlanView.index, name='index'),
	path('create', BenefitPlanView.get, name='create'),
	path('get/<int:benefitPlanId>/', BenefitPlanView.get, name='get'),
	path('save', BenefitPlanView.save, name='save'),
	path('getAll', BenefitPlanView.getAll, name='getAll'),
	path('delete/<int:benefitPlanId>/', BenefitPlanView.delete, name='delete'),
	path('assignOrganization/<int:benefitPlanId>/<int:OrganizationId>/', BenefitPlanView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:benefitPlanId>/', BenefitPlanView.unassignOrganization, name='unassignOrganization'),
	path('addEnrollments/<int:benefitPlanId>/<EnrollmentsIds>/', BenefitPlanView.addEnrollments, name='addEnrollments'),
	path('removeEnrollments/<int:benefitPlanId>/<EnrollmentsIds>/', BenefitPlanView.removeEnrollments, name='removeEnrollments'),
]
