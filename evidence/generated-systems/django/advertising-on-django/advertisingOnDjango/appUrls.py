"""mainsite URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include
urlpatterns = [
    path('Agency/', include('advertisingOnDjango.urls.AgencyUrls')),
    path('Team/', include('advertisingOnDjango.urls.TeamUrls')),
    path('User/', include('advertisingOnDjango.urls.UserUrls')),
    path('Advertiser/', include('advertisingOnDjango.urls.AdvertiserUrls')),
    path('BillingProfile/', include('advertisingOnDjango.urls.BillingProfileUrls')),
    path('PaymentMethod/', include('advertisingOnDjango.urls.PaymentMethodUrls')),
    path('AdAccount/', include('advertisingOnDjango.urls.AdAccountUrls')),
    path('DSP/', include('advertisingOnDjango.urls.DSPUrls')),
    path('Campaign/', include('advertisingOnDjango.urls.CampaignUrls')),
    path('KPI/', include('advertisingOnDjango.urls.KPIUrls')),
    path('AudienceSegment/', include('advertisingOnDjango.urls.AudienceSegmentUrls')),
    path('DataProvider/', include('advertisingOnDjango.urls.DataProviderUrls')),
    path('LineItem/', include('advertisingOnDjango.urls.LineItemUrls')),
    path('TargetingProfile/', include('advertisingOnDjango.urls.TargetingProfileUrls')),
    path('DeviceCriterion/', include('advertisingOnDjango.urls.DeviceCriterionUrls')),
    path('BrandSafetyPolicy/', include('advertisingOnDjango.urls.BrandSafetyPolicyUrls')),
    path('ContentCategory/', include('advertisingOnDjango.urls.ContentCategoryUrls')),
    path('Publisher/', include('advertisingOnDjango.urls.PublisherUrls')),
    path('InventorySource/', include('advertisingOnDjango.urls.InventorySourceUrls')),
    path('AdSlot/', include('advertisingOnDjango.urls.AdSlotUrls')),
    path('Deal/', include('advertisingOnDjango.urls.DealUrls')),
    path('Placement/', include('advertisingOnDjango.urls.PlacementUrls')),
    path('CreativeAsset/', include('advertisingOnDjango.urls.CreativeAssetUrls')),
    path('CreativeFile/', include('advertisingOnDjango.urls.CreativeFileUrls')),
    path('CreativeVariation/', include('advertisingOnDjango.urls.CreativeVariationUrls')),
    path('CreativeApproval/', include('advertisingOnDjango.urls.CreativeApprovalUrls')),
    path('TrackingPixel/', include('advertisingOnDjango.urls.TrackingPixelUrls')),
    path('ConversionEvent/', include('advertisingOnDjango.urls.ConversionEventUrls')),
    path('PerformanceMetric/', include('advertisingOnDjango.urls.PerformanceMetricUrls')),
    path('Report/', include('advertisingOnDjango.urls.ReportUrls')),
    path('InsertionOrder/', include('advertisingOnDjango.urls.InsertionOrderUrls')),
    path('RateCard/', include('advertisingOnDjango.urls.RateCardUrls')),
    path('Rate/', include('advertisingOnDjango.urls.RateUrls')),
    path('Experiment/', include('advertisingOnDjango.urls.ExperimentUrls')),
    path('ExperimentVariant/', include('advertisingOnDjango.urls.ExperimentVariantUrls')),
    path('GeoRegion/', include('advertisingOnDjango.urls.GeoRegionUrls')),
    path('admin/', admin.site.urls),
    path('', admin.site.urls),
]