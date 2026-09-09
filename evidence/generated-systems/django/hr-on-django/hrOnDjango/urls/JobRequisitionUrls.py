from django.urls import path
from hrOnDjango.views import JobRequisitionView

urlpatterns = [
    path('', JobRequisitionView.index, name='index'),
	path('create', JobRequisitionView.get, name='create'),
	path('get/<int:jobRequisitionId>/', JobRequisitionView.get, name='get'),
	path('save', JobRequisitionView.save, name='save'),
	path('getAll', JobRequisitionView.getAll, name='getAll'),
	path('delete/<int:jobRequisitionId>/', JobRequisitionView.delete, name='delete'),
	path('assignDepartment/<int:jobRequisitionId>/<int:DepartmentId>/', JobRequisitionView.assignDepartment, name='assignDepartment'),
	path('unassignDepartment/<int:jobRequisitionId>/', JobRequisitionView.unassignDepartment, name='unassignDepartment'),
	path('assignHiringManager/<int:jobRequisitionId>/<int:HiringManagerId>/', JobRequisitionView.assignHiringManager, name='assignHiringManager'),
	path('unassignHiringManager/<int:jobRequisitionId>/', JobRequisitionView.unassignHiringManager, name='unassignHiringManager'),
	path('assignRecruiter/<int:jobRequisitionId>/<int:RecruiterId>/', JobRequisitionView.assignRecruiter, name='assignRecruiter'),
	path('unassignRecruiter/<int:jobRequisitionId>/', JobRequisitionView.unassignRecruiter, name='unassignRecruiter'),
	path('assignJobProfile/<int:jobRequisitionId>/<int:JobProfileId>/', JobRequisitionView.assignJobProfile, name='assignJobProfile'),
	path('unassignJobProfile/<int:jobRequisitionId>/', JobRequisitionView.unassignJobProfile, name='unassignJobProfile'),
	path('addCandidates/<int:jobRequisitionId>/<CandidatesIds>/', JobRequisitionView.addCandidates, name='addCandidates'),
	path('removeCandidates/<int:jobRequisitionId>/<CandidatesIds>/', JobRequisitionView.removeCandidates, name='removeCandidates'),
	path('addInterviews/<int:jobRequisitionId>/<InterviewsIds>/', JobRequisitionView.addInterviews, name='addInterviews'),
	path('removeInterviews/<int:jobRequisitionId>/<InterviewsIds>/', JobRequisitionView.removeInterviews, name='removeInterviews'),
	path('addOffers/<int:jobRequisitionId>/<OffersIds>/', JobRequisitionView.addOffers, name='addOffers'),
	path('removeOffers/<int:jobRequisitionId>/<OffersIds>/', JobRequisitionView.removeOffers, name='removeOffers'),
]
