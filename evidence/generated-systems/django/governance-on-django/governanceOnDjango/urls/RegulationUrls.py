from django.urls import path
from governanceOnDjango.views import RegulationView

urlpatterns = [
    path('', RegulationView.index, name='index'),
	path('create', RegulationView.get, name='create'),
	path('get/<int:regulationId>/', RegulationView.get, name='get'),
	path('save', RegulationView.save, name='save'),
	path('getAll', RegulationView.getAll, name='getAll'),
	path('delete/<int:regulationId>/', RegulationView.delete, name='delete'),
	path('addObligations/<int:regulationId>/<ObligationsIds>/', RegulationView.addObligations, name='addObligations'),
	path('removeObligations/<int:regulationId>/<ObligationsIds>/', RegulationView.removeObligations, name='removeObligations'),
	path('addCompliancePrograms/<int:regulationId>/<ComplianceProgramsIds>/', RegulationView.addCompliancePrograms, name='addCompliancePrograms'),
	path('removeCompliancePrograms/<int:regulationId>/<ComplianceProgramsIds>/', RegulationView.removeCompliancePrograms, name='removeCompliancePrograms'),
]
