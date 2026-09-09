from django.urls import path
from governanceOnDjango.views import ControlTest_View

urlpatterns = [
    path('', ControlTest_View.index, name='index'),
	path('create', ControlTest_View.get, name='create'),
	path('get/<int:controlTest_Id>/', ControlTest_View.get, name='get'),
	path('save', ControlTest_View.save, name='save'),
	path('getAll', ControlTest_View.getAll, name='getAll'),
	path('delete/<int:controlTest_Id>/', ControlTest_View.delete, name='delete'),
	path('assignControl/<int:controlTest_Id>/<int:ControlId>/', ControlTest_View.assignControl, name='assignControl'),
	path('unassignControl/<int:controlTest_Id>/', ControlTest_View.unassignControl, name='unassignControl'),
	path('assignEngagement/<int:controlTest_Id>/<int:EngagementId>/', ControlTest_View.assignEngagement, name='assignEngagement'),
	path('unassignEngagement/<int:controlTest_Id>/', ControlTest_View.unassignEngagement, name='unassignEngagement'),
	path('addEvidence/<int:controlTest_Id>/<EvidenceIds>/', ControlTest_View.addEvidence, name='addEvidence'),
	path('removeEvidence/<int:controlTest_Id>/<EvidenceIds>/', ControlTest_View.removeEvidence, name='removeEvidence'),
]
