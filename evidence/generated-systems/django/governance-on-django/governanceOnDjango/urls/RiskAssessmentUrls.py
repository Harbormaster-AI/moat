from django.urls import path
from governanceOnDjango.views import RiskAssessmentView

urlpatterns = [
    path('', RiskAssessmentView.index, name='index'),
	path('create', RiskAssessmentView.get, name='create'),
	path('get/<int:riskAssessmentId>/', RiskAssessmentView.get, name='get'),
	path('save', RiskAssessmentView.save, name='save'),
	path('getAll', RiskAssessmentView.getAll, name='getAll'),
	path('delete/<int:riskAssessmentId>/', RiskAssessmentView.delete, name='delete'),
	path('assignRisk/<int:riskAssessmentId>/<int:RiskId>/', RiskAssessmentView.assignRisk, name='assignRisk'),
	path('unassignRisk/<int:riskAssessmentId>/', RiskAssessmentView.unassignRisk, name='unassignRisk'),
]
