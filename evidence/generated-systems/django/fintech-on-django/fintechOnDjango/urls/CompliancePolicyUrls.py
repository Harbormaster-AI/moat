from django.urls import path
from fintechOnDjango.views import CompliancePolicyView

urlpatterns = [
    path('', CompliancePolicyView.index, name='index'),
	path('create', CompliancePolicyView.get, name='create'),
	path('get/<int:compliancePolicyId>/', CompliancePolicyView.get, name='get'),
	path('save', CompliancePolicyView.save, name='save'),
	path('getAll', CompliancePolicyView.getAll, name='getAll'),
	path('delete/<int:compliancePolicyId>/', CompliancePolicyView.delete, name='delete'),
	path('assignInstitution/<int:compliancePolicyId>/<int:InstitutionId>/', CompliancePolicyView.assignInstitution, name='assignInstitution'),
	path('unassignInstitution/<int:compliancePolicyId>/', CompliancePolicyView.unassignInstitution, name='unassignInstitution'),
]
