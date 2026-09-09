import datetime

from django.test import TestCase
from django.utils import timezone
from governanceOnDjango.models.PrivacyNotice import PrivacyNotice
from governanceOnDjango.delegates.PrivacyNoticeDelegate import PrivacyNoticeDelegate

 #======================================================================
# 
# Encapsulates data for model PrivacyNotice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PrivacyNoticeTest Declaration
#======================================================================
class PrivacyNoticeTest (TestCase) :
	def test_crud(self) :
		privacyNotice = PrivacyNotice()
		privacyNotice.title = "default title field value"
		privacyNotice.audience = "default audience field value"
		privacyNotice.versionLabel = "default versionLabel field value"
		privacyNotice.publicationDate = datetime.datetime.now()
		privacyNotice.publicationUrl = "default publicationUrl field value"
		privacyNotice.status = "default status field value"
		
		delegate = PrivacyNoticeDelegate()
		responseObj = delegate.create(privacyNotice)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


