import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.CorrectiveAction import CorrectiveAction
from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

 #======================================================================
# 
# Encapsulates data for model CorrectiveAction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CorrectiveActionTest Declaration
#======================================================================
class CorrectiveActionTest (TestCase) :
	def test_crud(self) :
		correctiveAction = CorrectiveAction()
		correctiveAction.actionTitle = "default actionTitle field value"
		correctiveAction.owner = "default owner field value"
		correctiveAction.targetDate = datetime.datetime.now()
		correctiveAction.status = "default status field value"
		
		delegate = CorrectiveActionDelegate()
		responseObj = delegate.create(correctiveAction)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


