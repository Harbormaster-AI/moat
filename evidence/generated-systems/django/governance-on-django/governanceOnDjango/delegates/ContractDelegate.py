from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Contract import Contract
from governanceOnDjango.models.ThirdParty import ThirdParty
from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.Matter import Matter
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Contract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContractDelegate Declaration
#======================================================================
class ContractDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, contractId ):
		try:	
			contract = Contract.objects.filter(id=contractId)
			return contract.first();
		except Contract.DoesNotExist:
			raise ProcessingError("Contract with id " + str(contractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, contract):
		for model in serializers.deserialize("json", contract):
			model.save()
			return model;

	def create(self, contract):
		contract.save()
		return contract;

	def saveFromJson(self, contract):
		for model in serializers.deserialize("json", contract):
			model.save()
			return contract;
	
	def save(self, contract):
		contract.save()
		return contract;
	
	def delete(self, contractId ):
		errMsg = "Failed to delete Contract from db using id " + str(contractId)
		
		try:
			contract = Contract.objects.get(id=contractId)
			contract.delete()
			return True
		except Contract.DoesNotExist:
			raise ProcessingError("Contract with id " + str(contractId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Contract.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Contract from db")
		except Exception:
			return None;
		
	def assignThirdParty( self, contractId, thirdPartyId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to assign element " + str(thirdPartyId) + " for ThirdParty on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# get the ThirdParty from db
			thirdParty = ThirdPartyDelegate().get(thirdPartyId).first();
			
			# assign the ThirdParty		
			contract.thirdParty = thirdParty
			
			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignThirdParty( self, contractId ):
		errMsg = "Failed to unassign element " + str(thirdPartyId) + " for ThirdParty on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# assign to None for unassignment
			contract.thirdParty = None			

			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMatter( self, contractId, matterId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.MatterDelegate import MatterDelegate

		errMsg = "Failed to assign element " + str(matterId) + " for Matter on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# get the Matter from db
			matter = MatterDelegate().get(matterId).first();
			
			# assign the Matter		
			contract.matter = matter
			
			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMatter( self, contractId ):
		errMsg = "Failed to unassign element " + str(matterId) + " for Matter on Contract"

		try:
			# get the Contract from db
			contract = self.get( contractId ).first()	
			
			# assign to None for unassignment
			contract.matter = None			

			#save it
			contract.save()

			# reload and return the appropriate version					
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Exception:
			return None;
		
	def addObligations( self, contractId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to add elements " + str(obligationsIds) + " for Obligations on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				contract.obligations.add(obligation)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObligations( self, contractId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to remove elements " + str(obligationsIds) + " for Obligations on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				contract.obligations.remove(obligation)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataProcessingActivities( self, contractId, dataProcessingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(dataProcessingActivitiesIds) + " for DataProcessingActivities on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = dataProcessingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				contract.dataProcessingActivities.add(dataProcessingActivity)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataProcessingActivities( self, contractId, dataProcessingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(dataProcessingActivitiesIds) + " for DataProcessingActivities on Contract"

		try:
			# get the Contract
			contract = self.get( contractId ).first()
				
			# split on a comma with no spaces
			idList = dataProcessingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				contract.dataProcessingActivities.remove(dataProcessingActivity)
				
			# save it		
			contract.save()
			
			# reload and return the appropriate version
			return self.get( contractId );
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract with id " + str(contractId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
