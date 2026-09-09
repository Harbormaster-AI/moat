from django.urls import path
from hrOnDjango.views import ApprovalView

urlpatterns = [
    path('', ApprovalView.index, name='index'),
	path('create', ApprovalView.get, name='create'),
	path('get/<int:approvalId>/', ApprovalView.get, name='get'),
	path('save', ApprovalView.save, name='save'),
	path('getAll', ApprovalView.getAll, name='getAll'),
	path('delete/<int:approvalId>/', ApprovalView.delete, name='delete'),
	path('assignApprover/<int:approvalId>/<int:ApproverId>/', ApprovalView.assignApprover, name='assignApprover'),
	path('unassignApprover/<int:approvalId>/', ApprovalView.unassignApprover, name='unassignApprover'),
	path('assignTimesheet/<int:approvalId>/<int:TimesheetId>/', ApprovalView.assignTimesheet, name='assignTimesheet'),
	path('unassignTimesheet/<int:approvalId>/', ApprovalView.unassignTimesheet, name='unassignTimesheet'),
	path('assignLeaveRequest/<int:approvalId>/<int:LeaveRequestId>/', ApprovalView.assignLeaveRequest, name='assignLeaveRequest'),
	path('unassignLeaveRequest/<int:approvalId>/', ApprovalView.unassignLeaveRequest, name='unassignLeaveRequest'),
]
