from django.urls import path
from hrOnDjango.views import JobApplicationView

urlpatterns = [
    path('', JobApplicationView.index, name='index'),
	path('create', JobApplicationView.get, name='create'),
	path('get/<int:jobApplicationId>/', JobApplicationView.get, name='get'),
	path('save', JobApplicationView.save, name='save'),
	path('getAll', JobApplicationView.getAll, name='getAll'),
	path('delete/<int:jobApplicationId>/', JobApplicationView.delete, name='delete'),
	path('assignCandidate/<int:jobApplicationId>/<int:CandidateId>/', JobApplicationView.assignCandidate, name='assignCandidate'),
	path('unassignCandidate/<int:jobApplicationId>/', JobApplicationView.unassignCandidate, name='unassignCandidate'),
	path('assignRequisition/<int:jobApplicationId>/<int:RequisitionId>/', JobApplicationView.assignRequisition, name='assignRequisition'),
	path('unassignRequisition/<int:jobApplicationId>/', JobApplicationView.unassignRequisition, name='unassignRequisition'),
	path('addScreenings/<int:jobApplicationId>/<ScreeningsIds>/', JobApplicationView.addScreenings, name='addScreenings'),
	path('removeScreenings/<int:jobApplicationId>/<ScreeningsIds>/', JobApplicationView.removeScreenings, name='removeScreenings'),
]
