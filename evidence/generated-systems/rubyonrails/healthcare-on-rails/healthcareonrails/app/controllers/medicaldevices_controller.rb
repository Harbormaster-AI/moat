class MedicalDevicesController < ApplicationController
  def index
    @medicalDevices = MedicalDevice.all
  end
 
  def show
    @medicalDevice = MedicalDevice.find(params[:id])
  end
 
  def new
    @medicalDevice = MedicalDevice.new
  end
 
  def edit
    @medicalDevice = MedicalDevice.find(params[:id])
  end
 
  def create
    @medicalDevice = MedicalDevice.new(medicalDevice_params)
 
    if @medicalDevice.save
      redirect_to medicalDevices_path
    else
      render 'new'
    end
  end
 
  def update
    @medicalDevice = MedicalDevice.find(params[:id])
 
    if @medicalDevice.update(medicalDevice_params)
      redirect_to medicalDevices_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @medicalDevice = MedicalDevice.find(params[:id])
    @medicalDevice.destroy
    redirect_to medicalDevices_path
  end

 
  private
    def medicalDevice_params
      params.require(:medicalDevice).permit(:udi, :manufacturer, :DeviceType, :ConnectivityStatus)
    end
end