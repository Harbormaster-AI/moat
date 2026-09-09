import datetime

from django.test import TestCase
from django.utils import timezone
from analyticsOnDjango.models.Anomaly import Anomaly
from analyticsOnDjango.delegates.AnomalyDelegate import AnomalyDelegate

 #======================================================================
# 
# Encapsulates data for model Anomaly
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnomalyTest Declaration
#======================================================================
class AnomalyTest (TestCase) :
	def test_crud(self) :
		anomaly = Anomaly()
		anomaly.occurredAt = datetime.datetime.now()
		anomaly.details = "default details field value"
		anomaly.anomalyType = "default anomalyType field value"
		anomaly.severity = "default severity field value"
		
		delegate = AnomalyDelegate()
		responseObj = delegate.create(anomaly)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


