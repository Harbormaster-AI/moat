from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.DataProvider import DataProvider
from advertisingOnDjango.models.AudienceSegment import AudienceSegment
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataProviderDelegate Declaration
#======================================================================
class DataProviderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataProviderId ):
		try:	
			dataProvider = DataProvider.objects.filter(id=dataProviderId)
			return dataProvider.first();
		except DataProvider.DoesNotExist:
			raise ProcessingError("DataProvider with id " + str(dataProviderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataProvider):
		for model in serializers.deserialize("json", dataProvider):
			model.save()
			return model;

	def create(self, dataProvider):
		dataProvider.save()
		return dataProvider;

	def saveFromJson(self, dataProvider):
		for model in serializers.deserialize("json", dataProvider):
			model.save()
			return dataProvider;
	
	def save(self, dataProvider):
		dataProvider.save()
		return dataProvider;
	
	def delete(self, dataProviderId ):
		errMsg = "Failed to delete DataProvider from db using id " + str(dataProviderId)
		
		try:
			dataProvider = DataProvider.objects.get(id=dataProviderId)
			dataProvider.delete()
			return True
		except DataProvider.DoesNotExist:
			raise ProcessingError("DataProvider with id " + str(dataProviderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataProvider.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataProvider from db")
		except Exception:
			return None;
		
	def addAudienceSegments( self, dataProviderId, audienceSegmentsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

		errMsg = "Failed to add elements " + str(audienceSegmentsIds) + " for AudienceSegments on DataProvider"

		try:
			# get the DataProvider
			dataProvider = self.get( dataProviderId ).first()
				
			# split on a comma with no spaces
			idList = audienceSegmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AudienceSegment		
				audienceSegment = AudienceSegmentDelegate().get(id).first();	
				# add the AudienceSegment
				dataProvider.audienceSegments.add(audienceSegment)
				
			# save it		
			dataProvider.save()
			
			# reload and return the appropriate version
			return self.get( dataProviderId );
		except DataProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProvider with id " + str(dataProviderId) + " does not exist.")
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAudienceSegments( self, dataProviderId, audienceSegmentsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AudienceSegmentDelegate import AudienceSegmentDelegate

		errMsg = "Failed to remove elements " + str(audienceSegmentsIds) + " for AudienceSegments on DataProvider"

		try:
			# get the DataProvider
			dataProvider = self.get( dataProviderId ).first()
				
			# split on a comma with no spaces
			idList = audienceSegmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AudienceSegment		
				audienceSegment = AudienceSegmentDelegate().get(id).first();	
				# add the AudienceSegment
				dataProvider.audienceSegments.remove(audienceSegment)
				
			# save it		
			dataProvider.save()
			
			# reload and return the appropriate version
			return self.get( dataProviderId );
		except DataProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProvider with id " + str(dataProviderId) + " does not exist.")
		except AudienceSegment.DoesNotExist:
			raise ProcessingError(errMsg + " : AudienceSegment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
