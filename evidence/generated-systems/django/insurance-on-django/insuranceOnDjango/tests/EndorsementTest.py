import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Endorsement import Endorsement
from insuranceOnDjango.delegates.EndorsementDelegate import EndorsementDelegate

 #======================================================================
# 
# Encapsulates data for model Endorsement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EndorsementTest Declaration
#======================================================================
class EndorsementTest (TestCase) :
	def test_crud(self) :
		endorsement = Endorsement()
		endorsement.endorsementNumber = "default endorsementNumber field value"
		endorsement.effectiveDate = datetime.datetime.now()
		endorsement.description = "default description field value"
		
		delegate = EndorsementDelegate()
		responseObj = delegate.create(endorsement)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


