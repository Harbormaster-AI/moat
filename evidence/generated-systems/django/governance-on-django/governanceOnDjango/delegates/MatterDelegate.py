from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Matter import Matter
from governanceOnDjango.models.LegalHold import LegalHold
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.DataBreach import DataBreach
from governanceOnDjango.models.Contract import Contract
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Matter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MatterDelegate Declaration
#======================================================================
class MatterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, matterId ):
		try:	
			matter = Matter.objects.filter(id=matterId)
			return matter.first();
		except Matter.DoesNotExist:
			raise ProcessingError("Matter with id " + str(matterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, matter):
		for model in serializers.deserialize("json", matter):
			model.save()
			return model;

	def create(self, matter):
		matter.save()
		return matter;

	def saveFromJson(self, matter):
		for model in serializers.deserialize("json", matter):
			model.save()
			return matter;
	
	def save(self, matter):
		matter.save()
		return matter;
	
	def delete(self, matterId ):
		errMsg = "Failed to delete Matter from db using id " + str(matterId)
		
		try:
			matter = Matter.objects.get(id=matterId)
			matter.delete()
			return True
		except Matter.DoesNotExist:
			raise ProcessingError("Matter with id " + str(matterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Matter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Matter from db")
		except Exception:
			return None;
		
	def assignOrganization( self, matterId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Matter"

		try:
			# get the Matter from db
			matter = self.get( matterId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			matter.organization = organization
			
			#save it
			matter.save()

			# reload and return the appropriate version					
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, matterId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Matter"

		try:
			# get the Matter from db
			matter = self.get( matterId ).first()	
			
			# assign to None for unassignment
			matter.organization = None			

			#save it
			matter.save()

			# reload and return the appropriate version					
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except Exception:
			return None;
		
	def addLegalHolds( self, matterId, legalHoldsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

		errMsg = "Failed to add elements " + str(legalHoldsIds) + " for LegalHolds on Matter"

		try:
			# get the Matter
			matter = self.get( matterId ).first()
				
			# split on a comma with no spaces
			idList = legalHoldsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LegalHold		
				legalHold = LegalHoldDelegate().get(id).first();	
				# add the LegalHold
				matter.legalHolds.add(legalHold)
				
			# save it		
			matter.save()
			
			# reload and return the appropriate version
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLegalHolds( self, matterId, legalHoldsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.LegalHoldDelegate import LegalHoldDelegate

		errMsg = "Failed to remove elements " + str(legalHoldsIds) + " for LegalHolds on Matter"

		try:
			# get the Matter
			matter = self.get( matterId ).first()
				
			# split on a comma with no spaces
			idList = legalHoldsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LegalHold		
				legalHold = LegalHoldDelegate().get(id).first();	
				# add the LegalHold
				matter.legalHolds.remove(legalHold)
				
			# save it		
			matter.save()
			
			# reload and return the appropriate version
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except LegalHold.DoesNotExist:
			raise ProcessingError(errMsg + " : LegalHold does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataBreaches( self, matterId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to add elements " + str(dataBreachesIds) + " for DataBreaches on Matter"

		try:
			# get the Matter
			matter = self.get( matterId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				matter.dataBreaches.add(dataBreach)
				
			# save it		
			matter.save()
			
			# reload and return the appropriate version
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataBreaches( self, matterId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to remove elements " + str(dataBreachesIds) + " for DataBreaches on Matter"

		try:
			# get the Matter
			matter = self.get( matterId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				matter.dataBreaches.remove(dataBreach)
				
			# save it		
			matter.save()
			
			# reload and return the appropriate version
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContracts( self, matterId, contractsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on Matter"

		try:
			# get the Matter
			matter = self.get( matterId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				matter.contracts.add(contract)
				
			# save it		
			matter.save()
			
			# reload and return the appropriate version
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, matterId, contractsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on Matter"

		try:
			# get the Matter
			matter = self.get( matterId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				matter.contracts.remove(contract)
				
			# save it		
			matter.save()
			
			# reload and return the appropriate version
			return self.get( matterId );
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
