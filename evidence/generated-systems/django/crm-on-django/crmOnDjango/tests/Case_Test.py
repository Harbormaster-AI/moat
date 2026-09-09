import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.delegates.Case_Delegate import Case_Delegate

 #======================================================================
# 
# Encapsulates data for model Case_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Case_Test Declaration
#======================================================================
class Case_Test (TestCase) :
	def test_crud(self) :
		case_ = Case_()
		case_.caseNumber = "default caseNumber field value"
		case_.subject = "default subject field value"
		case_.description = "default description field value"
		case_.slaDue = "default slaDue field value"
		case_.status = "default status field value"
		case_.priority = "default priority field value"
		case_.origin = "default origin field value"
		case_.severity = "default severity field value"
		
		delegate = Case_Delegate()
		responseObj = delegate.create(case_)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


