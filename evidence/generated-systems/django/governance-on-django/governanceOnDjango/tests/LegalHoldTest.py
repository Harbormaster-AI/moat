import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.LegalHold import LegalHold
from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

 #======================================================================
# 
# Encapsulates data for model LegalHold
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LegalHoldTest Declaration
#======================================================================
class LegalHoldTest (TestCase) :
	def test_crud(self) :
		legalHold = LegalHold()
		legalHold.name = "default name field value"
		legalHold.reason = "default reason field value"
		legalHold.issuedDate = datetime.datetime.now()
		legalHold.releaseDate = datetime.datetime.now()
		legalHold.holdStatus = "default holdStatus field value"
		
		delegate = LegalHoldDelegate()
		responseObj = delegate.create(legalHold)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


