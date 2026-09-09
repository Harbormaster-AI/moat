from django.urls import path
from governanceOnDjango.views import ControlView

urlpatterns = [
    path('', ControlView.index, name='index'),
	path('create', ControlView.get, name='create'),
	path('get/<int:controlId>/', ControlView.get, name='get'),
	path('save', ControlView.save, name='save'),
	path('getAll', ControlView.getAll, name='getAll'),
	path('delete/<int:controlId>/', ControlView.delete, name='delete'),
	path('assignPolicy/<int:controlId>/<int:PolicyId>/', ControlView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:controlId>/', ControlView.unassignPolicy, name='unassignPolicy'),
	path('addControlTests/<int:controlId>/<ControlTestsIds>/', ControlView.addControlTests, name='addControlTests'),
	path('removeControlTests/<int:controlId>/<ControlTestsIds>/', ControlView.removeControlTests, name='removeControlTests'),
	path('addEvidence/<int:controlId>/<EvidenceIds>/', ControlView.addEvidence, name='addEvidence'),
	path('removeEvidence/<int:controlId>/<EvidenceIds>/', ControlView.removeEvidence, name='removeEvidence'),
	path('addRisks/<int:controlId>/<RisksIds>/', ControlView.addRisks, name='addRisks'),
	path('removeRisks/<int:controlId>/<RisksIds>/', ControlView.removeRisks, name='removeRisks'),
	path('addObligations/<int:controlId>/<ObligationsIds>/', ControlView.addObligations, name='addObligations'),
	path('removeObligations/<int:controlId>/<ObligationsIds>/', ControlView.removeObligations, name='removeObligations'),
	path('addProcedures/<int:controlId>/<ProceduresIds>/', ControlView.addProcedures, name='addProcedures'),
	path('removeProcedures/<int:controlId>/<ProceduresIds>/', ControlView.removeProcedures, name='removeProcedures'),
	path('addIssues/<int:controlId>/<IssuesIds>/', ControlView.addIssues, name='addIssues'),
	path('removeIssues/<int:controlId>/<IssuesIds>/', ControlView.removeIssues, name='removeIssues'),
]
