from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Forecast import Forecast
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.TimeSeries import TimeSeries
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Forecast
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastDelegate Declaration
#======================================================================
class ForecastDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, forecastId ):
		try:	
			forecast = Forecast.objects.filter(id=forecastId)
			return forecast.first();
		except Forecast.DoesNotExist:
			raise ProcessingError("Forecast with id " + str(forecastId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, forecast):
		for model in serializers.deserialize("json", forecast):
			model.save()
			return model;

	def create(self, forecast):
		forecast.save()
		return forecast;

	def saveFromJson(self, forecast):
		for model in serializers.deserialize("json", forecast):
			model.save()
			return forecast;
	
	def save(self, forecast):
		forecast.save()
		return forecast;
	
	def delete(self, forecastId ):
		errMsg = "Failed to delete Forecast from db using id " + str(forecastId)
		
		try:
			forecast = Forecast.objects.get(id=forecastId)
			forecast.delete()
			return True
		except Forecast.DoesNotExist:
			raise ProcessingError("Forecast with id " + str(forecastId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Forecast.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Forecast from db")
		except Exception:
			return None;
		
	def assignModelVersion( self, forecastId, modelVersionId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to assign element " + str(modelVersionId) + " for ModelVersion on Forecast"

		try:
			# get the Forecast from db
			forecast = self.get( forecastId ).first()	
			
			# get the ModelVersion from db
			modelVersion = ModelVersionDelegate().get(modelVersionId).first();
			
			# assign the ModelVersion		
			forecast.modelVersion = modelVersion
			
			#save it
			forecast.save()

			# reload and return the appropriate version					
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModelVersion( self, forecastId ):
		errMsg = "Failed to unassign element " + str(modelVersionId) + " for ModelVersion on Forecast"

		try:
			# get the Forecast from db
			forecast = self.get( forecastId ).first()	
			
			# assign to None for unassignment
			forecast.modelVersion = None			

			#save it
			forecast.save()

			# reload and return the appropriate version					
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTimeSeries( self, forecastId, timeSeriesId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TimeSeriesDelegate import TimeSeriesDelegate

		errMsg = "Failed to assign element " + str(timeSeriesId) + " for TimeSeries on Forecast"

		try:
			# get the Forecast from db
			forecast = self.get( forecastId ).first()	
			
			# get the TimeSeries from db
			timeSeries = TimeSeriesDelegate().get(timeSeriesId).first();
			
			# assign the TimeSeries		
			forecast.timeSeries = timeSeries
			
			#save it
			forecast.save()

			# reload and return the appropriate version					
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTimeSeries( self, forecastId ):
		errMsg = "Failed to unassign element " + str(timeSeriesId) + " for TimeSeries on Forecast"

		try:
			# get the Forecast from db
			forecast = self.get( forecastId ).first()	
			
			# assign to None for unassignment
			forecast.timeSeries = None			

			#save it
			forecast.save()

			# reload and return the appropriate version					
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except Exception:
			return None;
		
	def addDatasets( self, forecastId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Forecast"

		try:
			# get the Forecast
			forecast = self.get( forecastId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				forecast.datasets.add(dataSet)
				
			# save it		
			forecast.save()
			
			# reload and return the appropriate version
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, forecastId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Forecast"

		try:
			# get the Forecast
			forecast = self.get( forecastId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				forecast.datasets.remove(dataSet)
				
			# save it		
			forecast.save()
			
			# reload and return the appropriate version
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
