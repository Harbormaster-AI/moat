import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Activity import Activity
from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

 #======================================================================
# 
# Encapsulates data for model Activity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ActivityTest Declaration
#======================================================================
class ActivityTest (TestCase) :
	def test_crud(self) :
		activity = Activity()
		activity.subject = "default subject field value"
		activity.dueDate = datetime.datetime.now()
		activity.startAt = "default startAt field value"
		activity.endAt = "default endAt field value"
		activity.location = "default location field value"
		activity.activityType = "default activityType field value"
		activity.status = "default status field value"
		activity.priority = "default priority field value"
		
		delegate = ActivityDelegate()
		responseObj = delegate.create(activity)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


