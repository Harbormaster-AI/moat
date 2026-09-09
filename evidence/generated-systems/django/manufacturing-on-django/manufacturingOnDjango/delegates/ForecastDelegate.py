from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Forecast import Forecast
from manufacturingOnDjango.models.ForecastLine import ForecastLine
from manufacturingOnDjango.exceptions import Exceptions

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
		
	def addLines( self, forecastId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ForecastLineDelegate import ForecastLineDelegate

		errMsg = "Failed to add elements " + str(linesIds) + " for Lines on Forecast"

		try:
			# get the Forecast
			forecast = self.get( forecastId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ForecastLine		
				forecastLine = ForecastLineDelegate().get(id).first();	
				# add the ForecastLine
				forecast.lines.add(forecastLine)
				
			# save it		
			forecast.save()
			
			# reload and return the appropriate version
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except ForecastLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ForecastLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLines( self, forecastId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ForecastLineDelegate import ForecastLineDelegate

		errMsg = "Failed to remove elements " + str(linesIds) + " for Lines on Forecast"

		try:
			# get the Forecast
			forecast = self.get( forecastId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ForecastLine		
				forecastLine = ForecastLineDelegate().get(id).first();	
				# add the ForecastLine
				forecast.lines.remove(forecastLine)
				
			# save it		
			forecast.save()
			
			# reload and return the appropriate version
			return self.get( forecastId );
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except ForecastLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ForecastLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
