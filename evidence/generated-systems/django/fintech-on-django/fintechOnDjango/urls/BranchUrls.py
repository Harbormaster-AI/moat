from django.urls import path
from fintechOnDjango.views import BranchView

urlpatterns = [
    path('', BranchView.index, name='index'),
	path('create', BranchView.get, name='create'),
	path('get/<int:branchId>/', BranchView.get, name='get'),
	path('save', BranchView.save, name='save'),
	path('getAll', BranchView.getAll, name='getAll'),
	path('delete/<int:branchId>/', BranchView.delete, name='delete'),
	path('assignInstitution/<int:branchId>/<int:InstitutionId>/', BranchView.assignInstitution, name='assignInstitution'),
	path('unassignInstitution/<int:branchId>/', BranchView.unassignInstitution, name='unassignInstitution'),
]
