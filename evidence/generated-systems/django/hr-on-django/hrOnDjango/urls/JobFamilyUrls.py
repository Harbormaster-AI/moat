from django.urls import path
from hrOnDjango.views import JobFamilyView

urlpatterns = [
    path('', JobFamilyView.index, name='index'),
	path('create', JobFamilyView.get, name='create'),
	path('get/<int:jobFamilyId>/', JobFamilyView.get, name='get'),
	path('save', JobFamilyView.save, name='save'),
	path('getAll', JobFamilyView.getAll, name='getAll'),
	path('delete/<int:jobFamilyId>/', JobFamilyView.delete, name='delete'),
	path('assignOrganization/<int:jobFamilyId>/<int:OrganizationId>/', JobFamilyView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:jobFamilyId>/', JobFamilyView.unassignOrganization, name='unassignOrganization'),
	path('addJobProfiles/<int:jobFamilyId>/<JobProfilesIds>/', JobFamilyView.addJobProfiles, name='addJobProfiles'),
	path('removeJobProfiles/<int:jobFamilyId>/<JobProfilesIds>/', JobFamilyView.removeJobProfiles, name='removeJobProfiles'),
]
