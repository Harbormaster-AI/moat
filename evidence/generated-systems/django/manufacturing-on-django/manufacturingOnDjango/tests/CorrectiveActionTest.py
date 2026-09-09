import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.CorrectiveAction import CorrectiveAction
from manufacturingOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

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
		correctiveAction.capaNumber = "default capaNumber field value"
		correctiveAction.rootCause = "default rootCause field value"
		correctiveAction.correctiveAction = "default correctiveAction field value"
		correctiveAction.verificationDate = datetime.datetime.now()
		correctiveAction.status = "default status field value"
		
		delegate = CorrectiveActionDelegate()
		responseObj = delegate.create(correctiveAction)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


