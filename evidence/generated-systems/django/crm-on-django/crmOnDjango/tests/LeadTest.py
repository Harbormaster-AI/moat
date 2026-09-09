import datetime

from django.test import TestCase
from django.utils import timezone
from crmOnDjango.models.Lead import Lead
from crmOnDjango.delegates.LeadDelegate import LeadDelegate

 #======================================================================
# 
# Encapsulates data for model Lead
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeadTest Declaration
#======================================================================
class LeadTest (TestCase) :
	def test_crud(self) :
		lead = Lead()
		lead.firstName = "default firstName field value"
		lead.lastName = "default lastName field value"
		lead.company = "default company field value"
		lead.email = "default email field value"
		lead.phone = "default phone field value"
		lead.converted = False
		lead.status = "default status field value"
		lead.source = "default source field value"
		lead.rating = "default rating field value"
		
		delegate = LeadDelegate()
		responseObj = delegate.create(lead)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


