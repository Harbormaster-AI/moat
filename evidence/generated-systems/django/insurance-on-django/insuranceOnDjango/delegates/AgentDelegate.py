from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Agent import Agent
from insuranceOnDjango.models.Distributor import Distributor
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Agent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgentDelegate Declaration
#======================================================================
class AgentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, agentId ):
		try:	
			agent = Agent.objects.filter(id=agentId)
			return agent.first();
		except Agent.DoesNotExist:
			raise ProcessingError("Agent with id " + str(agentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, agent):
		for model in serializers.deserialize("json", agent):
			model.save()
			return model;

	def create(self, agent):
		agent.save()
		return agent;

	def saveFromJson(self, agent):
		for model in serializers.deserialize("json", agent):
			model.save()
			return agent;
	
	def save(self, agent):
		agent.save()
		return agent;
	
	def delete(self, agentId ):
		errMsg = "Failed to delete Agent from db using id " + str(agentId)
		
		try:
			agent = Agent.objects.get(id=agentId)
			agent.delete()
			return True
		except Agent.DoesNotExist:
			raise ProcessingError("Agent with id " + str(agentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Agent.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Agent from db")
		except Exception:
			return None;
		
	def assignDistributor( self, agentId, distributorId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.DistributorDelegate import DistributorDelegate

		errMsg = "Failed to assign element " + str(distributorId) + " for Distributor on Agent"

		try:
			# get the Agent from db
			agent = self.get( agentId ).first()	
			
			# get the Distributor from db
			distributor = DistributorDelegate().get(distributorId).first();
			
			# assign the Distributor		
			agent.distributor = distributor
			
			#save it
			agent.save()

			# reload and return the appropriate version					
			return self.get( agentId );
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent with id " + str(agentId) + " does not exist.")
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDistributor( self, agentId ):
		errMsg = "Failed to unassign element " + str(distributorId) + " for Distributor on Agent"

		try:
			# get the Agent from db
			agent = self.get( agentId ).first()	
			
			# assign to None for unassignment
			agent.distributor = None			

			#save it
			agent.save()

			# reload and return the appropriate version					
			return self.get( agentId );
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent with id " + str(agentId) + " does not exist.")
		except Exception:
			return None;
		
	def addPolicies( self, agentId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on Agent"

		try:
			# get the Agent
			agent = self.get( agentId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				agent.policies.add(policy)
				
			# save it		
			agent.save()
			
			# reload and return the appropriate version
			return self.get( agentId );
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent with id " + str(agentId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, agentId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on Agent"

		try:
			# get the Agent
			agent = self.get( agentId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				agent.policies.remove(policy)
				
			# save it		
			agent.save()
			
			# reload and return the appropriate version
			return self.get( agentId );
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent with id " + str(agentId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCustomers( self, agentId, customersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to add elements " + str(customersIds) + " for Customers on Agent"

		try:
			# get the Agent
			agent = self.get( agentId ).first()
				
			# split on a comma with no spaces
			idList = customersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Customer		
				customer = CustomerDelegate().get(id).first();	
				# add the Customer
				agent.customers.add(customer)
				
			# save it		
			agent.save()
			
			# reload and return the appropriate version
			return self.get( agentId );
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent with id " + str(agentId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCustomers( self, agentId, customersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to remove elements " + str(customersIds) + " for Customers on Agent"

		try:
			# get the Agent
			agent = self.get( agentId ).first()
				
			# split on a comma with no spaces
			idList = customersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Customer		
				customer = CustomerDelegate().get(id).first();	
				# add the Customer
				agent.customers.remove(customer)
				
			# save it		
			agent.save()
			
			# reload and return the appropriate version
			return self.get( agentId );
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent with id " + str(agentId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
