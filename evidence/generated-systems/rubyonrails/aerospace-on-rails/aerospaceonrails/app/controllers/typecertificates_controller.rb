class TypeCertificatesController < ApplicationController
  def index
    @typeCertificates = TypeCertificate.all
  end
 
  def show
    @typeCertificate = TypeCertificate.find(params[:id])
  end
 
  def new
    @typeCertificate = TypeCertificate.new
  end
 
  def edit
    @typeCertificate = TypeCertificate.find(params[:id])
  end
 
  def create
    @typeCertificate = TypeCertificate.new(typeCertificate_params)
 
    if @typeCertificate.save
      redirect_to typeCertificates_path
    else
      render 'new'
    end
  end
 
  def update
    @typeCertificate = TypeCertificate.find(params[:id])
 
    if @typeCertificate.update(typeCertificate_params)
      redirect_to typeCertificates_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @typeCertificate = TypeCertificate.find(params[:id])
    @typeCertificate.destroy
    redirect_to typeCertificates_path
  end

 
  private
    def typeCertificate_params
      params.require(:typeCertificate).permit(:certificateNumber, :authority)
    end
end