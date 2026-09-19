class TaxWithholdingsController < ApplicationController
  def index
    @taxWithholdings = TaxWithholding.all
  end
 
  def show
    @taxWithholding = TaxWithholding.find(params[:id])
  end
 
  def new
    @taxWithholding = TaxWithholding.new
  end
 
  def edit
    @taxWithholding = TaxWithholding.find(params[:id])
  end
 
  def create
    @taxWithholding = TaxWithholding.new(taxWithholding_params)
 
    if @taxWithholding.save
      redirect_to taxWithholdings_path
    else
      render 'new'
    end
  end
 
  def update
    @taxWithholding = TaxWithholding.find(params[:id])
 
    if @taxWithholding.update(taxWithholding_params)
      redirect_to taxWithholdings_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @taxWithholding = TaxWithholding.find(params[:id])
    @taxWithholding.destroy
    redirect_to taxWithholdings_path
  end

 
  private
    def taxWithholding_params
      params.require(:taxWithholding).permit(:taxId, :allowances, :additionalAmount, :FilingStatus)
    end
end