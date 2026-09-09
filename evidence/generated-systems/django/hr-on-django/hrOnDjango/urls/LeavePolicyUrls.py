from django.urls import path
from hrOnDjango.views import LeavePolicyView

urlpatterns = [
    path('', LeavePolicyView.index, name='index'),
	path('create', LeavePolicyView.get, name='create'),
	path('get/<int:leavePolicyId>/', LeavePolicyView.get, name='get'),
	path('save', LeavePolicyView.save, name='save'),
	path('getAll', LeavePolicyView.getAll, name='getAll'),
	path('delete/<int:leavePolicyId>/', LeavePolicyView.delete, name='delete'),
	path('assignOrganization/<int:leavePolicyId>/<int:OrganizationId>/', LeavePolicyView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:leavePolicyId>/', LeavePolicyView.unassignOrganization, name='unassignOrganization'),
	path('addLeaveRequests/<int:leavePolicyId>/<LeaveRequestsIds>/', LeavePolicyView.addLeaveRequests, name='addLeaveRequests'),
	path('removeLeaveRequests/<int:leavePolicyId>/<LeaveRequestsIds>/', LeavePolicyView.removeLeaveRequests, name='removeLeaveRequests'),
]
