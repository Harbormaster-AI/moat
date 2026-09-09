from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.DataBreach import DataBreach
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.DataCategory import DataCategory
from governanceOnDjango.models.ThirdParty import ThirdParty
from governanceOnDjango.models.Matter import Matter
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataBreach
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataBreachDelegate Declaration
#======================================================================
class DataBreachDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataBreachId ):
		try:	
			dataBreach = DataBreach.objects.filter(id=dataBreachId)
			return dataBreach.first();
		except DataBreach.DoesNotExist:
			raise ProcessingError("DataBreach with id " + str(dataBreachId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataBreach):
		for model in serializers.deserialize("json", dataBreach):
			model.save()
			return model;

	def create(self, dataBreach):
		dataBreach.save()
		return dataBreach;

	def saveFromJson(self, dataBreach):
		for model in serializers.deserialize("json", dataBreach):
			model.save()
			return dataBreach;
	
	def save(self, dataBreach):
		dataBreach.save()
		return dataBreach;
	
	def delete(self, dataBreachId ):
		errMsg = "Failed to delete DataBreach from db using id " + str(dataBreachId)
		
		try:
			dataBreach = DataBreach.objects.get(id=dataBreachId)
			dataBreach.delete()
			return True
		except DataBreach.DoesNotExist:
			raise ProcessingError("DataBreach with id " + str(dataBreachId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataBreach.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataBreach from db")
		except Exception:
			return None;
		
	def assignOrganization( self, dataBreachId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on DataBreach"

		try:
			# get the DataBreach from db
			dataBreach = self.get( dataBreachId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			dataBreach.organization = organization
			
			#save it
			dataBreach.save()

			# reload and return the appropriate version					
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, dataBreachId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on DataBreach"

		try:
			# get the DataBreach from db
			dataBreach = self.get( dataBreachId ).first()	
			
			# assign to None for unassignment
			dataBreach.organization = None			

			#save it
			dataBreach.save()

			# reload and return the appropriate version					
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMatter( self, dataBreachId, matterId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.MatterDelegate import MatterDelegate

		errMsg = "Failed to assign element " + str(matterId) + " for Matter on DataBreach"

		try:
			# get the DataBreach from db
			dataBreach = self.get( dataBreachId ).first()	
			
			# get the Matter from db
			matter = MatterDelegate().get(matterId).first();
			
			# assign the Matter		
			dataBreach.matter = matter
			
			#save it
			dataBreach.save()

			# reload and return the appropriate version					
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except Matter.DoesNotExist:
			raise ProcessingError(errMsg + " : Matter with id " + str(matterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMatter( self, dataBreachId ):
		errMsg = "Failed to unassign element " + str(matterId) + " for Matter on DataBreach"

		try:
			# get the DataBreach from db
			dataBreach = self.get( dataBreachId ).first()	
			
			# assign to None for unassignment
			dataBreach.matter = None			

			#save it
			dataBreach.save()

			# reload and return the appropriate version					
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except Exception:
			return None;
		
	def addProcessingActivities( self, dataBreachId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on DataBreach"

		try:
			# get the DataBreach
			dataBreach = self.get( dataBreachId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				dataBreach.processingActivities.add(dataProcessingActivity)
				
			# save it		
			dataBreach.save()
			
			# reload and return the appropriate version
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, dataBreachId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on DataBreach"

		try:
			# get the DataBreach
			dataBreach = self.get( dataBreachId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				dataBreach.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			dataBreach.save()
			
			# reload and return the appropriate version
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataCategories( self, dataBreachId, dataCategoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

		errMsg = "Failed to add elements " + str(dataCategoriesIds) + " for DataCategories on DataBreach"

		try:
			# get the DataBreach
			dataBreach = self.get( dataBreachId ).first()
				
			# split on a comma with no spaces
			idList = dataCategoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataCategory		
				dataCategory = DataCategoryDelegate().get(id).first();	
				# add the DataCategory
				dataBreach.dataCategories.add(dataCategory)
				
			# save it		
			dataBreach.save()
			
			# reload and return the appropriate version
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataCategories( self, dataBreachId, dataCategoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

		errMsg = "Failed to remove elements " + str(dataCategoriesIds) + " for DataCategories on DataBreach"

		try:
			# get the DataBreach
			dataBreach = self.get( dataBreachId ).first()
				
			# split on a comma with no spaces
			idList = dataCategoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataCategory		
				dataCategory = DataCategoryDelegate().get(id).first();	
				# add the DataCategory
				dataBreach.dataCategories.remove(dataCategory)
				
			# save it		
			dataBreach.save()
			
			# reload and return the appropriate version
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addThirdParties( self, dataBreachId, thirdPartiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to add elements " + str(thirdPartiesIds) + " for ThirdParties on DataBreach"

		try:
			# get the DataBreach
			dataBreach = self.get( dataBreachId ).first()
				
			# split on a comma with no spaces
			idList = thirdPartiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ThirdParty		
				thirdParty = ThirdPartyDelegate().get(id).first();	
				# add the ThirdParty
				dataBreach.thirdParties.add(thirdParty)
				
			# save it		
			dataBreach.save()
			
			# reload and return the appropriate version
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeThirdParties( self, dataBreachId, thirdPartiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to remove elements " + str(thirdPartiesIds) + " for ThirdParties on DataBreach"

		try:
			# get the DataBreach
			dataBreach = self.get( dataBreachId ).first()
				
			# split on a comma with no spaces
			idList = thirdPartiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ThirdParty		
				thirdParty = ThirdPartyDelegate().get(id).first();	
				# add the ThirdParty
				dataBreach.thirdParties.remove(thirdParty)
				
			# save it		
			dataBreach.save()
			
			# reload and return the appropriate version
			return self.get( dataBreachId );
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach with id " + str(dataBreachId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
