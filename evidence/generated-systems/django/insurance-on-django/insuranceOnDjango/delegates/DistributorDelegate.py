from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Distributor import Distributor
from insuranceOnDjango.models.Insurer import Insurer
from insuranceOnDjango.models.Agent import Agent
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Distributor
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DistributorDelegate Declaration
#======================================================================
class DistributorDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, distributorId ):
		try:	
			distributor = Distributor.objects.filter(id=distributorId)
			return distributor.first();
		except Distributor.DoesNotExist:
			raise ProcessingError("Distributor with id " + str(distributorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, distributor):
		for model in serializers.deserialize("json", distributor):
			model.save()
			return model;

	def create(self, distributor):
		distributor.save()
		return distributor;

	def saveFromJson(self, distributor):
		for model in serializers.deserialize("json", distributor):
			model.save()
			return distributor;
	
	def save(self, distributor):
		distributor.save()
		return distributor;
	
	def delete(self, distributorId ):
		errMsg = "Failed to delete Distributor from db using id " + str(distributorId)
		
		try:
			distributor = Distributor.objects.get(id=distributorId)
			distributor.delete()
			return True
		except Distributor.DoesNotExist:
			raise ProcessingError("Distributor with id " + str(distributorId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Distributor.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Distributor from db")
		except Exception:
			return None;
		
	def addInsurers( self, distributorId, insurersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

		errMsg = "Failed to add elements " + str(insurersIds) + " for Insurers on Distributor"

		try:
			# get the Distributor
			distributor = self.get( distributorId ).first()
				
			# split on a comma with no spaces
			idList = insurersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Insurer		
				insurer = InsurerDelegate().get(id).first();	
				# add the Insurer
				distributor.insurers.add(insurer)
				
			# save it		
			distributor.save()
			
			# reload and return the appropriate version
			return self.get( distributorId );
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInsurers( self, distributorId, insurersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsurerDelegate import InsurerDelegate

		errMsg = "Failed to remove elements " + str(insurersIds) + " for Insurers on Distributor"

		try:
			# get the Distributor
			distributor = self.get( distributorId ).first()
				
			# split on a comma with no spaces
			idList = insurersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Insurer		
				insurer = InsurerDelegate().get(id).first();	
				# add the Insurer
				distributor.insurers.remove(insurer)
				
			# save it		
			distributor.save()
			
			# reload and return the appropriate version
			return self.get( distributorId );
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAgents( self, distributorId, agentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.AgentDelegate import AgentDelegate

		errMsg = "Failed to add elements " + str(agentsIds) + " for Agents on Distributor"

		try:
			# get the Distributor
			distributor = self.get( distributorId ).first()
				
			# split on a comma with no spaces
			idList = agentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Agent		
				agent = AgentDelegate().get(id).first();	
				# add the Agent
				distributor.agents.add(agent)
				
			# save it		
			distributor.save()
			
			# reload and return the appropriate version
			return self.get( distributorId );
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAgents( self, distributorId, agentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.AgentDelegate import AgentDelegate

		errMsg = "Failed to remove elements " + str(agentsIds) + " for Agents on Distributor"

		try:
			# get the Distributor
			distributor = self.get( distributorId ).first()
				
			# split on a comma with no spaces
			idList = agentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Agent		
				agent = AgentDelegate().get(id).first();	
				# add the Agent
				distributor.agents.remove(agent)
				
			# save it		
			distributor.save()
			
			# reload and return the appropriate version
			return self.get( distributorId );
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Agent.DoesNotExist:
			raise ProcessingError(errMsg + " : Agent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPolicies( self, distributorId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on Distributor"

		try:
			# get the Distributor
			distributor = self.get( distributorId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				distributor.policies.add(policy)
				
			# save it		
			distributor.save()
			
			# reload and return the appropriate version
			return self.get( distributorId );
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, distributorId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on Distributor"

		try:
			# get the Distributor
			distributor = self.get( distributorId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				distributor.policies.remove(policy)
				
			# save it		
			distributor.save()
			
			# reload and return the appropriate version
			return self.get( distributorId );
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor with id " + str(distributorId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
