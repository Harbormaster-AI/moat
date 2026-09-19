class CreditorsController < ApplicationController
  def index
    @creditors = Creditor.all
  end
 
  def show
    @creditor = Creditor.find(params[:id])
  end
 
  def new
    @creditor = Creditor.new
  end
 
  def edit
    @creditor = Creditor.find(params[:id])
  end
 
  def create
    @creditor = Creditor.new(creditor_params)
 
    if @creditor.save
      redirect_to creditors_path
    else
      render 'new'
    end
  end
 
  def update
    @creditor = Creditor.find(params[:id])
 
    if @creditor.update(creditor_params)
      redirect_to creditors_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @creditor = Creditor.find(params[:id])
    @creditor.destroy
    redirect_to creditors_path
  end

 
  private
    def creditor_params
      params.require(:creditor).permit(:name, :bic, :address)
    end
end