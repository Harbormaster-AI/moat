from django.urls import path
from fintechOnDjango.views import RiskAssessmentView

urlpatterns = [
    path('', RiskAssessmentView.index, name='index'),
	path('create', RiskAssessmentView.get, name='create'),
	path('get/<int:riskAssessmentId>/', RiskAssessmentView.get, name='get'),
	path('save', RiskAssessmentView.save, name='save'),
	path('getAll', RiskAssessmentView.getAll, name='getAll'),
	path('delete/<int:riskAssessmentId>/', RiskAssessmentView.delete, name='delete'),
	path('assignApplication/<int:riskAssessmentId>/<int:ApplicationId>/', RiskAssessmentView.assignApplication, name='assignApplication'),
	path('unassignApplication/<int:riskAssessmentId>/', RiskAssessmentView.unassignApplication, name='unassignApplication'),
]
