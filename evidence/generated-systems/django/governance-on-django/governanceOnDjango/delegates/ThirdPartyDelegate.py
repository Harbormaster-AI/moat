from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.ThirdParty import ThirdParty
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.ThirdPartyAssessment import ThirdPartyAssessment
from governanceOnDjango.models.Contract import Contract
from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.models.DataBreach import DataBreach
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ThirdParty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyDelegate Declaration
#======================================================================
class ThirdPartyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, thirdPartyId ):
		try:	
			thirdParty = ThirdParty.objects.filter(id=thirdPartyId)
			return thirdParty.first();
		except ThirdParty.DoesNotExist:
			raise ProcessingError("ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, thirdParty):
		for model in serializers.deserialize("json", thirdParty):
			model.save()
			return model;

	def create(self, thirdParty):
		thirdParty.save()
		return thirdParty;

	def saveFromJson(self, thirdParty):
		for model in serializers.deserialize("json", thirdParty):
			model.save()
			return thirdParty;
	
	def save(self, thirdParty):
		thirdParty.save()
		return thirdParty;
	
	def delete(self, thirdPartyId ):
		errMsg = "Failed to delete ThirdParty from db using id " + str(thirdPartyId)
		
		try:
			thirdParty = ThirdParty.objects.get(id=thirdPartyId)
			thirdParty.delete()
			return True
		except ThirdParty.DoesNotExist:
			raise ProcessingError("ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ThirdParty.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ThirdParty from db")
		except Exception:
			return None;
		
	def assignOrganization( self, thirdPartyId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on ThirdParty"

		try:
			# get the ThirdParty from db
			thirdParty = self.get( thirdPartyId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			thirdParty.organization = organization
			
			#save it
			thirdParty.save()

			# reload and return the appropriate version					
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, thirdPartyId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on ThirdParty"

		try:
			# get the ThirdParty from db
			thirdParty = self.get( thirdPartyId ).first()	
			
			# assign to None for unassignment
			thirdParty.organization = None			

			#save it
			thirdParty.save()

			# reload and return the appropriate version					
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Exception:
			return None;
		
	def addProcessingActivities( self, thirdPartyId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				thirdParty.processingActivities.add(dataProcessingActivity)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, thirdPartyId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				thirdParty.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAssessments( self, thirdPartyId, assessmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyAssessmentDelegate import ThirdPartyAssessmentDelegate

		errMsg = "Failed to add elements " + str(assessmentsIds) + " for Assessments on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = assessmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ThirdPartyAssessment		
				thirdPartyAssessment = ThirdPartyAssessmentDelegate().get(id).first();	
				# add the ThirdPartyAssessment
				thirdParty.assessments.add(thirdPartyAssessment)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdPartyAssessment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAssessments( self, thirdPartyId, assessmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyAssessmentDelegate import ThirdPartyAssessmentDelegate

		errMsg = "Failed to remove elements " + str(assessmentsIds) + " for Assessments on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = assessmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ThirdPartyAssessment		
				thirdPartyAssessment = ThirdPartyAssessmentDelegate().get(id).first();	
				# add the ThirdPartyAssessment
				thirdParty.assessments.remove(thirdPartyAssessment)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except ThirdPartyAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdPartyAssessment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContracts( self, thirdPartyId, contractsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				thirdParty.contracts.add(contract)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, thirdPartyId, contractsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				thirdParty.contracts.remove(contract)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addObligations( self, thirdPartyId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to add elements " + str(obligationsIds) + " for Obligations on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				thirdParty.obligations.add(obligation)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObligations( self, thirdPartyId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to remove elements " + str(obligationsIds) + " for Obligations on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				thirdParty.obligations.remove(obligation)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataBreaches( self, thirdPartyId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to add elements " + str(dataBreachesIds) + " for DataBreaches on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				thirdParty.dataBreaches.add(dataBreach)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataBreaches( self, thirdPartyId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to remove elements " + str(dataBreachesIds) + " for DataBreaches on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				thirdParty.dataBreaches.remove(dataBreach)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
