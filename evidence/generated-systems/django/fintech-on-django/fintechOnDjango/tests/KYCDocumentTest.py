import datetime

from django.test import TestCase
from django.utils import timezone
from fintechOnDjango.models.KYCDocument import KYCDocument
from fintechOnDjango.delegates.KYCDocumentDelegate import KYCDocumentDelegate

 #======================================================================
# 
# Encapsulates data for model KYCDocument
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCDocumentTest Declaration
#======================================================================
class KYCDocumentTest (TestCase) :
	def test_crud(self) :
		kYCDocument = KYCDocument()
		kYCDocument.reference = "default reference field value"
		kYCDocument.issuedCountry = "default issuedCountry field value"
		kYCDocument.expirationDate = datetime.datetime.now()
		kYCDocument.documentType = "default documentType field value"
		kYCDocument.status = "default status field value"
		
		delegate = KYCDocumentDelegate()
		responseObj = delegate.create(kYCDocument)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


