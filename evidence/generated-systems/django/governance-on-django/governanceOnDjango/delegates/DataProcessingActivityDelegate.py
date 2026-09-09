from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.DataCategory import DataCategory
from governanceOnDjango.models.System_ import System_
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.models.PrivacyNotice import PrivacyNotice
from governanceOnDjango.models.ThirdParty import ThirdParty
from governanceOnDjango.models.Consent import Consent
from governanceOnDjango.models.DataBreach import DataBreach
from governanceOnDjango.models.DataSubjectRequest import DataSubjectRequest
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataProcessingActivity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProcessingActivityDelegate Declaration
#======================================================================
class DataProcessingActivityDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataProcessingActivityId ):
		try:	
			dataProcessingActivity = DataProcessingActivity.objects.filter(id=dataProcessingActivityId)
			return dataProcessingActivity.first();
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError("DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataProcessingActivity):
		for model in serializers.deserialize("json", dataProcessingActivity):
			model.save()
			return model;

	def create(self, dataProcessingActivity):
		dataProcessingActivity.save()
		return dataProcessingActivity;

	def saveFromJson(self, dataProcessingActivity):
		for model in serializers.deserialize("json", dataProcessingActivity):
			model.save()
			return dataProcessingActivity;
	
	def save(self, dataProcessingActivity):
		dataProcessingActivity.save()
		return dataProcessingActivity;
	
	def delete(self, dataProcessingActivityId ):
		errMsg = "Failed to delete DataProcessingActivity from db using id " + str(dataProcessingActivityId)
		
		try:
			dataProcessingActivity = DataProcessingActivity.objects.get(id=dataProcessingActivityId)
			dataProcessingActivity.delete()
			return True
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError("DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataProcessingActivity.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataProcessingActivity from db")
		except Exception:
			return None;
		
	def assignOrganization( self, dataProcessingActivityId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on DataProcessingActivity"

		try:
			# get the DataProcessingActivity from db
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			dataProcessingActivity.organization = organization
			
			#save it
			dataProcessingActivity.save()

			# reload and return the appropriate version					
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, dataProcessingActivityId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on DataProcessingActivity"

		try:
			# get the DataProcessingActivity from db
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()	
			
			# assign to None for unassignment
			dataProcessingActivity.organization = None			

			#save it
			dataProcessingActivity.save()

			# reload and return the appropriate version					
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except Exception:
			return None;
		
	def addDataCategories( self, dataProcessingActivityId, dataCategoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

		errMsg = "Failed to add elements " + str(dataCategoriesIds) + " for DataCategories on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = dataCategoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataCategory		
				dataCategory = DataCategoryDelegate().get(id).first();	
				# add the DataCategory
				dataProcessingActivity.dataCategories.add(dataCategory)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataCategories( self, dataProcessingActivityId, dataCategoriesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataCategoryDelegate import DataCategoryDelegate

		errMsg = "Failed to remove elements " + str(dataCategoriesIds) + " for DataCategories on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = dataCategoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataCategory		
				dataCategory = DataCategoryDelegate().get(id).first();	
				# add the DataCategory
				dataProcessingActivity.dataCategories.remove(dataCategory)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except DataCategory.DoesNotExist:
			raise ProcessingError(errMsg + " : DataCategory does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSystems( self, dataProcessingActivityId, systemsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.System_Delegate import System_Delegate

		errMsg = "Failed to add elements " + str(systemsIds) + " for Systems on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = systemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the System_		
				system_ = System_Delegate().get(id).first();	
				# add the System_
				dataProcessingActivity.systems.add(system_)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSystems( self, dataProcessingActivityId, systemsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.System_Delegate import System_Delegate

		errMsg = "Failed to remove elements " + str(systemsIds) + " for Systems on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = systemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the System_		
				system_ = System_Delegate().get(id).first();	
				# add the System_
				dataProcessingActivity.systems.remove(system_)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except System_.DoesNotExist:
			raise ProcessingError(errMsg + " : System_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRecords( self, dataProcessingActivityId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to add elements " + str(recordsIds) + " for Records on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				dataProcessingActivity.records.add(record_)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRecords( self, dataProcessingActivityId, recordsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to remove elements " + str(recordsIds) + " for Records on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = recordsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Record_		
				record_ = Record_Delegate().get(id).first();	
				# add the Record_
				dataProcessingActivity.records.remove(record_)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPrivacyNotices( self, dataProcessingActivityId, privacyNoticesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PrivacyNoticeDelegate import PrivacyNoticeDelegate

		errMsg = "Failed to add elements " + str(privacyNoticesIds) + " for PrivacyNotices on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = privacyNoticesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PrivacyNotice		
				privacyNotice = PrivacyNoticeDelegate().get(id).first();	
				# add the PrivacyNotice
				dataProcessingActivity.privacyNotices.add(privacyNotice)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePrivacyNotices( self, dataProcessingActivityId, privacyNoticesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PrivacyNoticeDelegate import PrivacyNoticeDelegate

		errMsg = "Failed to remove elements " + str(privacyNoticesIds) + " for PrivacyNotices on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = privacyNoticesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PrivacyNotice		
				privacyNotice = PrivacyNoticeDelegate().get(id).first();	
				# add the PrivacyNotice
				dataProcessingActivity.privacyNotices.remove(privacyNotice)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addThirdParties( self, dataProcessingActivityId, thirdPartiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to add elements " + str(thirdPartiesIds) + " for ThirdParties on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = thirdPartiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ThirdParty		
				thirdParty = ThirdPartyDelegate().get(id).first();	
				# add the ThirdParty
				dataProcessingActivity.thirdParties.add(thirdParty)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeThirdParties( self, dataProcessingActivityId, thirdPartiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

		errMsg = "Failed to remove elements " + str(thirdPartiesIds) + " for ThirdParties on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = thirdPartiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ThirdParty		
				thirdParty = ThirdPartyDelegate().get(id).first();	
				# add the ThirdParty
				dataProcessingActivity.thirdParties.remove(thirdParty)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addConsents( self, dataProcessingActivityId, consentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to add elements " + str(consentsIds) + " for Consents on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				dataProcessingActivity.consents.add(consent)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeConsents( self, dataProcessingActivityId, consentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to remove elements " + str(consentsIds) + " for Consents on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				dataProcessingActivity.consents.remove(consent)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataBreaches( self, dataProcessingActivityId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to add elements " + str(dataBreachesIds) + " for DataBreaches on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				dataProcessingActivity.dataBreaches.add(dataBreach)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataBreaches( self, dataProcessingActivityId, dataBreachesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataBreachDelegate import DataBreachDelegate

		errMsg = "Failed to remove elements " + str(dataBreachesIds) + " for DataBreaches on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = dataBreachesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataBreach		
				dataBreach = DataBreachDelegate().get(id).first();	
				# add the DataBreach
				dataProcessingActivity.dataBreaches.remove(dataBreach)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except DataBreach.DoesNotExist:
			raise ProcessingError(errMsg + " : DataBreach does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataSubjectRequests( self, dataProcessingActivityId, dataSubjectRequestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataSubjectRequestDelegate import DataSubjectRequestDelegate

		errMsg = "Failed to add elements " + str(dataSubjectRequestsIds) + " for DataSubjectRequests on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = dataSubjectRequestsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSubjectRequest		
				dataSubjectRequest = DataSubjectRequestDelegate().get(id).first();	
				# add the DataSubjectRequest
				dataProcessingActivity.dataSubjectRequests.add(dataSubjectRequest)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataSubjectRequests( self, dataProcessingActivityId, dataSubjectRequestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataSubjectRequestDelegate import DataSubjectRequestDelegate

		errMsg = "Failed to remove elements " + str(dataSubjectRequestsIds) + " for DataSubjectRequests on DataProcessingActivity"

		try:
			# get the DataProcessingActivity
			dataProcessingActivity = self.get( dataProcessingActivityId ).first()
				
			# split on a comma with no spaces
			idList = dataSubjectRequestsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSubjectRequest		
				dataSubjectRequest = DataSubjectRequestDelegate().get(id).first();	
				# add the DataSubjectRequest
				dataProcessingActivity.dataSubjectRequests.remove(dataSubjectRequest)
				
			# save it		
			dataProcessingActivity.save()
			
			# reload and return the appropriate version
			return self.get( dataProcessingActivityId );
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity with id " + str(dataProcessingActivityId) + " does not exist.")
		except DataSubjectRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSubjectRequest does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
