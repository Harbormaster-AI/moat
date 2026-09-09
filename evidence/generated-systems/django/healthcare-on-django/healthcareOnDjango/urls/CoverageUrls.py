from django.urls import path
from healthcareOnDjango.views import CoverageView

urlpatterns = [
    path('', CoverageView.index, name='index'),
	path('create', CoverageView.get, name='create'),
	path('get/<int:coverageId>/', CoverageView.get, name='get'),
	path('save', CoverageView.save, name='save'),
	path('getAll', CoverageView.getAll, name='getAll'),
	path('delete/<int:coverageId>/', CoverageView.delete, name='delete'),
	path('assignPatient/<int:coverageId>/<int:PatientId>/', CoverageView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:coverageId>/', CoverageView.unassignPatient, name='unassignPatient'),
	path('assignPlan/<int:coverageId>/<int:PlanId>/', CoverageView.assignPlan, name='assignPlan'),
	path('unassignPlan/<int:coverageId>/', CoverageView.unassignPlan, name='unassignPlan'),
	path('addClaims/<int:coverageId>/<ClaimsIds>/', CoverageView.addClaims, name='addClaims'),
	path('removeClaims/<int:coverageId>/<ClaimsIds>/', CoverageView.removeClaims, name='removeClaims'),
	path('addAuthorizations/<int:coverageId>/<AuthorizationsIds>/', CoverageView.addAuthorizations, name='addAuthorizations'),
	path('removeAuthorizations/<int:coverageId>/<AuthorizationsIds>/', CoverageView.removeAuthorizations, name='removeAuthorizations'),
]
