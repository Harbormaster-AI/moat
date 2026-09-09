from django.urls import path
from hrOnDjango.views import LeaveRequestView

urlpatterns = [
    path('', LeaveRequestView.index, name='index'),
	path('create', LeaveRequestView.get, name='create'),
	path('get/<int:leaveRequestId>/', LeaveRequestView.get, name='get'),
	path('save', LeaveRequestView.save, name='save'),
	path('getAll', LeaveRequestView.getAll, name='getAll'),
	path('delete/<int:leaveRequestId>/', LeaveRequestView.delete, name='delete'),
	path('assignEmployee/<int:leaveRequestId>/<int:EmployeeId>/', LeaveRequestView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:leaveRequestId>/', LeaveRequestView.unassignEmployee, name='unassignEmployee'),
	path('assignLeavePolicy/<int:leaveRequestId>/<int:LeavePolicyId>/', LeaveRequestView.assignLeavePolicy, name='assignLeavePolicy'),
	path('unassignLeavePolicy/<int:leaveRequestId>/', LeaveRequestView.unassignLeavePolicy, name='unassignLeavePolicy'),
	path('addApprovals/<int:leaveRequestId>/<ApprovalsIds>/', LeaveRequestView.addApprovals, name='addApprovals'),
	path('removeApprovals/<int:leaveRequestId>/<ApprovalsIds>/', LeaveRequestView.removeApprovals, name='removeApprovals'),
]
