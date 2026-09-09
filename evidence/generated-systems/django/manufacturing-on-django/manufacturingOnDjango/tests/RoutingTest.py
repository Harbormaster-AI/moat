import datetime

from django.test import TestCase
from django.utils import timezone
from manufacturingOnDjango.models.Routing import Routing
from manufacturingOnDjango.delegates.RoutingDelegate import RoutingDelegate

 #======================================================================
# 
# Encapsulates data for model Routing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoutingTest Declaration
#======================================================================
class RoutingTest (TestCase) :
	def test_crud(self) :
		routing = Routing()
		routing.routingNumber = "default routingNumber field value"
		routing.revision = "default revision field value"
		routing.effectivityStart = datetime.datetime.now()
		routing.effectivityEnd = datetime.datetime.now()
		routing.routingType = "default routingType field value"
		routing.status = "default status field value"
		
		delegate = RoutingDelegate()
		responseObj = delegate.create(routing)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


