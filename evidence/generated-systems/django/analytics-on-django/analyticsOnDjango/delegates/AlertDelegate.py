from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.QualityRule import QualityRule
from analyticsOnDjango.models.Anomaly import Anomaly
from analyticsOnDjango.models.Subscriber import Subscriber
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Alert
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AlertDelegate Declaration
#======================================================================
class AlertDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, alertId ):
		try:	
			alert = Alert.objects.filter(id=alertId)
			return alert.first();
		except Alert.DoesNotExist:
			raise ProcessingError("Alert with id " + str(alertId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, alert):
		for model in serializers.deserialize("json", alert):
			model.save()
			return model;

	def create(self, alert):
		alert.save()
		return alert;

	def saveFromJson(self, alert):
		for model in serializers.deserialize("json", alert):
			model.save()
			return alert;
	
	def save(self, alert):
		alert.save()
		return alert;
	
	def delete(self, alertId ):
		errMsg = "Failed to delete Alert from db using id " + str(alertId)
		
		try:
			alert = Alert.objects.get(id=alertId)
			alert.delete()
			return True
		except Alert.DoesNotExist:
			raise ProcessingError("Alert with id " + str(alertId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Alert.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Alert from db")
		except Exception:
			return None;
		
	def assignMetric( self, alertId, metricId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to assign element " + str(metricId) + " for Metric on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# get the Metric from db
			metric = MetricDelegate().get(metricId).first();
			
			# assign the Metric		
			alert.metric = metric
			
			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMetric( self, alertId ):
		errMsg = "Failed to unassign element " + str(metricId) + " for Metric on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# assign to None for unassignment
			alert.metric = None			

			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDashboard( self, alertId, dashboardId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to assign element " + str(dashboardId) + " for Dashboard on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# get the Dashboard from db
			dashboard = DashboardDelegate().get(dashboardId).first();
			
			# assign the Dashboard		
			alert.dashboard = dashboard
			
			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDashboard( self, alertId ):
		errMsg = "Failed to unassign element " + str(dashboardId) + " for Dashboard on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# assign to None for unassignment
			alert.dashboard = None			

			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDataset( self, alertId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			alert.dataset = dataSet
			
			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, alertId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# assign to None for unassignment
			alert.dataSet = None			

			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRule( self, alertId, ruleId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.QualityRuleDelegate import QualityRuleDelegate

		errMsg = "Failed to assign element " + str(ruleId) + " for Rule on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# get the QualityRule from db
			qualityRule = QualityRuleDelegate().get(ruleId).first();
			
			# assign the Rule		
			alert.rule = qualityRule
			
			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule with id " + str(ruleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRule( self, alertId ):
		errMsg = "Failed to unassign element " + str(ruleId) + " for Rule on Alert"

		try:
			# get the Alert from db
			alert = self.get( alertId ).first()	
			
			# assign to None for unassignment
			alert.qualityRule = None			

			#save it
			alert.save()

			# reload and return the appropriate version					
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Exception:
			return None;
		
	def addAnomalies( self, alertId, anomaliesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnomalyDelegate import AnomalyDelegate

		errMsg = "Failed to add elements " + str(anomaliesIds) + " for Anomalies on Alert"

		try:
			# get the Alert
			alert = self.get( alertId ).first()
				
			# split on a comma with no spaces
			idList = anomaliesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Anomaly		
				anomaly = AnomalyDelegate().get(id).first();	
				# add the Anomaly
				alert.anomalies.add(anomaly)
				
			# save it		
			alert.save()
			
			# reload and return the appropriate version
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAnomalies( self, alertId, anomaliesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnomalyDelegate import AnomalyDelegate

		errMsg = "Failed to remove elements " + str(anomaliesIds) + " for Anomalies on Alert"

		try:
			# get the Alert
			alert = self.get( alertId ).first()
				
			# split on a comma with no spaces
			idList = anomaliesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Anomaly		
				anomaly = AnomalyDelegate().get(id).first();	
				# add the Anomaly
				alert.anomalies.remove(anomaly)
				
			# save it		
			alert.save()
			
			# reload and return the appropriate version
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Anomaly.DoesNotExist:
			raise ProcessingError(errMsg + " : Anomaly does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSubscribers( self, alertId, subscribersIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SubscriberDelegate import SubscriberDelegate

		errMsg = "Failed to add elements " + str(subscribersIds) + " for Subscribers on Alert"

		try:
			# get the Alert
			alert = self.get( alertId ).first()
				
			# split on a comma with no spaces
			idList = subscribersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Subscriber		
				subscriber = SubscriberDelegate().get(id).first();	
				# add the Subscriber
				alert.subscribers.add(subscriber)
				
			# save it		
			alert.save()
			
			# reload and return the appropriate version
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Subscriber.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscriber does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSubscribers( self, alertId, subscribersIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SubscriberDelegate import SubscriberDelegate

		errMsg = "Failed to remove elements " + str(subscribersIds) + " for Subscribers on Alert"

		try:
			# get the Alert
			alert = self.get( alertId ).first()
				
			# split on a comma with no spaces
			idList = subscribersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Subscriber		
				subscriber = SubscriberDelegate().get(id).first();	
				# add the Subscriber
				alert.subscribers.remove(subscriber)
				
			# save it		
			alert.save()
			
			# reload and return the appropriate version
			return self.get( alertId );
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert with id " + str(alertId) + " does not exist.")
		except Subscriber.DoesNotExist:
			raise ProcessingError(errMsg + " : Subscriber does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
