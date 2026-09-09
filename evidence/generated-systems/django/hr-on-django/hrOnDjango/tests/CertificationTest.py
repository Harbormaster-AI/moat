import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Certification import Certification
from hrOnDjango.delegates.CertificationDelegate import CertificationDelegate

 #======================================================================
# 
# Encapsulates data for model Certification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CertificationTest Declaration
#======================================================================
class CertificationTest (TestCase) :
	def test_crud(self) :
		certification = Certification()
		certification.name = "default name field value"
		certification.issuer = "default issuer field value"
		certification.validFrom = datetime.datetime.now()
		certification.validTo = datetime.datetime.now()
		certification.credentialId = "default credentialId field value"
		
		delegate = CertificationDelegate()
		responseObj = delegate.create(certification)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


