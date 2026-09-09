import datetime

from django.test import TestCase
from django.utils import timezone
from insuranceOnDjango.models.Agent import Agent
from insuranceOnDjango.delegates.AgentDelegate import AgentDelegate

 #======================================================================
# 
# Encapsulates data for model Agent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgentTest Declaration
#======================================================================
class AgentTest (TestCase) :
	def test_crud(self) :
		agent = Agent()
		agent.firstName = "default firstName field value"
		agent.lastName = "default lastName field value"
		agent.licenseId = "default licenseId field value"
		agent.status = "default status field value"
		
		delegate = AgentDelegate()
		responseObj = delegate.create(agent)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


