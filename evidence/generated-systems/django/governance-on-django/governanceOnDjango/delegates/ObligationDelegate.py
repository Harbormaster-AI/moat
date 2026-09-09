from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.models.Regulation import Regulation
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.Contract import Contract
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Obligation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObligationDelegate Declaration
#======================================================================
class ObligationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, obligationId ):
		try:	
			obligation = Obligation.objects.filter(id=obligationId)
			return obligation.first();
		except Obligation.DoesNotExist:
			raise ProcessingError("Obligation with id " + str(obligationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, obligation):
		for model in serializers.deserialize("json", obligation):
			model.save()
			return model;

	def create(self, obligation):
		obligation.save()
		return obligation;

	def saveFromJson(self, obligation):
		for model in serializers.deserialize("json", obligation):
			model.save()
			return obligation;
	
	def save(self, obligation):
		obligation.save()
		return obligation;
	
	def delete(self, obligationId ):
		errMsg = "Failed to delete Obligation from db using id " + str(obligationId)
		
		try:
			obligation = Obligation.objects.get(id=obligationId)
			obligation.delete()
			return True
		except Obligation.DoesNotExist:
			raise ProcessingError("Obligation with id " + str(obligationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Obligation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Obligation from db")
		except Exception:
			return None;
		
	def assignRegulation( self, obligationId, regulationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RegulationDelegate import RegulationDelegate

		errMsg = "Failed to assign element " + str(regulationId) + " for Regulation on Obligation"

		try:
			# get the Obligation from db
			obligation = self.get( obligationId ).first()	
			
			# get the Regulation from db
			regulation = RegulationDelegate().get(regulationId).first();
			
			# assign the Regulation		
			obligation.regulation = regulation
			
			#save it
			obligation.save()

			# reload and return the appropriate version					
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Regulation.DoesNotExist:
			raise ProcessingError(errMsg + " : Regulation with id " + str(regulationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRegulation( self, obligationId ):
		errMsg = "Failed to unassign element " + str(regulationId) + " for Regulation on Obligation"

		try:
			# get the Obligation from db
			obligation = self.get( obligationId ).first()	
			
			# assign to None for unassignment
			obligation.regulation = None			

			#save it
			obligation.save()

			# reload and return the appropriate version					
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Exception:
			return None;
		
	def addControls( self, obligationId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to add elements " + str(controlsIds) + " for Controls on Obligation"

		try:
			# get the Obligation
			obligation = self.get( obligationId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				obligation.controls.add(control)
				
			# save it		
			obligation.save()
			
			# reload and return the appropriate version
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControls( self, obligationId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to remove elements " + str(controlsIds) + " for Controls on Obligation"

		try:
			# get the Obligation
			obligation = self.get( obligationId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				obligation.controls.remove(control)
				
			# save it		
			obligation.save()
			
			# reload and return the appropriate version
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPolicies( self, obligationId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on Obligation"

		try:
			# get the Obligation
			obligation = self.get( obligationId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				obligation.policies.add(policy)
				
			# save it		
			obligation.save()
			
			# reload and return the appropriate version
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, obligationId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on Obligation"

		try:
			# get the Obligation
			obligation = self.get( obligationId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				obligation.policies.remove(policy)
				
			# save it		
			obligation.save()
			
			# reload and return the appropriate version
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContracts( self, obligationId, contractsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on Obligation"

		try:
			# get the Obligation
			obligation = self.get( obligationId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				obligation.contracts.add(contract)
				
			# save it		
			obligation.save()
			
			# reload and return the appropriate version
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, obligationId, contractsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on Obligation"

		try:
			# get the Obligation
			obligation = self.get( obligationId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				obligation.contracts.remove(contract)
				
			# save it		
			obligation.save()
			
			# reload and return the appropriate version
			return self.get( obligationId );
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
