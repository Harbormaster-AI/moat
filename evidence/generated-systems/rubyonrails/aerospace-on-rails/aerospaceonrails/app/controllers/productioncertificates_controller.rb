class ProductionCertificatesController < ApplicationController
  def index
    @productionCertificates = ProductionCertificate.all
  end
 
  def show
    @productionCertificate = ProductionCertificate.find(params[:id])
  end
 
  def new
    @productionCertificate = ProductionCertificate.new
  end
 
  def edit
    @productionCertificate = ProductionCertificate.find(params[:id])
  end
 
  def create
    @productionCertificate = ProductionCertificate.new(productionCertificate_params)
 
    if @productionCertificate.save
      redirect_to productionCertificates_path
    else
      render 'new'
    end
  end
 
  def update
    @productionCertificate = ProductionCertificate.find(params[:id])
 
    if @productionCertificate.update(productionCertificate_params)
      redirect_to productionCertificates_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @productionCertificate = ProductionCertificate.find(params[:id])
    @productionCertificate.destroy
    redirect_to productionCertificates_path
  end

 
  private
    def productionCertificate_params
      params.require(:productionCertificate).permit(:certificateNumber, :authority)
    end
end