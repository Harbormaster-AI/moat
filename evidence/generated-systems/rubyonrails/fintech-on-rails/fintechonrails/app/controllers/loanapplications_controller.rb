class LoanApplicationsController < ApplicationController
  def index
    @loanApplications = LoanApplication.all
  end
 
  def show
    @loanApplication = LoanApplication.find(params[:id])
  end
 
  def new
    @loanApplication = LoanApplication.new
  end
 
  def edit
    @loanApplication = LoanApplication.find(params[:id])
  end
 
  def create
    @loanApplication = LoanApplication.new(loanApplication_params)
 
    if @loanApplication.save
      redirect_to loanApplications_path
    else
      render 'new'
    end
  end
 
  def update
    @loanApplication = LoanApplication.find(params[:id])
 
    if @loanApplication.update(loanApplication_params)
      redirect_to loanApplications_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @loanApplication = LoanApplication.find(params[:id])
    @loanApplication.destroy
    redirect_to loanApplications_path
  end

 
  private
    def loanApplication_params
      params.require(:loanApplication).permit(:applicationNumber, :amountRequested, :termMonths, :submittedAt, :Product, :Purpose, :Status)
    end
end