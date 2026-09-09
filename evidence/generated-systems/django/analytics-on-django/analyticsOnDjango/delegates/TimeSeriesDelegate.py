from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.TimeSeries import TimeSeries
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Forecast import Forecast
from analyticsOnDjango.models.Anomaly import Anomaly
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TimeSeries
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimeSeriesDelegate Declaration
#======================================================================
class TimeSeriesDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, timeSeriesId ):
		try:	
			timeSeries = TimeSeries.objects.filter(id=timeSeriesId)
			return timeSeries.first();
		except TimeSeries.DoesNotExist:
			raise ProcessingError("TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, timeSeries):
		for model in serializers.deserialize("json", timeSeries):
			model.save()
			return model;

	def create(self, timeSeries):
		timeSeries.save()
		return timeSeries;

	def saveFromJson(self, timeSeries):
		for model in serializers.deserialize("json", timeSeries):
			model.save()
			return timeSeries;
	
	def save(self, timeSeries):
		timeSeries.save()
		return timeSeries;
	
	def delete(self, timeSeriesId ):
		errMsg = "Failed to delete TimeSeries from db using id " + str(timeSeriesId)
		
		try:
			timeSeries = TimeSeries.objects.get(id=timeSeriesId)
			timeSeries.delete()
			return True
		except TimeSeries.DoesNotExist:
			raise ProcessingError("TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TimeSeries.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TimeSeries from db")
		except Exception:
			return None;
		
	def addDatasets( self, timeSeriesId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on TimeSeries"

		try:
			# get the TimeSeries
			timeSeries = self.get( timeSeriesId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				timeSeries.datasets.add(dataSet)
				
			# save it		
			timeSeries.save()
			
			# reload and return the appropriate version
			return self.get( timeSeriesId );
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, timeSeriesId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on TimeSeries"

		try:
			# get the TimeSeries
			timeSeries = self.get( timeSeriesId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				timeSeries.datasets.remove(dataSet)
				
			# save it		
			timeSeries.save()
			
			# reload and return the appropriate version
			return self.get( timeSeriesId );
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addForecasts( self, timeSeriesId, forecastsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ForecastDelegate import ForecastDelegate

		errMsg = "Failed to add elements " + str(forecastsIds) + " for Forecasts on TimeSeries"

		try:
			# get the TimeSeries
			timeSeries = self.get( timeSeriesId ).first()
				
			# split on a comma with no spaces
			idList = forecastsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Forecast		
				forecast = ForecastDelegate().get(id).first();	
				# add the Forecast
				timeSeries.forecasts.add(forecast)
				
			# save it		
			timeSeries.save()
			
			# reload and return the appropriate version
			return self.get( timeSeriesId );
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeForecasts( self, timeSeriesId, forecastsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ForecastDelegate import ForecastDelegate

		errMsg = "Failed to remove elements " + str(forecastsIds) + " for Forecasts on TimeSeries"

		try:
			# get the TimeSeries
			timeSeries = self.get( timeSeriesId ).first()
				
			# split on a comma with no spaces
			idList = forecastsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Forecast		
				forecast = ForecastDelegate().get(id).first();	
				# add the Forecast
				timeSeries.forecasts.remove(forecast)
				
			# save it		
			timeSeries.save()
			
			# reload and return the appropriate version
			return self.get( timeSeriesId );
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAnomalies( self, timeSeriesId, anomaliesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnomalyDelegate import AnomalyDelegate

		errMsg = "Failed to add elements " + str(anomaliesIds) + " for Anomalies on TimeSeries"

		try:
			# get the TimeSeries
			timeSeries = self.get( timeSeriesId ).first()
				
			# split on a comma with no spaces
			idList = anomaliesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Anomaly		
				anomaly = AnomalyDelegate().get(id).first();	
				# add the Anomaly
				timeSeries.anomalies.add(anomaly)
				
			# save it		
			timeSeries.save()
			
			# reload and return the appropriate version
			return self.get( timeSeriesId );
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAnomalies( self, timeSeriesId, anomaliesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnomalyDelegate import AnomalyDelegate

		errMsg = "Failed to remove elements " + str(anomaliesIds) + " for Anomalies on TimeSeries"

		try:
			# get the TimeSeries
			timeSeries = self.get( timeSeriesId ).first()
				
			# split on a comma with no spaces
			idList = anomaliesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Anomaly		
				anomaly = AnomalyDelegate().get(id).first();	
				# add the Anomaly
				timeSeries.anomalies.remove(anomaly)
				
			# save it		
			timeSeries.save()
			
			# reload and return the appropriate version
			return self.get( timeSeriesId );
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
