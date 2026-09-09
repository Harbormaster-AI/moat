import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Attestation import Attestation
from governanceOnDjango.delegates.AttestationDelegate import AttestationDelegate

 #======================================================================
# 
# Encapsulates data for model Attestation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AttestationTest Declaration
#======================================================================
class AttestationTest (TestCase) :
	def test_crud(self) :
		attestation = Attestation()
		attestation.statement = "default statement field value"
		attestation.attestor = "default attestor field value"
		attestation.dateSigned = datetime.datetime.now()
		attestation.result = "default result field value"
		
		delegate = AttestationDelegate()
		responseObj = delegate.create(attestation)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


