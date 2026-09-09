import datetime

from django.test import TestCase
from django.utils import timezone
from aerospaceOnDjango.models.ProductionCertificate import ProductionCertificate
from aerospaceOnDjango.delegates.ProductionCertificateDelegate import ProductionCertificateDelegate

 #======================================================================
# 
# Encapsulates data for model ProductionCertificate
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionCertificateTest Declaration
#======================================================================
class ProductionCertificateTest (TestCase) :
	def test_crud(self) :
		productionCertificate = ProductionCertificate()
		productionCertificate.certificateNumber = "default certificateNumber field value"
		productionCertificate.authority = "default authority field value"
		
		delegate = ProductionCertificateDelegate()
		responseObj = delegate.create(productionCertificate)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


