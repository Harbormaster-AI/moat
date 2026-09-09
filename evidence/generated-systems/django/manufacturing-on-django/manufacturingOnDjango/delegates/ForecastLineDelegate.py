from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.ForecastLine import ForecastLine
from manufacturingOnDjango.models.Forecast import Forecast
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ForecastLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ForecastLineDelegate Declaration
#======================================================================
class ForecastLineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, forecastLineId ):
		try:	
			forecastLine = ForecastLine.objects.filter(id=forecastLineId)
			return forecastLine.first();
		except ForecastLine.DoesNotExist:
			raise ProcessingError("ForecastLine with id " + str(forecastLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, forecastLine):
		for model in serializers.deserialize("json", forecastLine):
			model.save()
			return model;

	def create(self, forecastLine):
		forecastLine.save()
		return forecastLine;

	def saveFromJson(self, forecastLine):
		for model in serializers.deserialize("json", forecastLine):
			model.save()
			return forecastLine;
	
	def save(self, forecastLine):
		forecastLine.save()
		return forecastLine;
	
	def delete(self, forecastLineId ):
		errMsg = "Failed to delete ForecastLine from db using id " + str(forecastLineId)
		
		try:
			forecastLine = ForecastLine.objects.get(id=forecastLineId)
			forecastLine.delete()
			return True
		except ForecastLine.DoesNotExist:
			raise ProcessingError("ForecastLine with id " + str(forecastLineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ForecastLine.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ForecastLine from db")
		except Exception:
			return None;
		
	def assignForecast( self, forecastLineId, forecastId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ForecastDelegate import ForecastDelegate

		errMsg = "Failed to assign element " + str(forecastId) + " for Forecast on ForecastLine"

		try:
			# get the ForecastLine from db
			forecastLine = self.get( forecastLineId ).first()	
			
			# get the Forecast from db
			forecast = ForecastDelegate().get(forecastId).first();
			
			# assign the Forecast		
			forecastLine.forecast = forecast
			
			#save it
			forecastLine.save()

			# reload and return the appropriate version					
			return self.get( forecastLineId );
		except ForecastLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ForecastLine with id " + str(forecastLineId) + " does not exist.")
		except Forecast.DoesNotExist:
			raise ProcessingError(errMsg + " : Forecast with id " + str(forecastId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignForecast( self, forecastLineId ):
		errMsg = "Failed to unassign element " + str(forecastId) + " for Forecast on ForecastLine"

		try:
			# get the ForecastLine from db
			forecastLine = self.get( forecastLineId ).first()	
			
			# assign to None for unassignment
			forecastLine.forecast = None			

			#save it
			forecastLine.save()

			# reload and return the appropriate version					
			return self.get( forecastLineId );
		except ForecastLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ForecastLine with id " + str(forecastLineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignItem( self, forecastLineId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on ForecastLine"

		try:
			# get the ForecastLine from db
			forecastLine = self.get( forecastLineId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			forecastLine.item = item
			
			#save it
			forecastLine.save()

			# reload and return the appropriate version					
			return self.get( forecastLineId );
		except ForecastLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ForecastLine with id " + str(forecastLineId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, forecastLineId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on ForecastLine"

		try:
			# get the ForecastLine from db
			forecastLine = self.get( forecastLineId ).first()	
			
			# assign to None for unassignment
			forecastLine.item = None			

			#save it
			forecastLine.save()

			# reload and return the appropriate version					
			return self.get( forecastLineId );
		except ForecastLine.DoesNotExist:
			raise ProcessingError(errMsg + " : ForecastLine with id " + str(forecastLineId) + " does not exist.")
		except Exception:
			return None;
		
