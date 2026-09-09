import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.Evidence import Evidence
from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

 #======================================================================
# 
# Encapsulates data for model Evidence
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EvidenceTest Declaration
#======================================================================
class EvidenceTest (TestCase) :
	def test_crud(self) :
		evidence = Evidence()
		evidence.title = "default title field value"
		evidence.locationUrl = "default locationUrl field value"
		evidence.receivedDate = datetime.datetime.now()
		evidence.evidenceType = "default evidenceType field value"
		
		delegate = EvidenceDelegate()
		responseObj = delegate.create(evidence)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


