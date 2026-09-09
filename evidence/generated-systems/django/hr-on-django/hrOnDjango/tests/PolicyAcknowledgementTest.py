import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.PolicyAcknowledgement import PolicyAcknowledgement
from hrOnDjango.delegates.PolicyAcknowledgementDelegate import PolicyAcknowledgementDelegate

 #======================================================================
# 
# Encapsulates data for model PolicyAcknowledgement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyAcknowledgementTest Declaration
#======================================================================
class PolicyAcknowledgementTest (TestCase) :
	def test_crud(self) :
		policyAcknowledgement = PolicyAcknowledgement()
		policyAcknowledgement.acknowledgementDate = datetime.datetime.now()
		policyAcknowledgement.status = "default status field value"
		
		delegate = PolicyAcknowledgementDelegate()
		responseObj = delegate.create(policyAcknowledgement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


