import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.TypeCertificate import TypeCertificate
from aerospaceOnDjango.delegates.TypeCertificateDelegate import TypeCertificateDelegate

 #======================================================================
# 
# Encapsulates data for model TypeCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TypeCertificateTest Declaration
#======================================================================
class TypeCertificateTest (TestCase) :
	def test_crud(self) :
		typeCertificate = TypeCertificate()
		typeCertificate.certificateNumber = "default certificateNumber field value"
		typeCertificate.authority = "default authority field value"
		
		delegate = TypeCertificateDelegate()
		responseObj = delegate.create(typeCertificate)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


