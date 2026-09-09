import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Issue import Issue
from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

 #======================================================================
# 
# Encapsulates data for model Issue
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IssueTest Declaration
#======================================================================
class IssueTest (TestCase) :
	def test_crud(self) :
		issue = Issue()
		issue.title = "default title field value"
		issue.openedDate = datetime.datetime.now()
		issue.closedDate = datetime.datetime.now()
		issue.issueType = "default issueType field value"
		issue.priority = "default priority field value"
		issue.status = "default status field value"
		
		delegate = IssueDelegate()
		responseObj = delegate.create(issue)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


