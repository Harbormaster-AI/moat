from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Anomaly import Anomaly
from analyticsOnDjango.models.TimeSeries import TimeSeries
from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Anomaly
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnomalyDelegate Declaration
#======================================================================
class AnomalyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, anomalyId ):
		try:	
			anomaly = Anomaly.objects.filter(id=anomalyId)
			return anomaly.first();
		except Anomaly.DoesNotExist:
			raise ProcessingError("Anomaly with id " + str(anomalyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, anomaly):
		for model in serializers.deserialize("json", anomaly):
			model.save()
			return model;

	def create(self, anomaly):
		anomaly.save()
		return anomaly;

	def saveFromJson(self, anomaly):
		for model in serializers.deserialize("json", anomaly):
			model.save()
			return anomaly;
	
	def save(self, anomaly):
		anomaly.save()
		return anomaly;
	
	def delete(self, anomalyId ):
		errMsg = "Failed to delete Anomaly from db using id " + str(anomalyId)
		
		try:
			anomaly = Anomaly.objects.get(id=anomalyId)
			anomaly.delete()
			return True
		except Anomaly.DoesNotExist:
			raise ProcessingError("Anomaly with id " + str(anomalyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Anomaly.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Anomaly from db")
		except Exception:
			return None;
		
	def assignTimeSeries( self, anomalyId, timeSeriesId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TimeSeriesDelegate import TimeSeriesDelegate

		errMsg = "Failed to assign element " + str(timeSeriesId) + " for TimeSeries on Anomaly"

		try:
			# get the Anomaly from db
			anomaly = self.get( anomalyId ).first()	
			
			# get the TimeSeries from db
			timeSeries = TimeSeriesDelegate().get(timeSeriesId).first();
			
			# assign the TimeSeries		
			anomaly.timeSeries = timeSeries
			
			#save it
			anomaly.save()

			# reload and return the appropriate version					
			return self.get( anomalyId );
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly with id " + str(anomalyId) + " does not exist.")
		except TimeSeries.DoesNotExist:
			raise ProcessingError(errMsg + " : TimeSeries with id " + str(timeSeriesId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTimeSeries( self, anomalyId ):
		errMsg = "Failed to unassign element " + str(timeSeriesId) + " for TimeSeries on Anomaly"

		try:
			# get the Anomaly from db
			anomaly = self.get( anomalyId ).first()	
			
			# assign to None for unassignment
			anomaly.timeSeries = None			

			#save it
			anomaly.save()

			# reload and return the appropriate version					
			return self.get( anomalyId );
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly with id " + str(anomalyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAlert( self, anomalyId, alertId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to assign element " + str(alertId) + " for Alert on Anomaly"

		try:
			# get the Anomaly from db
			anomaly = self.get( anomalyId ).first()	
			
			# get the Alert from db
			alert = AlertDelegate().get(alertId).first();
			
			# assign the Alert		
			anomaly.alert = alert
			
			#save it
			anomaly.save()

			# reload and return the appropriate version					
			return self.get( anomalyId );
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly with id " + str(anomalyId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAlert( self, anomalyId ):
		errMsg = "Failed to unassign element " + str(alertId) + " for Alert on Anomaly"

		try:
			# get the Anomaly from db
			anomaly = self.get( anomalyId ).first()	
			
			# assign to None for unassignment
			anomaly.alert = None			

			#save it
			anomaly.save()

			# reload and return the appropriate version					
			return self.get( anomalyId );
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly with id " + str(anomalyId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDataset( self, anomalyId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on Anomaly"

		try:
			# get the Anomaly from db
			anomaly = self.get( anomalyId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			anomaly.dataset = dataSet
			
			#save it
			anomaly.save()

			# reload and return the appropriate version					
			return self.get( anomalyId );
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly with id " + str(anomalyId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, anomalyId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on Anomaly"

		try:
			# get the Anomaly from db
			anomaly = self.get( anomalyId ).first()	
			
			# assign to None for unassignment
			anomaly.dataSet = None			

			#save it
			anomaly.save()

			# reload and return the appropriate version					
			return self.get( anomalyId );
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly with id " + str(anomalyId) + " does not exist.")
		except Exception:
			return None;
		
