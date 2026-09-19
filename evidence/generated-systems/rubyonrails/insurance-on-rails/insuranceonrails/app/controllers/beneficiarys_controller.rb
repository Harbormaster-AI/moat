class BeneficiarysController < ApplicationController
  def index
    @beneficiarys = Beneficiary.all
  end
 
  def show
    @beneficiary = Beneficiary.find(params[:id])
  end
 
  def new
    @beneficiary = Beneficiary.new
  end
 
  def edit
    @beneficiary = Beneficiary.find(params[:id])
  end
 
  def create
    @beneficiary = Beneficiary.new(beneficiary_params)
 
    if @beneficiary.save
      redirect_to beneficiarys_path
    else
      render 'new'
    end
  end
 
  def update
    @beneficiary = Beneficiary.find(params[:id])
 
    if @beneficiary.update(beneficiary_params)
      redirect_to beneficiarys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @beneficiary = Beneficiary.find(params[:id])
    @beneficiary.destroy
    redirect_to beneficiarys_path
  end

 
  private
    def beneficiary_params
      params.require(:beneficiary).permit(:name, :share, :Relationship)
    end
end