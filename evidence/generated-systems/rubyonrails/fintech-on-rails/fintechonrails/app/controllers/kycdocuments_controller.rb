class KYCDocumentsController < ApplicationController
  def index
    @kYCDocuments = KYCDocument.all
  end
 
  def show
    @kYCDocument = KYCDocument.find(params[:id])
  end
 
  def new
    @kYCDocument = KYCDocument.new
  end
 
  def edit
    @kYCDocument = KYCDocument.find(params[:id])
  end
 
  def create
    @kYCDocument = KYCDocument.new(kYCDocument_params)
 
    if @kYCDocument.save
      redirect_to kYCDocuments_path
    else
      render 'new'
    end
  end
 
  def update
    @kYCDocument = KYCDocument.find(params[:id])
 
    if @kYCDocument.update(kYCDocument_params)
      redirect_to kYCDocuments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @kYCDocument = KYCDocument.find(params[:id])
    @kYCDocument.destroy
    redirect_to kYCDocuments_path
  end

 
  private
    def kYCDocument_params
      params.require(:kYCDocument).permit(:reference, :issuedCountry, :expirationDate, :DocumentType, :Status)
    end
end